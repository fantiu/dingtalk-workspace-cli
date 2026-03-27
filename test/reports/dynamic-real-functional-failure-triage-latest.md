# 动态CLI失败项分层清单（latest）

- 来源报告: `/Users/tianlei.qjb/Documents/my_python_project/cli/test/reports/dynamic-real-functional-full-rerun-latest.md`
- 失败总数: `33`

| 分类 | 数量 |
|---|---:|
| 其他待分析 | 14 |
| 测试数据依赖不足 | 6 |
| 服务端能力缺失 | 6 |
| 业务规则限制 | 3 |
| 缺少真实机器人配置 | 3 |
| 命令参数映射问题 | 1 |

## 其他待分析

| 命令 | 类型 | 退出码 | 摘要 |
|---|---|---:|---|
| `dws aitable base get` | `fail_business` | `0` | {   "data": {},   "error": {     "code": "BASE_NOT_FOUND",     "message": "Specified base does not exist, has been deleted, or is inaccessible",     "retryable": false,     "type": |
| `dws aitable base update` | `fail_business` | `0` | {   "data": {},   "error": {     "code": "BASE_NOT_FOUND",     "message": "Specified base does not exist, has been deleted, or is inaccessible",     "retryable": false,     "type": |
| `dws aitable field delete` | `fail_business` | `0` | {   "data": {},   "error": {     "code": "404",     "message": "getDentryDTO returns null",     "retryable": true,     "type": "SYSTEM_ERROR"   },   "meta": {},   "status": "error" |
| `dws aitable field get` | `fail_business` | `0` | {   "data": {},   "error": {     "code": "GET_FIELDS_ERROR",     "message": "Failed to load field schema from sheet service",     "retryable": true,     "type": "SYSTEM_ERROR"   }, |
| `dws aitable record create` | `fail_business` | `0` | {   "data": {},   "error": {     "code": "404",     "message": "getDentryDTO returns null",     "retryable": true,     "type": "SYSTEM_ERROR"   },   "meta": {},   "status": "error" |
| `dws aitable record delete` | `fail_business` | `0` | {   "data": {},   "error": {     "code": "404",     "message": "getDentryDTO returns null",     "retryable": true,     "type": "SYSTEM_ERROR"   },   "meta": {},   "status": "error" |
| `dws aitable table delete` | `fail_business` | `0` | {   "data": {},   "error": {     "code": "BASE_NOT_FOUND",     "message": "Specified base does not exist, has been deleted, or is inaccessible",     "retryable": false,     "type": |
| `dws aitable table get` | `fail_business` | `0` | {   "data": {},   "error": {     "code": "BASE_NOT_FOUND",     "message": "Specified base does not exist, has been deleted, or is inaccessible",     "retryable": false,     "type": |
| `dws attendance summary` | `fail_business` | `0` | {   "code": "C0002",   "message": "统计类型错误",   "result": {},   "success": false,   "trace_id": "0bab027317745976626518395e0780" } |
| `dws chat message recall-by-bot` | `fail_business` | `0` | {   "code": "PARAM_ERROR",   "error": "参数错误",   "message": "参数错误。请检查输入参数是否正确，参考工具说明调整参数后重试",   "retryable": false,   "success": false,   "technical_detail": "调用失败，原因: 调用失败，原因: 调用远程 |
| `dws chat message send-by-bot` | `fail_business` | `0` | {   "code": "PARAM_ERROR",   "error": "参数错误",   "message": "参数错误。请检查输入参数是否正确，参考工具说明调整参数后重试",   "retryable": false,   "success": false,   "technical_detail": "调用失败，原因: 调用失败，原因: 调用远程 |
| `dws report create` | `fail_business` | `0` | {   "code": "PARAM_ERROR",   "error": "参数错误",   "message": "参数错误。请检查输入参数是否正确，参考工具说明调整参数后重试",   "retryable": false,   "success": false,   "technical_detail": "调用失败，原因: 调用失败，原因: 调用远程 |
| `dws report stats` | `fail_business` | `0` | {   "code": "PARAM_ERROR",   "error": "参数错误",   "message": "参数错误。请检查输入参数是否正确，参考工具说明调整参数后重试",   "retryable": false,   "success": false,   "technical_detail": "调用失败，原因: 调用失败，原因: 调用远程 |
| `dws todo task get` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected |

## 测试数据依赖不足

| 命令 | 类型 | 退出码 | 摘要 |
|---|---|---:|---|
| `dws aitable attachment upload` | `fail_business` | `0` | {   "data": {},   "error": {     "code": "INVALID_PREPARE_ATTACHMENT_UPLOAD_REQUEST",     "message": "fileName is required",     "retryable": false,     "type": "INPUT_ERROR"   },  |
| `dws aitable base create` | `fail_business` | `0` | {   "data": {},   "error": {     "code": "INVALID_NAME",     "message": "baseName is required and cannot be empty or only whitespace",     "retryable": false,     "type": "INPUT_ER |
| `dws aitable field update` | `fail_business` | `0` | {   "data": {},   "error": {     "code": "INVALID_PARAMS",     "message": "At least one of newFieldName, config or aiConfig must be provided for update",     "retryable": false,    |
| `dws aitable record query` | `fail_business` | `0` | {   "data": {},   "error": {     "code": "INVALID_BASE_ID",     "message": "failed to resolve docId from baseId",     "retryable": false,     "type": "INPUT_ERROR"   },   "meta": { |
| `dws aitable record update` | `fail_business` | `0` | {   "data": {},   "error": {     "code": "INVALID_RECORDS",     "message": "records must contain at least one writable record",     "retryable": false,     "type": "INPUT_ERROR"    |
| `dws aitable table create` | `fail_business` | `0` | {   "data": {},   "error": {     "code": "INVALID_PARAMS",     "message": "tableName is required",     "retryable": false,     "type": "INPUT_ERROR"   },   "meta": {},   "status":  |

## 服务端能力缺失

| 命令 | 类型 | 退出码 | 摘要 |
|---|---|---:|---|
| `dws calendar event list-mine` | `fail_business` | `0` | {   "code": "PARAM_ERROR",   "error": "参数错误",   "message": "参数错误。请检查输入参数是否正确，参考工具说明调整参数后重试",   "retryable": false,   "success": false,   "technical_detail": "Tool metadata API erro |
| `dws chat group create` | `fail_business` | `0` | {   "code": "PARAM_ERROR",   "error": "参数错误",   "message": "参数错误。请检查输入参数是否正确，参考工具说明调整参数后重试",   "retryable": false,   "success": false,   "technical_detail": "Tool metadata API erro |
| `dws chat group members add` | `fail_business` | `0` | {   "code": "PARAM_ERROR",   "error": "参数错误",   "message": "参数错误。请检查输入参数是否正确，参考工具说明调整参数后重试",   "retryable": false,   "success": false,   "technical_detail": "Tool metadata API erro |
| `dws chat group members remove` | `fail_business` | `0` | {   "code": "PARAM_ERROR",   "error": "参数错误",   "message": "参数错误。请检查输入参数是否正确，参考工具说明调整参数后重试",   "retryable": false,   "success": false,   "technical_detail": "Tool metadata API erro |
| `dws chat group rename` | `fail_business` | `0` | {   "code": "PARAM_ERROR",   "error": "参数错误",   "message": "参数错误。请检查输入参数是否正确，参考工具说明调整参数后重试",   "retryable": false,   "success": false,   "technical_detail": "Tool metadata API erro |
| `dws chat message list` | `fail_business` | `0` | {   "code": "PARAM_ERROR",   "error": "参数错误",   "message": "参数错误。请检查输入参数是否正确，参考工具说明调整参数后重试",   "retryable": false,   "success": false,   "technical_detail": "Tool metadata API erro |

## 业务规则限制

| 命令 | 类型 | 退出码 | 摘要 |
|---|---|---:|---|
| `dws calendar event update` | `fail_business` | `0` | {   "arguments": [],   "errorCode": "300000",   "errorMsg": "code: 300000, developerMessage: Cannot patch event that is not in 'confirmed' status.",   "result": {},   "success": fa |
| `dws calendar participant delete` | `fail_business` | `0` | {   "arguments": [],   "errorCode": "300000",   "errorMsg": "code: 300000, developerMessage: can not remove organizer.",   "result": {},   "success": false,   "trace_id": "0bab0273 |
| `dws calendar room add` | `fail_business` | `0` | {   "arguments": [],   "errorCode": "300000",   "errorMsg": "code: 300000, developerMessage: Cannot patch event that is not in 'confirmed' status.",   "result": {},   "success": fa |

## 缺少真实机器人配置

| 命令 | 类型 | 退出码 | 摘要 |
|---|---|---:|---|
| `dws chat group members add-bot` | `fail_business` | `0` | {   "errorCode": "500025",   "errorMsg": "robotCode不存在",   "success": false,   "trace_id": "213ee25c17745976757057733e0b11" } |
| `dws ding message recall` | `fail_business` | `0` | {   "dingOpenErrcode": 400100,   "errorMsg": "robotCode is illegal",   "result": {},   "success": false,   "trace_id": "0bab027317745976874644796e07a4" } |
| `dws ding message send` | `fail_business` | `0` | {   "dingOpenErrcode": 400100,   "errorMsg": "robotCode is illegal",   "result": {},   "success": false,   "trace_id": "2132f5ca17745976882365774e0a19" } |

## 命令参数映射问题

| 命令 | 类型 | 退出码 | 摘要 |
|---|---|---:|---|
| `dws chat message send` | `fail_exit` | `3` | {   "error": {     "category": "validation",     "code": 3,     "message": "--group and --user are mutually exclusive"   } } |
