package yaml

import (
	"fmt"

	"github.com/r0kyi/glua/core"
	lua "github.com/r0kyi/gopher-lua"
)

func (y *Yaml) String() string {
	return fmt.Sprintf("glua.yaml: %p", y)
}

func (y *Yaml) Type() lua.LValueType {
	return lua.LTObject
}

func (y *Yaml) AssertFunction() (*lua.LFunction, bool) {
	return nil, false
}

func (y *Yaml) encodeL(L *lua.LState) int {
	yaml := L.CheckTable(1)
	yaml_, err := core.LTableToMap[any](yaml)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	raw, err := y.encode(yaml_)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	L.Push(lua.LString(raw))
	L.Push(lua.LNil)

	return 2
}

func (y *Yaml) decodeL(L *lua.LState) int {
	raw := L.CheckString(1)

	yaml, err := y.decode(raw)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	L.Push(core.MapToLTable(L, yaml))
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

func Preload() lua.LValue {
	y := &Yaml{}

	return y
}
