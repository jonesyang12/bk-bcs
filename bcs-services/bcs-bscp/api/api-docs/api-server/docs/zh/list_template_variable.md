#### API基本信息

API名称：list_template_variable

API Path：/api/v1/config/biz/{biz_id}/template_variables

Method：GET

#### 描述

该接口提供版本：v1.0.0+

查询模版变量列表

#### 输入参数

| 参数名称   | 参数类型 | 必选 | 描述                   |
| ---------- | -------- | ---- | ---------------------- |
| biz_id     | uint32   | 是   | 业务ID                 |
| search_key | string   | 否   | 搜索条件（名称、描述） |
| start      | uint32   | 是   | 分页起始值             |
| limit      | uint32   | 是   | 分页大小               |

#### 调用示例

```json

```

#### 响应示例

```json
{
  "data": {
    "count": 1,
    "details": [
      {
        "id": 1,
        "spec": {
          "name": "bk_bscp_variable001",
          "memo": "my first template variable"
        },
        "attachment": {
          "biz_id": 2
        },
        "revision": {
          "creator": "bscp_local_tester",
          "reviser": "bscp_local_tester",
          "create_at": "2023-05-05 10:34:15",
          "update_at": "2023-05-05 10:34:15"
        }
      }
    ]
  }
}
```

#### 响应参数说明

| 参数名称 | 参数类型 | 描述     |
| -------- | -------- | -------- |
| data     | object   | 响应数据 |

#### data

| 参数名称 | 参数类型 | 描述                         |
| -------- | -------- | ---------------------------- |
| count    | uint32   | 当前规则能匹配到的总记录条数 |
| detail   | array    | 查询返回的数据               |

#### data.detail[n]

| 参数名称   | 参数类型 | 描述     |
| ---------- | -------- | -------- |
| id         | uint32   | 应用ID   |
| biz_id     | uint32   | 业务ID   |
| spec       | object   | 资源信息 |
| attachment | object   | 关联信息 |
| revision   | object   | 修改信息 |

#### spec

| 参数名称    | 参数类型 | 描述                                         |
| ----------- | -------- | -------------------------------------------- |
| name        | string   | 模版变量名称                                 |
| type        | string   | 模版变量类型（枚举值：string、number、bool） |
| default_val | string   | 模版变量默认值的json串                       |
| memo        | string   | 模版变量描述                                 |

#### attachment

| 参数名称 | 参数类型 | 描述   |
| -------- | -------- | ------ |
| biz_id   | uint32   | 业务ID |

#### revision

| 参数名称  | 参数类型 | 描述                 |
| --------- | -------- | -------------------- |
| creator   | string   | 创建者               |
| reviser   | string   | 最后一次修改的修改者 |
| create_at | string   | 创建时间             |
| update_at | string   | 最后一次修改时间     |

