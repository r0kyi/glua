# uuid

通用唯一识别码模块

**参数值**

| 参数名称  | 参数类型 | 备注                                                         |
| --------- | -------- | ------------------------------------------------------------ |
| version   | string   | uuid 版本，非必须，可选值：<br />v1<br />v3<br />v4（默认）<br />v5<br />v6<br />v7 |
| namespace | string   | 命名空间，非必须，仅 v3 和 v5 需要<br />格式为：<br />xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx<br />urn:uuid:xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx |
| data      | string   | uuid 命名值，非必须，仅 v3 和 v5 需要                        |

**返回值**

| 参数名称 | 参数类型 | 备注        |
| -------- | -------- | ----------- |
| uuid     | string   | uuid 字符串 |
| err      | string   | 错误返回值  |

**demo**

```lua
local uuid = glua.uuid

local u1, err = uuid("v1")
if err ~= nil then
    print("v1 err:", err)
end
print("v1: ", u1)

local u3, err = uuid("v3", "6ba7b810-9dad-11d1-80b4-00c04fd430c8", "user:10001")
if err ~= nil then
    print("v3 err:", err)
end
print("v3: ", u3)

local u4, err = uuid("v4")
if err ~= nil then
    print("v4 err:", err)
end
print("v4: ", u4)

local u5, err = uuid("v5", "6ba7b810-9dad-11d1-80b4-00c04fd430c8", "user:10001")
if err ~= nil then
    print("v5 err:", err)
end
print("v5: ", u5)

local u6, err = uuid("v6")
if err ~= nil then
    print("v6 err:", err)
end
print("v6: ", u6)

local u7, err = uuid("v7")
if err ~= nil then
    print("v7 err:", err)
end
print("v7: ", u7)

local u, err = uuid()
if err ~= nil then
    print("v4 err:", err)
end
print("v4: ", u)
```

