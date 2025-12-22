# validator

数据验证模块

## email

验证是否符合邮箱格式

**参数值**

| 参数名称 | 参数类型 | 备注               |
| -------- | -------- | ------------------ |
| value    | string   | 需要验证的值，必须 |

**返回值**

| 参数名称 | 参数类型 | 备注       |
| -------- | -------- | ---------- |
| valid    | boolean  | 验证返回值 |

**demo**

```lua
local validator = glua.validator

print(validator.email("user@example.com"))
```

## cidr

验证是否符合 cidr 格式

**参数值**

| 参数名称 | 参数类型 | 备注               |
| -------- | -------- | ------------------ |
| value    | string   | 需要验证的值，必须 |

**返回值**

| 参数名称 | 参数类型 | 备注       |
| -------- | -------- | ---------- |
| valid    | boolean  | 验证返回值 |

**demo**

```lua
local validator = glua.validator

print(validator.cidr("10.0.0.0/8"))
print(validator.cidr("::1/128"))
```

## cidrv4

验证是否符合 cidrv4 格式

**参数值**

| 参数名称 | 参数类型 | 备注               |
| -------- | -------- | ------------------ |
| value    | string   | 需要验证的值，必须 |

**返回值**

| 参数名称 | 参数类型 | 备注       |
| -------- | -------- | ---------- |
| valid    | boolean  | 验证返回值 |

**demo**

```lua
local validator = glua.validator

print(validator.cidrv4("10.0.0.0/8"))
```

## cidrv6

验证是否符合 cidrv6 格式

**参数值**

| 参数名称 | 参数类型 | 备注               |
| -------- | -------- | ------------------ |
| value    | string   | 需要验证的值，必须 |

**返回值**

| 参数名称 | 参数类型 | 备注       |
| -------- | -------- | ---------- |
| valid    | boolean  | 验证返回值 |

**demo**

```lua
local validator = glua.validator

print(validator.cidrv6("::1/128"))
```

## ip

验证是否符合 ip 格式

**参数值**

| 参数名称 | 参数类型 | 备注               |
| -------- | -------- | ------------------ |
| value    | string   | 需要验证的值，必须 |

**返回值**

| 参数名称 | 参数类型 | 备注       |
| -------- | -------- | ---------- |
| valid    | boolean  | 验证返回值 |

**demo**

```lua
local validator = glua.validator

print(validator.ip("127.0.0.1"))
print(validator.ip("::1"))
```

## ipv4

验证是否符合 ipv4 格式

**参数值**

| 参数名称 | 参数类型 | 备注               |
| -------- | -------- | ------------------ |
| value    | string   | 需要验证的值，必须 |

**返回值**

| 参数名称 | 参数类型 | 备注       |
| -------- | -------- | ---------- |
| valid    | boolean  | 验证返回值 |

**demo**

```lua
local validator = glua.validator

print(validator.ipv4("127.0.0.1"))
```

## ipv6

验证是否符合 ipv6 格式

**参数值**

| 参数名称 | 参数类型 | 备注               |
| -------- | -------- | ------------------ |
| value    | string   | 需要验证的值，必须 |

**返回值**

| 参数名称 | 参数类型 | 备注       |
| -------- | -------- | ---------- |
| valid    | boolean  | 验证返回值 |

**demo**

```lua
local validator = glua.validator

print(validator.ipv4("127.0.0.1"))
```

## url

验证是否符合 url 格式

**参数值**

| 参数名称 | 参数类型 | 备注               |
| -------- | -------- | ------------------ |
| value    | string   | 需要验证的值，必须 |

**返回值**

| 参数名称 | 参数类型 | 备注       |
| -------- | -------- | ---------- |
| valid    | boolean  | 验证返回值 |

**demo**

```lua
local validator = glua.validator

print(validator.url("https://example.com/1.html"))
```

## hostname

验证是否符合 hostname 格式

**参数值**

| 参数名称 | 参数类型 | 备注               |
| -------- | -------- | ------------------ |
| value    | string   | 需要验证的值，必须 |

**返回值**

| 参数名称 | 参数类型 | 备注       |
| -------- | -------- | ---------- |
| valid    | boolean  | 验证返回值 |

**demo**

```lua
local validator = glua.validator

print(validator.hostname("example.com"))
print(validator.hostname("localhost"))
print(validator.hostname("name-01"))
```

## mac

验证是否符合 mac 格式

**参数值**

| 参数名称 | 参数类型 | 备注               |
| -------- | -------- | ------------------ |
| value    | string   | 需要验证的值，必须 |

**返回值**

| 参数名称 | 参数类型 | 备注       |
| -------- | -------- | ---------- |
| valid    | boolean  | 验证返回值 |

**demo**

```lua
local validator = glua.validator

print(validator.mac("FF-FF-FF-FF-FF-FF"))
print(validator.mac("FF:FF:FF:FF:FF:FF"))
```

## alpha

验证是否符合 alpha 格式

只允许：`A–Z`、`a–z`

**参数值**

| 参数名称 | 参数类型 | 备注               |
| -------- | -------- | ------------------ |
| value    | string   | 需要验证的值，必须 |

**返回值**

| 参数名称 | 参数类型 | 备注       |
| -------- | -------- | ---------- |
| valid    | boolean  | 验证返回值 |

**demo**

```lua
local validator = glua.validator

print(validator.alpha("alpha"))
```

## alphanum

验证是否符合 alphanum 格式

只允许：`A–Z`、`a–z`、`0–9`

**参数值**

| 参数名称 | 参数类型 | 备注               |
| -------- | -------- | ------------------ |
| value    | string   | 需要验证的值，必须 |

**返回值**

| 参数名称 | 参数类型 | 备注       |
| -------- | -------- | ---------- |
| valid    | boolean  | 验证返回值 |

**demo**

```lua
local validator = glua.validator

print(validator.alphanum("alpha09"))
```
