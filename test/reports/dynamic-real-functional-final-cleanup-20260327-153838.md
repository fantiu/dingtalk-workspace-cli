# CLI 动态功能测试最终收口报告

| 指标 | 值 |
|---|---|
| 时间 | 2026-03-27 15:39:06 |
| 可修复项重跑数 | `18` |
| 可修复项 pass | `10` |
| 可修复项 fail_exit | `2` |
| 可修复项 fail_business | `6` |

## 关键上下文

- uid: `061978`
- event-id: `UkFPcCtZNWtTbFozRG9lSDBsOUFPdz09`
- room-id: `bdff146f0172e49511c91904cf69b186003e6834fdd7fee7`
- base-id: `9bN7RYPWdPxOBgrQubjeYd74WZd1wyK0`
- table-id: `z9JU482`
- task-id: `51470771039`
- template-id: `18b46f75236c558cc63be07419f8d65f`
- report-id: `(无)`

## 可修复项结果

| 步骤 | exit | 语义 | 摘要 |
|---|---:|---|---|
| `fix.aitable.base.get` | `0` | `pass` | {   "data": {     "baseId": "9bN7RYPWdPxOBgrQubjeYd74WZd1wyK0",     "baseName": "final-base-20260327-153838",     "dashboards": [],     "tables": [       {         "tableId": "z9JU482",         "tableName": "final-table- |
| `fix.aitable.base.update` | `0` | `pass` | {   "data": {     "baseId": "9bN7RYPWdPxOBgrQubjeYd74WZd1wyK0",     "baseName": "final-base-upd-20260327-153838",     "updatedAt": "2026-03-27T15:38:50+08:00"   },   "error": {},   "meta": {},   "status": "success",   "s |
| `fix.aitable.table.get` | `0` | `pass` | {   "data": {     "tables": [       {         "fields": [           {             "fieldId": "bIJb2LD",             "fieldName": "标题",             "type": "primaryDoc"           },           {             "fieldId": "h4D |
| `fix.aitable.table.delete` | `0` | `fail_business` | {   "data": {},   "error": {     "code": "InvalidRequest.Forbidden",     "message": "cannot delete the last sheet",     "retryable": false,     "type": "USER_ERROR"   },   "meta": {},   "status": "error",   "summary": "F |
| `fix.aitable.record.create` | `0` | `pass` | {   "data": {     "newRecordIds": [       "BO49HAmJEs"     ]   },   "error": {},   "meta": {},   "status": "success",   "summary": "Successfully created 1 record(s) in table z9JU482 of base 9bN7RYPWdPxOBgrQubjeYd74WZd1wy |
| `fix.aitable.record.query` | `0` | `pass` | {   "data": {     "nextCursor": "BO49HAmJEs",     "records": [       {         "cells": {           "h4DK3d5": "ok"         },         "recordId": "BO49HAmJEs"       }     ]   },   "error": {},   "meta": {},   "status":  |
| `fix.aitable.record.update` | `0` | `fail_business` | {   "data": {},   "error": {     "code": "InvalidRequest.ResourceNotFound",     "message": "fail to find the record 'dummy'",     "retryable": true,     "type": "SYSTEM_ERROR"   },   "meta": {},   "status": "error",   "s |
| `fix.aitable.field.get` | `0` | `pass` | {   "data": {     "fields": [       {         "fieldId": "bIJb2LD",         "fieldName": "标题",         "type": "primaryDoc"       },       {         "fieldId": "h4DK3d5",         "fieldName": "自动化字段FINAL",         "type" |
| `fix.attendance.summary` | `0` | `fail_business` | {   "code": "C0002",   "message": "统计类型错误",   "result": {},   "success": false,   "trace_id": "2127d89817745971408062765e0aa2" }  |
| `fix.calendar.event.update` | `0` | `pass` | {   "arguments": [],   "errorCode": null,   "errorMsg": null,   "result": {     "attendees": [       {         "displayName": "天雷",         "optional": false,         "responseStatus": "accepted",         "self": true    |
| `fix.calendar.participant.delete` | `0` | `fail_business` | {   "arguments": [],   "errorCode": "300000",   "errorMsg": "code: 300000, developerMessage: can not remove organizer.",   "result": {},   "success": false,   "trace_id": "0bb7c36217745971422665047e04a6" }  |
| `fix.calendar.room.add` | `0` | `pass` | {   "arguments": [],   "errorCode": null,   "errorMsg": null,   "result": {},   "success": true }  |
| `fix.chat.message.send` | `0` | `fail_business` | {   "code": "PARAM_ERROR",   "error": "参数错误",   "message": "参数错误。请检查输入参数是否正确，参考工具说明调整参数后重试",   "retryable": false,   "success": false,   "technical_detail": "Tool metadata API error: PARAM_ERROR - 未找到指定工具",   "trace_id": |
| `fix.chat.message.send-by-bot` | `5` | `fail_exit` | {   "error": {     "category": "internal",     "code": 5,     "message": "unknown flag: --users"   } }  |
| `fix.chat.message.recall-by-bot` | `5` | `fail_exit` | {   "error": {     "category": "internal",     "code": 5,     "message": "unknown flag: --id"   } }  |
| `fix.report.create` | `0` | `fail_business` | {   "code": "PARAM_ERROR",   "error": "参数错误",   "message": "参数错误。请检查输入参数是否正确，参考工具说明调整参数后重试",   "retryable": false,   "success": false,   "technical_detail": "调用失败，原因: 调用失败，原因: 调用远程服务业务异常, 出参校验不通过。root.success当前值：false， 预 |
| `fix.report.list` | `0` | `pass` | {   "errcode": 0,   "errorMsg": "ok",   "result": {     "hasMore": false,     "nextCursor": 0,     "report_list": [],     "size": 10   },   "success": true }  |
| `fix.todo.task.get` | `0` | `pass` | {   "arguments": [],   "errorCode": null,   "errorMsg": null,   "result": {     "todoDetailModel": {       "activities": [         {           "action": "task.create",           "activityId": "12796115088",           "cr |

## 不可控阻塞项

| 项目 | 原因 |
|---|---|
| `blocked.calendar.event.list-mine` | `server_tool_missing` |
| `blocked.chat.group.*` | `server_tool_missing` |
| `blocked.ding.message.*` | `needs_real_robot_code` |
