package web

import (
	"fmt"

	"github.com/r0kyi/glua/core"
	lua "github.com/r0kyi/gopher-lua"
)

func (c *Context) String() string {
	return fmt.Sprintf("glua.web.context: %p", c)
}

func (c *Context) Type() lua.LValueType {
	return lua.LTObject
}

func (c *Context) AssertFunction() (*lua.LFunction, bool) {
	return nil, false
}

func (c *Context) jsonL(L *lua.LState) int {
	statusCode := L.CheckNumber(1)
	obj := L.CheckTable(2)

	obj_, err := core.LTableToMap[any](obj)
	if err != nil {
		L.RaiseError(err.Error())
		return 0
	}

	c.json(int(statusCode), obj_)

	return 0
}

func (c *Context) asciiJsonL(L *lua.LState) int {
	statusCode := L.CheckNumber(1)
	obj := L.CheckTable(2)

	obj_, err := core.LTableToMap[any](obj)
	if err != nil {
		L.RaiseError(err.Error())
		return 0
	}

	c.asciiJson(int(statusCode), obj_)

	return 0
}

func (c *Context) stringL(L *lua.LState) int {
	statusCode := L.CheckNumber(1)
	format := L.CheckString(2)

	values := make([]any, 0)
	for i := 3; i <= L.GetTop(); i++ {
		values = append(values, L.CheckAny(i))
	}
	c.string(int(statusCode), format, values...)

	return 0
}

func (c *Context) htmlL(L *lua.LState) int {
	statusCode := L.CheckNumber(1)
	name := L.CheckString(2)
	obj := L.CheckTable(3)

	obj_, err := core.LTableToMap[any](obj)
	if err != nil {
		c.json(500, map[string]any{})
		return 0
	}

	c.html(int(statusCode), name, obj_)

	return 0
}

func (c *Context) getCookieL(L *lua.LState) int {
	name := L.CheckString(1)

	L.Push(lua.LString(c.getCookie(name)))

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

	c.setCookie(name, value, maxAge, path, domain, secure, httpOnly)

	return 0
}

func (c *Context) getHeaderL(L *lua.LState) int {
	key := L.CheckString(1)

	L.Push(lua.LString(c.getHeader(key)))

	return 1
}

func (c *Context) setHeaderL(L *lua.LState) int {
	key := L.CheckString(1)
	value := L.CheckString(2)

	c.setHeader(key, value)

	return 0
}

func (c *Context) getQueryL(L *lua.LState) int {
	key := L.CheckString(1)

	L.Push(lua.LString(c.getQuery(key)))

	return 1
}

func (c *Context) getFormL(L *lua.LState) int {
	key := L.CheckString(1)

	L.Push(lua.LString(c.getForm(key)))

	return 1
}

func (c *Context) getParamL(L *lua.LState) int {
	key := L.CheckString(1)

	L.Push(lua.LString(c.getParam(key)))

	return 1
}

func (c *Context) bodyL() string {
	return c.body()
}

func (c *Context) methodL() string {
	return c.method()
}

func (c *Context) pathL() string {
	return c.path()
}

func (c *Context) uriL() string {
	return c.uri()
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
