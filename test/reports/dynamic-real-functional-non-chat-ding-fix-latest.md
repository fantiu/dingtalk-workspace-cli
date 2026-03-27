# 非Chat/Ding问题继续修复回归报告

| 指标 | 值 |
|---|---|
| 时间 | 2026-03-27 15:59:25 |
| 重跑项数 | `14` |
| pass | `9` |
| fail_exit | `0` |
| fail_business | `5` |

- uid: `061978`
- base-id: `1zknDm0WRz0NZeXQu2zY5PLaWBQEx5rG`
- table-id: `qfszJ1P`
- event-id: `OG0xR0FVMThsSEN6Y2dwdTBxa1ZMUT09`
- room-id: `d173a1976db46927cf7a9d7fa900870a003e6834fdd7fee7`
- task-id: `51683596509`
- template-id: `18b46f75236c558cc63be07419f8d65f`

## 白名单说明

- chat.* 与 ding.* 当前归入白名单阻塞（本报告未纳入修复目标）。
- calendar event list-mine 已继续探测，不按白名单处理。

| 步骤 | exit | 语义 | 摘要 |
|---|---:|---|---|
| `fix.aitable.base.get` | `0` | `pass` | {   "data": {     "baseId": "1zknDm0WRz0NZeXQu2zY5PLaWBQEx5rG",     "baseName": "non-chat-ding-base",     "dashboards": [],     "tables": [       {         "tableId": "qfszJ1P",         "tableName": "non-chat-ding-table" |
| `fix.aitable.base.update` | `0` | `pass` | {   "data": {     "baseId": "1zknDm0WRz0NZeXQu2zY5PLaWBQEx5rG",     "baseName": "non-chat-ding-base-upd",     "updatedAt": "2026-03-27T15:59:13+08:00"   },   "error": {},   "meta": {},   "status": "success",   "summary": |
| `fix.aitable.table.get` | `0` | `pass` | {   "data": {     "tables": [       {         "fields": [           {             "fieldId": "XeI8OyF",             "fieldName": "标题",             "type": "primaryDoc"           },           {             "fieldId": "FoE |
| `fix.aitable.record.create` | `0` | `pass` | {   "data": {     "newRecordIds": [       "0j1cfP7VjU"     ]   },   "error": {},   "meta": {},   "status": "success",   "summary": "Successfully created 1 record(s) in table qfszJ1P of base 1zknDm0WRz0NZeXQu2zY5PLaWBQEx5 |
| `fix.aitable.record.query` | `0` | `pass` | {   "data": {     "nextCursor": "0j1cfP7VjU",     "records": [       {         "cells": {           "FoE2ouB": "ok"         },         "recordId": "0j1cfP7VjU"       }     ]   },   "error": {},   "meta": {},   "status":  |
| `fix.aitable.record.delete` | `0` | `pass` | {   "data": {     "deletedCount": 0   },   "error": {},   "meta": {},   "status": "success",   "summary": "Deleted 0 record(s) from table qfszJ1P. Skipped already missing recordIds (1): dummy" }  |
| `fix.attendance.summary` | `0` | `fail_business` | {   "code": "C0002",   "message": "统计类型错误",   "result": {},   "success": false,   "trace_id": "0bab027317745983602097877e0828" }  |
| `fix.calendar.list-mine` | `0` | `fail_business` | {   "code": "PARAM_ERROR",   "error": "参数错误",   "message": "参数错误。请检查输入参数是否正确，参考工具说明调整参数后重试",   "retryable": false,   "success": false,   "technical_detail": "Tool metadata API error: PARAM_ERROR - 未找到指定工具",   "trace_id": |
| `fix.calendar.event.update` | `0` | `pass` | {   "arguments": [],   "errorCode": null,   "errorMsg": null,   "result": {     "attendees": [       {         "displayName": "天雷",         "optional": false,         "responseStatus": "accepted",         "self": true    |
| `fix.calendar.participant.delete` | `0` | `fail_business` | {   "arguments": [],   "errorCode": "300000",   "errorMsg": "code: 300000, developerMessage: can not remove organizer.",   "result": {},   "success": false,   "trace_id": "2106d98117745983619102665e0a69" }  |
| `fix.calendar.room.add` | `0` | `fail_business` | {   "arguments": [],   "errorCode": "321001",   "errorMsg": "code: 321001, developerMessage: [{\"roomName\":\"钉钉展厅-电视\",\"text\":\"根据管理员设置，同一时间段内只允许预定一个会议室，请移除多余的会议室\"}]",   "result": {},   "success": false,   "trace_id" |
| `fix.todo.task.get` | `0` | `pass` | {   "arguments": [],   "errorCode": null,   "errorMsg": null,   "result": {     "todoDetailModel": {       "activities": [         {           "action": "task.create",           "activityId": "12796095969",           "cr |
| `fix.report.create` | `0` | `fail_business` | {   "code": "PARAM_ERROR",   "error": "参数错误",   "message": "参数错误。请检查输入参数是否正确，参考工具说明调整参数后重试",   "retryable": false,   "success": false,   "technical_detail": "调用失败，原因: 调用失败，原因: 调用远程服务业务异常, 出参校验不通过。root.success当前值：false， 预 |
| `fix.report.list` | `0` | `pass` | {   "errcode": 0,   "errorMsg": "ok",   "result": {     "hasMore": false,     "nextCursor": 0,     "report_list": [],     "size": 10   },   "success": true }  |
