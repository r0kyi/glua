# ini

ini 配置

## load

加载 ini 文件，解析配置

**参数值**

| 参数名称 | 参数类型 | 备注                      |
| -------- | -------- | ------------------------- |
| filename | string   | 需要加载的 ini 文件，必须 |

**返回值**

| 参数名称 | 参数类型 | 备注                                                         |
| -------- | -------- | ------------------------------------------------------------ |
| ok       | boolean  | 函数是否执行成功                                             |
| cfg      | table    | 函数执行成功时返回 **配置**，函数执行失败时返回 **错误信息** |

**demo**

```lua
local ini = glua.ini

local ok, cfg = ini.load("cfg.ini")

if not ok then
    print(cfg)
    return
end

print(cfg.DEFAULT.username)
print(cfg.database.host)
```

ini 文件内容

```ini
username = admin
password = admin

[database]
host = 127.0.0.1
port = 3306
```

## save

将配置保存至 ini 文件

**参数值**

| 参数名称 | 参数类型 | 备注                    |
| -------- | -------- | ----------------------- |
| filename | string   | 保存 ini 文件路径，必须 |
| cfg      | table    | 配置，必须              |

**返回值**

| 参数名称 | 参数类型 | 备注             |
| -------- | -------- | ---------------- |
| ok       | boolean  | 函数是否执行成功 |
| result | string   | 错误信息       |

**demo**

```lua
local ini = glua.ini

local cfg = {
    ["DEFAULT"] = {
        ["username"] = "admin",
        ["password"] = "admin",
    },
    ["database"] = {
        ["host"] = "127.0.0.1",
        ["port"] = 3306,
    }
}

local ok, result = ini.save("cfg.ini", cfg)

if not ok then
    print(result)
    return
end
```
