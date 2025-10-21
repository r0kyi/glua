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
| raw      | string   | 转化之后的字符串 |
| err      | string   | 错误返回值       |

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

local raw, err = xml.encode(x)

if err ~= nil then
    print(err)
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
| xml      | table    | 转化之后的表 |
| err      | string   | 错误返回值   |

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

local x, err = xml.decode(raw)

if err ~= nil then
    print(err)
    return
end

print(x["note"]["to"])
```

