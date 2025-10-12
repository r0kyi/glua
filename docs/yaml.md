# yaml

yaml 解析模块

## encode

将 table 转化为 yaml 字符串

**参数值**

| 参数名称 | 参数类型 | 备注               |
| -------- | -------- | ------------------ |
| yaml     | table    | 需要转化的表，必须 |

**返回值**

| 参数名称 | 参数类型 | 备注             |
| -------- | -------- | ---------------- |
| raw      | string   | 转化之后的字符串 |

**demo**

```lua
local yaml = glua.yaml

local y = {
    ["name"] = "test",
    ["arr"] = {1.0, 2, 3, "4"},
    ["xxx"] = {
        ["1"] = 1,
    }
}
local raw, err = yaml.encode(y)

if err ~= nil then
    print(err)
    return
end

print(raw)
```

## decode

将 yaml 字符串转化为 table

**参数值**

| 参数名称 | 参数类型 | 备注                   |
| -------- | -------- | ---------------------- |
| raw      | string   | 需要转化的字符串，必须 |

**返回值**

| 参数名称 | 参数类型 | 备注         |
| -------- | -------- | ------------ |
| yaml     | table    | 转化之后的表 |

**demo**

```lua
local yaml = glua.yaml

local raw = [[
arr:
- 1.0
- 2.0
- 3.0
- "4"
name: test
xxx:
  "1": 1.0
]]

local y, err = yaml.decode(raw)

if err ~= nil then
    print(err)
    return
end

print(y.name)
```

