package json

import (
	"github.com/r0kyi/glua/core"
	lua "github.com/yuin/gopher-lua"
)

func (j *Json) String() string {
	return "json"
}

func (j *Json) AssertFunction() lua.LGFunction {
	return nil
}

func (j *Json) MetatableName() string {
	return "lua.table.json"
}

func (j *Json) encodeL(L *lua.LState) int {
	json := L.CheckTable(1)
	j.json = core.LTableToMap(json)

	err := j.encode()
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	L.Push(lua.LString(j.raw))
	L.Push(lua.LNil)

	return 2
}

func (j *Json) decodeL(L *lua.LState) int {
	raw := L.CheckString(1)
	j.raw = raw

	err := j.decode()
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	L.Push(core.MapToLTable(L, j.json))
	L.Push(lua.LNil)

	return 2
}
func (j *Json) Index(L *lua.LState, key string) lua.LValue {
	switch key {
	case "encode":
		return L.NewFunction(j.encodeL)
	case "decode":
		return L.NewFunction(j.decodeL)
	default:
		return lua.LNil
	}
}

func Preload(L *lua.LState) lua.LValue {
	j := &Json{}
	ud := core.NewUserData(L, j)

	return ud
}
