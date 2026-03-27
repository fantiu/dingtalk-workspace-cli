# CLI 动态发现真实功能测试报告

| 项目 | 值 |
|---|---|
| 时间 | 2026-03-27 14:59:01 |
| 登录命令退出码 | `0` |
| 动态发现服务数 | `10` |
| 枚举命令节点数(`--help`) | `118` |
| `--help` 失败数 | `0` |
| 叶子命令真实执行数 | `83` |
| 执行成功数 | `75` |
| 执行失败数 | `8` |

## 登录结果

```text
[OK] Token 有效，无需重新登录
企业 ID:          ding8196cd9a2b2405da24f2f5cc6abecb85
有效期:            30 天后
Token 将自动刷新，无需重复登录
```

## 同产品包命令依赖关系分析

### aitable

| 模块 | 动作 | 依赖分析 |
|---|---|---|
| `attachment` | `upload` | 参数 `base-id` 依赖上游实体标识 |
| `base` | `create` | 写操作通常依赖 list/get/search 先拿到目标 ID 再执行；参数 `template-id` 依赖上游实体标识 |
| `base` | `delete` | 写操作通常依赖 list/get/search 先拿到目标 ID 再执行；参数 `base-id` 依赖上游实体标识 |
| `base` | `get` | 参数 `base-id` 依赖上游实体标识 |
| `base` | `list` | 以查询类命令为主，通常仅依赖关键词/分页参数 |
| `base` | `search` | 以查询类命令为主，通常仅依赖关键词/分页参数 |
| `base` | `update` | 写操作通常依赖 list/get/search 先拿到目标 ID 再执行；参数 `base-id` 依赖上游实体标识 |
| `field` | `create` | 写操作通常依赖 list/get/search 先拿到目标 ID 再执行；参数 `base-id` 依赖上游实体标识；参数 `table-id` 依赖上游实体标识 |
| `field` | `delete` | 写操作通常依赖 list/get/search 先拿到目标 ID 再执行；参数 `base-id` 依赖上游实体标识；参数 `field-id` 依赖上游实体标识；参数 `table-id` 依赖上游实体标识 |
| `field` | `get` | 参数 `base-id` 依赖上游实体标识；参数 `field-ids` 依赖上游实体标识；参数 `table-id` 依赖上游实体标识 |
| `field` | `update` | 写操作通常依赖 list/get/search 先拿到目标 ID 再执行；参数 `base-id` 依赖上游实体标识；参数 `field-id` 依赖上游实体标识；参数 `table-id` 依赖上游实体标识 |
| `record` | `create` | 写操作通常依赖 list/get/search 先拿到目标 ID 再执行；参数 `base-id` 依赖上游实体标识；参数 `table-id` 依赖上游实体标识 |
| `record` | `delete` | 写操作通常依赖 list/get/search 先拿到目标 ID 再执行；参数 `base-id` 依赖上游实体标识；参数 `record-ids` 依赖上游实体标识；参数 `table-id` 依赖上游实体标识 |
| `record` | `query` | 参数 `base-id` 依赖上游实体标识；参数 `field-ids` 依赖上游实体标识；参数 `record-ids` 依赖上游实体标识；参数 `table-id` 依赖上游实体标识 |
| `record` | `update` | 写操作通常依赖 list/get/search 先拿到目标 ID 再执行；参数 `base-id` 依赖上游实体标识；参数 `table-id` 依赖上游实体标识 |
| `table` | `create` | 写操作通常依赖 list/get/search 先拿到目标 ID 再执行；参数 `base-id` 依赖上游实体标识 |
| `table` | `delete` | 写操作通常依赖 list/get/search 先拿到目标 ID 再执行；参数 `base-id` 依赖上游实体标识；参数 `table-id` 依赖上游实体标识 |
| `table` | `get` | 参数 `base-id` 依赖上游实体标识；参数 `table-ids` 依赖上游实体标识 |
| `table` | `update` | 写操作通常依赖 list/get/search 先拿到目标 ID 再执行；参数 `base-id` 依赖上游实体标识；参数 `table-id` 依赖上游实体标识 |
| `template` | `search` | 以查询类命令为主，通常仅依赖关键词/分页参数 |

