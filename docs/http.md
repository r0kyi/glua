# http

http 请求模块

**初始化**

| 参数名称 | 参数类型 | 备注                                                 |
| -------- | -------- | ---------------------------------------------------- |
| headers  | table    | http 请求头，非必须<br />同一 key 下可设置多个 value |
| proxy    | string   | 代理，非必须<br />支持 http/socks                    |
| timeout  | number   | 超时时间，非必须<br />单位：秒                       |

**返回值**

| 参数名称 | 参数类型 | 备注      |
| -------- | -------- | --------- |
| http     | object   | http 对象 |

**demo**

```lua
local http = glua.http

local h = http{
    headers = {
        ["User-Agent"] = {
            "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0",
            "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/140.0.0.0 Safari/537.36",
            "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/26.0 Safari/605.1.15",
        },
        ["Content-Type"] = "text/html; charset=utf-8",
    },
    proxy = "http://127.0.0.1",
    timeout = 5,
}

-- local h = http()
```

## get

get 请求

**参数值**

| 参数名称 | 参数类型 | 备注                       |
| -------- | -------- | -------------------------- |
| url      | string   | 请求链接，必须             |
| args     | table    | 查询参数，非必须，不可重复 |

**返回值**

| 参数名称 | 参数类型 | 备注                                                         |
| -------- | -------- | ------------------------------------------------------------ |
| ok       | boolean  | 函数是否执行成功                                             |
| response | object   | 函数执行成功时返回 **请求响应**，共有三个属性<br />1. status_code: 响应码，int<br />2. headers: 响应头，table<br />3. body: 响应体，string<br />函数执行失败时返回 **错误信息** |

**demo**

```lua
local http = glua.http

local h = http{
    headers = {
        ["User-Agent"] = {
            "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0",
            "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/140.0.0.0 Safari/537.36",
            "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/26.0 Safari/605.1.15",
        },
        ["Content-Type"] = "text/html; charset=utf-8",
    },
    proxy = "http://127.0.0.1",
    timeout = 5,
}

local args = {
    ["username"] = "username",
    ["password"] = "password",
}

local ok, response = h.get("https://www.example.com", args)

if not ok then
    print(response)
    return
end

print(response.status_code)

for k, v in pairs(response.headers) do
    for _, vv in pairs(v) do
        print(k, vv)
    end
end

print(response.body)
```

## post

post 请求

**参数值**

| 参数名称 | 参数类型 | 备注                       |
| -------- | -------- | -------------------------- |
| url      | string   | 请求链接，必须             |
| args     | table    | 查询参数，非必须，不可重复 |
| body     | string   | 请求体，非必须             |

**返回值**

| 参数名称 | 参数类型 | 备注                                                         |
| -------- | -------- | ------------------------------------------------------------ |
| ok       | boolean  | 函数是否执行成功                                             |
| response | object   | 函数执行成功时返回 **请求响应**，共有三个属性<br />1. status_code: 响应码，int<br />2. headers: 响应头，table<br />3. body: 响应体，string<br />函数执行失败时返回 **错误信息** |

**demo**

```lua
local http = glua.http

local h = http{
    headers = {
        ["User-Agent"] = {
            "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0",
            "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/140.0.0.0 Safari/537.36",
            "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/26.0 Safari/605.1.15",
        },
        ["Content-Type"] = "text/html; charset=utf-8",
    },
    proxy = "http://127.0.0.1",
    timeout = 5,
}

local args = {
    ["username"] = "username",
    ["password"] = "password",
}

local ok, response = h.post("https://www.example.com", args, "username=username&password=password&code=123456")

if not ok then
    print(response)
    return
end

print(response.status_code)

for k, v in pairs(response.headers) do
    for _, vv in pairs(v) do
        print(k, vv)
    end
end

print(response.body)
```

## put

put 请求

**参数值**

| 参数名称 | 参数类型 | 备注                       |
| -------- | -------- | -------------------------- |
| url      | string   | 请求链接，必须             |
| args     | table    | 查询参数，非必须，不可重复 |
| body     | string   | 请求体，非必须             |

**返回值**

| 参数名称 | 参数类型 | 备注                                                         |
| -------- | -------- | ------------------------------------------------------------ |
| ok       | boolean  | 函数是否执行成功                                             |
| response | object   | 函数执行成功时返回 **请求响应**，共有三个属性<br />1. status_code: 响应码，int<br />2. headers: 响应头，table<br />3. body: 响应体，string<br />函数执行失败时返回 **错误信息** |

**demo**

```lua
local http = glua.http

local h = http{
    headers = {
        ["User-Agent"] = {
            "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0",
            "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/140.0.0.0 Safari/537.36",
            "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/26.0 Safari/605.1.15",
        },
        ["Content-Type"] = "text/html; charset=utf-8",
    },
    proxy = "http://127.0.0.1",
    timeout = 5,
}

local args = {
    ["username"] = "username",
    ["password"] = "password",
}

local ok, response = h.put("https://www.example.com", args, "username=username&password=password&code=123456")

if not ok then
    print(response)
    return
end

print(response.status_code)

for k, v in pairs(response.headers) do
    for _, vv in pairs(v) do
        print(k, vv)
    end
end

print(response.body)
```

## delete

delete 请求

**参数值**

