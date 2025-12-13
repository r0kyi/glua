package xml

import (
	"fmt"

	"github.com/r0kyi/glua/core"
	lua "github.com/r0kyi/gopher-lua"
)

func (x *Xml) String() string {
	return fmt.Sprintf("glua.xml: %p", x)
}

func (x *Xml) Type() lua.LValueType {
	return lua.LTObject
}

func (x *Xml) AssertFunction() (*lua.LFunction, bool) {
	return nil, false
}

func (x *Xml) encodeL(L *lua.LState) int {
	xml := L.CheckTable(1)
	xml_, err := core.LTableToMap[any](xml)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	raw, err := x.encode(xml_)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	L.Push(lua.LString(raw))
	L.Push(lua.LNil)

	return 2
}

func (x *Xml) decodeL(L *lua.LState) int {
	raw := L.CheckString(1)

	xml, err := x.decode(raw)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	L.Push(core.MapToLTable(L, xml))
	L.Push(lua.LNil)

	return 2
}

func (x *Xml) Index(L *lua.LState, key string) lua.LValue {
	switch key {
	case "encode":
		return L.NewFunction(x.encodeL)
	case "decode":
		return L.NewFunction(x.decodeL)
	default:
		return lua.LNil
	}
}

func Preload() lua.LValue {
	x := &Xml{}

	return x
}
