package xml

import (
	"github.com/r0kyi/glua/core"
	lua "github.com/yuin/gopher-lua"
)

func (x *Xml) String() string {
	return "xml"
}

func (x *Xml) AssertFunction() lua.LGFunction {
	return nil
}

func (x *Xml) MetatableName() string {
	return "lua.table.xml"
}

func (x *Xml) encodeL(L *lua.LState) int {
	xml := L.CheckTable(1)
	x.xml = core.LTableToMap(xml)

	err := x.encode()
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	L.Push(lua.LString(x.raw))
	L.Push(lua.LNil)

	return 2
}

func (x *Xml) decodeL(L *lua.LState) int {
	raw := L.CheckString(1)
	x.raw = raw

	err := x.decode()
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	L.Push(core.MapToLTable(L, x.xml))
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

func Preload(L *lua.LState) lua.LValue {
	x := &Xml{}
	ud := core.NewUserData(L, x)

	return ud
}
