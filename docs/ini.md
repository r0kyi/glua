# ini

ini 配置

## load

加载 ini 文件，解析配置

**参数值**

| 参数名称 | 参数类型 | 备注                      |
| -------- | -------- | ------------------------- |
| filename | string   | 需要加载的 ini 文件，必须 |

**返回值**

| 参数名称 | 参数类型 | 备注       |
| -------- | -------- | ---------- |
| cfg      | table    | 配置       |
| err      | string   | 错误返回值 |

**demo**

```lua
local ini = glua.ini

local cfg, err = ini.load("cfg.ini")

if err ~= nil then
    print(err)
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

| 参数名称 | 参数类型 | 备注       |
| -------- | -------- | ---------- |
| err      | string   | 错误返回值 |

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

local err = ini.save("cfg.ini", cfg)

if err ~= nil then
    print(err)
    return
end
```
