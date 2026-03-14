# jwt

jwt 解析模块

## sign

将 table 转化为 jwt 字符串

**参数值**

| 参数名称 | 参数类型 | 备注                                                         |
| -------- | -------- | ------------------------------------------------------------ |
| key      | string   | 签名密钥，必须                                               |
| alg      | string   | 签名算法，必须，可选值<br />HS256（默认）<br />HS384<br />HS512<br />RS256<br />RS384<br />RS512<br />PS256<br />PS384<br />PS512<br />ES256<br />ES384<br />ES512<br />EdDSA |
| jwt      | table    | 需要转化的表，必须                                           |

**返回值**

| 参数名称 | 参数类型 | 备注             |
| -------- | -------- | ---------------- |
| ok | boolean | 函数是否执行成功 |
| raw      | string   | 函数执行成功时返回 **转化之后的字符串**，函数执行失败时返回 **错误信息** |

**demo**

```lua
local jwt = glua.jwt

local key = "0123456789"

local j = {
    ["username"] = "admin"
}

local ok, raw = jwt.sign(key, "HS256", j)

if not ok then
    print(raw)
    return
end

print(raw)
```

## verify

将 jwt 字符串转化为 table

**参数值**

| 参数名称 | 参数类型 | 备注                   |
| -------- | -------- | ---------------------- |
| key      | string   | 签名密钥，必须         |
| raw      | string   | 需要转化的字符串，必须 |

**返回值**

| 参数名称 | 参数类型 | 备注         |
| -------- | -------- | ------------ |
| ok | boolean | 函数是否执行成功 |
| jwt      | table    | 函数执行成功时返回 **转化之后的表**，函数执行失败时返回 **错误信息** |

**demo**

```lua
local jwt = glua.jwt

local key = "0123456789"

local raw = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIn0.WoMb_yIEbJ2wncbLtNVtqjdsraroEc6wzX_qvvxzAD8"

local ok, j = jwt.verify(key, raw)

if not ok then
    print(j)
    return
end

print(j.username)
```

