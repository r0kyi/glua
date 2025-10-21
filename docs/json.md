# json

json 解析模块

## encode

将 table 转化为 json 字符串

**参数值**

| 参数名称 | 参数类型 | 备注               |
| -------- | -------- | ------------------ |
| json     | table    | 需要转化的表，必须 |

**返回值**

| 参数名称 | 参数类型 | 备注             |
| -------- | -------- | ---------------- |
| raw      | string   | 转化之后的字符串 |
| err      | string   | 错误返回值       |

**demo**

```lua
local json = glua.json

local j = {
    ["name"] = "test",
    ["arr"] = {1.0, 2, 3, "4"},
    ["xxx"] = {
        ["1"] = 1,
    }
}
local raw, err = json.encode(j)

if err ~= nil then
    print(err)
    return
end

print(raw)
```

## decode

将 json 字符串转化为 table

**参数值**

| 参数名称 | 参数类型 | 备注                   |
| -------- | -------- | ---------------------- |
| raw      | string   | 需要转化的字符串，必须 |

**返回值**

| 参数名称 | 参数类型 | 备注         |
| -------- | -------- | ------------ |
| json     | table    | 转化之后的表 |
| err      | string   | 错误返回值   |

**demo**

```lua
local json = glua.json

local raw = '{"arr":[1,2,3,"4"],"name":"test","xxx":{"1":1}}'

local j, err = json.decode(raw)

if err ~= nil then
    print(err)
    return
end

print(j.name)
```

