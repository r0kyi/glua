package re

import (
	"fmt"
	"regexp"

	"github.com/r0kyi/glua/core"
	lua "github.com/r0kyi/gopher-lua"
)

func (r *Re) String() string {
	return fmt.Sprintf("glua.re: %p", r)
}

func (r *Re) Type() lua.LValueType {
	return lua.LTObject
}

func (r *Re) AssertFunction() (*lua.LFunction, bool) {
	return core.NewFunction(newReL), true
}

func (r *Re) matchStringL(L *lua.LState) int {
	src := L.CheckString(1)

	L.Push(lua.LBool(r.matchString(src)))

	return 1
}

func (r *Re) findStringL(L *lua.LState) int {
	src := L.CheckString(1)

	L.Push(lua.LString(r.findString(src)))

	return 1
}

func (r *Re) findAllStringL(L *lua.LState) int {
	src := L.CheckString(1)
	tbl := L.NewTable()

	arr := r.findAllString(src)
	for _, str := range arr {
		tbl.Append(lua.LString(str))
	}

	L.Push(tbl)

	return 1
}

func (r *Re) replaceAllStringL(L *lua.LState) int {
	src := L.CheckString(1)
	repl := L.CheckString(2)

	L.Push(lua.LString(r.replaceAllString(src, repl)))

	return 1
}

func (r *Re) splitL(L *lua.LState) int {
	src := L.CheckString(1)
	tbl := L.NewTable()

	arr := r.split(src)
	for _, str := range arr {
		tbl.Append(lua.LString(str))
	}

	L.Push(tbl)

	return 1
}

func (r *Re) Index(L *lua.LState, key string) lua.LValue {
	switch key {
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

func newReL(L *lua.LState) int {
	r := &Re{}

	pattern := L.CheckString(1)
	compiled, err := regexp.Compile(pattern)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	r.regexp = compiled

	L.Push(r)
	L.Push(lua.LNil)

	return 2
}

func Preload() lua.LValue {
	r := &Re{}

	return r
}
