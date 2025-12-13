package http

import (
	"fmt"

	"github.com/r0kyi/glua/core"
	lua "github.com/r0kyi/gopher-lua"
)

func (r *Response) String() string {
	return fmt.Sprintf("glua.http.response: %p", r)
}

func (r *Response) Type() lua.LValueType {
	return lua.LTObject
}

func (r *Response) AssertFunction() (*lua.LFunction, bool) {
	return nil, false
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
