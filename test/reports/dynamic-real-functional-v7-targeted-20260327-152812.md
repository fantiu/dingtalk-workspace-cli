# CLI 第七轮定向修复报告

| 指标 | 值 |
|---|---|
| 时间 | 2026-03-27 15:28:17 |
| 样本数 | `23` |
| pass | `0` |
| fail_exit | `23` |
| fail_business | `0` |

- `uid`: `061978`
- `todo priority` 可用值: `(未探测到)`
- `task-id`: `(无)`
- `report template`: `日报`
- `template-id`: `1`
- `calendar event-id`: `(无)`
- `calendar room-id`: `(无)`

| 步骤 | exit | 语义 | 摘要 |
|---|---:|---|---|
| `contact.user.get-self` | `1` | `fail_exit` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `todo.task.create(priority=1)` | `1` | `fail_exit` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `todo.task.create(priority=2)` | `1` | `fail_exit` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `todo.task.create(priority=3)` | `1` | `fail_exit` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `todo.task.create(priority=4)` | `1` | `fail_exit` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `todo.task.create(priority=P0)` | `1` | `fail_exit` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `todo.task.create(priority=P1)` | `1` | `fail_exit` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `todo.task.create(priority=P2)` | `1` | `fail_exit` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `todo.task.create(priority=P3)` | `1` | `fail_exit` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `todo.task.create(priority=HIGH)` | `1` | `fail_exit` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `todo.task.create(priority=MEDIUM)` | `1` | `fail_exit` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `todo.task.create(priority=LOW)` | `1` | `fail_exit` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `todo.task.create(priority=URGENT)` | `1` | `fail_exit` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `todo.task.create(priority=NORMAL)` | `1` | `fail_exit` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `report.template.list` | `1` | `fail_exit` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `report.template.detail` | `1` | `fail_exit` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `report.create(with_template_fields)` | `1` | `fail_exit` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `calendar.event.create(confirmed_seed)` | `1` | `fail_exit` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `calendar.room.search(96)` | `1` | `fail_exit` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `calendar.room.search(41)` | `1` | `fail_exit` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `calendar.room.search(74)` | `1` | `fail_exit` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `ding.message.send(robot_placeholder)` | `1` | `fail_exit` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `ding.message.recall(robot_placeholder)` | `1` | `fail_exit` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
