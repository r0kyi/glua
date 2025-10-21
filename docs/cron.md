# cron

计划任务模块

## job

使用 cron 表达式生成计划任务

**参数值**

| 参数名称       | 参数类型 | 备注                                         |
| -------------- | -------- | -------------------------------------------- |
| cronExpression | string   | cron 表达式，必须                            |
| fn             | function | 需要执行的函数，必须<br />**只支持无参函数** |

**返回值**

| 参数名称 | 参数类型 | 备注       |
| -------- | -------- | ---------- |
| err      | string   | 错误返回值 |

**demo**

```lua
local cron = glua.cron

local c = cron()

local err = c.job("*/1 * * * *", function()
    print("hello")
end)

if err ~= nil then
    print(err)
    return
end

c.start_block()
```

## seconds

间隔执行计划任务

**参数值**

| 参数名称 | 参数类型 | 备注                                         |
| -------- | -------- | -------------------------------------------- |
| interval | number   | 间隔秒数，必须                               |
| fn       | function | 需要执行的函数，必须<br />**只支持无参函数** |

**返回值**

| 参数名称 | 参数类型 | 备注       |
| -------- | -------- | ---------- |
| err      | string   | 错误返回值 |

**demo**

```lua
local cron = glua.cron

local c = cron()

local err = c.seconds(1, function()
    print("hello")
end)

if err ~= nil then
    print(err)
    return
end

c.start_block()
```

## minutes

间隔执行计划任务

**参数值**

| 参数名称 | 参数类型 | 备注                                         |
| -------- | -------- | -------------------------------------------- |
| interval | number   | 间隔分钟数，必须                             |
| fn       | function | 需要执行的函数，必须<br />**只支持无参函数** |

**返回值**

| 参数名称 | 参数类型 | 备注       |
| -------- | -------- | ---------- |
| err      | string   | 错误返回值 |

**demo**

```lua
local cron = glua.cron

local c = cron()

local err = c.minutes(1, function()
    print("hello")
end)

if err ~= nil then
    print(err)
    return
end

c.start_block()
```

## hours

间隔执行计划任务

**参数值**

| 参数名称 | 参数类型 | 备注                                         |
| -------- | -------- | -------------------------------------------- |
| interval | number   | 间隔小时数，必须                             |
| fn       | function | 需要执行的函数，必须<br />**只支持无参函数** |

**返回值**

| 参数名称 | 参数类型 | 备注       |
| -------- | -------- | ---------- |
| err      | string   | 错误返回值 |

**demo**

```lua
local cron = glua.cron

local c = cron()

local err = c.hours(1, function()
    print("hello")
end)

if err ~= nil then
    print(err)
    return
end

c.start_block()
```

## days

间隔执行计划任务

**参数值**

| 参数名称 | 参数类型 | 备注                                         |
| -------- | -------- | -------------------------------------------- |
| interval | number   | 间隔天数，必须                               |
| fn       | function | 需要执行的函数，必须<br />**只支持无参函数** |

**返回值**

| 参数名称 | 参数类型 | 备注       |
| -------- | -------- | ---------- |
| err      | string   | 错误返回值 |

**demo**

```lua
local cron = glua.cron

local c = cron()

local err = c.days(1, function()
    print("hello")
end)

if err ~= nil then
    print(err)
    return
end

c.start_block()
```

## weeks

间隔执行计划任务

**参数值**

| 参数名称 | 参数类型 | 备注                                         |
| -------- | -------- | -------------------------------------------- |
| interval | number   | 间隔周数，必须                               |
| fn       | function | 需要执行的函数，必须<br />**只支持无参函数** |

**返回值**

| 参数名称 | 参数类型 | 备注       |
| -------- | -------- | ---------- |
| err      | string   | 错误返回值 |

**demo**

```lua
local cron = glua.cron

local c = cron()

local err = c.weeks(1, function()
    print("hello")
end)

if err ~= nil then
    print(err)
    return
end

c.start_block()
```

## start_block

以阻塞线程的形式运行计划任务

**必须放在所有逻辑的最后，在其后方的代码不会运行**

**demo**

```lua
local cron = glua.cron

local c = cron()

local err = c.seconds(1, function()
    print("hello")
end)

if err ~= nil then
    print(err)
    return
end

c.start_block()
```

## start_async

以异步的方式运行计划任务

**除非你能确保脚本不会退出，否则不要调用此函数**

```lua
local cron = glua.cron

local c = cron()

local err = c.seconds(1, function()
    print("hello")
end)

if err ~= nil then
    print(err)
    return
end

c.start_async()

while true do
    os.execute("sleep 1")
end
```
