# CLI 动态发现真实功能测试报告（第三轮定向强化）

| 项目 | 值 |
|---|---|
| 时间 | 2026-03-27 15:11:22 |
| 登录步骤退出码 | `0` |
| 动态服务数 | `10` |
| 命令节点(`--help`) | `118` |
| `--help` 失败 | `0` |
| 叶子命令执行数 | `83` |
| 语义成功(pass) | `8` |
| 退出码失败 | `75` |
| 业务失败 | `0` |

## 本轮增强点

- 补齐位置参数：`completion zsh`、`chat message send <message>`。
- `calendar` 子命令优先注入 `--event`，并尝试注入真实 `--rooms`。
- `report create --contents` 改为 JSON 字符串。
- `todo task create -> task-id` 回填 `todo task get`。

## 关键上下文

- `base-id`: `PwkYGxZV3L9AdX7YSvpndwYeJAgozOKL`
- `table-id`: `A1Pl1DO`
- `event-id`: `TTJHTTlLUG9valZvSE5HQkQ5aGVOZz09`
- `room-id`: `(未获取)`
- `task-id`: `(未获取)`

## 失败明细

| 命令 | 分类 | 退出码 | 摘要 |
|---|---|---:|---|
| `dws aitable attachment upload` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws aitable base create` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws aitable base delete` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws aitable base get` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws aitable base list` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws aitable base search` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws aitable base update` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws aitable field create` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws aitable field delete` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws aitable field get` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws aitable field update` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws aitable record create` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws aitable record delete` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws aitable record query` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws aitable record update` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws aitable table create` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws aitable table delete` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws aitable table get` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws aitable table update` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws aitable template search` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws attendance record get` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws attendance rules` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws attendance shift list` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws attendance summary` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws calendar busy search` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws calendar event create` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws calendar event delete` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws calendar event get` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws calendar event list` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws calendar event list-mine` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws calendar event update` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws calendar participant add` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws calendar participant delete` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws calendar participant list` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws calendar room add` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws calendar room delete` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws calendar room list-groups` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws calendar room search` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws chat bot search` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws chat group create` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws chat group members add` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws chat group members add-bot` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws chat group members remove` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws chat group rename` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws chat message list` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws chat message recall-by-bot` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws chat message send` | `fail_exit` | `3` | {   "error": {     "category": "validation",     "code": 3,     "message": "--group and --user are mutually exclusive"   } }  |
| `dws chat message send-by-bot` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws chat message send-by-webhook` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws chat search` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws contact dept list-children` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws contact dept list-members` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws contact dept search` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws contact user get` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws contact user get-self` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws contact user search` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws contact user search-mobile` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws devdoc article search` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws ding message recall` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws ding message send` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws report create` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws report detail` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws report list` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws report sent` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws report stats` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws report template detail` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws report template list` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws todo task create` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws todo task delete` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws todo task done` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws todo task get` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws todo task list` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws todo task update` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws workbench app get` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
| `dws workbench app list` | `fail_exit` | `1` | {   "error": {     "actions": [       "Check authentication, permissions, and parameters, then retry"     ],     "category": "api",     "code": 1,     "hint": "Request was rejected by upstream service; check parameters,  |
