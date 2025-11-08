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
| raw      | string   | 转化之后的字符串 |
| err      | string   | 错误返回值       |

**demo**

```lua
local jwt = glua.jwt

local key = "0123456789"

local j = {
    ["username"] = "admin"
}

local raw, err = jwt.sign(key, "HS256", j)

if err ~= nil then
    print(err)
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
| jwt      | table    | 转化之后的表 |
| err      | string   | 错误返回值   |

**demo**

```lua
local jwt = glua.jwt

local key = "0123456789"

local raw = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIn0.WoMb_yIEbJ2wncbLtNVtqjdsraroEc6wzX_qvvxzAD8"

local j, err = jwt.verify(key, raw)

if err ~= nil then
    print(err)
    return
end

print(j.username)
```

