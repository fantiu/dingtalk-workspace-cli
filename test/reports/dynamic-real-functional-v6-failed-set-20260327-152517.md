# CLI 第六轮失败集修复回归报告

| 指标 | 值 |
|---|---|
| 时间 | 2026-03-27 15:25:41 |
| 样本数 | `33` |
| pass | `23` |
| fail_exit | `0` |
| fail_business | `10` |

- `uid`: `061978`
- `event-id`: `d1RYWVpPZ1MwSVIxWnRRUnNFbWtJUT09`
- `room-id`: `bdff146f0172e49511c91904cf69b186003e6834fdd7fee7`
- `base-id`: `Exel2BLV5gOAdXr1uDpB0mrB8gk9rpMq`
- `table-id`: `i6sTftc`
- `report-id`: `(无)`
- `task-id`: `(无)`

| 步骤 | exit | 语义 | 摘要 |
|---|---:|---|---|
| `auth.status` | `0` | `pass` | 状态:             已登录 ✅ 有效期:            2026-03-27T09:17:05Z  |
| `contact.user.get-self` | `0` | `pass` | {   "result": [     {       "isAdmin": true,       "orgEmployeeModel": {         "corpId": "ding8196cd9a2b2405da24f2f5cc6abecb85",         "depts": [           {             "deptId": 1037045267,             "deptName":  |
| `calendar.event.create(seed)` | `0` | `pass` | {   "arguments": [],   "result": {     "attendees": [       {         "displayName": "天雷",         "optional": false,         "responseStatus": "accepted",         "self": true       }     ],     "created": 1774596318405 |
| `calendar.room.search(seed-96)` | `0` | `pass` | {   "arguments": [],   "result": {     "result": [       {         "capacity": null,         "fullGroupPath": "阿里中心-杭州未科",         "labels": [],         "roomId": "bdff146f0172e49511c91904cf69b186003e6834fdd7fee7",       |
| `aitable.base.create(seed)` | `0` | `pass` | {   "data": {     "baseId": "Exel2BLV5gOAdXr1uDpB0mrB8gk9rpMq",     "baseName": "v6-base-20260327-152517"   },   "error": {},   "meta": {},   "status": "success",   "summary": "Created base 'v6-base-20260327-152517' (id= |
| `aitable.table.create(seed)` | `0` | `pass` | {   "data": {     "baseId": "Exel2BLV5gOAdXr1uDpB0mrB8gk9rpMq",     "tableId": "i6sTftc",     "tableName": "v6-table-20260327-152517"   },   "error": {},   "meta": {},   "status": "success",   "summary": "create_table su |
| `aitable.field.create(seed)` | `0` | `pass` | {   "data": {     "failedCount": 0,     "results": [       {         "fieldId": "zKUcAsV",         "fieldName": "自动化字段V6",         "success": true       }     ],     "successCount": 1   },   "error": {},   "meta": {},    |
| `report.template.list(seed)` | `0` | `pass` | {   "items": [     {       "last_modified_time": 1697702695000,       "report_template_id": "18b46f75236c558cc63be07419f8d65f",       "report_template_logo_url": "https://landray.dingtalkapps.com/alid/app/report/images/i |
| `attendance.record.get` | `0` | `pass` | {   "code": "0",   "message": "success",   "result": {     "approveList": [],     "group": {       "name": "创新项目-QM",       "type": "FIXED"     },     "isHasSchedule": false,     "isRest": false,     "isUnSigned": false, |
| `attendance.summary` | `0` | `fail_business` | {   "code": "C0002",   "message": "统计类型错误",   "result": {},   "success": false,   "trace_id": "2132f5ca17745963243085228e0a19" }  |
| `calendar.event.get` | `0` | `pass` | {   "arguments": [],   "result": {     "attendees": [       {         "displayName": "天雷",         "optional": false,         "responseStatus": "accepted",         "self": true       }     ],     "categories": [],     "c |
| `calendar.event.update` | `0` | `pass` | {   "arguments": [],   "errorCode": null,   "errorMsg": null,   "result": {     "attendees": [       {         "displayName": "天雷",         "optional": false,         "responseStatus": "accepted",         "self": true    |
| `calendar.event.delete` | `0` | `pass` | {   "arguments": [],   "errorCode": null,   "errorMsg": null,   "result": {},   "success": true }  |
| `calendar.participant.add` | `0` | `pass` | {   "arguments": [],   "errorCode": null,   "errorMsg": null,   "success": true }  |
| `calendar.participant.delete` | `0` | `fail_business` | {   "arguments": [],   "errorCode": "300000",   "errorMsg": "code: 300000, developerMessage: can not remove organizer.",   "result": {},   "success": false,   "trace_id": "0bab027317745963274351337e0761" }  |
| `calendar.participant.list` | `0` | `pass` | {   "arguments": [],   "result": {     "attendees": [       {         "displayName": "天雷",         "optional": false,         "responseStatus": "accepted",         "self": true       }     ]   },   "success": true }  |
| `calendar.room.add` | `0` | `fail_business` | {   "arguments": [],   "errorCode": "300000",   "errorMsg": "code: 300000, developerMessage: Cannot patch event that is not in 'confirmed' status.",   "result": {},   "success": false,   "trace_id": "2132f5ca177459632895 |
| `calendar.room.delete` | `0` | `pass` | {   "arguments": [],   "errorCode": null,   "errorMsg": null,   "result": {},   "success": true }  |
| `aitable.base.get` | `0` | `pass` | {   "data": {     "baseId": "Exel2BLV5gOAdXr1uDpB0mrB8gk9rpMq",     "baseName": "v6-base-20260327-152517",     "dashboards": [],     "tables": [       {         "tableId": "i6sTftc",         "tableName": "v6-table-202603 |
| `aitable.base.update` | `0` | `pass` | {   "data": {     "baseId": "Exel2BLV5gOAdXr1uDpB0mrB8gk9rpMq",     "baseName": "v6-base-upd-20260327-152517",     "updatedAt": "2026-03-27T15:25:31+08:00"   },   "error": {},   "meta": {},   "status": "success",   "summ |
| `aitable.table.get` | `0` | `pass` | {   "data": {     "tables": [       {         "fields": [           {             "fieldId": "UmGsxF5",             "fieldName": "标题",             "type": "primaryDoc"           },           {             "fieldId": "zKU |
| `aitable.table.update` | `0` | `pass` | {   "data": {     "baseId": "Exel2BLV5gOAdXr1uDpB0mrB8gk9rpMq",     "tableId": "i6sTftc",     "tableName": "v6-table-upd-20260327-152517",     "updatedAt": "2026-03-27T15:25:33+08:00"   },   "error": {},   "meta": {},    |
| `aitable.record.create` | `0` | `pass` | {   "data": {     "newRecordIds": [       "ZkEj7TNe5G"     ]   },   "error": {},   "meta": {},   "status": "success",   "summary": "Successfully created 1 record(s) in table i6sTftc of base Exel2BLV5gOAdXr1uDpB0mrB8gk9rp |
| `aitable.record.query` | `0` | `pass` | {   "data": {     "nextCursor": "ZkEj7TNe5G",     "records": [       {         "cells": {           "zKUcAsV": "ok"         },         "recordId": "ZkEj7TNe5G"       }     ]   },   "error": {},   "meta": {},   "status":  |
| `chat.group.create` | `0` | `fail_business` | {   "code": "PARAM_ERROR",   "error": "参数错误",   "message": "参数错误。请检查输入参数是否正确，参考工具说明调整参数后重试",   "retryable": false,   "success": false,   "technical_detail": "Tool metadata API error: PARAM_ERROR - 未找到指定工具",   "trace_id": |
| `chat.message.send` | `0` | `fail_business` | {   "code": "PARAM_ERROR",   "error": "参数错误",   "message": "参数错误。请检查输入参数是否正确，参考工具说明调整参数后重试",   "retryable": false,   "success": false,   "technical_detail": "Tool metadata API error: PARAM_ERROR - 未找到指定工具",   "trace_id": |
| `chat.message.list` | `0` | `fail_business` | {   "code": "PARAM_ERROR",   "error": "参数错误",   "message": "参数错误。请检查输入参数是否正确，参考工具说明调整参数后重试",   "retryable": false,   "success": false,   "technical_detail": "Tool metadata API error: PARAM_ERROR - 未找到指定工具",   "trace_id": |
| `report.template.detail` | `0` | `pass` | {   "errcode": 0,   "errorMsg": "ok",   "result": {     "allowAddReceivers": 0,     "lastModifiedTime": 1650879711000,     "report_template_fields": [       {         "field_name": "今日完成工作",         "field_sort": 0,      |
| `report.create` | `0` | `fail_business` | {   "code": "PARAM_ERROR",   "error": "参数错误",   "message": "参数错误。请检查输入参数是否正确，参考工具说明调整参数后重试",   "retryable": false,   "success": false,   "technical_detail": "调用失败，原因: 调用失败，原因: 调用远程服务业务异常, 出参校验不通过。root.success当前值：false， 预 |
| `report.list` | `0` | `pass` | {   "errcode": 0,   "errorMsg": "ok",   "result": {     "hasMore": false,     "nextCursor": 0,     "report_list": [],     "size": 10   },   "success": true }  |
| `todo.task.create` | `0` | `fail_business` | {   "arguments": [],   "errorCode": "500",   "errorMsg": "TodoServiceImpl#createTask: Unrecognized priority type: 1",   "success": false,   "trace_id": "0bab027317745963398617211e07a4" }  |
| `ding.message.send` | `0` | `fail_business` | {   "dingOpenErrcode": 400009,   "errorMsg": "robotCode is in not valid or not in the org",   "result": {},   "success": false,   "trace_id": "2132f5ca17745963404804743e0a7c" }  |
| `ding.message.recall` | `0` | `fail_business` | {   "dingOpenErrcode": 500100,   "errorMsg": "system error",   "result": {},   "success": false,   "trace_id": "2106d98117745963413161346e0b34" }  |
