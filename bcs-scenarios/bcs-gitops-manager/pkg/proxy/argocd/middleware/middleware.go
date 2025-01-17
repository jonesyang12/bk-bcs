/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package middleware defines the middleware for gitops
package middleware

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	traceconst "github.com/Tencent/bk-bcs/bcs-common/pkg/otel/trace/constants"
	"github.com/Tencent/bk-bcs/bcs-scenarios/bcs-gitops-manager/pkg/metric"
	"github.com/Tencent/bk-bcs/bcs-scenarios/bcs-gitops-manager/pkg/proxy"
	"github.com/Tencent/bk-bcs/bcs-scenarios/bcs-gitops-manager/pkg/proxy/argocd/session"
	"github.com/Tencent/bk-bcs/bcs-scenarios/bcs-gitops-manager/pkg/utils"
	"github.com/Tencent/bk-bcs/bcs-services/cluster-resources/pkg/tracing"
)

type httpHandler func(ctx context.Context, r *http.Request) *HttpResponse

type httpWrapper struct {
	handler       httpHandler
	handlerName   string
	option        *proxy.GitOpsOptions
	argoSession   *session.ArgoSession
	secretSession *session.SecretSession
}

// HttpResponse 定义了返回信息，根据返回信息 httpWrapper 做对应处理
type HttpResponse struct {
	respType   responseType
	obj        interface{}
	statusCode int
	err        error
}

type contextKey string

const (
	ctxKeyUser contextKey = "user"
)

// RequestID return the requestID of context
func RequestID(ctx context.Context) string {
	return ctx.Value(traceconst.RequestIDHeaderKey).(string)
}

// User return user info of context
func User(ctx context.Context) *proxy.UserInfo {
	return ctx.Value(ctxKeyUser).(*proxy.UserInfo)
}

func (p *httpWrapper) setContext(rw http.ResponseWriter, r *http.Request) (context.Context, string) {
	// 获取 RequestID 信息，并重新存入上下文
	var requestID string
	requestIDHeader := r.Context().Value(traceconst.RequestIDHeaderKey)
	if v, ok := requestIDHeader.(string); ok && v != "" {
		requestID = v
	} else {
		requestID = uuid.New().String()
	}
	ctx := context.WithValue(r.Context(), traceconst.RequestIDHeaderKey, requestID)
	ctx = tracing.ContextWithRequestID(ctx, requestID)

	// 统一获取 User 信息，并存入上下文
	user, err := proxy.GetJWTInfo(r, p.option.JWTDecoder)
	if err != nil || user == nil {
		http.Error(rw, errors.Wrapf(err, "get user info failed").Error(), http.StatusUnauthorized)
		return nil, requestID
	}
	if user.ClientID != "" {
		blog.Infof("RequestID[%s] manager received user '%s' with client '%s' serve [%s/%s]",
			requestID, user.GetUser(), user.ClientID, r.Method, r.URL.Path)
	} else {
		blog.Infof("RequestID[%s] manager received user '%s' serve [%s/%s]",
			requestID, user.GetUser(), r.Method, r.URL.Path)
	}
	ctx = context.WithValue(ctx, ctxKeyUser, user)
	return ctx, requestID
}

