# web

web 服务模块

**初始化**

| 参数名称        | 参数类型 | 备注                                                         |
| --------------- | -------- | ------------------------------------------------------------ |
| addr            | string   | web 服务监听的 ip 和端口<br />格式：ip:port，必须            |
| mode            | string   | 模式，非必须，可选值：<br />1. debug，默认<br />2. release<br />3. test |
| pattern         | string   | html 渲染模板的路径，非必须                                  |
| static          | table    | 目录映射到一个 url 前缀                                      |
| trusted_proxies | table    | 设置信任代理 ip                                              |

**demo**

```lua
local web = glua.web

local r = web{
    addr = "127.0.0.1:8080",
    mode = "debug",
    pattern = "templates/*",
    static = {
        ["/assets"] = "./static",
        ["/images"] = "./images"
    },
  	trusted_proxies = {"127.0.0.1", "192.168.1.1"},
}

r.get("/", function(c)
    c.html(200, "index.tmpl", {
        ["title"] = "Main website",
    })
end)

r.run()
```

## get

路由注册方法，声明一个 get 请求的处理逻辑

**参数值**

| 参数名称 | 参数类型 | 备注                                                         |
| -------- | -------- | ------------------------------------------------------------ |
| path     | string   | 路径，必须                                                   |
| fn       | function | 处理函数，必须<br />参数为 context，类型为 userdata<br />方法请看下方的 context 部分 |

**demo**

```lua
local web = glua.web

local r = web{
    addr = "127.0.0.1:8080",
    mode = "debug"
}

r.get("/", function(c)
    c.json(200, {
        ["code"] = 0,
        ["msg"] = "hello",
    })
end)

r.run()
```

## post

路由注册方法，声明一个 post 请求的处理逻辑

**参数值**

| 参数名称 | 参数类型 | 备注                                                         |
| -------- | -------- | ------------------------------------------------------------ |
| path     | string   | 路径，必须                                                   |
| fn       | function | 处理函数，必须<br />参数为 context，类型为 userdata<br />方法请看下方的 context 部分 |

**demo**

```lua
local web = glua.web

local r = web{
    addr = "127.0.0.1:8080",
    mode = "debug"
}

r.post("/", function(c)
    c.json(200, {
        ["code"] = 0,
        ["msg"] = "hello",
    })
end)

r.run()
```

## put

路由注册方法，声明一个 put 请求的处理逻辑

**参数值**

| 参数名称 | 参数类型 | 备注                                                         |
| -------- | -------- | ------------------------------------------------------------ |
| path     | string   | 路径，必须                                                   |
| fn       | function | 处理函数，必须<br />参数为 context，类型为 userdata<br />方法请看下方的 context 部分 |

**demo**

```lua
local web = glua.web

local r = web{
    addr = "127.0.0.1:8080",
    mode = "debug"
}

r.put("/", function(c)
    c.json(200, {
        ["code"] = 0,
        ["msg"] = "hello",
    })
end)

r.run()
```

## delete

路由注册方法，声明一个 delete 请求的处理逻辑

**参数值**

| 参数名称 | 参数类型 | 备注                                                         |
| -------- | -------- | ------------------------------------------------------------ |
| path     | string   | 路径，必须                                                   |
| fn       | function | 处理函数，必须<br />参数为 context，类型为 userdata<br />方法请看下方的 context 部分 |

**demo**

```lua
local web = glua.web

local r = web{
    addr = "127.0.0.1:8080",
    mode = "debug"
}

r.delete("/", function(c)
    c.json(200, {
        ["code"] = 0,
        ["msg"] = "hello",
    })
end)

r.run()
```

## patch

路由注册方法，声明一个 patch 请求的处理逻辑

**参数值**

| 参数名称 | 参数类型 | 备注                                                         |
| -------- | -------- | ------------------------------------------------------------ |
| path     | string   | 路径，必须                                                   |
| fn       | function | 处理函数，必须<br />参数为 context，类型为 userdata<br />方法请看下方的 context 部分 |

**demo**

```lua
local web = glua.web

local r = web{
    addr = "127.0.0.1:8080",
    mode = "debug"
}

r.patch("/", function(c)
    c.json(200, {
        ["code"] = 0,
        ["msg"] = "hello",
    })
end)

r.run()
```

## options

路由注册方法，声明一个 options 请求的处理逻辑

**参数值**

| 参数名称 | 参数类型 | 备注                                                         |
| -------- | -------- | ------------------------------------------------------------ |
| path     | string   | 路径，必须                                                   |
| fn       | function | 处理函数，必须<br />参数为 context，类型为 userdata<br />方法请看下方的 context 部分 |

**demo**

