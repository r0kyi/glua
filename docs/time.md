# time

时间模块

## now

获取当前本地时间

**返回值**

| 参数名称 | 参数类型 | 备注     |
| -------- | -------- | -------- |
| time     | userdata | 时间模块 |

**demo**

```lua
local time = glua.time

local t = time.now()

print(t.year())
```

## date

生成指定日期和时间

**参数值**

| 参数名称 | 参数类型 | 备注                                                       |
| -------- | -------- | ---------------------------------------------------------- |
| year     | number   | 年，必须                                                   |
| month    | number   | 月，必须                                                   |
| day      | number   | 日，必须                                                   |
| hour     | number   | 小时，必须                                                 |
| min      | number   | 分，必须                                                   |
| sec      | number   | 秒，必须                                                   |
| nsec     | number   | 纳秒，必须                                                 |
| loc      | string   | 时区，必须<br />如果传入的字符串不合法，则使用系统本地时区 |

**返回值**

| 参数名称 | 参数类型 | 备注     |
| -------- | -------- | -------- |
| time     | userdata | 时间模块 |

```lua
local time = glua.time

local t = time.date(2006, 01, 02, 15, 04, 05, 0, "Asia/Shanghai")

print(t.year())
```

## unix

根据 unix 时间戳构造时间

**参数值**

| 参数名称 | 参数类型 | 备注       |
| -------- | -------- | ---------- |
| sec      | number   | 秒，必须   |
| nsec     | number   | 纳秒，必须 |

**返回值**

| 参数名称 | 参数类型 | 备注     |
| -------- | -------- | -------- |
| time     | userdata | 时间模块 |

**demo**

```lua
local time = glua.time

local t = time.unix(1136185445, 0)

print(t.year())
```

## parse

解析时间字符串

**参数值**

| 参数名称 | 参数类型 | 备注                     |
| -------- | -------- | ------------------------ |
| layout   | string   | 时间格式模板，必须       |
| value    | string   | 需要解析的时间字符，必须 |

**返回值**

| 参数名称 | 参数类型 | 备注     |
| -------- | -------- | -------- |
| time     | userdata | 时间模块 |

**demo**

```lua
local time = glua.time

local t = time.parse("2006-01-02 15:04:05", "2025-10-03 14:30:00")

print(t.year())
```

## parse_in_location

指定时区解析时间字符串

**参数值**

| 参数名称 | 参数类型 | 备注                                                       |
| -------- | -------- | ---------------------------------------------------------- |
| layout   | string   | 时间格式模板，必须                                         |
| value    | string   | 需要解析的时间字符，必须                                   |
| loc      | string   | 时区，必须<br />如果传入的字符串不合法，则使用系统本地时区 |

**返回值**

| 参数名称 | 参数类型 | 备注     |
| -------- | -------- | -------- |
| time     | userdata | 时间模块 |

**demo**

```lua
local time = glua.time

local t = time.parse_in_location("2006-01-02 15:04:05", "2025-10-03 14:30:00", "Asia/Shanghai")

print(t.year())
```

## year

获取当前时间对象的年份

**返回值**

| 参数名称 | 参数类型 | 备注 |
| -------- | -------- | ---- |
| year     | number   | 年份 |

**demo**

```lua
local time = glua.time

local t = time.now()

print(t.year())
```

## month

获取当前时间对象的月份

**返回值**

| 参数名称 | 参数类型 | 备注 |
| -------- | -------- | ---- |
| month    | number   | 月份 |

**demo**

```lua
local time = glua.time

local t = time.now()

print(t.month())
```

## day

获取当前时间对象的日

**返回值**

| 参数名称 | 参数类型 | 备注 |
| -------- | -------- | ---- |
| day      | number   | 日   |

**demo**

```lua
local time = glua.time

local t = time.now()

print(t.day())
```

## hour

获取当前时间对象的小时

**返回值**

| 参数名称 | 参数类型 | 备注 |
| -------- | -------- | ---- |
| hour     | number   | 小时 |

**demo**

```lua
local time = glua.time

local t = time.now()

print(t.hour())
```

## minute

获取当前时间对象的分钟

**返回值**

| 参数名称 | 参数类型 | 备注 |
| -------- | -------- | ---- |
| minute   | number   | 分钟 |

**demo**

```lua
local time = glua.time

local t = time.now()

print(t.year())
```

## second

获取当前时间对象的秒

**返回值**

| 参数名称 | 参数类型 | 备注 |
| -------- | -------- | ---- |
| second   | number   | 秒   |

**demo**

```lua
local time = glua.time

local t = time.now()

print(t.second())
```

## nanosecond

获取当前时间对象的纳秒

**返回值**

| 参数名称   | 参数类型 | 备注 |
| ---------- | -------- | ---- |
| nanosecond | number   | 纳秒 |

**demo**

```lua
local time = glua.time

local t = time.now()

print(t.nanosecond())
```

## format

将时间对象格式化成字符串

**参数值**

| 参数名称 | 参数类型 | 备注               |
| -------- | -------- | ------------------ |
| layout   | string   | 时间格式模板，必须 |

**返回值**

| 参数名称  | 参数类型 | 备注       |
| --------- | -------- | ---------- |
| formatted | string   | 时间字符串 |

**demo**

```lua
local time = glua.time

local t = time.now()

print(t.format("2006-01-02 15:04:05"))
```

