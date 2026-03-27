# CLI 动态服务发现全量功能测试报告（完整重跑）

| 指标 | 值 |
|---|---|
| 时间 | 2026-03-27 15:34:46 |
| 动态服务数 | `10` |
| 命令节点(`--help`) | `118` |
| `--help` 失败 | `0` |
| 叶子命令执行数 | `83` |
| 语义成功(pass) | `50` |
| 退出码失败 | `3` |
| 业务失败 | `30` |

## 关键上下文

- `uid`: `061978`
- `event-id`: `ZWVVS1Y2S3VLalpwZFJPK2l1YitKZz09`
- `room-id`: `bdff146f0172e49511c91904cf69b186003e6834fdd7fee7`
- `base-id`: `7dx2rn0Jbakrq9O5t6NRK1DRVMGjLRb3`
- `table-id`: `iFiXjY8`
- `task-id`: `51681032441`
- `template-id`: `18b46f75236c558cc63be07419f8d65f`

## 失败明细

| 命令 | 分类 | 退出码 | 摘要 |
|---|---|---:|---|
| `dws aitable attachment upload` | `fail_business` | `0` | {   "data": {},   "error": {     "code": "INVALID_PREPARE_ATTACHMENT_UPLOAD_REQUEST",     "message": "fileName is required",     "retryable": false,     "type": "INPUT_ERROR"   },   "meta": {},   "status": "error",   "su |
| `dws aitable base create` | `fail_business` | `0` | {   "data": {},   "error": {     "code": "INVALID_NAME",     "message": "baseName is required and cannot be empty or only whitespace",     "retryable": false,     "type": "INPUT_ERROR"   },   "meta": {},   "status": "err |
| `dws aitable base get` | `fail_business` | `0` | {   "data": {},   "error": {     "code": "BASE_NOT_FOUND",     "message": "Specified base does not exist, has been deleted, or is inaccessible",     "retryable": false,     "type": "INPUT_ERROR"   },   "meta": {},   "sta |
| `dws aitable base update` | `fail_business` | `0` | {   "data": {},   "error": {     "code": "BASE_NOT_FOUND",     "message": "Specified base does not exist, has been deleted, or is inaccessible",     "retryable": false,     "type": "INPUT_ERROR"   },   "meta": {},   "sta |
| `dws aitable field delete` | `fail_business` | `0` | {   "data": {},   "error": {     "code": "404",     "message": "getDentryDTO returns null",     "retryable": true,     "type": "SYSTEM_ERROR"   },   "meta": {},   "status": "error",   "summary": "Failed to get current fi |
| `dws aitable field get` | `fail_business` | `0` | {   "data": {},   "error": {     "code": "GET_FIELDS_ERROR",     "message": "Failed to load field schema from sheet service",     "retryable": true,     "type": "SYSTEM_ERROR"   },   "meta": {},   "status": "error",   "s |
| `dws aitable field update` | `fail_business` | `0` | {   "data": {},   "error": {     "code": "INVALID_PARAMS",     "message": "At least one of newFieldName, config or aiConfig must be provided for update",     "retryable": false,     "type": "INPUT_ERROR"   },   "meta": { |
| `dws aitable record create` | `fail_business` | `0` | {   "data": {},   "error": {     "code": "404",     "message": "getDentryDTO returns null",     "retryable": true,     "type": "SYSTEM_ERROR"   },   "meta": {},   "status": "error",   "summary": "Failed to create records |
| `dws aitable record delete` | `fail_exit` | `3` | {   "error": {     "category": "validation",     "code": 3,     "message": "--record-ids is required"   } }  |
| `dws aitable record query` | `fail_business` | `0` | {   "data": {},   "error": {     "code": "INVALID_BASE_ID",     "message": "failed to resolve docId from baseId",     "retryable": false,     "type": "INPUT_ERROR"   },   "meta": {},   "status": "error",   "summary": "Fa |
| `dws aitable record update` | `fail_business` | `0` | {   "data": {},   "error": {     "code": "INVALID_RECORDS",     "message": "records must contain at least one writable record",     "retryable": false,     "type": "INPUT_ERROR"   },   "meta": {},   "status": "error",    |
| `dws aitable table create` | `fail_business` | `0` | {   "data": {},   "error": {     "code": "INVALID_PARAMS",     "message": "tableName is required",     "retryable": false,     "type": "INPUT_ERROR"   },   "meta": {},   "status": "error",   "summary": "Failed to create  |
| `dws aitable table delete` | `fail_business` | `0` | {   "data": {},   "error": {     "code": "BASE_NOT_FOUND",     "message": "Specified base does not exist, has been deleted, or is inaccessible",     "retryable": false,     "type": "INPUT_ERROR"   },   "meta": {},   "sta |
| `dws aitable table get` | `fail_business` | `0` | {   "data": {},   "error": {     "code": "BASE_NOT_FOUND",     "message": "Specified base does not exist, has been deleted, or is inaccessible",     "retryable": false,     "type": "INPUT_ERROR"   },   "meta": {},   "sta |
| `dws attendance summary` | `fail_business` | `0` | {   "code": "C0002",   "message": "统计类型错误",   "result": {},   "success": false,   "trace_id": "0bb7c36217745968417782683e04a7" }  |
| `dws calendar event list-mine` | `fail_business` | `0` | {   "code": "PARAM_ERROR",   "error": "参数错误",   "message": "参数错误。请检查输入参数是否正确，参考工具说明调整参数后重试",   "retryable": false,   "success": false,   "technical_detail": "Tool metadata API error: PARAM_ERROR - 未找到指定工具",   "trace_id": |
| `dws calendar event update` | `fail_business` | `0` | {   "arguments": [],   "errorCode": "300000",   "errorMsg": "code: 300000, developerMessage: Cannot patch event that is not in 'confirmed' status.",   "result": {},   "success": false,   "trace_id": "2127d898177459684755 |
| `dws calendar participant delete` | `fail_business` | `0` | {   "arguments": [],   "errorCode": "300000",   "errorMsg": "code: 300000, developerMessage: can not remove organizer.",   "result": {},   "success": false,   "trace_id": "2127d89817745968487392353e09fc" }  |
| `dws calendar room add` | `fail_business` | `0` | {   "arguments": [],   "errorCode": "300000",   "errorMsg": "code: 300000, developerMessage: Cannot patch event that is not in 'confirmed' status.",   "result": {},   "success": false,   "trace_id": "2104a64c177459685002 |
| `dws chat group create` | `fail_business` | `0` | {   "code": "PARAM_ERROR",   "error": "参数错误",   "message": "参数错误。请检查输入参数是否正确，参考工具说明调整参数后重试",   "retryable": false,   "success": false,   "technical_detail": "Tool metadata API error: PARAM_ERROR - 未找到指定工具",   "trace_id": |
| `dws chat group members add` | `fail_business` | `0` | {   "code": "PARAM_ERROR",   "error": "参数错误",   "message": "参数错误。请检查输入参数是否正确，参考工具说明调整参数后重试",   "retryable": false,   "success": false,   "technical_detail": "Tool metadata API error: PARAM_ERROR - 未找到指定工具",   "trace_id": |
| `dws chat group members add-bot` | `fail_business` | `0` | {   "errorCode": "500025",   "errorMsg": "robotCode不存在",   "success": false,   "trace_id": "2127d89817745968545635804e0aa2" }  |
| `dws chat group members remove` | `fail_business` | `0` | {   "code": "PARAM_ERROR",   "error": "参数错误",   "message": "参数错误。请检查输入参数是否正确，参考工具说明调整参数后重试",   "retryable": false,   "success": false,   "technical_detail": "Tool metadata API error: PARAM_ERROR - 未找到指定工具",   "trace_id": |
| `dws chat group rename` | `fail_business` | `0` | {   "code": "PARAM_ERROR",   "error": "参数错误",   "message": "参数错误。请检查输入参数是否正确，参考工具说明调整参数后重试",   "retryable": false,   "success": false,   "technical_detail": "Tool metadata API error: PARAM_ERROR - 未找到指定工具",   "trace_id": |
| `dws chat message list` | `fail_business` | `0` | {   "code": "PARAM_ERROR",   "error": "参数错误",   "message": "参数错误。请检查输入参数是否正确，参考工具说明调整参数后重试",   "retryable": false,   "success": false,   "technical_detail": "Tool metadata API error: PARAM_ERROR - 未找到指定工具",   "trace_id": |
| `dws chat message recall-by-bot` | `fail_business` | `0` | {   "code": "PARAM_ERROR",   "error": "参数错误",   "message": "参数错误。请检查输入参数是否正确，参考工具说明调整参数后重试",   "retryable": false,   "success": false,   "technical_detail": "调用失败，原因: 调用失败，原因: 调用远程服务业务异常, 出参校验不通过。DingOpenResult.success当前 |
| `dws chat message send` | `fail_exit` | `3` | {   "error": {     "category": "validation",     "code": 3,     "message": "--group and --user are mutually exclusive"   } }  |
| `dws chat message send-by-bot` | `fail_business` | `0` | {   "code": "PARAM_ERROR",   "error": "参数错误",   "message": "参数错误。请检查输入参数是否正确，参考工具说明调整参数后重试",   "retryable": false,   "success": false,   "technical_detail": "调用失败，原因: 调用失败，原因: 调用远程服务业务异常, 出参校验不通过。DingOpenResult.success当前 |
| `dws ding message recall` | `fail_business` | `0` | {   "dingOpenErrcode": 400100,   "errorMsg": "robotCode is illegal",   "result": {},   "success": false,   "trace_id": "2127d89817745968660356774e0a02" }  |
| `dws ding message send` | `fail_business` | `0` | {   "dingOpenErrcode": 400100,   "errorMsg": "robotCode is illegal",   "result": {},   "success": false,   "trace_id": "2104a64c17745968669262499e0a07" }  |
| `dws report create` | `fail_business` | `0` | {   "code": "PARAM_ERROR",   "error": "参数错误",   "message": "参数错误。请检查输入参数是否正确，参考工具说明调整参数后重试",   "retryable": false,   "success": false,   "technical_detail": "调用失败，原因: 调用失败，原因: 调用远程服务业务异常, 出参校验不通过。root.success当前值：false， 预 |
| `dws report stats` | `fail_business` | `0` | {   "code": "PARAM_ERROR",   "error": "参数错误",   "message": "参数错误。请检查输入参数是否正确，参考工具说明调整参数后重试",   "retryable": false,   "success": false,   "technical_detail": "调用失败，原因: 调用失败，原因: 调用远程服务业务异常, 出参校验不通过。DingOpenResult.success当前 |
| `dws todo task get` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
