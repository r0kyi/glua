# xml

xml 解析模块

## encode

将 table 转化为 xml 字符串

**参数值**

| 参数名称 | 参数类型 | 备注               |
| -------- | -------- | ------------------ |
| xml      | table    | 需要转化的表，必须 |

**返回值**

| 参数名称 | 参数类型 | 备注             |
| -------- | -------- | ---------------- |
| ok | boolean | 函数是否执行成功 |
| raw      | string   | 函数执行成功时返回 **转化之后的字符串**，函数执行失败时返回 **错误信息** |

**demo**

```lua
local xml = glua.xml

local x = {
    ["note"] = {
        ["to"] = "Tove",
        ["from"] = "Jani",
        ["heading"] = "Reminder",
        ["body"] = "Don't forget me this weekend!",
    }
}

local ok, raw = xml.encode(x)

if not ok then
    print(raw)
    return
end

print(raw)
```

## decode

将 xml 字符串转化为 table

**参数值**

| 参数名称 | 参数类型 | 备注                   |
| -------- | -------- | ---------------------- |
| raw      | string   | 需要转化的字符串，必须 |

**返回值**

| 参数名称 | 参数类型 | 备注         |
| -------- | -------- | ------------ |
| ok | boolean | 函数是否执行成功 |
| xml      | table    | 函数执行成功时返回 **转化之后的表**，函数执行失败时返回 **错误信息** |

**demo**

```lua
local xml = glua.xml

local raw = [[
<?xml version="1.0" encoding="UTF-8"?>
<note>
  <to>Tove</to>
  <from>Jani</from>
  <heading>Reminder</heading>
  <body>Don't forget me this weekend!</body>
</note>
]]

local ok, x = xml.decode(raw)

if not ok then
    print(x)
    return
end

print(x["note"]["to"])
```