```lua
local web = glua.web

local r = web{
    addr = "127.0.0.1:8080",
    mode = "debug"
}

r.options("/", function(c)
    c.string(200, "")
end)

r.run()
```

## head

路由注册方法，声明一个 head 请求的处理逻辑

**参数值**

| 参数名称 | 参数类型 | 备注                                                         |
| -------- | -------- | ------------------------------------------------------------ |
| path     | string   | 路径，必须                                                   |
| fn       | function | 处理函数，必须<br />参数为 context，类型为 userdata<br />方法请看下方的 context 部分 |

**demo**

```lua
local web = glua.web

local r = web{
    addr = "127.0.0.1:8080",
    mode = "debug"
}

r.head("/", function(c)
    c.string(200, "")
end)

r.run()
```

## use

**参数值**

| 参数名称 | 参数类型 | 备注                                           |
| -------- | -------- | ---------------------------------------------- |
| session  | userdata | session，必须<br />方法请看下方的 session 部分 |

**demo**

```lua
local web = glua.web
local session = glua.web.session

local r = web{
    addr = "127.0.0.1:8080"
}

local s = session{
    keys = {"123456", "1234567890123456"},
    name = "session"
}

r.use(s)

r.get("/:name/:id", function(c)
    s.default(c)
    local name = c.get_param("name")
    local id = c.get_param("id")
    s.set("name", name)
    s.set("id", id)
    s.save()
    c.json(200, {
        ["code"] = 0,
        ["msg"] = "good",
    })
end)

r.get("/", function(c)
    s.default(c)
    local name = s.get("name")
    local id = s.get("id")
    c.json(200, {
        ["code"] = 0,
        ["name"] = name,
        ["id"] = id,
    })
end)

r.run()
```

## run

启动 web 服务，使用详情看具体的 demo

## context

请求响应的上下文对象

### json

用于返回 json 响应给客户端

**参数值**

| 参数名称 | 参数类型 | 备注              |
| -------- | -------- | ----------------- |
| code     | number   | 状态码，必须      |
| obj      | table    | 响应的 json，必须 |

**demo**

```lua
local web = glua.web

local r = web{
    addr = "127.0.0.1:8080"
}

r.get("/", function(c)
    c.json(200, {
        ["code"] = 0,
        ["msg"] = "hello",
    })
end)

r.run()
```

### ascii_json

用于返回 json 响应给客户端，将非 ascii 字符（如中文）转义成 `\uxxxx` 格式，保证响应是纯 ascii

**参数值**

| 参数名称 | 参数类型 | 备注              |
| -------- | -------- | ----------------- |
| code     | number   | 状态码，必须      |
| obj      | table    | 响应的 json，必须 |

**demo**

```lua
local web = glua.web

local r = web{
    addr = "127.0.0.1:8080"
}

r.get("/", function(c)
    c.ascii_json(200, {
        ["code"] = 0,
        ["msg"] = "hello",
    })
end)

r.run()
```

### string

纯文本响应

**参数值**

| 参数名称 | 参数类型                                 | 备注                   |
| -------- | ---------------------------------------- | ---------------------- |
| code     | number                                   | 状态码，必须           |
| format   | table                                    | 格式化字符串，必须     |
| values   | nil<br />boolean<br />number<br />string | 任意数量的变量，非必须 |

**demo**

```lua
local web = glua.web

local r = web{
    addr = "127.0.0.1:8080"
}

r.get("/", function(c)
    c.string(200, "hello")
end)

r.get("/hello", function(c)
    c.string(200, "%s", "world")
end)

r.run()
```

### html

返回 html 页面

**参数值**

| 参数名称 | 参数类型 | 备注                     |
| -------- | -------- | ------------------------ |
| code     | number   | 状态码，必须             |
| name     | string   | 模板的名称，必须         |
| obj      | table    | 模板里面对应的参数，必须 |

**demo**

```lua
local web = glua.web

local r = web{
    addr = "127.0.0.1:8080",
    pattern = "templates/*"
}

r.get("/", function(c)
    c.html(200, "index.tmpl", {
        ["title"] = "Main website",
    })
end)

r.run()
```

对应的 `index.tmpl` 为

```html
<html>
  <h1>
    {{ .title }}
  </h1>
</html>
```

### get_cookie

读取客户端请求里携带的 cookie

**参数值**

| 参数名称 | 参数类型 | 备注                     |
| -------- | -------- | ------------------------ |
| name     | string   | 请求 cookie 的名称，必须 |

**返回值**

| 参数名称 | 参数类型 | 备注             |
| -------- | -------- | ---------------- |
| value    | string   | 请求 cookie 的值 |

**demo**

