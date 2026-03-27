# 会议室ID实采链路专项报告

| 指标 | 值 |
|---|---|
| 时间 | 2026-03-27 15:22:35 |
| 步骤数 | `32` |
| pass | `32` |
| fail_exit | `0` |
| fail_business | `0` |

- `event-id`: `bjRTVjdjSlliZHpiblp5TVB4TDlGQT09`
- `发现的 roomId 数量`: `2`
- `成功验证 roomId`: `bdff146f0172e49511c91904cf69b186003e6834fdd7fee7`

| 步骤 | exit | 语义 | 摘要 |
|---|---:|---|---|
| `calendar.event.create` | `0` | `pass` | {   "arguments": [],   "result": {     "attendees": [       {         "displayName": "天雷",         "optional": false,         "responseStatus": "accepted",         "self": true       }     ],     "created": 1774596132153 |
| `calendar.room.list-groups` | `0` | `pass` | {   "result": {     "groupList": [       {         "groupId": 41,         "groupName": "绿城·未来park",         "parentId": 0       },       {         "groupId": 74,         "groupName": "西溪C区",         "parentId": 0       } |
| `calendar.room.search(group=41)` | `0` | `pass` | {   "arguments": [],   "result": {     "result": [       {         "labels": null       }     ]   },   "success": true }  |
| `calendar.room.search(group=74)` | `0` | `pass` | {   "arguments": [],   "result": {     "result": [       {         "labels": null       }     ]   },   "success": true }  |
| `calendar.room.search(group=96)` | `0` | `pass` | {   "arguments": [],   "result": {     "result": [       {         "capacity": null,         "fullGroupPath": "阿里中心-杭州未科",         "labels": [],         "roomId": "bdff146f0172e49511c91904cf69b186003e6834fdd7fee7",       |
| `calendar.room.search(group=42)` | `0` | `pass` | {   "arguments": [],   "result": {     "result": [       {         "labels": null       }     ]   },   "success": true }  |
| `calendar.room.search(group=43)` | `0` | `pass` | {   "arguments": [],   "result": {     "result": [       {         "labels": null       }     ]   },   "success": true }  |
| `calendar.room.search(group=44)` | `0` | `pass` | {   "arguments": [],   "result": {     "result": [       {         "labels": null       }     ]   },   "success": true }  |
| `calendar.room.search(group=45)` | `0` | `pass` | {   "arguments": [],   "result": {     "result": [       {         "labels": null       }     ]   },   "success": true }  |
| `calendar.room.search(group=50)` | `0` | `pass` | {   "arguments": [],   "result": {     "result": [       {         "labels": null       }     ]   },   "success": true }  |
| `calendar.room.search(group=51)` | `0` | `pass` | {   "arguments": [],   "result": {     "result": [       {         "labels": null       }     ]   },   "success": true }  |
| `calendar.room.search(group=52)` | `0` | `pass` | {   "arguments": [],   "result": {     "result": [       {         "labels": null       }     ]   },   "success": true }  |
| `calendar.room.search(group=53)` | `0` | `pass` | {   "arguments": [],   "result": {     "result": [       {         "labels": null       }     ]   },   "success": true }  |
| `calendar.room.search(group=54)` | `0` | `pass` | {   "arguments": [],   "result": {     "result": [       {         "labels": null       }     ]   },   "success": true }  |
| `calendar.room.search(group=55)` | `0` | `pass` | {   "arguments": [],   "result": {     "result": [       {         "labels": null       }     ]   },   "success": true }  |
| `calendar.room.search(group=58)` | `0` | `pass` | {   "arguments": [],   "result": {     "result": [       {         "labels": null       }     ]   },   "success": true }  |
| `calendar.room.search(group=59)` | `0` | `pass` | {   "arguments": [],   "result": {     "result": [       {         "labels": null       }     ]   },   "success": true }  |
| `calendar.room.search(group=60)` | `0` | `pass` | {   "arguments": [],   "result": {     "result": [       {         "labels": null       }     ]   },   "success": true }  |
| `calendar.room.search(group=61)` | `0` | `pass` | {   "arguments": [],   "result": {     "result": [       {         "labels": null       }     ]   },   "success": true }  |
| `calendar.room.search(group=62)` | `0` | `pass` | {   "arguments": [],   "result": {     "result": [       {         "labels": null       }     ]   },   "success": true }  |
| `calendar.room.search(group=63)` | `0` | `pass` | {   "arguments": [],   "result": {     "result": [       {         "labels": null       }     ]   },   "success": true }  |
| `calendar.room.search(group=64)` | `0` | `pass` | {   "arguments": [],   "result": {     "result": [       {         "labels": null       }     ]   },   "success": true }  |
| `calendar.room.search(group=65)` | `0` | `pass` | {   "arguments": [],   "result": {     "result": [       {         "labels": null       }     ]   },   "success": true }  |
| `calendar.room.search(group=46)` | `0` | `pass` | {   "arguments": [],   "result": {     "result": [       {         "labels": null       }     ]   },   "success": true }  |
| `calendar.room.search(group=48)` | `0` | `pass` | {   "arguments": [],   "result": {     "result": [       {         "labels": null       }     ]   },   "success": true }  |
| `calendar.room.search(group=75)` | `0` | `pass` | {   "arguments": [],   "result": {     "result": [       {         "labels": null       }     ]   },   "success": true }  |
| `calendar.room.search(group=76)` | `0` | `pass` | {   "arguments": [],   "result": {     "result": [       {         "labels": null       }     ]   },   "success": true }  |
| `calendar.room.search(group=77)` | `0` | `pass` | {   "arguments": [],   "result": {     "result": [       {         "labels": null       }     ]   },   "success": true }  |
| `calendar.room.search(group=78)` | `0` | `pass` | {   "arguments": [],   "result": {     "result": [       {         "labels": null       }     ]   },   "success": true }  |
| `calendar.room.add(room=bdff146f0172e49511c91904cf69b186003e6834fdd7fee7)` | `0` | `pass` | {   "arguments": [],   "errorCode": null,   "errorMsg": null,   "result": {},   "success": true }  |
| `calendar.room.delete(room=bdff146f0172e49511c91904cf69b186003e6834fdd7fee7)` | `0` | `pass` | {   "arguments": [],   "errorCode": null,   "errorMsg": null,   "result": {},   "success": true }  |
| `calendar.event.delete` | `0` | `pass` | {   "arguments": [],   "errorCode": null,   "errorMsg": null,   "result": {},   "success": true }  |
