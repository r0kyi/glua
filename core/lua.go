package core

import lua "github.com/r0kyi/gopher-lua"

// NewFunction for AssertFunction
// run without LState's env
func NewFunction(fn lua.LGFunction) *lua.LFunction {
	return &lua.LFunction{
		IsG:       true,
		Proto:     nil,
		GFunction: fn,
	}
}

func SubModIndex(L *lua.LState, key string, objs ...lua.LValue) lua.LValue {
	for _, obj := range objs {
		if v := obj.Index(L, key); v != lua.LNil {
			return v
		}
	}

	return lua.LNil
}