| 参数名称 | 参数类型 | 备注                       |
| -------- | -------- | -------------------------- |
| url      | string   | 请求链接，必须             |
| args     | table    | 查询参数，非必须，不可重复 |
| body     | string   | 请求体，非必须             |

**返回值**

| 参数名称 | 参数类型 | 备注                                                         |
| -------- | -------- | ------------------------------------------------------------ |
| ok       | boolean  | 函数是否执行成功                                             |
| response | object   | 函数执行成功时返回 **请求响应**，共有三个属性<br />1. status_code: 响应码，int<br />2. headers: 响应头，table<br />3. body: 响应体，string<br />函数执行失败时返回 **错误信息** |

**demo**

```lua
local http = glua.http

local h = http{
    headers = {
        ["User-Agent"] = {
            "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0",
            "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/140.0.0.0 Safari/537.36",
            "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/26.0 Safari/605.1.15",
        },
        ["Content-Type"] = "text/html; charset=utf-8",
    },
    proxy = "http://127.0.0.1",
    timeout = 5,
}

local args = {
    ["username"] = "username",
    ["password"] = "password",
}

local ok, response = h.delete("https://www.example.com", args, "username=username&password=password&code=123456")

if not ok then
    print(response)
    return
end

print(response.status_code)

for k, v in pairs(response.headers) do
    for _, vv in pairs(v) do
        print(k, vv)
    end
end

print(response.body)
```

## patch

patch 请求

**参数值**

| 参数名称 | 参数类型 | 备注                       |
| -------- | -------- | -------------------------- |
| url      | string   | 请求链接，必须             |
| args     | table    | 查询参数，非必须，不可重复 |
| body     | string   | 请求体，非必须             |

**返回值**

| 参数名称 | 参数类型 | 备注                                                         |
| -------- | -------- | ------------------------------------------------------------ |
| ok       | boolean  | 函数是否执行成功                                             |
| response | object   | 函数执行成功时返回 **请求响应**，共有三个属性<br />1. status_code: 响应码，int<br />2. headers: 响应头，table<br />3. body: 响应体，string<br />函数执行失败时返回 **错误信息** |

**demo**

```lua
local http = glua.http

local h = http{
    headers = {
        ["User-Agent"] = {
            "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0",
            "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/140.0.0.0 Safari/537.36",
            "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/26.0 Safari/605.1.15",
        },
        ["Content-Type"] = "text/html; charset=utf-8",
    },
    proxy = "http://127.0.0.1",
    timeout = 5,
}

local args = {
    ["username"] = "username",
    ["password"] = "password",
}

local ok, response = h.patch("https://www.example.com", args, "username=username&password=password&code=123456")

if not ok then
    print(response)
    return
end

print(response.status_code)

for k, v in pairs(response.headers) do
    for _, vv in pairs(v) do
        print(k, vv)
    end
end

print(response.body)
```

## options

options 请求

**参数值**

| 参数名称 | 参数类型 | 备注                       |
| -------- | -------- | -------------------------- |
| url      | string   | 请求链接，必须             |
| args     | table    | 查询参数，非必须，不可重复 |

**返回值**

| 参数名称 | 参数类型 | 备注                                                         |
| -------- | -------- | ------------------------------------------------------------ |
| ok       | boolean  | 函数是否执行成功                                             |
| response | object   | 函数执行成功时返回 **请求响应**，共有两个属性<br />1. status_code: 响应码，int<br />2. headers: 响应头，table<br />函数执行失败时返回 **错误信息** |

**demo**

```lua
local http = glua.http

local h = http{
    headers = {
        ["User-Agent"] = {
            "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0",
            "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/140.0.0.0 Safari/537.36",
            "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/26.0 Safari/605.1.15",
        },
        ["Content-Type"] = "text/html; charset=utf-8",
    },
    proxy = "http://127.0.0.1",
    timeout = 5,
}

local args = {
    ["username"] = "username",
    ["password"] = "password",
}

local ok, response = h.options("https://www.example.com", args)

if not ok then
    print(response)
    return
end

print(response.status_code)

for k, v in pairs(response.headers) do
    for _, vv in pairs(v) do
        print(k, vv)
    end
end
```

## head

head 请求

**参数值**

| 参数名称 | 参数类型 | 备注                       |
| -------- | -------- | -------------------------- |
| url      | string   | 请求链接，必须             |
| args     | table    | 查询参数，非必须，不可重复 |

**返回值**

| 参数名称 | 参数类型 | 备注                                                         |
| -------- | -------- | ------------------------------------------------------------ |
| ok       | boolean  | 函数是否执行成功                                             |
| response | object   | 函数执行成功时返回 **请求响应**，共有两个属性<br />1. status_code: 响应码，int<br />2. headers: 响应头，table<br />函数执行失败时返回 **错误信息** |

**demo**

```lua
local http = glua.http

local h = http{
    headers = {
        ["User-Agent"] = {
            "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0",
            "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/140.0.0.0 Safari/537.36",
            "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/26.0 Safari/605.1.15",
        },
        ["Content-Type"] = "text/html; charset=utf-8",
    },
    proxy = "http://127.0.0.1",
    timeout = 5,
}

local args = {
    ["username"] = "username",
    ["password"] = "password",
}

local ok, response = h.head("https://www.example.com", args)

if not ok then
    print(response)
    return
end

print(response.status_code)

for k, v in pairs(response.headers) do
    for _, vv in pairs(v) do
        print(k, vv)
    end
end
```

