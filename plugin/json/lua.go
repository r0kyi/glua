package json

import (
	"fmt"

	"github.com/r0kyi/glua/core"
	lua "github.com/r0kyi/gopher-lua"
)

func (j *Json) String() string {
	return fmt.Sprintf("glua.json: %p", j)
}

func (j *Json) Type() lua.LValueType {
	return lua.LTObject
}

func (j *Json) AssertFunction() (*lua.LFunction, bool) {
	return nil, false
}

func (j *Json) encodeL(L *lua.LState) int {
	json := L.CheckTable(1)
	json_, err := core.LTableToMap[any](json)

	raw, err := j.encode(json_)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	L.Push(lua.LString(raw))
	L.Push(lua.LNil)

	return 2
}

func (j *Json) decodeL(L *lua.LState) int {
	raw := L.CheckString(1)

	json, err := j.decode(raw)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	L.Push(core.MapToLTable(L, json))
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

func Preload() lua.LValue {
	j := &Json{}

	return j
}