### attendance

| 模块 | 动作 | 依赖分析 |
|---|---|---|
| `record` | `get` | 参数 `date` 与时间窗口强相关；参数 `user` 依赖组织/人员上下文 |
| `rules` | `rules` | 参数 `date` 与时间窗口强相关 |
| `shift` | `list` | 参数 `end` 与时间窗口强相关；参数 `start` 与时间窗口强相关；参数 `users` 依赖组织/人员上下文 |
| `summary` | `summary` | 参数 `date` 与时间窗口强相关；参数 `user` 依赖组织/人员上下文 |

### auth

| 模块 | 动作 | 依赖分析 |
|---|---|---|
| `login` | `login` | 参数 `login-timeout` 与时间窗口强相关；参数 `token-url` 依赖外部凭证；参数 `token` 依赖外部凭证 |
| `logout` | `logout` | 以查询类命令为主，通常仅依赖关键词/分页参数 |
| `reset` | `reset` | 以查询类命令为主，通常仅依赖关键词/分页参数 |
| `status` | `status` | 以查询类命令为主，通常仅依赖关键词/分页参数 |

### cache

| 模块 | 动作 | 依赖分析 |
|---|---|---|
| `refresh` | `refresh` | 以查询类命令为主，通常仅依赖关键词/分页参数 |
| `status` | `status` | 以查询类命令为主，通常仅依赖关键词/分页参数 |

### calendar

| 模块 | 动作 | 依赖分析 |
|---|---|---|
| `busy` | `search` | 参数 `end` 与时间窗口强相关；参数 `start` 与时间窗口强相关；参数 `users` 依赖组织/人员上下文 |
| `event` | `create` | 写操作通常依赖 list/get/search 先拿到目标 ID 再执行；参数 `end` 与时间窗口强相关；参数 `start` 与时间窗口强相关 |
| `event` | `delete` | 写操作通常依赖 list/get/search 先拿到目标 ID 再执行；参数 `id` 依赖上游实体标识 |
| `event` | `get` | 参数 `id` 依赖上游实体标识 |
| `event` | `list` | 参数 `end` 与时间窗口强相关；参数 `start` 与时间窗口强相关 |
| `event` | `list-mine` | 参数 `end` 与时间窗口强相关；参数 `start` 与时间窗口强相关 |
| `event` | `update` | 写操作通常依赖 list/get/search 先拿到目标 ID 再执行；参数 `end` 与时间窗口强相关；参数 `id` 依赖上游实体标识；参数 `start` 与时间窗口强相关 |
| `participant` | `add` | 写操作通常依赖 list/get/search 先拿到目标 ID 再执行；参数 `users` 依赖组织/人员上下文 |
| `participant` | `delete` | 写操作通常依赖 list/get/search 先拿到目标 ID 再执行；参数 `users` 依赖组织/人员上下文 |
| `participant` | `list` | 以查询类命令为主，通常仅依赖关键词/分页参数 |
| `room` | `add` | 写操作通常依赖 list/get/search 先拿到目标 ID 再执行 |
| `room` | `delete` | 写操作通常依赖 list/get/search 先拿到目标 ID 再执行 |
| `room` | `list-groups` | 以查询类命令为主，通常仅依赖关键词/分页参数 |
| `room` | `search` | 参数 `end` 与时间窗口强相关；参数 `group-id` 依赖上游实体标识；参数 `start` 与时间窗口强相关 |

### chat

