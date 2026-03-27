# 可修复项定向回归报告（latest）

| 指标 | 值 |
|---|---|
| 时间 | 2026-03-27 15:49:46 |
| 重跑项数 | `4` |
| pass | `0` |
| fail_exit | `4` |
| fail_business | `0` |

- uid: `061978`
- base-id: `(无)`
- table-id: `(无)`
- event-id: `(无)`
- room-id: `(无)`

| 步骤 | exit | 语义 | 摘要 |
|---|---:|---|---|
| `fix.attendance.summary` | `1` | `fail_exit` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `fix.chat.message.send` | `1` | `fail_exit` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `fix.chat.message.send-by-bot` | `5` | `fail_exit` | {   "error": {     "category": "internal",     "code": 5,     "message": "unknown flag: --user"   } }  |
| `fix.chat.message.recall-by-bot` | `5` | `fail_exit` | {   "error": {     "category": "internal",     "code": 5,     "message": "unknown flag: --open-ding-id"   } }  |
