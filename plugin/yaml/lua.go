package yaml

import (
	"fmt"

	. "github.com/r0kyi/glua/core"
	lua "github.com/yuin/gopher-lua"
)

func (y *Yaml) String() string {
	return fmt.Sprintf("glua.yaml: %p", y)
}

func (y *Yaml) AssertFunction() lua.LGFunction {
	return nil
}

func (y *Yaml) MetatableName() string {
	return "lua.table.yaml"
}

func (y *Yaml) encodeL(L *lua.LState) int {
	yaml := L.CheckTable(1)
	y.yaml = LTableToMap(yaml)

	err := y.encode()
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	L.Push(lua.LString(y.raw))
	L.Push(lua.LNil)

	return 2
}

func (y *Yaml) decodeL(L *lua.LState) int {
	raw := L.CheckString(1)
	y.raw = raw

	err := y.decode()
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	L.Push(MapToLTable(L, y.yaml))
	L.Push(lua.LNil)

	return 2
}

func (y *Yaml) Index(L *lua.LState, key string) lua.LValue {
	switch key {
	case "encode":
		return L.NewFunction(y.encodeL)
	case "decode":
		return L.NewFunction(y.decodeL)
	default:
		return lua.LNil
	}
}

func Preload(L *lua.LState) lua.LValue {
	y := &Yaml{}
	ud := NewUserData(L, y)

	return ud
}