// ServeHTTP 接收请求的入口，根据返回的 type 类型做不同的操作
func (p *httpWrapper) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	start := time.Now()
	ctx, requestID := p.setContext(rw, r)
	if ctx == nil {
		return
	}
	defer func() {
		cost := time.Since(start)
		blog.Infof("RequestID[%s] handle request '%s' cost time: %v", requestID, r.URL.Path, cost)

		// 对于包含 stream/webhook 的请求过滤
		if !strings.Contains(r.URL.Path, "/api/v1/stream") &&
			!strings.Contains(r.URL.Path, "/api/webhook") &&
			!strings.Contains(r.URL.Path, "/clean") {
			if strings.Contains(r.URL.Path, "Service/") {
				metric.ManagerGRPCRequestTotal.WithLabelValues().Inc()
				metric.ManagerGRPCRequestDuration.WithLabelValues().Observe(cost.Seconds())
			} else {
				metric.ManagerHTTPRequestTotal.WithLabelValues().Inc()
				metric.ManagerHTTPRequestDuration.WithLabelValues().Observe(cost.Seconds())
			}
		}
	}()

	resp := p.handler(ctx, r)
	blog.V(5).Infof("RequestID[%s] handler '%s' cost time: %v", requestID, p.handlerName, time.Since(start))
	if resp == nil {
		blog.Warnf("RequestID[%s] response should not be nil", requestID)
		resp = &HttpResponse{
			respType: reverseArgo,
		}
	}
	if resp.statusCode >= 500 {
		if !utils.IsContextCanceled(resp.err) {
			metric.ManagerReturnErrorNum.WithLabelValues().Inc()
		}
	}
	switch resp.respType {
	case reverseArgo:
		p.argoSession.ServeHTTP(rw, r)
	case reverseSecret:
		p.secretSession.ServeHTTP(rw, r)
	case returnError:
		blog.Warnf("RequestID[%s] handler return code '%d': %s", requestID, resp.statusCode, resp.err.Error())
		http.Error(rw, resp.err.Error(), resp.statusCode)
	case returnGrpcError:
		blog.Warnf("RequestID[%s] handler grpc request return code '%d': %s",
			requestID, resp.statusCode, resp.err.Error())
		proxy.GRPCErrorResponse(rw, resp.statusCode, resp.err)
	case grpcResponse:
		proxy.GRPCResponse(rw, resp.obj)
	case directResponse:
		proxy.DirectlyResponse(rw, resp.obj)
	case jsonResponse:
		proxy.JSONResponse(rw, resp.obj)
	}
}

type responseType int

const (
	// reverseArgo 请求反向代理给 argoCD
	reverseArgo responseType = iota
	// reverseSecret 请求反向代理给 secret 服务
	reverseSecret
	// returnError 直接返回错误给客户端
	returnError
	// returnGrpcError 返回 grpc 的错误给客户端
	returnGrpcError
	// grpcResponse 返回特殊的 GRPC 给客户端
	grpcResponse
	// directResponse 直接返回给客户端（不做 JSON/GRPC 序列化，用于 metric proxy 代理）
	directResponse
	// jsonResponse 返回 JSON 信息给客户端
	jsonResponse
)

// ReturnArgoReverse will reverse to argocd
func ReturnArgoReverse() *HttpResponse {
	return &HttpResponse{
		respType: reverseArgo,
	}
}

// ReturnSecretReverse will reverse to secret server
func ReturnSecretReverse() *HttpResponse {
	return &HttpResponse{
		respType: reverseSecret,
	}
}

// ReturnErrorResponse will return error message to client
func ReturnErrorResponse(statusCode int, err error) *HttpResponse {
	return &HttpResponse{
		respType:   returnError,
		statusCode: statusCode,
		err:        err,
	}
}

// ReturnGRPCErrorResponse 返回 rpc 的错误给客户端
func ReturnGRPCErrorResponse(statusCode int, err error) *HttpResponse {
	return &HttpResponse{
		respType:   returnGrpcError,
		statusCode: statusCode,
		err:        err,
	}
}

// ReturnJSONResponse will return response to client with json marshal
func ReturnJSONResponse(obj interface{}) *HttpResponse {
	return &HttpResponse{
		respType: jsonResponse,
		obj:      obj,
	}
}

// ReturnDirectResponse will return object to client without marshal
func ReturnDirectResponse(obj interface{}) *HttpResponse {
	return &HttpResponse{
		respType: directResponse,
		obj:      obj,
	}
}

// ReturnGRPCResponse will return response to client with grpc marshal
func ReturnGRPCResponse(obj interface{}) *HttpResponse {
	return &HttpResponse{
		respType: grpcResponse,
		obj:      obj,
	}
}
