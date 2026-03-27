# CLI 第五轮清障报告（5条失败定向）

| 指标 | 值 |
|---|---|
| 时间 | 2026-03-27 15:20:17 |
| 命令数 | `17` |
| pass | `8` |
| fail_exit | `0` |
| fail_business | `9` |

- `event-id`: `UnlJR0l0YXZNMTVqSU9JNTljVkEzQT09`
- `base-id`: `7dx2rn0Jbakrq9O5t6NY541PVMGjLRb3`
- `table-id`: `i7zsObk`

| 步骤 | exit | 语义 | 摘要 |
|---|---:|---|---|
| `calendar.event.create` | `0` | `pass` | {   "arguments": [],   "result": {     "attendees": [       {         "displayName": "天雷",         "optional": false,         "responseStatus": "accepted",         "self": true       }     ],     "created": 1774596001303 |
| `calendar.participant.add(plain)` | `0` | `pass` | {   "arguments": [],   "errorCode": null,   "errorMsg": null,   "success": true }  |
| `calendar.participant.delete(plain)` | `0` | `fail_business` | {   "arguments": [],   "errorCode": "300000",   "errorMsg": "code: 300000, developerMessage: can not remove organizer.",   "result": {},   "success": false,   "trace_id": "0bb7c36217745960023866596e04c8" }  |
| `calendar.participant.add(csv)` | `0` | `pass` | {   "arguments": [],   "errorCode": null,   "errorMsg": null,   "success": true }  |
| `calendar.participant.delete(csv)` | `0` | `fail_business` | {   "arguments": [],   "errorCode": "300000",   "errorMsg": "code: 300000, developerMessage: can not remove organizer.",   "result": {},   "success": false,   "trace_id": "2127d89817745960035161714e0aa2" }  |
| `calendar.participant.add(jsonarr)` | `0` | `pass` | {   "arguments": [],   "errorCode": null,   "errorMsg": null,   "success": true }  |
| `calendar.participant.delete(jsonarr)` | `0` | `fail_business` | {   "arguments": [],   "errorCode": "300000",   "errorMsg": "code: 300000, developerMessage: can not remove organizer.",   "result": {},   "success": false,   "trace_id": "2104a64c17745960050053458e0acc" }  |
| `calendar.room.add(1001)` | `0` | `fail_business` | {   "arguments": [],   "errorCode": "300000",   "errorMsg": "code: 300000, developerMessage: roomId invalid",   "result": {},   "success": false,   "trace_id": "0bb7c36217745960058146858e054a" }  |
| `calendar.room.delete(1001)` | `0` | `fail_business` | {   "arguments": [],   "errorCode": "300000",   "errorMsg": "code: 300000, developerMessage: roomId invalid",   "result": {},   "success": false,   "trace_id": "2104a64c17745960064643041e0a69" }  |
| `calendar.room.add(1)` | `0` | `fail_business` | {   "arguments": [],   "errorCode": "300000",   "errorMsg": "code: 300000, developerMessage: roomId invalid",   "result": {},   "success": false,   "trace_id": "0bb7c36217745960071086603e04a5" }  |
| `calendar.room.delete(1)` | `0` | `fail_business` | {   "arguments": [],   "errorCode": "300000",   "errorMsg": "code: 300000, developerMessage: roomId invalid",   "result": {},   "success": false,   "trace_id": "2104a64c17745960080577877e0a8b" }  |
| `calendar.room.add(A1001)` | `0` | `fail_business` | {   "arguments": [],   "errorCode": "300000",   "errorMsg": "code: 300000, developerMessage: roomId invalid",   "result": {},   "success": false,   "trace_id": "2127d89817745960087162927e09fc" }  |
| `calendar.room.delete(A1001)` | `0` | `fail_business` | {   "arguments": [],   "errorCode": "300000",   "errorMsg": "code: 300000, developerMessage: roomId invalid",   "result": {},   "success": false,   "trace_id": "0b5deb3217745960096788390e0c07" }  |
| `aitable.base.create` | `0` | `pass` | {   "data": {     "baseId": "7dx2rn0Jbakrq9O5t6NY541PVMGjLRb3",     "baseName": "v5-base-20260327-152000"   },   "error": {},   "meta": {},   "status": "success",   "summary": "Created base 'v5-base-20260327-152000' (id= |
| `aitable.table.create` | `0` | `pass` | {   "data": {     "baseId": "7dx2rn0Jbakrq9O5t6NY541PVMGjLRb3",     "tableId": "i7zsObk",     "tableName": "v5-table-20260327-152000"   },   "error": {},   "meta": {},   "status": "success",   "summary": "create_table su |
| `aitable.field.create` | `0` | `pass` | {   "data": {     "failedCount": 0,     "results": [       {         "fieldId": "ZfncUO3",         "fieldName": "自动化字段B",         "success": true       }     ],     "successCount": 1   },   "error": {},   "meta": {},   " |
| `aitable.record.create(cells)` | `0` | `pass` | {   "data": {     "newRecordIds": [       "Ji0akBZIwI"     ]   },   "error": {},   "meta": {},   "status": "success",   "summary": "Successfully created 1 record(s) in table i7zsObk of base 7dx2rn0Jbakrq9O5t6NY541PVMGjLR |
