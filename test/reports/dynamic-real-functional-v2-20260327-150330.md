# CLI 动态发现真实功能测试报告（第二轮强化）

| 项目 | 值 |
|---|---|
| 时间 | 2026-03-27 15:05:18 |
| 登录退出码 | `0` |
| 动态服务数 | `10` |
| 命令节点(`--help`) | `118` |
| `--help` 失败 | `0` |
| 叶子命令执行数 | `83` |
| 语义成功(pass) | `37` |
| 退出码失败 | `4` |
| 业务失败(success=false/错误体) | `42` |

## 依赖链增强策略

- `aitable`：先 `base create` 获取 `baseId`，再复用到 table/field/record 子命令。
- `calendar`：先 `event create/list` 采集 `eventId`，再复用到 event get/update/delete、participant 相关命令。
- 日期时间按命令帮助约束分流：`YYYY-MM-DD` 与 `yyyy-MM-dd HH:mm:ss`。
- 人员参数优先填充真实用户 `manager5450`（来自第一轮接口返回上下文）。

## 关键上下文ID

- `base-id`: `Exel2BLV5gOAdXr1uDpNXZN38gk9rpMq`
- `table-id`: `F0k4Z0j`
- `event-id`: `UzFMazluYWE5bkI2dElacTgvdlJwUT09`

## 失败明细（需进一步人工业务校验）

