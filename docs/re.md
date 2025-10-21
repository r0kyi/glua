# re

正则模块

## compile

将正则表达式字符串编译成一个可复用的正则对象

**参数值**

| 参数名称 | 参数类型 | 备注       |
| -------- | -------- | ---------- |
| pattern  | string   | 模式，必须 |

**返回值**

| 参数名称 | 参数类型 | 备注               |
| -------- | -------- | ------------------ |
| regexp   | userdata | 编译后的正则表达式 |
| err      | string   | 错误返回值         |

## match

检测某个字符串是否符合正则表达式

**参数值**

| 参数名称 | 参数类型 | 备注             |
| -------- | -------- | ---------------- |
| src      | string   | 原始字符串，必须 |

**返回值**

| 参数名称 | 参数类型 | 备注               |
| -------- | -------- | ------------------ |
| match    | boolean  | 是否命中正则表达式 |

**demo**

```lua
local re = glua.re

local r, err = re.compile("\\d+")

if err ~= nil then
    print(err)
    return
end

print(r.match("123456"))
print(r.match("re"))

-- true
-- false
```

## find

查找第一个符合正则的子串

**参数值**

| 参数名称 | 参数类型 | 备注             |
| -------- | -------- | ---------------- |
| src      | string   | 原始字符串，必须 |

**返回值**

| 参数名称 | 参数类型 | 备注         |
| -------- | -------- | ------------ |
| dst      | string   | 第一个匹配项 |

**demo**

```lua
local re = glua.re

local r, err = re.compile("\\d+")

if err ~= nil then
    print(err)
    return
end

print(r.find("123re456"))
print(r.find("re123"))

-- 123
-- 123
```

## find_all

查找所有符合正则的子串

**参数值**

| 参数名称 | 参数类型 | 备注             |
| -------- | -------- | ---------------- |
| src      | string   | 原始字符串，必须 |

**返回值**

| 参数名称 | 参数类型 | 备注       |
| -------- | -------- | ---------- |
| dst      | table    | 所有匹配项 |

**demo**

```lua
local re = glua.re

local r, err = re.compile("\\d+")

if err ~= nil then
    print(err)
    return
end

local dst = r.find_all("123re456")

for _, v in pairs(dst) do
    print(v)
end

-- 123
-- 456
```

## replace

将匹配正则的部分替换成新的内容

**参数值**

| 参数名称 | 参数类型 | 备注             |
| -------- | -------- | ---------------- |
| src      | string   | 原始字符串，必须 |

**返回值**

| 参数名称 | 参数类型 | 备注               |
| -------- | -------- | ------------------ |
| dst      | string   | 替换匹配后的字符串 |

**demo**

```lua
local re = glua.re

local r, err = re.compile("\\d+")

if err ~= nil then
    print(err)
    return
end

print(r.replace("123re456", "re"))

-- rerere
```

## split

用正则表达式来切分字符串

**参数值**

| 参数名称 | 参数类型 | 备注             |
| -------- | -------- | ---------------- |
| src      | string   | 原始字符串，必须 |

**返回值**

| 参数名称 | 参数类型 | 备注                     |
| -------- | -------- | ------------------------ |
| dst      | table    | 正则表达式分割后的字符串 |

**demo**

```lua
local re = glua.re

local r, err = re.compile("\\d+")

if err ~= nil then
    print(err)
    return
end

local dst = r.split("re12345re")

for i, v in pairs(dst) do
    print(v)
end

-- re
-- re
```