| 模块 | 动作 | 依赖分析 |
|---|---|---|
| `bot` | `search` | 以查询类命令为主，通常仅依赖关键词/分页参数 |
| `group` | `create` | 写操作通常依赖 list/get/search 先拿到目标 ID 再执行；参数 `users` 依赖组织/人员上下文 |
| `group` | `add` | 写操作通常依赖 list/get/search 先拿到目标 ID 再执行；参数 `id` 依赖上游实体标识；参数 `users` 依赖组织/人员上下文 |
| `group` | `add-bot` | 参数 `id` 依赖上游实体标识 |
| `group` | `remove` | 写操作通常依赖 list/get/search 先拿到目标 ID 再执行；参数 `id` 依赖上游实体标识；参数 `users` 依赖组织/人员上下文 |
| `group` | `rename` | 参数 `id` 依赖上游实体标识 |
| `message` | `list` | 参数 `time` 与时间窗口强相关；参数 `user` 依赖组织/人员上下文 |
| `message` | `recall-by-bot` | 以查询类命令为主，通常仅依赖关键词/分页参数 |
| `message` | `send` | 写操作通常依赖 list/get/search 先拿到目标 ID 再执行；参数 `user` 依赖组织/人员上下文 |
| `message` | `send-by-bot` | 以查询类命令为主，通常仅依赖关键词/分页参数 |
| `message` | `send-by-webhook` | 参数 `at-users` 依赖组织/人员上下文；参数 `token` 依赖外部凭证 |
| `search` | `search` | 以查询类命令为主，通常仅依赖关键词/分页参数 |

### completion

| 模块 | 动作 | 依赖分析 |
|---|---|---|
| `(root)` | `completion` | 以查询类命令为主，通常仅依赖关键词/分页参数 |

### contact

| 模块 | 动作 | 依赖分析 |
|---|---|---|
| `dept` | `list-children` | 参数 `id` 依赖上游实体标识 |
| `dept` | `list-members` | 参数 `ids` 依赖上游实体标识 |
| `dept` | `search` | 以查询类命令为主，通常仅依赖关键词/分页参数 |
| `user` | `get` | 参数 `ids` 依赖上游实体标识 |
| `user` | `get-self` | 以查询类命令为主，通常仅依赖关键词/分页参数 |
| `user` | `search` | 以查询类命令为主，通常仅依赖关键词/分页参数 |
| `user` | `search-mobile` | 以查询类命令为主，通常仅依赖关键词/分页参数 |

### devdoc

| 模块 | 动作 | 依赖分析 |
|---|---|---|
| `article` | `search` | 以查询类命令为主，通常仅依赖关键词/分页参数 |

### ding

| 模块 | 动作 | 依赖分析 |
|---|---|---|
| `message` | `recall` | 参数 `id` 依赖上游实体标识 |
| `message` | `send` | 写操作通常依赖 list/get/search 先拿到目标 ID 再执行；参数 `users` 依赖组织/人员上下文 |

### report

| 模块 | 动作 | 依赖分析 |
|---|---|---|
| `create` | `create` | 写操作通常依赖 list/get/search 先拿到目标 ID 再执行；参数 `template-id` 依赖上游实体标识；参数 `to-user-ids` 依赖上游实体标识；参数 `to-user-ids` 依赖组织/人员上下文 |
| `detail` | `detail` | 参数 `report-id` 依赖上游实体标识 |
| `list` | `list` | 参数 `end` 与时间窗口强相关；参数 `start` 与时间窗口强相关 |
| `sent` | `sent` | 参数 `end` 与时间窗口强相关；参数 `modified-end` 与时间窗口强相关；参数 `modified-start` 与时间窗口强相关；参数 `start` 与时间窗口强相关 |
| `stats` | `stats` | 参数 `report-id` 依赖上游实体标识 |
| `template` | `detail` | 以查询类命令为主，通常仅依赖关键词/分页参数 |
| `template` | `list` | 以查询类命令为主，通常仅依赖关键词/分页参数 |

### todo

