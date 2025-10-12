package web

import (
	"github.com/r0kyi/glua/core"
	lua "github.com/yuin/gopher-lua"
)

func (c *Context) String() string {
	return "web.context"
}

func (c *Context) AssertFunction() lua.LGFunction {
	return nil
}

func (c *Context) MetatableName() string {
	return "lua.table.web.context"
}

func (c *Context) jsonL(L *lua.LState) int {
	statusCode := L.CheckNumber(1)
	obj := L.CheckTable(2)
	c.response.statusCode = int(statusCode)
	c.response.obj = core.LTableToMap(obj)
	c.json()

	return 0
}

func (c *Context) asciiJsonL(L *lua.LState) int {
	statusCode := L.CheckNumber(1)
	obj := L.CheckTable(2)
	c.response.statusCode = int(statusCode)
	c.response.obj = core.LTableToMap(obj)
	c.asciiJson()

	return 0
}

func (c *Context) stringL(L *lua.LState) int {
	statusCode := L.CheckNumber(1)
	format := L.CheckString(2)
	c.response.statusCode = int(statusCode)
	c.response.format = format
	for i := 3; i <= L.GetTop(); i++ {
		c.response.values = append(c.response.values, L.CheckAny(i))
	}
	c.string()

	return 0
}

func (c *Context) htmlL(L *lua.LState) int {
	statusCode := L.CheckNumber(1)
	name := L.CheckString(2)
	obj := L.CheckTable(3)
	c.response.statusCode = int(statusCode)
	c.response.name = name
	c.response.obj = core.LTableToMap(obj)
	c.html()

	return 0
}

func (c *Context) getCookieL(L *lua.LState) int {
	name := L.CheckString(1)
	c.response.name = name
	c.getCookie()
	L.Push(lua.LString(c.cookie.value))

	return 1
}

func (c *Context) setCookieL(L *lua.LState) int {
	name := L.CheckString(1)
	value := L.CheckString(2)
	maxAge := L.CheckInt(3)
	path := L.CheckString(4)
	domain := L.CheckString(5)
	secure := L.CheckBool(6)
	httpOnly := L.CheckBool(7)

	c.cookie.name = name
	c.cookie.value = value
	c.cookie.maxAge = maxAge
	c.cookie.path = path
	c.cookie.domain = domain
	c.cookie.secure = secure
	c.cookie.httpOnly = httpOnly
	c.setCookie()

	return 0
}

func (c *Context) getHeaderL(L *lua.LState) int {
	key := L.CheckString(1)
	c.header.key = key
	c.getHeader()
	L.Push(lua.LString(c.header.value))

	return 1
}

func (c *Context) setHeaderL(L *lua.LState) int {
	key := L.CheckString(1)
	value := L.CheckString(2)
	c.header.key = key
	c.header.value = value
	c.setHeader()

	return 0
}

func (c *Context) getQueryL(L *lua.LState) int {
	key := L.CheckString(1)
	c.query.key = key
	c.getQuery()
	L.Push(lua.LString(c.query.value))

	return 1
}

func (c *Context) getFormL(L *lua.LState) int {
	key := L.CheckString(1)
	c.form.key = key
	c.getForm()
	L.Push(lua.LString(c.form.value))

	return 1
}

func (c *Context) getParamL(L *lua.LState) int {
	key := L.CheckString(1)
	c.param.key = key
	c.getParam()
	L.Push(lua.LString(c.param.value))

	return 1
}

func (c *Context) bodyL() string {
	c.body()

	return c.request.body
}

func (c *Context) methodL() string {
	c.method()

	return c.request.method
}

func (c *Context) pathL() string {
	c.path()

	return c.request.path
}

func (c *Context) uriL() string {
	c.uri()

	return c.request.uri
}

func (c *Context) Index(L *lua.LState, key string) lua.LValue {
	switch key {
	case "json":
		return L.NewFunction(c.jsonL)
	case "ascii_json":
		return L.NewFunction(c.asciiJsonL)
	case "string":
		return L.NewFunction(c.stringL)
	case "html":
		return L.NewFunction(c.htmlL)
	case "get_cookie":
		return L.NewFunction(c.getCookieL)
	case "set_cookie":
		return L.NewFunction(c.setCookieL)
	case "get_header":
		return L.NewFunction(c.getHeaderL)
	case "set_header":
		return L.NewFunction(c.setHeaderL)
	case "get_query":
		return L.NewFunction(c.getQueryL)
	case "get_form":
		return L.NewFunction(c.getFormL)
	case "get_param":
		return L.NewFunction(c.getParamL)
	case "body":
		return lua.LString(c.bodyL())
	case "method":
		return lua.LString(c.methodL())
	case "path":
		return lua.LString(c.pathL())
	case "uri":
		return lua.LString(c.uriL())
	default:
		return lua.LNil
	}
}
