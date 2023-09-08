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
 */

package terminalRecord

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"k8s.io/klog/v2"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-webconsole/console/audit/asciinema"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-webconsole/console/config"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-webconsole/console/types"
)

const (
	dateTimeFormat = "2006-01-02"
	dayTimeFormat  = "150405"

	replayFilenameSuffix = ".cast"
)

// ReplyInfo 回访记录初始信息
type ReplyInfo struct {
	Width     uint16
	Height    uint16
	TimeStamp time.Time
}

// ReplyRecorder 终端回放记录器
type ReplyRecorder struct {
	SessionID   string
	Info        *ReplyInfo
	absFilePath string
	//Target      string
	Writer *asciinema.Writer
	Err    error
	ctx    context.Context

	file *os.File
	once sync.Once
}

func fileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func ensureDirExist(name string) error {
	if !fileExists(name) {
		return os.MkdirAll(name, os.ModePerm)
	}
	return nil
}

// NewReplayRecord 初始化Recorder
// 确认是否开启终端记录 / 创建记录文件 / 初始记录信息
func NewReplayRecord(ctx context.Context, podCtx *types.PodContext, originTerminalSize *ReplyInfo) *ReplyRecorder {
	if !config.G.Audit.Enabled {
		return nil
	}
	recorder := &ReplyRecorder{
		ctx:       ctx,
		SessionID: podCtx.SessionId,
		Info:      originTerminalSize,
	}
	date := time.Now().Format(dateTimeFormat)
	path := config.G.Audit.DataDir
	path = filepath.Join(path, date)
	err := ensureDirExist(path)
	if err != nil {
		klog.Errorf("Create dir %s error: %s\n", path, err)
		recorder.Err = err
		return recorder
	}
	d := time.Now().Format(dayTimeFormat)
	f := fmt.Sprintf("%s_%s_%s_%s", d, podCtx.ClusterId, podCtx.Username, podCtx.SessionId[:6])
	//f := fmt.Sprintf("%s_%s_%s_%s_%s_%s_%s", date, podCtx.Username, podCtx.ClusterId, podCtx.Namespace, podCtx.PodName,
	//	podCtx.ContainerName, podCtx.SessionId[:6])
	filename := f + replayFilenameSuffix
	absFilePath := filepath.Join(path, filename)
	recorder.absFilePath = absFilePath
	fd, err := os.Create(recorder.absFilePath)
	if err != nil {
		klog.Errorf("Create replay file %s error: %s\n", recorder.absFilePath, err)
		recorder.Err = err
		return recorder
	}
	recorder.file = fd
	options := make([]asciinema.Option, 0, 3)
	options = append(options, asciinema.WithHeight(originTerminalSize.Height))
	options = append(options, asciinema.WithWidth(originTerminalSize.Width))
	options = append(options, asciinema.WithTimestamp(originTerminalSize.TimeStamp))
	recorder.Writer = asciinema.NewWriter(recorder.file, podCtx, options...)
	return recorder
}

// isNullError 记录异常
func (r *ReplyRecorder) isNullError() bool {
	if r.Err != nil {
		r.once.Do(func() {
			//异常退出: 直接关闭文件
			r.file.Close()
		})
		return true
	}
	return false
}

// Record 记录终端输出信息
func Record(r *ReplyRecorder, p []byte, event asciinema.EventType) {
	//不开启terminal recorder时, ReplyRecorder返回nil
	if r == nil {
		return
	}
	//有错误异常就退出本次记录
	if r.isNullError() {
		return
	}
	if len(p) > 0 {
		r.once.Do(func() {
			if err := r.Writer.WriteHeader(); err != nil {
				r.Err = err
				klog.Errorf("Session %s write replay header failed: %s", r.SessionID, err)
			}
		})
		if err := r.Writer.WriteRow(p, event); err != nil {
			r.Err = err
			klog.Errorf("Session %s write replay row failed: %s", r.SessionID, err)
		}
	}
}

// End 正常退出: 关闭缓存和文件
func (r *ReplyRecorder) End() {
	if r == nil {
		return
	} else {
		//关闭前将剩余缓冲区数据写入
		r.Writer.Write(r.Writer.WriteBuff)
		r.file.Close()
		return
	}
}

// GracefulShutdownRecorder 关闭文件
func (r *ReplyRecorder) GracefulShutdownRecorder() {
	r.Writer.Write(r.Writer.WriteBuff)
	r.file.Close()
}
