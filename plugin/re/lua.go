package re

import (
	"github.com/r0kyi/glua/core"
	lua "github.com/yuin/gopher-lua"
)

func (r *Re) String() string {
	return "re"
}

func (r *Re) AssertFunction() lua.LGFunction {
	return nil
}

func (r *Re) MetatableName() string {
	return "lua.table.re"
}

func (r *Re) compileL(L *lua.LState) int {
	r.pattern = L.CheckString(1)

	err := r.compile()
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	L.Push(core.NewUserData(L, r))
	L.Push(lua.LNil)

	return 2
}

func (r *Re) matchStringL(L *lua.LState) int {
	r.src = L.CheckString(1)
	L.Push(lua.LBool(r.matchString()))

	return 1
}

func (r *Re) findStringL(L *lua.LState) int {
	r.src = L.CheckString(1)
	L.Push(lua.LString(r.findString()))

	return 1
}

func (r *Re) findAllStringL(L *lua.LState) int {
	r.src = L.CheckString(1)
	tbl := L.NewTable()

	arr := r.findAllString()
	for _, str := range arr {
		tbl.Append(lua.LString(str))
	}
	L.Push(tbl)

	return 1
}

func (r *Re) replaceAllStringL(L *lua.LState) int {
	r.src = L.CheckString(1)
	r.repl = L.CheckString(2)
	L.Push(lua.LString(r.replaceAllString()))

	return 1
}

func (r *Re) splitL(L *lua.LState) int {
	r.src = L.CheckString(1)
	tbl := L.NewTable()

	arr := r.split()
	for _, str := range arr {
		tbl.Append(lua.LString(str))
	}
	L.Push(tbl)

	return 1
}

func (r *Re) Index(L *lua.LState, key string) lua.LValue {
	switch key {
	case "compile":
		return L.NewFunction(r.compileL)
	case "match":
		return L.NewFunction(r.matchStringL)
	case "find":
		return L.NewFunction(r.findStringL)
	case "find_all":
		return L.NewFunction(r.findAllStringL)
	case "replace":
		return L.NewFunction(r.replaceAllStringL)
	case "split":
		return L.NewFunction(r.splitL)
	default:
		return lua.LNil
	}
}

func Preload(L *lua.LState) lua.LValue {
	r := &Re{}
	ud := core.NewUserData(L, r)

	return ud
}
