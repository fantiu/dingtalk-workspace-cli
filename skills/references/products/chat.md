# 机器人消息 (chat) 命令参考

## 命令总览

| 子命令 | 用途 |
|-------|------|
| `bot search` | 搜索我的机器人 |
| `group members add-bot` | 添加机器人到群 |
| `message send-by-bot` | 机器人发消息（群聊或批量单聊） |
| `message recall-by-bot` | 机器人撤回消息（群聊或批量单聊） |
| `message send-by-webhook` | 自定义机器人 Webhook 发消息 |

---

## bot search — 搜索我的机器人

```
Usage:
  dws chat bot search [flags]
Example:
  dws chat bot search --name "考勤" --format json
Flags:
      --name string   机器人名称（模糊搜索）
      --page string   页码（默认 1）
      --size string   每页数量（默认 10）
```

---

## group members add-bot — 添加机器人到群

```
Usage:
  dws chat group members add-bot [flags]
Example:
  dws chat group members add-bot --id <openConversationId> --robot-code <robotCode> --format json
Flags:
      --id string           群会话 ID (必填)
      --robot-code string   机器人 code (必填)
```

---

## message send-by-bot — 机器人发消息

支持两种模式：群聊发送 和 批量单聊发送，通过 `--group` 和 `--users` 互斥区分。

### 群聊发送
```
Usage:
  dws chat message send-by-bot [flags]
Example:
  dws chat message send-by-bot --robot-code <code> --group <openConversationId> \
    --title "日报提醒" --text "请提交今日日报" --format json
Flags:
      --robot-code string   机器人 code (必填)
      --group string        群会话 ID (必填，与 --users 互斥)
      --title string        消息标题 (必填)
      --text string         消息内容，支持 Markdown (必填)
```

### 批量单聊发送
```
Usage:
  dws chat message send-by-bot [flags]
Example:
  dws chat message send-by-bot --robot-code <code> --users "user1,user2" \
    --title "通知" --text "会议已取消" --format json
Flags:
      --robot-code string   机器人 code (必填)
      --users string        用户 ID 列表，逗号分隔 (必填，与 --group 互斥)
      --title string        消息标题 (必填)
      --text string         消息内容，支持 Markdown (必填)
```

> ⚠️ `--group` 和 `--users` 互斥：群聊用 `--group`，单聊用 `--users`，不能同时传。

---

## message recall-by-bot — 机器人撤回消息

支持两种模式：群聊撤回 和 批量单聊撤回。

### 群聊撤回
```
Usage:
  dws chat message recall-by-bot [flags]
Example:
  dws chat message recall-by-bot --robot-code <code> --group <openConversationId> \
    --keys "key1,key2" --format json
Flags:
      --robot-code string   机器人 code (必填)
      --group string        群会话 ID (必填，与批量单聊互斥)
      --keys string         消息 key 列表，逗号分隔 (必填)
```

### 批量单聊撤回
```
Usage:
  dws chat message recall-by-bot [flags]
Example:
  dws chat message recall-by-bot --robot-code <code> --keys "key1,key2" --format json
Flags:
      --robot-code string   机器人 code (必填)
      --keys string         消息 key 列表，逗号分隔 (必填)
```

> ⚠️ 消息 key 从 `send-by-bot` 返回结果中提取。

---

## message send-by-webhook — 自定义机器人 Webhook 发消息

```
Usage:
  dws chat message send-by-webhook [flags]
Example:
  dws chat message send-by-webhook --token <robotToken> \
    --title "告警" --text "CPU 使用率超过 90%" --format json
Flags:
      --token string        自定义机器人 Webhook Token (必填)
      --title string        消息标题 (必填)
      --text string         消息内容 (必填)
      --at-all              @所有人
      --at-mobiles string   @指定手机号列表，逗号分隔
      --at-users string     @指定用户 ID 列表，逗号分隔
```

---

## 意图判断

- 用户说"让机器人在群里发通知" → `message send-by-bot --group`
- 用户说"机器人给张三发消息" → 先 `contact user search` 获取 userId，再 `message send-by-bot --users`
- 用户说"通过 Webhook 发告警" / 用户有 Webhook Token → `message send-by-webhook`
- 用户说"撤回机器人消息" → `message recall-by-bot`
- 用户说"查一下我的机器人" → `bot search`
- 用户说"把机器人加到群里" → `group members add-bot`

**关键区分**: `send-by-bot`(企业内部机器人，需 robotCode) vs `send-by-webhook`(自定义机器人 Webhook，需 token)

## 核心工作流

```bash
# ── 工作流: 机器人群发消息 ──

# 1. 搜索可用机器人
dws chat bot search --format json

# 2. 发送群消息
dws chat message send-by-bot --robot-code <code> --group <groupId> \
  --title "通知" --text "内容" --format json
```

```bash
# ── 工作流: Webhook 告警 ──

# 直接通过 Webhook Token 发送
dws chat message send-by-webhook --token <token> \
  --title "告警" --text "服务异常" --at-all --format json
```

## 上下文传递表

| 操作 | 从返回中提取 | 用于 |
|------|-------------|------|
| `bot search` | robotCode | `send-by-bot` / `recall-by-bot` / `add-bot` |
| `message send-by-bot` | processQueryKey | `recall-by-bot --keys` |
