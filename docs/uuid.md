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
| ok | boolean | 函数是否执行成功 |
| uuid     | string   | 函数执行成功时返回 **uuid 字符串**，函数执行失败时返回 **错误信息** |

**demo**

```lua
local uuid = glua.uuid

local ok, u1 = uuid("v1")
if not ok then
    print("v1 err:", u1)
    return
end
print("v1: ", u1)

local ok, u3 = uuid("v3", "6ba7b810-9dad-11d1-80b4-00c04fd430c8", "user:10001")
if not ok then
    print("v3 err:", u3)
    return
end
print("v3: ", u3)

local ok, u4 = uuid("v4")
if not ok then
    print("v4 err:", u4)
    return
end
print("v4: ", u4)

local ok, u5 = uuid("v5", "6ba7b810-9dad-11d1-80b4-00c04fd430c8", "user:10001")
if not ok then
    print("v5 err:", u5)
    return
end
print("v5: ", u5)

local ok, u6 = uuid("v6")
if not ok then
    print("v6 err:", u6)
    return
end
print("v6: ", u6)

local ok, u7 = uuid("v7")
if not ok then
    print("v7 err:", u7)
    return
end
print("v7: ", u7)

local ok, u = uuid()
if not ok then
    print("v4 err:", u)
    return
end
print("v4: ", u)
```