```lua
local web = glua.web

local r = web{
    addr = "127.0.0.1:8080"
}

r.get("/", function(c)
    local session = c.get_cookie("session")
    c.string(200, session)
end)

r.run()
```

### set_cookie

给客户端设置 cookie

**参数值**

| 参数名称 | 参数类型 | 备注                                                         |
| -------- | -------- | ------------------------------------------------------------ |
| name     | string   | 响应 cookie 的名称，必须                                     |
| value    | string   | 响应 cookie 的值，必须                                       |
| maxAge   | number   | 响应 cookie 的有效期，必须                                   |
| path     | string   | 响应 cookie 的作用路径，必须                                 |
| domain   | string   | 响应 cookie 的作用域域名，必须                               |
| secure   | boolean  | 如果为 true，则只有在 https 连接时才会发送该 cookie<br/><br/>如果为 false，http 和 http 都会带上 |
| httpOnly | boolean  | 如果为 true，cookie 只能通过 http 请求访问，js 无法读取（防止 xss 攻击）<br/>如果为 false，js 也能通过 document.cookie 访问 |

**demo**

```lua
local web = glua.web

local r = web{
    addr = "127.0.0.1:8080"
}

r.get("/", function(c)
    c.set_cookie("session", "123456", 3600, "/", "127.0.0.1", false, false)
    c.string(200, "hello")
end)

r.run()
```

### get_header

读取客户端请求里携带的 header

**参数值**

| 参数名称 | 参数类型 | 备注                     |
| -------- | -------- | ------------------------ |
| key      | string   | 请求 header 的键名，必须 |

**返回值**

| 参数名称 | 参数类型 | 备注             |
| -------- | -------- | ---------------- |
| value    | string   | 请求 header 的值 |

**demo**

```lua
local web = glua.web

local r = web{
    addr = "127.0.0.1:8080"
}

r.get("/", function(c)
    local ua = c.get_header("user-agent")
    c.string(200, ua)
end)

r.run()
```

### set_header

设置响应头

**参数值**

| 参数名称 | 参数类型 | 备注                     |
| -------- | -------- | ------------------------ |
| key      | string   | 响应 header 的键名，必须 |
| value    | string   | 响应 header 的值         |

**demo**

```lua
local web = glua.web

local r = web{
    addr = "127.0.0.1:8080"
}

r.get("/", function(c)
    c.set_header("new-header", "hello")
    c.string(200, "good")
end)

r.run()
```

### get_query

获取 url 查询参数

**参数值**

| 参数名称 | 参数类型 | 备注                 |
| -------- | -------- | -------------------- |
| key      | string   | 查询参数的键名，必须 |

**返回值**

| 参数名称 | 参数类型 | 备注                                           |
| -------- | -------- | ---------------------------------------------- |
| value    | string   | 查询参数的值<br />如果参数不存在，返回空字符串 |

**demo**

```lua
local web = glua.web

local r = web{
    addr = "127.0.0.1:8080"
}

r.get("/", function(c)
    local username = c.get_query("username")
    c.string(200, username)
end)

r.run()
```

### get_form

获取表单参数

**参数值**

| 参数名称 | 参数类型 | 备注                 |
| -------- | -------- | -------------------- |
| key      | string   | 表单字段的键名，必须 |

**返回值**

| 参数名称 | 参数类型 | 备注                                           |
| -------- | -------- | ---------------------------------------------- |
| value    | string   | 表单字段的值<br />如果参数不存在，返回空字符串 |

**demo**

```lua
local web = glua.web

local r = web{
    addr = "127.0.0.1:8080"
}

r.post("/", function(c)
    local username = c.get_form("username")
    c.string(200, username)
end)

r.run()
```

### get_param

获取路径参数

**参数值**

| 参数名称 | 参数类型 | 备注                   |
| -------- | -------- | ---------------------- |
| key      | string   | 路径中定义的键名，必须 |

**返回值**

| 参数名称 | 参数类型 | 备注                                             |
| -------- | -------- | ------------------------------------------------ |
| value    | string   | 路径中定义的值<br />如果参数不存在，返回空字符串 |

**demo**

```lua
local web = glua.web

local r = web{
    addr = "127.0.0.1:8080"
}

r.get("/:username/:id", function(c)
    local username = c.get_param("username")
    local id = c.get_param("id")
    c.string(200, "%s: %s", username, id)
end)

r.run()
```

### body

请求体

```lua
local web = glua.web

local r = web{
    addr = "127.0.0.1:8080"
}

r.post("/", function(c)
    c.string(200, c.body)
end)

r.run()
```

### method

请求方法

```lua
local web = glua.web

local r = web{
    addr = "127.0.0.1:8080"
}

r.post("/", function(c)
    c.string(200, c.method)
end)

r.run()
```

