# CLI 第四轮专项报告（calendar + aitable）

| 指标 | 值 |
|---|---|
| 执行时间 | 2026-03-27 15:18:15 |
| 命令总数 | `18` |
| pass | `13` |
| fail_exit | `0` |
| fail_business | `5` |

## 依赖链上下文

- `event-id`: `NEhic05DdkczVXc0TDhUc3hUaDZtQT09`
- `base-id`: `wva2dxOW4kO6dpRacBYe6XA78bkz3BRL`
- `table-id`: `h6pp9yr`

## 关键发现

- `calendar event create/update` 在 RFC3339 格式（含时区）下可成功。
- `participant/room` 需严格使用 `--event` 参数（不是 `--id`）。
- `aitable field.create` 对 `fields` JSON 字段名敏感（`fieldName` 必填）。

## 执行明细

| 步骤 | exit | 语义 | 摘要 |
|---|---:|---|---|
| `auth.login` | `0` | `pass` |  [OK] Token 有效，无需重新登录 企业 ID:          ding8196cd9a2b2405da24f2f5cc6abecb85 有效期:            30 天后 Token 将自动刷新，无需重复登录  |
| `calendar.event.create` | `0` | `pass` | {   "arguments": [],   "result": {     "attendees": [       {         "displayName": "天雷",         "optional": false,         "responseStatus": "accepted",         "self": true     |
| `calendar.event.get` | `0` | `pass` | {   "arguments": [],   "result": {     "attendees": [       {         "displayName": "天雷",         "optional": false,         "responseStatus": "accepted",         "self": true     |
| `calendar.participant.list` | `0` | `pass` | {   "arguments": [],   "result": {     "attendees": [       {         "displayName": "天雷",         "optional": false,         "responseStatus": "accepted",         "self": true     |
| `calendar.participant.add` | `0` | `fail_business` | {   "arguments": [],   "errorCode": "300000",   "errorMsg": "code: 300000, developerMessage: user id is required",   "success": false,   "trace_id": "2127d89817745958829492933e09fc |
| `calendar.participant.delete` | `0` | `fail_business` | {   "arguments": [],   "errorCode": "300000",   "errorMsg": "code: 300000, developerMessage: user id is required",   "result": {},   "success": false,   "trace_id": "2104a64c177459 |
| `calendar.room.add` | `0` | `fail_business` | {   "arguments": [],   "errorCode": "300000",   "errorMsg": "code: 300000, developerMessage: roomId invalid",   "result": {},   "success": false,   "trace_id": "0bb7c36217745958843 |
| `calendar.room.delete` | `0` | `fail_business` | {   "arguments": [],   "errorCode": "300000",   "errorMsg": "code: 300000, developerMessage: roomId invalid",   "result": {},   "success": false,   "trace_id": "0b5deb3217745958849 |
| `calendar.event.update` | `0` | `pass` | {   "arguments": [],   "errorCode": null,   "errorMsg": null,   "result": {     "attendees": [       {         "displayName": "天雷",         "optional": false,         "responseStat |
| `calendar.event.delete` | `0` | `pass` | {   "arguments": [],   "errorCode": null,   "errorMsg": null,   "result": {},   "success": true }  |
| `calendar.busy.search` | `0` | `pass` | {   "result": [     {       "scheduleItems": []     }   ],   "success": true }  |
| `aitable.base.create` | `0` | `pass` | {   "data": {     "baseId": "wva2dxOW4kO6dpRacBYe6XA78bkz3BRL",     "baseName": "v4-base-20260327-151758"   },   "error": {},   "meta": {},   "status": "success",   "summary": "Cre |
| `aitable.base.get` | `0` | `pass` | {   "data": {     "baseId": "wva2dxOW4kO6dpRacBYe6XA78bkz3BRL",     "baseName": "v4-base-20260327-151758",     "dashboards": [],     "tables": []   },   "error": {},   "meta": {},  |
| `aitable.table.create` | `0` | `pass` | {   "data": {     "baseId": "wva2dxOW4kO6dpRacBYe6XA78bkz3BRL",     "tableId": "h6pp9yr",     "tableName": "v4-table-20260327-151758"   },   "error": {},   "meta": {},   "status":  |
| `aitable.table.get` | `0` | `pass` | {   "data": {     "tables": [       {         "fields": [           {             "fieldId": "agH1bGR",             "fieldName": "标题",             "type": "primaryDoc"           }  |
| `aitable.field.create` | `0` | `pass` | {   "data": {     "failedCount": 0,     "results": [       {         "fieldId": "UpAI9fd",         "fieldName": "自动化字段A",         "success": true       }     ],     "successCount": |
| `aitable.record.create` | `0` | `fail_business` | {   "data": {},   "error": {     "code": "INVALID_CELLS",     "message": "records[0].cells must be an object mapping field ids or names to values",     "retryable": false,     "typ |
| `aitable.record.query` | `0` | `pass` | {   "data": {},   "error": {},   "meta": {},   "status": "success",   "summary": "Successfully queried 0 record(s) from table h6pp9yr in base wva2dxOW4kO6dpRacBYe6XA78bkz3BRL" }  |
