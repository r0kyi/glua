# database

数据库连接

**参数值**

| 参数名称         | 参数类型 | 备注                            |
| ---------------- | -------- | ------------------------------- |
| driver_name      | string   | 驱动名，必须，可选值<br />mysql |
| data_source_name | string   | 数据源名称，必须                |

**返回值**

| 参数名称 | 参数类型 | 备注       |
| -------- | -------- | ---------- |
| db       | object   | 数据库     |
| err      | string   | 错误返回值 |

**demo**

```lua
local database = glua.database

local db, err = database{
    driver_name = "mysql",
    data_source_name = "root:password@tcp(127.0.0.1:3306)/testdb"
}

if err ~= nil then
    print(err)
end
```

## exec

执行无返回值的 sql 语句

**参数值**

| 参数名称 | 参数类型                                 | 备注               |
| -------- | ---------------------------------------- | ------------------ |
| query    | string                                   | sql 语句，必须     |
| args     | nil<br />boolean<br />number<br />string | sql 参数值，非必须 |

**返回值**

| 参数名称 | 参数类型 | 备注       |
| -------- | -------- | ---------- |
| err      | string   | 错误返回值 |

**demo**

```lua
local database = glua.database

local db, err = database{
    driver_name = "mysql",
    data_source_name = "root:password@tcp(127.0.0.1:3306)/testdb"
}

if err ~= nil then
    print(err)
end

local err = db.exec("INSERT INTO `users` (`username`, `password`) VALUES (?, ?)", "tom", "xxxx")

if err ~= nil then
    print(err)
end
```

## query

执行查询的 sql 语句

**参数值**

| 参数名称 | 参数类型                                 | 备注               |
| -------- | ---------------------------------------- | ------------------ |
| query    | string                                   | sql 语句，必须     |
| args     | nil<br />boolean<br />number<br />string | sql 参数值，非必须 |

**返回值**

| 参数名称 | 参数类型 | 备注               |
| -------- | -------- | ------------------ |
| rows     | table    | sql 语句查询返回值 |
| err      | string   | 错误返回值         |

**demo**

```lua
local database = glua.database

local db, err = database{
    driver_name = "mysql",
    data_source_name = "root:password@tcp(127.0.0.1:3306)/testdb"
}

if err ~= nil then
    print(err)
end

local rows, err = db.query("SELECT * FROM `users` WHERE `id` = ?;", 1)
if err ~= nil then
    print(err)
end
print(rows[1]["username"])
```