### path

请求路径

```lua
local web = glua.web

local r = web{
    addr = "127.0.0.1:8080"
}

r.get("/", function(c)
    c.string(200, c.path)
end)

r.run()
```

### uri

完整的请求

```lua
local web = glua.web

local r = web{
    addr = "127.0.0.1:8080"
}

r.get("/", function(c)
    c.string(200, c.uri)
end)

r.run()
```

## session

跨请求保存用户状态

**初始化**

| 参数名称 | 参数类型 | 备注                                                         |
| -------- | -------- | ------------------------------------------------------------ |
| name     | string   | session 的名称，必须<br />一般会作为 cookie 的 `name` 出现在浏览器里 |
| key      | table    | 密钥，必须<br />在此查看其[定义](https://github.com/gin-contrib/sessions/blob/master/cookie/cookie.go#L12) |

**demo**

```lua
local web = glua.web
local session = glua.web.session

local r = web{
    addr = "127.0.0.1:8080"
}

local s = session{
    keys = {"123456", "1234567890123456"},
    name = "session"
}

r.use(s)
```

### default

每次请求必须调用此函数

**参数值**

| 参数名称 | 参数类型 | 备注                                                    |
| -------- | -------- | ------------------------------------------------------- |
| context  | userdata | 本次请求的上下文，必须<br />方法请看上方的 context 部分 |

**demo**

```lua
local web = glua.web
local session = glua.web.session

local r = web{
    addr = "127.0.0.1:8080"
}

local s = session{
    keys = {"123456", "1234567890123456"},
    name = "session"
}

r.use(s)

r.get("/", function(c)
    s.default(c)
    c.string(200, "hello")
end)

r.run()
```

### get

获取指定 key 的 session 值

**参数值**

| 参数名称 | 参数类型 | 备注                       |
| -------- | -------- | -------------------------- |
| key      | string   | 存在 cookie 里的键名，必须 |

**返回值**

| 参数名称 | 参数类型 | 备注                                                 |
| -------- | -------- | ---------------------------------------------------- |
| value    | string   | 存在 cookie 里的值<br />如果参数不存在，返回空字符串 |

**demo**

```lua
local web = glua.web
local session = glua.web.session

local r = web{
    addr = "127.0.0.1:8080"
}

local s = session{
    keys = {"123456", "1234567890123456"},
    name = "session"
}

r.use(s)

r.get("/", function(c)
    s.default(c)
    local username = s.get("username")
    local id = s.get("id")
    c.string(200, "%s: %s", username, id)
end)

r.run()
```

### set

设置或更新 session 中的 key-value

**参数值**

| 参数名称 | 参数类型 | 备注                       |
| -------- | -------- | -------------------------- |
| key      | string   | 存在 cookie 里的键名，必须 |
| value    | string   | 存在 cookie 里的值，必须   |

**demo**

```lua
local web = glua.web
local session = glua.web.session

local r = web{
    addr = "127.0.0.1:8080"
}

local s = session{
    keys = {"123456", "1234567890123456"},
    name = "session"
}

r.use(s)

r.get("/", function(c)
    s.default(c)
    s.set("username", "admin")
    s.save()
    c.string(200, "hello")
end)

r.run()
```

### delete

删除指定 key 的 session 值

**参数值**

| 参数名称 | 参数类型 | 备注                       |
| -------- | -------- | -------------------------- |
| key      | string   | 存在 cookie 里的键名，必须 |

**demo**

```lua
local web = glua.web
local session = glua.web.session

local r = web{
    addr = "127.0.0.1:8080"
}

local s = session{
    keys = {"123456", "1234567890123456"},
    name = "session"
}

r.use(s)

r.get("/", function(c)
    s.default(c)
    s.delete("username")
    s.save()
    c.string(200, "hello")
end)

r.run()
```

### clear

清空当前 session 中的所有数据

```lua
local web = glua.web
local session = glua.web.session

local r = web{
    addr = "127.0.0.1:8080"
}

local s = session{
    keys = {"123456", "1234567890123456"},
    name = "session"
}

r.use(s)

r.get("/", function(c)
    s.default(c)
    s.clear()
    s.save()
    c.string(200, "hello")
end)

r.run()
```

### save

保存对 session 的修改（如 set、delete、clear）

```lua
local web = glua.web
local session = glua.web.session

local r = web{
    addr = "127.0.0.1:8080"
}

local s = session{
    keys = {"123456", "1234567890123456"},
    name = "session"
}

r.use(s)

r.get("/", function(c)
    s.default(c)
    s.set("username", "admin")
    s.save()
    c.string(200, "hello")
end)

r.run()
```