| 命令 | 分类 | 退出码 | 摘要 |
|---|---|---:|---|
| `dws aitable attachment upload` | `fail_business` | `0` | {   "data": {},   "error": {     "code": "INVALID_PREPARE_ATTACHMENT_UPLOAD_REQUEST",     "message": "fileName must include an extension",     "retryable": false,     "type": "INPUT_ERROR"   },   "meta": {},   "status":  |
| `dws aitable base get` | `fail_business` | `0` | {   "data": {},   "error": {     "code": "BASE_NOT_FOUND",     "message": "Specified base does not exist, has been deleted, or is inaccessible",     "retryable": false,     "type": "INPUT_ERROR"   },   "meta": {},   "sta |
| `dws aitable base update` | `fail_business` | `0` | {   "data": {},   "error": {     "code": "BASE_NOT_FOUND",     "message": "Specified base does not exist, has been deleted, or is inaccessible",     "retryable": false,     "type": "INPUT_ERROR"   },   "meta": {},   "sta |
| `dws aitable field create` | `fail_business` | `0` | {   "data": {},   "error": {     "code": "INVALID_PARAMS",     "message": "fields are required",     "retryable": false,     "type": "INPUT_ERROR"   },   "meta": {},   "status": "error",   "summary": "Failed to create fi |
| `dws aitable field delete` | `fail_business` | `0` | {   "data": {},   "error": {     "code": "404",     "message": "getDentryDTO returns null",     "retryable": true,     "type": "SYSTEM_ERROR"   },   "meta": {},   "status": "error",   "summary": "Failed to get current fi |
| `dws aitable field get` | `fail_business` | `0` | {   "data": {},   "error": {     "code": "INVALID_FIELD_ID",     "message": "fieldIds[0] must contain only letters and digits",     "retryable": false,     "type": "INPUT_ERROR"   },   "meta": {},   "status": "error",    |
| `dws aitable field update` | `fail_business` | `0` | {   "data": {},   "error": {     "code": "404",     "message": "getDentryDTO returns null",     "retryable": true,     "type": "SYSTEM_ERROR"   },   "meta": {},   "status": "error",   "summary": "Failed to get current fi |
| `dws aitable record create` | `fail_business` | `0` | {   "data": {},   "error": {     "code": "INVALID_RECORDS",     "message": "records must contain at least one record",     "retryable": false,     "type": "INPUT_ERROR"   },   "meta": {},   "status": "error",   "summary" |
| `dws aitable record delete` | `fail_business` | `0` | {   "data": {},   "error": {     "code": "INVALID_RECORD_ID",     "message": "recordIds[0] must contain only letters and digits",     "retryable": false,     "type": "INPUT_ERROR"   },   "meta": {},   "status": "error",  |
| `dws aitable record query` | `fail_business` | `0` | {   "data": {},   "error": {     "code": "INVALID_RECORD_ID",     "message": "recordIds[0] must contain only letters and digits",     "retryable": false,     "type": "INPUT_ERROR"   },   "meta": {},   "status": "error",  |
| `dws aitable record update` | `fail_business` | `0` | {   "data": {},   "error": {     "code": "INVALID_RECORDS",     "message": "records must contain at least one record",     "retryable": false,     "type": "INPUT_ERROR"   },   "meta": {},   "status": "error",   "summary" |
| `dws aitable table create` | `fail_business` | `0` | {   "data": {},   "error": {     "code": "404",     "message": "getDentryDTO returns null",     "retryable": true,     "type": "SYSTEM_ERROR"   },   "meta": {},   "status": "error",   "summary": "Failed to create table ' |
| `dws aitable table delete` | `fail_business` | `0` | {   "data": {},   "error": {     "code": "BASE_NOT_FOUND",     "message": "Specified base does not exist, has been deleted, or is inaccessible",     "retryable": false,     "type": "INPUT_ERROR"   },   "meta": {},   "sta |
| `dws aitable table get` | `fail_business` | `0` | {   "data": {},   "error": {     "code": "INVALID_TABLE_ID",     "message": "tableIds[0] must contain only letters and digits",     "retryable": false,     "type": "INPUT_ERROR"   },   "meta": {},   "status": "error",    |
| `dws aitable table update` | `fail_business` | `0` | {   "data": {},   "error": {     "code": "404",     "message": "getDentryDTO returns null",     "retryable": true,     "type": "SYSTEM_ERROR"   },   "meta": {},   "status": "error",   "summary": "Failed to update table F |
| `dws attendance record get` | `fail_business` | `0` | {   "code": "ERROR",   "error": "操作失败",   "message": "操作失败。发生错误，建议稍后重试",   "retryable": true,   "success": false,   "technical_detail": "调用失败，原因: 调用失败，原因: 调用远程服务业务异常, 出参校验不通过。root.success当前值：false， 预期值：true，所有值：{\"code\" |
| `dws attendance summary` | `fail_business` | `0` | {   "code": "C0002",   "message": "统计类型错误",   "result": {},   "success": false,   "trace_id": "0bab027317745950775382548e0761" }  |
| `dws calendar event create` | `fail_business` | `0` | {   "arguments": [],   "errorCode": "300000",   "errorMsg": "code: 300000, developerMessage: Invalid start dateTime format.",   "result": {     "attendees": null,     "reminders": null   },   "success": false,   "trace_i |
| `dws calendar event list-mine` | `fail_business` | `0` | {   "code": "PARAM_ERROR",   "error": "参数错误",   "message": "参数错误。请检查输入参数是否正确，参考工具说明调整参数后重试",   "retryable": false,   "success": false,   "technical_detail": "Tool metadata API error: PARAM_ERROR - 未找到指定工具",   "trace_id": |
| `dws calendar event update` | `fail_business` | `0` | {   "arguments": [],   "errorCode": "300000",   "errorMsg": "code: 300000, developerMessage: Invalid start dateTime format.",   "result": {},   "success": false,   "trace_id": "213ee25c17745950825128586e0b32" }  |
| `dws calendar participant add` | `fail_business` | `0` | {   "arguments": [],   "errorCode": "300000",   "errorMsg": "code: 300000, developerMessage: eventId cannot be blank",   "success": false,   "trace_id": "2132f5ca17745950833098855e09f8" }  |
| `dws calendar participant delete` | `fail_business` | `0` | {   "arguments": [],   "errorCode": "300000",   "errorMsg": "code: 300000, developerMessage: eventId cannot be blank",   "result": {},   "success": false,   "trace_id": "0bab027317745950840092994e0761" }  |
| `dws calendar participant list` | `fail_business` | `0` | {   "arguments": [],   "errorCode": "300000",   "errorMsg": "code: 300000, developerMessage: eventId cannot be blank",   "result": {     "attendees": null   },   "success": false,   "trace_id": "213ee25c17745950846354940 |
| `dws calendar room add` | `fail_business` | `0` | {   "arguments": [],   "errorCode": "300000",   "errorMsg": "code: 300000, developerMessage: eventId cannot be blank",   "result": {},   "success": false,   "trace_id": "0bab027317745950853753133e0761" }  |
| `dws calendar room delete` | `fail_business` | `0` | {   "arguments": [],   "errorCode": "300000",   "errorMsg": "code: 300000, developerMessage: eventId cannot be blank",   "result": {},   "success": false,   "trace_id": "2132f5ca17745950860591744e0a4e" }  |
| `dws calendar room search` | `fail_business` | `0` | {   "arguments": [],   "errorCode": "458019",   "errorMsg": "查询范围内的会议室数量，超过上限100，请选择更小范围的分组进行查询。",   "result": {     "result": [       {         "labels": null       }     ]   },   "success": false,   "trace_id": "2132f5 |
| `dws chat group create` | `fail_business` | `0` | {   "code": "PARAM_ERROR",   "error": "参数错误",   "message": "参数错误。请检查输入参数是否正确，参考工具说明调整参数后重试",   "retryable": false,   "success": false,   "technical_detail": "Tool metadata API error: PARAM_ERROR - 未找到指定工具",   "trace_id": |
| `dws chat group members add` | `fail_business` | `0` | {   "code": "PARAM_ERROR",   "error": "参数错误",   "message": "参数错误。请检查输入参数是否正确，参考工具说明调整参数后重试",   "retryable": false,   "success": false,   "technical_detail": "Tool metadata API error: PARAM_ERROR - 未找到指定工具",   "trace_id": |
| `dws chat group members add-bot` | `fail_business` | `0` | {   "errorCode": "500025",   "errorMsg": "robotCode不存在",   "success": false,   "trace_id": "2106d98117745950903798357e0b34" }  |
| `dws chat group members remove` | `fail_business` | `0` | {   "code": "PARAM_ERROR",   "error": "参数错误",   "message": "参数错误。请检查输入参数是否正确，参考工具说明调整参数后重试",   "retryable": false,   "success": false,   "technical_detail": "Tool metadata API error: PARAM_ERROR - 未找到指定工具",   "trace_id": |
| `dws chat group rename` | `fail_business` | `0` | {   "code": "PARAM_ERROR",   "error": "参数错误",   "message": "参数错误。请检查输入参数是否正确，参考工具说明调整参数后重试",   "retryable": false,   "success": false,   "technical_detail": "Tool metadata API error: PARAM_ERROR - 未找到指定工具",   "trace_id": |
| `dws chat message list` | `fail_exit` | `3` | {   "error": {     "category": "validation",     "code": 3,     "message": "--group and --user are mutually exclusive"   } }  |
| `dws chat message recall-by-bot` | `fail_business` | `0` | {   "code": "PARAM_ERROR",   "error": "参数错误",   "message": "参数错误。请检查输入参数是否正确，参考工具说明调整参数后重试",   "retryable": false,   "success": false,   "technical_detail": "调用失败，原因: 调用失败，原因: 调用远程服务业务异常, 出参校验不通过。DingOpenResult.success当前 |
| `dws chat message send` | `fail_exit` | `5` | {   "error": {     "category": "internal",     "code": 5,     "message": "accepts 1 arg(s), received 0"   } }  |
| `dws chat message send-by-bot` | `fail_business` | `0` | {   "code": "PARAM_ERROR",   "error": "参数错误",   "message": "参数错误。请检查输入参数是否正确，参考工具说明调整参数后重试",   "retryable": false,   "success": false,   "technical_detail": "调用失败，原因: 调用失败，原因: 调用远程服务业务异常, 出参校验不通过。DingOpenResult.success当前 |
| `dws contact user search-mobile` | `fail_business` | `0` | {   "arguments": [],   "errorCode": "11001",   "errorMsg": "手机格式不正确，请重新输入",   "success": false,   "trace_id": "2106d98117745950982347064e0ab0" }  |
| `dws ding message recall` | `fail_business` | `0` | {   "dingOpenErrcode": 400100,   "errorMsg": "robotCode is illegal",   "result": {},   "success": false,   "trace_id": "0bab027317745950995614026e0761" }  |
| `dws ding message send` | `fail_business` | `0` | {   "dingOpenErrcode": 400100,   "errorMsg": "robotCode is illegal",   "result": {},   "success": false,   "trace_id": "2106d98117745951002636917e0b13" }  |
| `dws report create` | `fail_exit` | `3` | {   "error": {     "category": "validation",     "code": 3,     "message": "json_parse: invalid JSON: invalid character 'ç' looking for beginning of value"   } }  |
| `dws report list` | `fail_business` | `0` | {   "errcode": 40035,   "errorMsg": "不合法的参数",   "result": {     "report_list": null,     "size": 0   },   "success": false,   "trace_id": "213ee25c17745951014526150e0af0" }  |
| `dws report sent` | `fail_business` | `0` | {   "code": "ERROR",   "error": "操作失败",   "message": "操作失败。发生错误，建议稍后重试",   "retryable": true,   "success": false,   "technical_detail": "调用失败，原因: 调用失败，原因: 调用远程服务业务异常, 出参校验不通过。root.success当前值：null， 预期值：true，所有值：{\"stackTr |
| `dws report stats` | `fail_business` | `0` | {   "code": "PARAM_ERROR",   "error": "参数错误",   "message": "参数错误。请检查输入参数是否正确，参考工具说明调整参数后重试",   "retryable": false,   "success": false,   "technical_detail": "调用失败，原因: 调用失败，原因: 调用远程服务业务异常, 出参校验不通过。DingOpenResult.success当前 |
| `dws report template detail` | `fail_business` | `0` | {   "code": "PARAM_ERROR",   "error": "参数错误",   "message": "参数错误。请检查输入参数是否正确，参考工具说明调整参数后重试",   "retryable": false,   "success": false,   "technical_detail": "调用失败，原因: 调用失败，原因: 调用远程服务业务异常, 出参校验不通过。root.success当前值：false， 预 |
| `dws todo task create` | `fail_business` | `0` | {   "arguments": [],   "errorCode": "999",   "errorMsg": "system error: java.lang.NullPointerException",   "success": false,   "trace_id": "213ee25c17745951058325426e0acd" }  |
| `dws todo task get` | `fail_business` | `0` | {   "arguments": [],   "errorCode": "500",   "errorMsg": "NewTodoService#getTodoDetail taskId is null",   "result": {},   "success": false,   "trace_id": "0bab027317745951078651445e07c5" }  |
| `dws completion` | `fail_exit` | `5` | {   "error": {     "category": "internal",     "code": 5,     "message": "accepts 1 arg(s), received 0"   } }  |

## 全量执行明细

| 命令 | 语义结果 | 退出码 | 自动补参个数 |
|---|---|---:|---:|
| `dws aitable attachment upload` | `fail_business` | `0` | `2` |
| `dws aitable base create` | `pass` | `0` | `2` |
| `dws aitable base delete` | `pass` | `0` | `1` |
| `dws aitable base get` | `fail_business` | `0` | `1` |
| `dws aitable base list` | `pass` | `0` | `0` |
| `dws aitable base search` | `pass` | `0` | `1` |
| `dws aitable base update` | `fail_business` | `0` | `3` |
| `dws aitable field create` | `fail_business` | `0` | `2` |
| `dws aitable field delete` | `fail_business` | `0` | `3` |
| `dws aitable field get` | `fail_business` | `0` | `3` |
| `dws aitable field update` | `fail_business` | `0` | `4` |
| `dws aitable record create` | `fail_business` | `0` | `2` |
| `dws aitable record delete` | `fail_business` | `0` | `3` |
| `dws aitable record query` | `fail_business` | `0` | `5` |
| `dws aitable record update` | `fail_business` | `0` | `2` |
| `dws aitable table create` | `fail_business` | `0` | `2` |
| `dws aitable table delete` | `fail_business` | `0` | `2` |
| `dws aitable table get` | `fail_business` | `0` | `2` |
| `dws aitable table update` | `fail_business` | `0` | `3` |
| `dws aitable template search` | `pass` | `0` | `1` |
| `dws attendance record get` | `fail_business` | `0` | `2` |
| `dws attendance rules` | `pass` | `0` | `1` |
| `dws attendance shift list` | `pass` | `0` | `3` |
| `dws attendance summary` | `fail_business` | `0` | `2` |
| `dws calendar busy search` | `pass` | `0` | `3` |
| `dws calendar event create` | `fail_business` | `0` | `4` |
| `dws calendar event delete` | `pass` | `0` | `1` |
| `dws calendar event get` | `pass` | `0` | `1` |
| `dws calendar event list` | `pass` | `0` | `2` |
| `dws calendar event list-mine` | `fail_business` | `0` | `2` |
| `dws calendar event update` | `fail_business` | `0` | `4` |
| `dws calendar participant add` | `fail_business` | `0` | `1` |
| `dws calendar participant delete` | `fail_business` | `0` | `1` |
| `dws calendar participant list` | `fail_business` | `0` | `0` |
| `dws calendar room add` | `fail_business` | `0` | `0` |
| `dws calendar room delete` | `fail_business` | `0` | `0` |
| `dws calendar room list-groups` | `pass` | `0` | `0` |
| `dws calendar room search` | `fail_business` | `0` | `3` |
| `dws chat bot search` | `pass` | `0` | `3` |
| `dws chat group create` | `fail_business` | `0` | `2` |
| `dws chat group members add` | `fail_business` | `0` | `2` |
| `dws chat group members add-bot` | `fail_business` | `0` | `2` |
| `dws chat group members remove` | `fail_business` | `0` | `2` |
| `dws chat group rename` | `fail_business` | `0` | `2` |
| `dws chat message list` | `fail_exit` | `3` | `4` |
| `dws chat message recall-by-bot` | `fail_business` | `0` | `0` |
| `dws chat message send` | `fail_exit` | `5` | `3` |
| `dws chat message send-by-bot` | `fail_business` | `0` | `2` |
| `dws chat message send-by-webhook` | `pass` | `0` | `4` |
| `dws chat search` | `pass` | `0` | `1` |
| `dws contact dept list-children` | `pass` | `0` | `1` |
| `dws contact dept list-members` | `pass` | `0` | `1` |
| `dws contact dept search` | `pass` | `0` | `1` |
| `dws contact user get` | `pass` | `0` | `1` |
| `dws contact user get-self` | `pass` | `0` | `0` |
| `dws contact user search` | `pass` | `0` | `1` |
| `dws contact user search-mobile` | `fail_business` | `0` | `0` |
| `dws devdoc article search` | `pass` | `0` | `1` |
| `dws ding message recall` | `fail_business` | `0` | `1` |
| `dws ding message send` | `fail_business` | `0` | `2` |
| `dws report create` | `fail_exit` | `3` | `3` |
| `dws report detail` | `pass` | `0` | `1` |
| `dws report list` | `fail_business` | `0` | `2` |
| `dws report sent` | `fail_business` | `0` | `3` |
| `dws report stats` | `fail_business` | `0` | `1` |
| `dws report template detail` | `fail_business` | `0` | `1` |
| `dws report template list` | `pass` | `0` | `0` |
| `dws todo task create` | `fail_business` | `0` | `1` |
| `dws todo task delete` | `pass` | `0` | `1` |
| `dws todo task done` | `pass` | `0` | `1` |
| `dws todo task get` | `fail_business` | `0` | `1` |
| `dws todo task list` | `pass` | `0` | `1` |
| `dws todo task update` | `pass` | `0` | `2` |
| `dws workbench app get` | `pass` | `0` | `1` |
| `dws workbench app list` | `pass` | `0` | `0` |
| `dws auth login` | `pass` | `0` | `3` |
| `dws auth logout` | `pass` | `0` | `0` |
| `dws auth reset` | `pass` | `0` | `0` |
| `dws auth status` | `pass` | `0` | `0` |
| `dws cache refresh` | `pass` | `0` | `0` |
| `dws cache status` | `pass` | `0` | `0` |
| `dws completion` | `fail_exit` | `5` | `0` |
| `dws version` | `pass` | `0` | `0` |
