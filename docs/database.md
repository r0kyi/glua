# database

数据库连接，具体的数据库使用方法请参考对应的库

- [mysql](https://github.com/go-sql-driver/mysql)
- [postgres](https://github.com/jackc/pgx/)
- [sqlserver](https://github.com/denisenkom/go-mssqldb)
- [sqlite](https://pkg.go.dev/modernc.org/sqlite)

**参数值**

| 参数名称         | 参数类型 | 备注                                                         |
| ---------------- | -------- | ------------------------------------------------------------ |
| driver_name      | string   | 驱动名，必须，可选值：<br />mysql<br />pgx<br />sqlserver<br />sqlite |
| data_source_name | string   | 数据源名称，必须                                             |

**返回值**

| 参数名称 | 参数类型 | 备注                         |
| -------- | -------- | ---------------------------- |
| ok       | boolean  | 函数是否执行成功             |
| db       | object   | 函数执行成功时返回 **数据库**，函数执行失败时返回 **错误信息** |

**demo**

```lua
local database = glua.database

local ok, db = database{
    driver_name = "mysql",
    data_source_name = "root:password@tcp(127.0.0.1:3306)/testdb"
}

if not ok then
    print(db)
    return
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

| 参数名称 | 参数类型 | 备注             |
| -------- | -------- | ---------------- |
| ok       | boolean  | 函数是否执行成功 |
| result   | string   | 错误信息       |

**demo**

```lua
local database = glua.database

local ok, db = database{
    driver_name = "mysql",
    data_source_name = "root:password@tcp(127.0.0.1:3306)/testdb"
}

if not ok then
    print(db)
    return
end

local ok, result = db.exec("INSERT INTO `users` (`username`, `password`) VALUES (?, ?)", "tom", "xxxx")

if not ok then
    print(result)
    return
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

| 参数名称 | 参数类型 | 备注                                                         |
| -------- | -------- | ------------------------------------------------------------ |
| ok       | boolean  | 函数是否执行成功                                             |
| rows     | table    | 函数执行成功时返回  **sql 语句查询返回值**，函数执行失败时返回 **错误信息** |

**demo**

```lua
local database = glua.database

local ok, db = database{
    driver_name = "mysql",
    data_source_name = "root:password@tcp(127.0.0.1:3306)/testdb"
}

if not ok then
    print(db)
    return
end

local ok, rows = db.query("SELECT * FROM `users` WHERE `id` = ?;", 1)
if not ok then
    print(rows)
    return
end
print(rows[1]["username"])
```

## prepare

执行查询的 sql 语句

**参数值**

| 参数名称 | 参数类型 | 备注           |
| -------- | -------- | -------------- |
| query    | string   | sql 语句，必须 |

**返回值**

| 参数名称  | 参数类型 | 备注                                    |
| --------- | -------- | --------------------------------------- |
| ok        | boolean  | 函数是否执行成功                        |
| statement | object   | 函数执行成功时返回 **预编译 sql 的封装**，函数执行失败时返回 **错误信息** |

**demo**

```lua
local database = glua.database

local ok, db = database{
    driver_name = "mysql",
    data_source_name = "root:password@tcp(127.0.0.1:3306)/testdb"
}

if not ok then
    print(db)
    return
end

local ok, stmt = db.prepare("SELECT * FROM `users` WHERE `id` = ?;")
if not ok then
    print(stmt)
    return
end

local ok, rows = stmt.query(1)
if not ok then
    print(rows)
    return
end
stmt.close()

print(rows[1]["username"])

db.close()
```

## close

关闭数据库连接

**返回值**

| 参数名称 | 参数类型 | 备注             |
| -------- | -------- | ---------------- |
| ok       | boolean  | 函数是否执行成功 |
| result   | string   | 错误信息       |

**demo**

```lua
local database = glua.database

local ok, db = database{
    driver_name = "mysql",
    data_source_name = "root:password@tcp(127.0.0.1:3306)/testdb"
}

local ok, result = db.close()
if not ok then
    print(result)
end
```

## statement

### query

执行预处理查询的 sql 语句

**参数值**

| 参数名称 | 参数类型                                 | 备注               |
| -------- | ---------------------------------------- | ------------------ |
| args     | nil<br />boolean<br />number<br />string | sql 参数值，非必须 |

**返回值**

| 参数名称 | 参数类型 | 备注                                                         |
| -------- | -------- | ------------------------------------------------------------ |
| ok       | boolean  | 函数是否执行成功                                             |
| rows     | table    | 函数执行成功时返回  **sql 语句查询返回值**，函数执行失败时返回 **错误信息** |

**demo**

```lua
local database = glua.database

local ok, db = database{
    driver_name = "mysql",
    data_source_name = "root:password@tcp(127.0.0.1:3306)/testdb"
}

if not ok then
    print(db)
    return
end

local ok, stmt = db.prepare("SELECT * FROM `users` WHERE `id` = ?;")
if not ok then
    print(stmt)
    return
end

local ok, rows = stmt.query(1)
if not ok then
    print(rows)
    return
end
stmt.close()

print(rows[1]["username"])

db.close()
```

### close

关闭预编译 sql

**返回值**

| 参数名称 | 参数类型 | 备注             |
| -------- | -------- | ---------------- |
| ok       | boolean  | 函数是否执行成功 |
| result   | string   | 错误信息       |

**demo**

```lua
local database = glua.database

local ok, db = database{
    driver_name = "mysql",
    data_source_name = "root:password@tcp(127.0.0.1:3306)/testdb"
}

if not ok then
    print(db)
    return
end

local ok, stmt = db.prepare("SELECT * FROM `users` WHERE `id` = ?;")
if not ok then
    print(stmt)
    return
end

local ok, rows = stmt.query(1)
if not ok then
    print(rows)
    return
end

local ok, result = stmt.close()
if not ok then
    print(result)
    return
end

print(rows[1]["username"])

db.close()
```
