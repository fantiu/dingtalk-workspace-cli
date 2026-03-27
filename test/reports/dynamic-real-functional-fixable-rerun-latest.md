# 可修复项定向回归报告（latest）

| 指标 | 值 |
|---|---|
| 时间 | 2026-03-27 15:51:28 |
| 重跑项数 | `12` |
| pass | `6` |
| fail_exit | `0` |
| fail_business | `6` |

- uid: `061978`
- base-id: `QBnd5ExVEaQd4meXt06QbeQgVyeZqMmz`
- table-id: `MoptjEc`
- event-id: `SGlWMTJLdlAzVEp3RVhaSmhxc3hLZz09`
- room-id: `d173a1976db46927cf7a9d7fa900870a003e6834fdd7fee7`

| 步骤 | exit | 语义 | 摘要 |
|---|---:|---|---|
| `fix.aitable.base.get` | `0` | `pass` | {   "data": {     "baseId": "QBnd5ExVEaQd4meXt06QbeQgVyeZqMmz",     "baseName": "fixable-155104",     "dashboards": [],     "tables": [       {         "tableId": "MoptjEc",         "tableName": "fixable-table"       }   |
| `fix.aitable.base.update` | `0` | `pass` | {   "data": {     "baseId": "QBnd5ExVEaQd4meXt06QbeQgVyeZqMmz",     "baseName": "fixable-base-upd",     "updatedAt": "2026-03-27T15:51:15+08:00"   },   "error": {},   "meta": {},   "status": "success",   "summary": "Succ |
| `fix.aitable.table.get` | `0` | `pass` | {   "data": {     "tables": [       {         "fields": [           {             "fieldId": "iizMpn1",             "fieldName": "标题",             "type": "primaryDoc"           },           {             "fieldId": "912 |
| `fix.aitable.record.create` | `0` | `pass` | {   "data": {     "newRecordIds": [       "1W22tVCf9b"     ]   },   "error": {},   "meta": {},   "status": "success",   "summary": "Successfully created 1 record(s) in table MoptjEc of base QBnd5ExVEaQd4meXt06QbeQgVyeZqM |
| `fix.aitable.record.query` | `0` | `pass` | {   "data": {     "nextCursor": "1W22tVCf9b",     "records": [       {         "cells": {           "912NdwK": "ok"         },         "recordId": "1W22tVCf9b"       }     ]   },   "error": {},   "meta": {},   "status":  |
| `fix.attendance.summary` | `0` | `fail_business` | {   "code": "C0002",   "message": "统计类型错误",   "result": {},   "success": false,   "trace_id": "213ee25c17745978819184321e0acd" }  |
| `fix.calendar.event.update` | `0` | `pass` | {   "arguments": [],   "errorCode": null,   "errorMsg": null,   "result": {     "attendees": [       {         "displayName": "天雷",         "optional": false,         "responseStatus": "accepted",         "self": true    |
| `fix.calendar.participant.delete` | `0` | `fail_business` | {   "arguments": [],   "errorCode": "300000",   "errorMsg": "code: 300000, developerMessage: can not remove organizer.",   "result": {},   "success": false,   "trace_id": "213ee25c17745978840783439e0b54" }  |
| `fix.calendar.room.add` | `0` | `fail_business` | {   "arguments": [],   "errorCode": "321001",   "errorMsg": "code: 321001, developerMessage: [{\"roomName\":\"钉钉展厅-电视\",\"text\":\"根据管理员设置，同一时间段内只允许预定一个会议室，请移除多余的会议室\"}]",   "result": {},   "success": false,   "trace_id" |
| `fix.chat.message.send` | `0` | `fail_business` | {   "code": "PARAM_ERROR",   "error": "参数错误",   "message": "参数错误。请检查输入参数是否正确，参考工具说明调整参数后重试",   "retryable": false,   "success": false,   "technical_detail": "Tool metadata API error: PARAM_ERROR - 未找到指定工具",   "trace_id": |
| `fix.chat.message.send-by-bot` | `0` | `fail_business` | {   "code": "PARAM_ERROR",   "error": "参数错误",   "message": "参数错误。请检查输入参数是否正确，参考工具说明调整参数后重试",   "retryable": false,   "success": false,   "technical_detail": "调用失败，原因: 调用失败，原因: 调用远程服务业务异常, 出参校验不通过。DingOpenResult.success当前 |
| `fix.chat.message.recall-by-bot` | `0` | `fail_business` | {   "code": "PARAM_ERROR",   "error": "参数错误",   "message": "参数错误。请检查输入参数是否正确，参考工具说明调整参数后重试",   "retryable": false,   "success": false,   "technical_detail": "调用失败，原因: 调用失败，原因: 调用远程服务业务异常, 出参校验不通过。DingOpenResult.success当前 |
