package http

import (
	"github.com/r0kyi/glua/core"
	lua "github.com/yuin/gopher-lua"
)

func (r *Response) String() string {
	return "http.response"
}

func (r *Response) AssertFunction() lua.LGFunction {
	return nil
}

func (r *Response) MetatableName() string {
	return "lua.table.http.response"
}

func (r *Response) Index(L *lua.LState, key string) lua.LValue {
	switch key {
	case "status_code":
		return lua.LNumber(r.statusCode)
	case "headers":
		return core.MapToLTable(L, r.headers)
	case "body":
		return lua.LString(r.body)
	default:
		return lua.LNil
	}
}
