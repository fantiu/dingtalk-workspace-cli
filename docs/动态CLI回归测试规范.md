# 动态 CLI 回归测试规范

本文档用于沉淀当前仓库中**基于动态服务发现**的 CLI 真实功能回归方法，目标是让后续回归可重复、可对比、可定位问题。

## 1. 适用范围

- 适用于 `dws --help` 动态发现出的全部服务与子命令。
- 覆盖真实调用（非 `go test`、非纯 `--dry-run`）。
- 默认输出为 JSON，便于自动判定与报告汇总。

## 2. 执行前准备

### 2.1 环境准备

- 在仓库根目录执行：

```bash
cd /Users/tianlei.qjb/Documents/my_python_project/cli
go build -o ./dws ./cmd
```

### 2.2 登录准备

- 每轮测试前必须先确认登录状态：

```bash
./dws auth status
```

- 若显示未登录，先执行：

```bash
./dws auth login --client-id <CLIENT_ID> --client-secret <CLIENT_SECRET>
```

> 注意：长时间回归过程中登录态可能失效。若出现大量 HTTP 400/统一网关拒绝，优先重新登录后再继续。

## 3. 标准执行口径

### 3.1 统一执行参数

- 所有命令默认附带：
  - `-f json`
  - `-y`
- 对写操作，优先使用实时采集到的真实依赖 ID（禁止仅用占位值）。

### 3.2 结果判定规则

- `pass`：命令退出码 `0`，且业务体未出现失败信号。
- `fail_exit`：命令退出码非 `0`。
- `fail_business`：退出码 `0` 但业务体 `success=false` 或错误体非空。

## 4. 前置依赖采集规范

以下依赖应在全量执行前统一采集并缓存到上下文变量中。

### 4.1 用户信息

```bash
./dws -f json -y contact user get-self
```

提取 `userId`，供 `attendance`、`todo`、`chat`、`report` 等命令复用。

### 4.2 calendar 依赖

1) 创建种子事件（RFC3339 时间格式，含时区）：

```bash
./dws -f json -y calendar event create --title "<title>" --start "2026-03-27T16:00:00+08:00" --end "2026-03-27T16:30:00+08:00"
```

提取 `eventId`。

2) 采集会议室：

```bash
./dws -f json -y calendar room list-groups
./dws -f json -y calendar room search --group-id <groupId> --start "<RFC3339>" --end "<RFC3339>" --available true
```

提取真实可用 `roomId`。

### 4.3 aitable 依赖

```bash
./dws -f json -y aitable base create --name "<base-name>"
./dws -f json -y aitable table create --base-id <baseId> --name "<table-name>"
./dws -f json -y aitable field create --base-id <baseId> --table-id <tableId> --fields '[{"fieldName":"自动化字段","type":"text"}]'
```

提取 `baseId`、`tableId`，并确保已存在至少一个可写字段。

### 4.4 report 依赖

```bash
./dws -f json -y report template list
./dws -f json -y report template detail --name "<template-name>"
```

提取 `template-id` 与模板字段定义，构造 `contents`。

#### report.create 修复要点（已验证可通过）

`report.create` 必须遵循“先模板、后内容”的构造顺序：

1) 先从 `report template list` 获取真实 `report_template_id`
2) 再从 `report template detail --name` 获取字段列表（例如 `今日完成工作`、`未完成工作`、`需协调工作`）
3) `--contents` 按字段逐项构造，且每项包含：
   - `content`
   - `sort`
   - `key`
   - `contentType`
   - `type`

已验证成功示例（可直接复用）：

```bash
./dws -f json -y report create \
  --template-id 180601946fdafa15424d9d84b5e8c33a \
  --contents '[{"content":"完成开发","sort":"0","key":"今日完成工作","contentType":"markdown","type":"1"},{"content":"无","sort":"1","key":"未完成工作","contentType":"markdown","type":"1"},{"content":"无","sort":"2","key":"需协调工作","contentType":"markdown","type":"1"}]' \
  --to-user-ids 061978 \
  --dd-from dws
```

成功返回示例：

```json
{"errorCode":0,"errorMessage":"ok","reportId":"<id>","success":true}
```

### 4.5 todo 依赖

优先级枚举建议优先使用 `LOW`（当前验证可通过）：

```bash
./dws -f json -y todo task create --title "<title>" --executors <userId> --due "2026-03-27 20:00:00" --priority LOW
```

提取 `taskId` 用于 `todo task get`。

## 5. 高风险参数模板（已验证）

### 5.1 calendar 时间参数

- `calendar event create/update/list`：优先使用 `RFC3339` 格式（含时区）。
- 示例：`2026-03-27T16:00:00+08:00`。

### 5.2 aitable record.create

`records` 必须是 `cells` 结构：

```json
[{"cells":{"自动化字段":"ok"}}]
```

### 5.3 report.create contents

建议按模板字段生成：

```json
[{"content":"完成开发","sort":"0","key":"今日完成工作","contentType":"markdown","type":"1"}]
```

### 5.4 todo.priority

- 建议优先：`LOW`
- 其他枚举值需在当前租户再次验证。

## 6. 失败分类与处置

### 6.1 可修复（测试数据/参数导致）

- 缺少依赖 ID（`eventId/baseId/tableId/roomId/taskId`）
- 时间格式不符合接口期望
- JSON 结构不符合字段协议

处置：补齐前置采集与参数映射后重跑。

### 6.2 业务规则型（预期内失败）

- 例如：
  - 删除 organizer 被拒绝
  - 删除最后一张 sheet 被拒绝
  - 事件状态不允许 patch

处置：在报告中标注为“业务规则限制”，不归类为脚本缺陷。

### 6.3 外部依赖型（需真实资源）

- 例如：
  - `ding` 需要真实有效 `robotCode`
  - 特定群/机器人资源不存在

处置：要求补充真实资源再复测。

### 6.4 服务端能力缺失

- 例如 `Tool metadata API error: PARAM_ERROR - 未找到指定工具`

处置：归档为“服务端能力阻塞”，等待平台侧修复。

## 7. 全量回归推荐流程

1. `auth status` 检查登录态。
2. 采集上下文依赖（用户、event/room、base/table、template、task）。
3. 动态遍历所有叶子命令并执行真实调用。
4. 对 `fail_exit/fail_business` 分类汇总。
5. 对可修复项二次重跑并生成收口报告。
6. 输出不可控阻塞清单（服务端缺失/真实资源缺失）。

## 8. 报告产物规范

每轮至少输出两份文件：

- 时间戳报告：便于历史比对。
- `latest` 报告：便于快速查看当前状态。

建议命名：

- 全量：`test/reports/dynamic-real-functional-full-rerun-<ts>.md`
- 收口：`test/reports/dynamic-real-functional-final-cleanup-<ts>.md`

## 9. 已知阻塞项（当前）

- `calendar.event.list-mine`：服务端工具缺失（metadata 不可用）。
- `chat.group.*` 部分命令：服务端工具缺失。
- `ding.message.*`：缺少组织内可用 `robotCode` 时无法通过。
- `attendance.summary`：当前租户下返回“统计类型错误”（业务规则/上下文限制）。
- `calendar.participant.delete`：删除组织者被业务规则拒绝。
- `calendar.room.add`：同时间段会议室冲突时会被业务规则拒绝。

