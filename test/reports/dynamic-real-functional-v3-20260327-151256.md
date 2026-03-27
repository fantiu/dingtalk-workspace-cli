# CLI 动态发现真实功能测试报告（第三轮定向强化）

| 项目 | 值 |
|---|---|
| 时间 | 2026-03-27 15:14:03 |
| 动态服务数 | `10` |
| 命令节点(`--help`) | `118` |
| `--help` 失败 | `0` |
| 叶子命令执行数 | `83` |
| 语义成功(pass) | `36` |
| 退出码失败 | `1` |
| 业务失败 | `46` |

## 本轮增强点

- 补齐位置参数（completion/chat message send）。
- 增强 event/base/table 依赖 ID 回填。
- report contents 统一 JSON 串。

## 关键上下文

- `base-id`: `(未获取)`
- `table-id`: `(未获取)`
- `event-id`: `(未获取)`

## 失败明细

| 命令 | 分类 | 退出码 | 摘要 |
|---|---|---:|---|
| `dws aitable attachment upload` | `fail_business` | `0` | {   "data": {},   "error": {     "code": "INVALID_PREPARE_ATTACHMENT_UPLOAD_REQUEST",     "message": "fileName must include an extension",     "retryable": false,     "type": "INPUT_ERROR"   },   "meta": {},   "status":  |
| `dws aitable base delete` | `fail_business` | `0` | {   "data": {},   "error": {     "code": "52600003",     "message": "Data not found",     "retryable": true,     "type": "SYSTEM_ERROR"   },   "meta": {},   "status": "error",   "summary": "Failed to delete base test001. |
| `dws aitable base get` | `fail_business` | `0` | {   "data": {},   "error": {     "code": "BASE_NOT_FOUND",     "message": "Specified base does not exist, has been deleted, or is inaccessible",     "retryable": false,     "type": "INPUT_ERROR"   },   "meta": {},   "sta |
| `dws aitable base update` | `fail_business` | `0` | {   "data": {},   "error": {     "code": "BASE_NOT_FOUND",     "message": "Specified base does not exist, has been deleted, or is inaccessible",     "retryable": false,     "type": "INPUT_ERROR"   },   "meta": {},   "sta |
| `dws aitable field create` | `fail_business` | `0` | {   "data": {},   "error": {     "code": "INVALID_PARAMS",     "message": "fields are required",     "retryable": false,     "type": "INPUT_ERROR"   },   "meta": {},   "status": "error",   "summary": "Failed to create fi |
| `dws aitable field delete` | `fail_business` | `0` | {   "data": {},   "error": {     "code": "404",     "message": "getDentryDTO returns null",     "retryable": true,     "type": "SYSTEM_ERROR"   },   "meta": {},   "status": "error",   "summary": "Failed to get current fi |
| `dws aitable field get` | `fail_business` | `0` | {   "data": {},   "error": {     "code": "GET_FIELDS_ERROR",     "message": "Failed to load field schema from sheet service",     "retryable": true,     "type": "SYSTEM_ERROR"   },   "meta": {},   "status": "error",   "s |
| `dws aitable field update` | `fail_business` | `0` | {   "data": {},   "error": {     "code": "404",     "message": "getDentryDTO returns null",     "retryable": true,     "type": "SYSTEM_ERROR"   },   "meta": {},   "status": "error",   "summary": "Failed to get current fi |
| `dws aitable record create` | `fail_business` | `0` | {   "data": {},   "error": {     "code": "INVALID_RECORDS",     "message": "records must contain at least one record",     "retryable": false,     "type": "INPUT_ERROR"   },   "meta": {},   "status": "error",   "summary" |
| `dws aitable record delete` | `fail_business` | `0` | {   "data": {},   "error": {     "code": "404",     "message": "getDentryDTO returns null",     "retryable": true,     "type": "SYSTEM_ERROR"   },   "meta": {},   "status": "error",   "summary": "Failed to delete records |
| `dws aitable record query` | `fail_business` | `0` | {   "data": {},   "error": {     "code": "INVALID_BASE_ID",     "message": "failed to resolve docId from baseId",     "retryable": false,     "type": "INPUT_ERROR"   },   "meta": {},   "status": "error",   "summary": "Fa |
| `dws aitable record update` | `fail_business` | `0` | {   "data": {},   "error": {     "code": "INVALID_RECORDS",     "message": "records must contain at least one record",     "retryable": false,     "type": "INPUT_ERROR"   },   "meta": {},   "status": "error",   "summary" |
| `dws aitable table create` | `fail_business` | `0` | {   "data": {},   "error": {     "code": "404",     "message": "getDentryDTO returns null",     "retryable": true,     "type": "SYSTEM_ERROR"   },   "meta": {},   "status": "error",   "summary": "Failed to create table ' |
| `dws aitable table delete` | `fail_business` | `0` | {   "data": {},   "error": {     "code": "BASE_NOT_FOUND",     "message": "Specified base does not exist, has been deleted, or is inaccessible",     "retryable": false,     "type": "INPUT_ERROR"   },   "meta": {},   "sta |
| `dws aitable table get` | `fail_business` | `0` | {   "data": {},   "error": {     "code": "BASE_NOT_FOUND",     "message": "Specified base does not exist, has been deleted, or is inaccessible",     "retryable": false,     "type": "INPUT_ERROR"   },   "meta": {},   "sta |
| `dws aitable table update` | `fail_business` | `0` | {   "data": {},   "error": {     "code": "404",     "message": "getDentryDTO returns null",     "retryable": true,     "type": "SYSTEM_ERROR"   },   "meta": {},   "status": "error",   "summary": "Failed to update table t |
| `dws attendance record get` | `fail_business` | `0` | {   "code": "ERROR",   "error": "操作失败",   "message": "操作失败。发生错误，建议稍后重试",   "retryable": true,   "success": false,   "technical_detail": "调用失败，原因: 调用失败，原因: 调用远程服务业务异常, 出参校验不通过。root.success当前值：false， 预期值：true，所有值：{\"code\" |
| `dws attendance summary` | `fail_business` | `0` | {   "code": "C0002",   "message": "统计类型错误",   "result": {},   "success": false,   "trace_id": "0b5deb3217745956001386733e0c8d" }  |
| `dws calendar event create` | `fail_business` | `0` | {   "arguments": [],   "errorCode": "300000",   "errorMsg": "code: 300000, developerMessage: Invalid start dateTime format.",   "result": {     "attendees": null,     "reminders": null   },   "success": false,   "trace_i |
| `dws calendar event delete` | `fail_business` | `0` | {   "arguments": [],   "errorCode": "300014",   "errorMsg": "Event not exist",   "result": {},   "success": false,   "trace_id": "0b5deb3217745956024626876e0c8d" }  |
| `dws calendar event get` | `fail_business` | `0` | {   "arguments": [],   "result": {     "attendees": null,     "categories": null,     "reminders": null   },   "success": false,   "trace_id": "0b5deb3217745956032378417e0be6" }  |
| `dws calendar event list-mine` | `fail_business` | `0` | {   "code": "PARAM_ERROR",   "error": "参数错误",   "message": "参数错误。请检查输入参数是否正确，参考工具说明调整参数后重试",   "retryable": false,   "success": false,   "technical_detail": "Tool metadata API error: PARAM_ERROR - 未找到指定工具",   "trace_id": |
| `dws calendar event update` | `fail_business` | `0` | {   "arguments": [],   "errorCode": "300000",   "errorMsg": "code: 300000, developerMessage: Invalid start dateTime format.",   "result": {},   "success": false,   "trace_id": "0bb7c36217745956048842137e04c8" }  |
| `dws calendar participant add` | `fail_business` | `0` | {   "arguments": [],   "errorCode": "300000",   "errorMsg": "code: 300000, developerMessage: eventId cannot be blank",   "success": false,   "trace_id": "0bb7c36217745956055183039e04a5" }  |
| `dws calendar participant delete` | `fail_business` | `0` | {   "arguments": [],   "errorCode": "300000",   "errorMsg": "code: 300000, developerMessage: eventId cannot be blank",   "result": {},   "success": false,   "trace_id": "0bb7c36217745956062986584e04a6" }  |
| `dws calendar participant list` | `fail_business` | `0` | {   "arguments": [],   "errorCode": "300000",   "errorMsg": "code: 300000, developerMessage: eventId cannot be blank",   "result": {     "attendees": null   },   "success": false,   "trace_id": "0b5deb3217745956068548731 |
| `dws calendar room add` | `fail_business` | `0` | {   "arguments": [],   "errorCode": "300000",   "errorMsg": "code: 300000, developerMessage: eventId cannot be blank",   "result": {},   "success": false,   "trace_id": "0bb7c36217745956076466697e04a6" }  |
| `dws calendar room delete` | `fail_business` | `0` | {   "arguments": [],   "errorCode": "300000",   "errorMsg": "code: 300000, developerMessage: eventId cannot be blank",   "result": {},   "success": false,   "trace_id": "2127d89817745956082684802e0aa2" }  |
| `dws calendar room search` | `fail_business` | `0` | {   "arguments": [],   "errorCode": "458019",   "errorMsg": "查询范围内的会议室数量，超过上限100，请选择更小范围的分组进行查询。",   "result": {     "result": [       {         "labels": null       }     ]   },   "success": false,   "trace_id": "0b5deb |
| `dws chat group create` | `fail_business` | `0` | {   "code": "PARAM_ERROR",   "error": "参数错误",   "message": "参数错误。请检查输入参数是否正确，参考工具说明调整参数后重试",   "retryable": false,   "success": false,   "technical_detail": "Tool metadata API error: PARAM_ERROR - 未找到指定工具",   "trace_id": |
| `dws chat group members add` | `fail_business` | `0` | {   "code": "PARAM_ERROR",   "error": "参数错误",   "message": "参数错误。请检查输入参数是否正确，参考工具说明调整参数后重试",   "retryable": false,   "success": false,   "technical_detail": "Tool metadata API error: PARAM_ERROR - 未找到指定工具",   "trace_id": |
| `dws chat group members add-bot` | `fail_business` | `0` | {   "errorCode": "500025",   "errorMsg": "robotCode不存在",   "success": false,   "trace_id": "0bb7c36217745956120232510e058c" }  |
| `dws chat group members remove` | `fail_business` | `0` | {   "code": "PARAM_ERROR",   "error": "参数错误",   "message": "参数错误。请检查输入参数是否正确，参考工具说明调整参数后重试",   "retryable": false,   "success": false,   "technical_detail": "Tool metadata API error: PARAM_ERROR - 未找到指定工具",   "trace_id": |
| `dws chat group rename` | `fail_business` | `0` | {   "code": "PARAM_ERROR",   "error": "参数错误",   "message": "参数错误。请检查输入参数是否正确，参考工具说明调整参数后重试",   "retryable": false,   "success": false,   "technical_detail": "Tool metadata API error: PARAM_ERROR - 未找到指定工具",   "trace_id": |
| `dws chat message list` | `fail_business` | `0` | {   "code": "PARAM_ERROR",   "error": "参数错误",   "message": "参数错误。请检查输入参数是否正确，参考工具说明调整参数后重试",   "retryable": false,   "success": false,   "technical_detail": "Tool metadata API error: PARAM_ERROR - 未找到指定工具",   "trace_id": |
| `dws chat message recall-by-bot` | `fail_business` | `0` | {   "code": "PARAM_ERROR",   "error": "参数错误",   "message": "参数错误。请检查输入参数是否正确，参考工具说明调整参数后重试",   "retryable": false,   "success": false,   "technical_detail": "调用失败，原因: 调用失败，原因: 调用远程服务业务异常, 出参校验不通过。DingOpenResult.success当前 |
| `dws chat message send` | `fail_exit` | `3` | {   "error": {     "category": "validation",     "code": 3,     "message": "--group and --user are mutually exclusive"   } }  |
| `dws chat message send-by-bot` | `fail_business` | `0` | {   "code": "PARAM_ERROR",   "error": "参数错误",   "message": "参数错误。请检查输入参数是否正确，参考工具说明调整参数后重试",   "retryable": false,   "success": false,   "technical_detail": "调用失败，原因: 调用失败，原因: 调用远程服务业务异常, 出参校验不通过。DingOpenResult.success当前 |
| `dws ding message recall` | `fail_business` | `0` | {   "dingOpenErrcode": 400100,   "errorMsg": "robotCode is illegal",   "result": {},   "success": false,   "trace_id": "0b5deb3217745956239415766e0c28" }  |
| `dws ding message send` | `fail_business` | `0` | {   "dingOpenErrcode": 400100,   "errorMsg": "robotCode is illegal",   "result": {},   "success": false,   "trace_id": "2104a64c17745956244614506e0a3b" }  |
| `dws report create` | `fail_business` | `0` | {   "code": "PARAM_ERROR",   "error": "参数错误",   "message": "参数错误。请检查输入参数是否正确，参考工具说明调整参数后重试",   "retryable": false,   "success": false,   "technical_detail": "调用失败，原因: 调用失败，原因: 调用远程服务业务异常, 出参校验不通过。root.success当前值：false， 预 |
| `dws report list` | `fail_business` | `0` | {   "errcode": 40035,   "errorMsg": "不合法的参数",   "result": {     "report_list": null,     "size": 0   },   "success": false,   "trace_id": "0b5deb3217745956268642260e0be6" }  |
| `dws report sent` | `fail_business` | `0` | {   "code": "ERROR",   "error": "操作失败",   "message": "操作失败。发生错误，建议稍后重试",   "retryable": true,   "success": false,   "technical_detail": "调用失败，原因: 调用失败，原因: 调用远程服务业务异常, 出参校验不通过。root.success当前值：null， 预期值：true，所有值：{\"stackTr |
| `dws report stats` | `fail_business` | `0` | {   "code": "PARAM_ERROR",   "error": "参数错误",   "message": "参数错误。请检查输入参数是否正确，参考工具说明调整参数后重试",   "retryable": false,   "success": false,   "technical_detail": "调用失败，原因: 调用失败，原因: 调用远程服务业务异常, 出参校验不通过。DingOpenResult.success当前 |
| `dws report template detail` | `fail_business` | `0` | {   "code": "PARAM_ERROR",   "error": "参数错误",   "message": "参数错误。请检查输入参数是否正确，参考工具说明调整参数后重试",   "retryable": false,   "success": false,   "technical_detail": "调用失败，原因: 调用失败，原因: 调用远程服务业务异常, 出参校验不通过。root.success当前值：false， 预 |
| `dws todo task create` | `fail_business` | `0` | {   "arguments": [],   "errorCode": "500",   "errorMsg": "TodoServiceImpl#createTask: Unrecognized priority type: 1",   "success": false,   "trace_id": "0bb7c36217745956307524175e04c8" }  |
| `dws todo task get` | `fail_business` | `0` | {   "arguments": [],   "errorCode": "500",   "errorMsg": "NewTodoService#getTodoDetail taskId is null",   "result": {},   "success": false,   "trace_id": "2127d89817745956327921364e09fc" }  |
