// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/cpu_request_usage": {
            "get": {
                "tags": [
                    "Metrics"
                ],
                "summary": "集群 CPU 装箱率",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/cpu_usage": {
            "get": {
                "tags": [
                    "Metrics"
                ],
                "summary": "集群 CPU 使用率",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/disk_usage": {
            "get": {
                "tags": [
                    "Metrics"
                ],
                "summary": "集群磁盘使用率",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/diskio_usage": {
            "get": {
                "tags": [
                    "Metrics"
                ],
                "summary": "集群磁盘IO使用率",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/log_collector/entrypoints": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "LogCollectors"
                ],
                "summary": "获取日志采集规则列表",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/logcollector.Entrypoint"
                        }
                    }
                }
            }
        },
        "/log_collector/rules": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "LogCollectors"
                ],
                "summary": "获取日志采集规则列表",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.LogCollector"
                            }
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "LogCollectors"
                ],
                "summary": "创建日志采集规则",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/log_collector/rules/:name": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "LogCollectors"
                ],
                "summary": "获取日志采集规则详情",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.LogCollector"
                        }
                    }
                }
            },
            "put": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "LogCollectors"
                ],
                "summary": "更新日志采集规则",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "LogCollectors"
                ],
                "summary": "删除日志采集规则",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/memory_request_usage": {
            "get": {
                "tags": [
                    "Metrics"
                ],
                "summary": "集群内存装箱率",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/memory_usage": {
            "get": {
                "tags": [
                    "Metrics"
                ],
                "summary": "集群内存使用率",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/namespaces/:namespace/pods/:pod/containers": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Logs"
                ],
                "summary": "获取 Pod 容器列表",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/k8sclient.Container"
                            }
                        }
                    }
                }
            }
        },
        "/namespaces/:namespace/pods/:pod/logs": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Logs"
                ],
                "summary": "查询容器日志",
                "parameters": [
                    {
                        "type": "string",
                        "description": "容器名称",
                        "name": "container_name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "是否使用上一次日志, 异常退出使用",
                        "name": "previous",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/k8sclient.Log"
                            }
                        }
                    }
                }
            }
        },
        "/namespaces/:namespace/pods/:pod/logs/download": {
            "get": {
                "produces": [
                    "application/octet-stream"
                ],
                "tags": [
                    "Logs"
                ],
                "summary": "下载日志",
                "parameters": [
                    {
                        "type": "string",
                        "description": "容器名称",
                        "name": "container_name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "是否使用上一次日志, 异常退出使用",
                        "name": "previous",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/namespaces/:namespace/pods/:pod/logs/stream": {
            "get": {
                "produces": [
                    "text/event-stream"
                ],
                "tags": [
                    "Logs"
                ],
                "summary": "SSE 实时日志流",
                "parameters": [
                    {
                        "type": "string",
                        "description": "容器名称",
                        "name": "container_name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "开始时间",
                        "name": "started_at",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/namespaces/:namespace/pods/cpu_usage": {
            "post": {
                "tags": [
                    "Metrics"
                ],
                "summary": "Pod CPU使用率",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/namespaces/:namespace/pods/memory_used": {
            "post": {
                "tags": [
                    "Metrics"
                ],
                "summary": "Pod 内存使用量",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/namespaces/:namespace/pods/network_receive": {
            "post": {
                "tags": [
                    "Metrics"
                ],
                "summary": "网络接收量",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/namespaces/:namespace/pods/network_transmit": {
            "post": {
                "tags": [
                    "Metrics"
                ],
                "summary": "Pod 网络发送量",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/namespaces/namespace/pods/:pod/containers/:container/cpu_limit": {
            "get": {
                "tags": [
                    "Metrics"
                ],
                "summary": "容器 CPU 限制",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/namespaces/namespace/pods/:pod/containers/:container/cpu_usage": {
            "get": {
                "tags": [
                    "Metrics"
                ],
                "summary": "容器 CPU 使用率",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/namespaces/namespace/pods/:pod/containers/:container/disk_read_total": {
            "get": {
                "tags": [
                    "Metrics"
                ],
                "summary": "容器磁盘读总量",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/namespaces/namespace/pods/:pod/containers/:container/disk_write_total": {
            "get": {
                "tags": [
                    "Metrics"
                ],
                "summary": "容器磁盘写总量",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/namespaces/namespace/pods/:pod/containers/:container/memory_limit": {
            "get": {
                "tags": [
                    "Metrics"
                ],
                "summary": "容器内存限制",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/namespaces/namespace/pods/:pod/containers/:container/memory_used": {
            "get": {
                "tags": [
                    "Metrics"
                ],
                "summary": "容器内存使用量",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/nodes/:node/cpu_request_usage": {
            "get": {
                "tags": [
                    "Metrics"
                ],
                "summary": "查询 CPU 装箱率",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/nodes/:node/cpu_usage": {
            "get": {
                "tags": [
                    "Metrics"
                ],
                "summary": "查询 CPU 使用率",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/nodes/:node/disk_usage": {
            "get": {
                "tags": [
                    "Metrics"
                ],
                "summary": "节点磁盘使用率",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/nodes/:node/diskio_usage": {
            "get": {
                "tags": [
                    "Metrics"
                ],
                "summary": "节点磁盘IO",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/nodes/:node/info": {
            "get": {
                "tags": [
                    "Metrics"
                ],
                "summary": "节点信息",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/nodes/:node/memory_request_usage": {
            "get": {
                "tags": [
                    "Metrics"
                ],
                "summary": "节点内存装箱率",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/nodes/:node/memory_usage": {
            "get": {
                "tags": [
                    "Metrics"
                ],
                "summary": "节点内存使用率",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/nodes/:node/network_receive": {
            "get": {
                "tags": [
                    "Metrics"
                ],
                "summary": "节点网络发送",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/nodes/:node/network_transmit": {
            "get": {
                "tags": [
                    "Metrics"
                ],
                "summary": "节点网络接收",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/nodes/:node/overview": {
            "get": {
                "tags": [
                    "Metrics"
                ],
                "summary": "查询节点概览",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/overview": {
            "get": {
                "tags": [
                    "Metrics"
                ],
                "summary": "集群概览数据",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.Expression": {
            "type": "object",
            "properties": {
                "key": {
                    "type": "string"
                },
                "operator": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "entity.LogCollector": {
            "type": "object",
            "properties": {
                "add_pod_label": {
                    "type": "boolean"
                },
                "cluster_id": {
                    "type": "string"
                },
                "config": {
                    "$ref": "#/definitions/entity.LogCollectorConfig"
                },
                "config_selected": {
                    "type": "string"
                },
                "created_at": {
                    "$ref": "#/definitions/utils.JSONTime"
                },
                "creator": {
                    "type": "string"
                },
                "deleted": {
                    "type": "boolean"
                },
                "description": {
                    "type": "string"
                },
                "extra_labels": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "file_index_set_id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "namespace": {
                    "type": "string"
                },
                "project_code": {
                    "type": "string"
                },
                "project_id": {
                    "type": "string"
                },
                "rule_id": {
                    "type": "integer"
                },
                "std_index_set_id": {
                    "type": "integer"
                },
                "updated_at": {
                    "$ref": "#/definitions/utils.JSONTime"
                },
                "updator": {
                    "type": "string"
                }
            }
        },
        "entity.LogCollectorConfig": {
            "type": "object",
            "properties": {
                "all_containers": {
                    "$ref": "#/definitions/entity.LogCollectorConfigAllContainers"
                },
                "label_selector": {
                    "$ref": "#/definitions/entity.LogCollectorConfigSelector"
                },
                "workload": {
                    "$ref": "#/definitions/entity.LogCollectorConfigWorkload"
                }
            }
        },
        "entity.LogCollectorConfigAllContainers": {
            "type": "object",
            "properties": {
                "data_encoding": {
                    "type": "string"
                },
                "enable_stdout": {
                    "type": "boolean"
                },
                "paths": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "entity.LogCollectorConfigContainer": {
            "type": "object",
            "properties": {
                "container_name": {
                    "type": "string"
                },
                "data_encoding": {
                    "type": "string"
                },
                "enable_stdout": {
                    "type": "boolean"
                },
                "paths": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "entity.LogCollectorConfigSelector": {
            "type": "object",
            "properties": {
                "data_encoding": {
                    "type": "string"
                },
                "enable_stdout": {
                    "type": "boolean"
                },
                "match_expressions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Expression"
                    }
                },
                "match_labels": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "paths": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "entity.LogCollectorConfigWorkload": {
            "type": "object",
            "properties": {
                "containers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.LogCollectorConfigContainer"
                    }
                },
                "kind": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "k8sclient.Container": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "k8sclient.Log": {
            "type": "object",
            "properties": {
                "log": {
                    "type": "string"
                },
                "time": {
                    "type": "string"
                }
            }
        },
        "logcollector.Entrypoint": {
            "type": "object",
            "properties": {
                "file_log_url": {
                    "type": "string"
                },
                "std_log_url": {
                    "type": "string"
                }
            }
        },
        "utils.JSONTime": {
            "type": "object",
            "properties": {
                "time.Time": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "/bcsapi/v4/monitor/api/projects/:projectId/clusters/:clusterId",
	Schemes:          []string{},
	Title:            "BCS-Monitor OpenAPI",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
