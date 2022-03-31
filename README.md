# DataServer

自用可配置数据服务

## 全局错误码与数据格式

```json
{
  "error": 0,
  "data": "some data"
}
```
| code | 说明    |
|------|-------|
| 0    | OK    |
| 1    | TLS错误 |

## 接口说明

| URL                                        | 方法  | 说明          | 状态  |
|--------------------------------------------|-----|-------------|-----|
| /api/v1/dataserver/xmu-daily-report-config | GET | 拉取每日打卡配置元数据 | 在线  |

### GET /api/v1/dataserver/xmu-daily-report-config

+ 参数：无
+ 访问权限：无

返回值：

```json
{
  "type": "string/number/email/dropdown etc.",
  "key": "生成打卡配置项的key",
  "name": "该配置项的显示名称",
  "values": "一个逗号分隔的字符串，为drop down的可选项"
}
```

## 使用指南

TODO