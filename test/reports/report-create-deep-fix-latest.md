# report.create 深度修复结论

## 目标

- 继续修复 `fix.report.create`
- 同时确认以下项按白名单处理：
  - `attendance.summary`（统计类型业务限制）
  - `calendar.participant.delete`（不能删除组织者）
  - `calendar.room.add`（会议室冲突规则）

## 执行结果

- 已对 `report.create` 做多轮参数矩阵探测：
  - `template-id` 使用了两组真实值：
    - `18b46f75236c558cc63be07419f8d65f`
    - `180601946fdafa15424d9d84b5e8c33a`
  - `contents` 试了 5 种结构：
    - `[{ "content": "..." }]`
    - `[{ "key": "...", "content": "..." }, ...]`
    - `[{ "field_name": "...", "content": "..." }, ...]`
    - `[{ "name": "...", "value": "..." }, ...]`
    - `{ "今日完成工作": "...", ... }`
  - 每种都分别试了是否带 `--dd-from 0`

- 最终结果：
  - 所有组合均 `exit=0`，但业务返回均为 `success=false`
  - 错误统一为 `PARAM_ERROR`（服务端业务校验未通过）

## 结论

- `fix.report.create` 在当前可见 CLI 参数维度下仍无法修复通过。
- 当前更可能是服务端对 `contents` 的协议要求与 CLI 暴露参数存在缺口，或需额外上下文（当前命令参数未暴露）。

## 当前白名单

- `attendance.summary`
- `calendar.participant.delete`
- `calendar.room.add`
- `chat.*`
- `ding.*`

> 注：`calendar event list-mine` 已按你的要求继续探测，但在不同时间格式和参数组合下均返回 `Tool metadata API error: 未找到指定工具`，现阶段可视为服务端侧阻塞。