| 模块 | 动作 | 依赖分析 |
|---|---|---|
| `task` | `create` | 写操作通常依赖 list/get/search 先拿到目标 ID 再执行 |
| `task` | `delete` | 写操作通常依赖 list/get/search 先拿到目标 ID 再执行；参数 `task-id` 依赖上游实体标识 |
| `task` | `done` | 参数 `task-id` 依赖上游实体标识 |
| `task` | `get` | 参数 `task-id` 依赖上游实体标识 |
| `task` | `list` | 以查询类命令为主，通常仅依赖关键词/分页参数 |
| `task` | `update` | 写操作通常依赖 list/get/search 先拿到目标 ID 再执行；参数 `task-id` 依赖上游实体标识 |

### version

| 模块 | 动作 | 依赖分析 |
|---|---|---|
| `(root)` | `version` | 以查询类命令为主，通常仅依赖关键词/分页参数 |

### workbench

| 模块 | 动作 | 依赖分析 |
|---|---|---|
| `app` | `get` | 参数 `ids` 依赖上游实体标识 |
| `app` | `list` | 以查询类命令为主，通常仅依赖关键词/分页参数 |

## 叶子命令真实执行明细

| 命令路径 | 退出码 | 自动补齐入参数量 | 摘要 |
|---|---:|---:|---|
| `dws aitable attachment upload` | `0` | `3` | {   "data": {},   "error": {     "code": "INVALID_BASE_ID",     "message": "baseId must contain only letters and digits",     "retryable": false,     "type": "INPUT_ERROR"   },   " |
| `dws aitable base create` | `0` | `3` | {   "data": {     "baseId": "3NwLYZXWyrmRolXLcQG7gxMdVkyEqBQm",     "baseName": "auto-test-name"   },   "error": {},   "meta": {},   "status": "success",   "summary": "Created base |
| `dws aitable base delete` | `0` | `2` | {   "data": {},   "error": {     "code": "INVALID_BASE_ID",     "message": "baseId must contain only letters and digits",     "retryable": false,     "type": "INPUT_ERROR"   },   " |
| `dws aitable base get` | `0` | `2` | {   "data": {},   "error": {     "code": "INVALID_BASE_ID",     "message": "baseId must contain only letters and digits",     "retryable": false,     "type": "INPUT_ERROR"   },   " |
| `dws aitable base list` | `0` | `1` | {   "data": {     "bases": [       {         "baseId": "3NwLYZXWyrmRolXLcQG7gxMdVkyEqBQm",         "baseName": "auto-test-name"       },       {         "baseId": "QPGYqjpJYRyab0nx |
| `dws aitable base search` | `0` | `2` | {   "data": {     "bases": [       {         "baseId": "14dA3GK8gkP4oNl0crRvZyXv89ekBD76",         "baseName": "测试AI 表格"       },       {         "baseId": "DnRL6jAJMNX9kAgycZdkqoY |
| `dws aitable base update` | `0` | `3` | {   "data": {},   "error": {     "code": "INVALID_BASE_ID",     "message": "baseId must contain only letters and digits",     "retryable": false,     "type": "INPUT_ERROR"   },   " |
| `dws aitable field create` | `0` | `3` | {   "data": {},   "error": {     "code": "INVALID_BASE_ID",     "message": "baseId must contain only letters and digits",     "retryable": false,     "type": "INPUT_ERROR"   },   " |
| `dws aitable field delete` | `0` | `4` | {   "data": {},   "error": {     "code": "INVALID_BASE_ID",     "message": "baseId must contain only letters and digits",     "retryable": false,     "type": "INPUT_ERROR"   },   " |
| `dws aitable field get` | `0` | `4` | {   "data": {},   "error": {     "code": "INVALID_BASE_ID",     "message": "baseId must contain only letters and digits",     "retryable": false,     "type": "INPUT_ERROR"   },   " |
| `dws aitable field update` | `0` | `5` | {   "data": {},   "error": {     "code": "INVALID_BASE_ID",     "message": "baseId must contain only letters and digits",     "retryable": false,     "type": "INPUT_ERROR"   },   " |
| `dws aitable record create` | `0` | `3` | {   "data": {},   "error": {     "code": "INVALID_BASE_ID",     "message": "baseId must contain only letters and digits",     "retryable": false,     "type": "INPUT_ERROR"   },   " |
| `dws aitable record delete` | `0` | `4` | {   "data": {},   "error": {     "code": "INVALID_BASE_ID",     "message": "baseId must contain only letters and digits",     "retryable": false,     "type": "INPUT_ERROR"   },   " |
| `dws aitable record query` | `0` | `6` | {   "data": {},   "error": {     "code": "INVALID_BASE_ID",     "message": "baseId must contain only letters and digits",     "retryable": false,     "type": "INPUT_ERROR"   },   " |
| `dws aitable record update` | `0` | `3` | {   "data": {},   "error": {     "code": "INVALID_BASE_ID",     "message": "baseId must contain only letters and digits",     "retryable": false,     "type": "INPUT_ERROR"   },   " |
| `dws aitable table create` | `0` | `3` | {   "data": {},   "error": {     "code": "INVALID_BASE_ID",     "message": "baseId must contain only letters and digits",     "retryable": false,     "type": "INPUT_ERROR"   },   " |
| `dws aitable table delete` | `0` | `3` | {   "data": {},   "error": {     "code": "INVALID_BASE_ID",     "message": "baseId must contain only letters and digits",     "retryable": false,     "type": "INPUT_ERROR"   },   " |
| `dws aitable table get` | `0` | `3` | {   "data": {},   "error": {     "code": "INVALID_BASE_ID",     "message": "baseId must contain only letters and digits",     "retryable": false,     "type": "INPUT_ERROR"   },   " |
| `dws aitable table update` | `0` | `4` | {   "data": {},   "error": {     "code": "INVALID_BASE_ID",     "message": "baseId must contain only letters and digits",     "retryable": false,     "type": "INPUT_ERROR"   },   " |
| `dws aitable template search` | `0` | `2` | {   "data": {     "hasMore": false,     "nextCursor": "10",     "templates": [       {         "description": "随堂小测试模板用于课堂即时检测，适用教师与学生，快速收集反馈，提升教学效率。",         "name": "随堂小测试",     |
| `dws attendance record get` | `3` | `3` | {   "error": {     "category": "validation",     "code": 3,     "message": "--date format error, use YYYY-MM-DD, e.g. 2026-03-08"   } }  |
| `dws attendance rules` | `3` | `2` | {   "error": {     "category": "validation",     "code": 3,     "message": "--date format error, use YYYY-MM-DD or yyyy-MM-dd HH:mm:ss: test"   } }  |
| `dws attendance shift list` | `3` | `4` | {   "error": {     "category": "validation",     "code": 3,     "message": "--start date format error, use YYYY-MM-DD"   } }  |
| `dws attendance summary` | `3` | `3` | {   "error": {     "category": "validation",     "code": 3,     "message": "--date format error, use yyyy-MM-dd HH:mm:ss"   } }  |
| `dws calendar busy search` | `0` | `2` | {   "errorCode": "300000",   "errorMsg": "code: 300000, developerMessage: startTime cannot be blank",   "result": [     {       "scheduleItems": null     }   ],   "success": false, |
| `dws calendar event create` | `0` | `2` | {   "arguments": [],   "errorCode": "300000",   "errorMsg": "code: 300000, developerMessage: Start time cannot be blank for a non-all day event.",   "result": {     "attendees": nu |
| `dws calendar event delete` | `0` | `2` | {   "arguments": [],   "errorCode": "300014",   "errorMsg": "Event not exist",   "result": {},   "success": false,   "trace_id": "213ee25c17745946765898223e0af0" }  |
| `dws calendar event get` | `0` | `2` | {   "arguments": [],   "result": {     "attendees": null,     "categories": null,     "reminders": null   },   "success": false,   "trace_id": "0bab027317745946772971785e07c5" }  |
| `dws calendar event list` | `0` | `1` | {   "arguments": [],   "result": {     "events": [       {         "attendees": [           {             "displayName": "尘舟",             "optional": false,             "responseS |
| `dws calendar event list-mine` | `0` | `1` | {   "code": "PARAM_ERROR",   "error": "参数错误",   "message": "参数错误。请检查输入参数是否正确，参考工具说明调整参数后重试",   "retryable": false,   "success": false,   "technical_detail": "Tool metadata API erro |
| `dws calendar event update` | `0` | `3` | {   "arguments": [],   "errorCode": "300000",   "errorMsg": "code: 300000, developerMessage: Event does not exist.",   "result": {},   "success": false,   "trace_id": "2106d9811774 |
| `dws calendar participant add` | `0` | `2` | {   "arguments": [],   "errorCode": "300000",   "errorMsg": "code: 300000, developerMessage: eventId cannot be blank",   "success": false,   "trace_id": "2127d89817745947042464823e |
| `dws calendar participant delete` | `0` | `2` | {   "arguments": [],   "errorCode": "300000",   "errorMsg": "code: 300000, developerMessage: eventId cannot be blank",   "result": {},   "success": false,   "trace_id": "2106d98117 |
| `dws calendar participant list` | `0` | `1` | {   "arguments": [],   "errorCode": "300000",   "errorMsg": "code: 300000, developerMessage: eventId cannot be blank",   "result": {     "attendees": null   },   "success": false,  |
| `dws calendar room add` | `0` | `1` | {   "arguments": [],   "errorCode": "300000",   "errorMsg": "code: 300000, developerMessage: eventId cannot be blank",   "result": {},   "success": false,   "trace_id": "2132f5ca17 |
| `dws calendar room delete` | `0` | `1` | {   "arguments": [],   "errorCode": "300000",   "errorMsg": "code: 300000, developerMessage: eventId cannot be blank",   "result": {},   "success": false,   "trace_id": "213ee25c17 |
| `dws calendar room list-groups` | `0` | `1` | {   "result": {     "groupList": [       {         "groupId": 41,         "groupName": "绿城·未来park",         "parentId": 0       },       {         "groupId": 74,         "groupName |
| `dws calendar room search` | `0` | `2` | {   "arguments": [],   "errorCode": "400002",   "errorMsg": "filterStartTime or filterEndTime error",   "result": {     "result": [       {         "labels": null       }     ]   } |
| `dws chat bot search` | `0` | `4` | {   "robotList": [],   "success": true }  |
| `dws chat group create` | `0` | `3` | {   "code": "PARAM_ERROR",   "error": "参数错误",   "message": "参数错误。请检查输入参数是否正确，参考工具说明调整参数后重试",   "retryable": false,   "success": false,   "technical_detail": "Tool metadata API erro |
| `dws chat group members add` | `0` | `3` | {   "code": "PARAM_ERROR",   "error": "参数错误",   "message": "参数错误。请检查输入参数是否正确，参考工具说明调整参数后重试",   "retryable": false,   "success": false,   "technical_detail": "Tool metadata API erro |
| `dws chat group members add-bot` | `0` | `3` | {   "errorCode": "500025",   "errorMsg": "robotCode不存在",   "success": false,   "trace_id": "0bab027317745947105987158e0780" }  |
| `dws chat group members remove` | `0` | `3` | {   "code": "PARAM_ERROR",   "error": "参数错误",   "message": "参数错误。请检查输入参数是否正确，参考工具说明调整参数后重试",   "retryable": false,   "success": false,   "technical_detail": "Tool metadata API erro |
| `dws chat group rename` | `0` | `3` | {   "code": "PARAM_ERROR",   "error": "参数错误",   "message": "参数错误。请检查输入参数是否正确，参考工具说明调整参数后重试",   "retryable": false,   "success": false,   "technical_detail": "Tool metadata API erro |
| `dws chat message list` | `3` | `4` | {   "error": {     "category": "validation",     "code": 3,     "message": "--time is required"   } }  |
| `dws chat message recall-by-bot` | `0` | `1` | {   "code": "PARAM_ERROR",   "error": "参数错误",   "message": "参数错误。请检查输入参数是否正确，参考工具说明调整参数后重试",   "retryable": false,   "success": false,   "technical_detail": "调用失败，原因: 调用失败，原因: 调用远程 |
| `dws chat message send` | `5` | `4` | {   "error": {     "category": "internal",     "code": 5,     "message": "accepts 1 arg(s), received 0"   } }  |
| `dws chat message send-by-bot` | `0` | `3` | {   "code": "PARAM_ERROR",   "error": "参数错误",   "message": "参数错误。请检查输入参数是否正确，参考工具说明调整参数后重试",   "retryable": false,   "success": false,   "technical_detail": "调用失败，原因: 调用失败，原因: 调用远程 |
| `dws chat message send-by-webhook` | `0` | `6` | {   "errcode": "300005",   "errmsg": "token is not exist" }  |
| `dws chat search` | `0` | `2` | {   "result": {     "hasMore": true,     "nextCursor": "abe74d73660ef3d46bc92185e05915f2f48689450699e0bf70849ba71758a903",     "total": 262,     "value": [       {         "extensi |
| `dws contact dept list-children` | `0` | `2` | {   "result": [],   "success": true }  |
| `dws contact dept list-members` | `0` | `2` | {   "deptUserList": [],   "success": true }  |
| `dws contact dept search` | `0` | `2` | {   "deptList": [],   "hasMore": false,   "totalCount": 0 }  |
| `dws contact user get` | `0` | `2` | {   "result": [     {       "orgEmployeeModel": {         "depts": null,         "labels": null,         "orgId": null       }     }   ],   "success": true }  |
| `dws contact user get-self` | `0` | `1` | {   "result": [     {       "isAdmin": true,       "orgEmployeeModel": {         "corpId": "ding8196cd9a2b2405da24f2f5cc6abecb85",         "depts": [           {             "deptI |
| `dws contact user search` | `0` | `2` | {   "userId": [     "033067356464903146",     "030037243111903146",     "025109313335903146"   ] }  |
| `dws contact user search-mobile` | `0` | `2` | {   "arguments": [],   "success": true }  |
| `dws devdoc article search` | `0` | `2` | {   "result": {     "currentPage": 1,     "hasMore": true,     "items": [       {         "desc": "标题: 创建知识库文档 - 开放平台\n内容: * * *\n例如，企业员工小明在测试知识库内，单击**新建** 按钮创建文档，与调用本接口实现效果一致，如下图： |
| `dws ding message recall` | `0` | `2` | {   "dingOpenErrcode": 400100,   "errorMsg": "robotCode is illegal",   "result": {},   "success": false,   "trace_id": "2106d98117745947215725878e0ad1" }  |
| `dws ding message send` | `0` | `3` | {   "dingOpenErrcode": 400100,   "errorMsg": "robotCode is illegal",   "result": {},   "success": false,   "trace_id": "2106d98117745947223905960e0ad1" }  |
| `dws report create` | `3` | `4` | {   "error": {     "category": "validation",     "code": 3,     "message": "json_parse: invalid JSON: invalid character 'è' looking for beginning of value"   } }  |
| `dws report detail` | `0` | `2` | {   "errorMsg": "ok",   "result": {     "contentV2": [       {         "images": null       }     ],     "images": [],     "report_Id": "test-001",     "report_content": [       {  |
| `dws report list` | `0` | `1` | {   "errcode": 40035,   "errorMsg": "不合法的参数",   "result": {     "report_list": null,     "size": 0   },   "success": false,   "trace_id": "0bab027317745947235068184e0780" }  |
| `dws report sent` | `0` | `2` | {   "code": "ERROR",   "error": "操作失败",   "message": "操作失败。发生错误，建议稍后重试",   "retryable": true,   "success": false,   "technical_detail": "调用失败，原因: 调用失败，原因: 调用远程服务业务异常, 出参校验不通过。root. |
| `dws report stats` | `0` | `2` | {   "code": "PARAM_ERROR",   "error": "参数错误",   "message": "参数错误。请检查输入参数是否正确，参考工具说明调整参数后重试",   "retryable": false,   "success": false,   "technical_detail": "调用失败，原因: 调用失败，原因: 调用远程 |
| `dws report template detail` | `0` | `2` | {   "code": "PARAM_ERROR",   "error": "参数错误",   "message": "参数错误。请检查输入参数是否正确，参考工具说明调整参数后重试",   "retryable": false,   "success": false,   "technical_detail": "调用失败，原因: 调用失败，原因: 调用远程 |
| `dws report template list` | `0` | `1` | {   "items": [     {       "last_modified_time": 1697702695000,       "report_template_id": "18b46f75236c558cc63be07419f8d65f",       "report_template_logo_url": "https://landray.d |
| `dws todo task create` | `0` | `2` | {   "arguments": [],   "errorCode": "999",   "errorMsg": "system error: java.lang.NullPointerException",   "success": false,   "trace_id": "2106d98117745947277007837e0b34" }  |
| `dws todo task delete` | `0` | `2` | {   "result": {     "success": false   } }  |
| `dws todo task done` | `0` | `2` | {   "result": {     "success": false   } }  |
| `dws todo task get` | `0` | `2` | {   "arguments": [],   "errorCode": "500",   "errorMsg": "NewTodoService#getTodoDetail taskId is null",   "result": {},   "success": false,   "trace_id": "213ee25c17745947295756384 |
| `dws todo task list` | `0` | `2` | {   "result": {     "todoCards": null   } }  |
| `dws todo task update` | `0` | `3` | {   "result": {     "success": false   } }  |
| `dws workbench app get` | `0` | `2` | {   "eventData": {     "activityType": "MULTI_PANEL",     "content": {       "assets": {         "description": "工作台扩展",         "entry": "https://dev.g.alicdn.com/dingding/dd-chat |
| `dws workbench app list` | `0` | `1` | {   "appList": [     {       "d": "商业保险",       "i": "0_2403230743",       "n": "商业保险"     },     {       "d": "助力组织文化活动快速落地",       "i": "0_2680655126",       "n": "组织文化管理平台"      |
| `dws auth login` | `0` | `7` |  [OK] Token 有效，无需重新登录 有效期:            24 小时后 Token 将自动刷新，无需重复登录  |
| `dws auth logout` | `0` | `1` | [OK] 已清除所有认证信息 请运行 dws auth login 重新登录  |
| `dws auth reset` | `0` | `1` | [OK] 认证信息已重置 请运行 dws auth login 重新登录  |
| `dws auth status` | `0` | `1` | 状态:             未登录 运行 dws auth login 进行登录  |
| `dws cache refresh` | `0` | `1` | [OK] 缓存刷新完成：已刷新 10 个服务，失败 10 个 缓存目录: /Users/tianlei.qjb/.dws/cache  |
| `dws cache status` | `0` | `1` | 缓存目录: /Users/tianlei.qjb/.dws/cache 文件数:   1   大小: 63214 字节  |
| `dws completion` | `5` | `1` | {   "error": {     "category": "internal",     "code": 5,     "message": "accepts 1 arg(s), received 0"   } }  |
| `dws version` | `0` | `1` | {   "go": "1.24+",   "version": "v1.0.0" }  |

## 执行口径说明

- 本次为**真实命令执行**（非 go test，非 dry-run）；为了避免破坏性写入，入参使用自动生成测试值，部分写操作可能因业务校验失败而返回非 0。
- 所有命令均通过动态发现入口进行递归枚举，覆盖当前可见服务树。
- 凭证未写入报告，报告仅保留退出码与脱敏输出片段。
