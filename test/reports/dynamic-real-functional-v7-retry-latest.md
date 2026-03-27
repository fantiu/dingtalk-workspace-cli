# CLI 第七轮定向修复重跑报告

| 指标 | 值 |
|---|---|
| 时间 | 2026-03-27 15:29:25 |
| 样本数 | `11` |
| pass | `9` |
| fail_exit | `0` |
| fail_business | `2` |

- `uid`: `061978`
- `todo priority` 命中值: `LOW`
- `task-id`: `51469483424`
- `template`: `日报` / `18b46f75236c558cc63be07419f8d65f`
- `event-id`: `bVBwVWRCZDhSenhzdmwvcDNzYkQxQT09`
- `room-id`: `bdff146f0172e49511c91904cf69b186003e6834fdd7fee7`

| 步骤 | exit | 语义 | 摘要 |
|---|---:|---|---|
| `contact.user.get-self` | `0` | `pass` | {   "result": [     {       "isAdmin": true,       "orgEmployeeModel": {         "corpId": "ding8196cd9a2b2405da24f2f5cc6abecb85",         "depts": [           {             "deptId": 1037045267,             "deptName":  |
| `todo.task.create(priority=LOW)` | `0` | `pass` | {   "arguments": [],   "result": {     "subject": "v7-retry-20260327-152915-LOW",     "taskId": "51469483424"   },   "success": true }  |
| `todo.task.get` | `0` | `pass` | {   "arguments": [],   "errorCode": null,   "errorMsg": null,   "result": {     "todoDetailModel": {       "activities": [         {           "action": "task.create",           "activityId": "12839806740",           "cr |
| `report.template.list` | `0` | `pass` | {   "items": [     {       "last_modified_time": 1697702695000,       "report_template_id": "18b46f75236c558cc63be07419f8d65f",       "report_template_logo_url": "https://landray.dingtalkapps.com/alid/app/report/images/i |
| `report.template.detail` | `0` | `pass` | {   "errcode": 0,   "errorMsg": "ok",   "result": {     "allowAddReceivers": 0,     "lastModifiedTime": 1650879711000,     "report_template_fields": [       {         "field_name": "今日完成工作",         "field_sort": 0,      |
| `report.create` | `0` | `fail_business` | {   "code": "PARAM_ERROR",   "error": "参数错误",   "message": "参数错误。请检查输入参数是否正确，参考工具说明调整参数后重试",   "retryable": false,   "success": false,   "technical_detail": "调用失败，原因: 调用失败，原因: 调用远程服务业务异常, 出参校验不通过。root.success当前值：false， 预 |
| `calendar.event.create` | `0` | `pass` | {   "arguments": [],   "result": {     "attendees": [       {         "displayName": "天雷",         "optional": false,         "responseStatus": "accepted",         "self": true       }     ],     "created": 1774596561911 |
| `calendar.room.search(96)` | `0` | `pass` | {   "arguments": [],   "result": {     "result": [       {         "capacity": null,         "fullGroupPath": "阿里中心-杭州未科",         "labels": [],         "roomId": "bdff146f0172e49511c91904cf69b186003e6834fdd7fee7",       |
| `calendar.room.add` | `0` | `pass` | {   "arguments": [],   "errorCode": null,   "errorMsg": null,   "result": {},   "success": true }  |
| `calendar.room.delete` | `0` | `pass` | {   "arguments": [],   "errorCode": null,   "errorMsg": null,   "result": {},   "success": true }  |
| `ding.message.send(placeholder)` | `0` | `fail_business` | {   "dingOpenErrcode": 400009,   "errorMsg": "robotCode is in not valid or not in the org",   "result": {},   "success": false,   "trace_id": "0bb7c36217745965657028317e058c" }  |
