package core

import lua "github.com/yuin/gopher-lua"

type LValueObject interface {
	String() string
	AssertFunction() lua.LGFunction
	MetatableName() string
	Index(L *lua.LState, key string) lua.LValue
}

func RegisterMetatable(L *lua.LState, mtName string) {
	if L.GetTypeMetatable(mtName) != lua.LNil {
		return
	}
	mt := L.NewTypeMetatable(mtName)

	L.SetField(mt, "__index", L.NewFunction(func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		key := L.CheckString(2)
		if obj, ok := ud.Value.(LValueObject); ok {
			L.Push(obj.Index(L, key))
			return 1
		}
		L.Push(lua.LNil)
		return 1
	}))

	L.SetField(mt, "__tostring", L.NewFunction(func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		if obj, ok := ud.Value.(LValueObject); ok {
			L.Push(lua.LString(obj.String()))
			return 1
		}
		L.Push(lua.LString("<userdata>"))
		return 1
	}))

	L.SetField(mt, "__call", L.NewFunction(func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		if obj, ok := ud.Value.(LValueObject); ok {
			return obj.AssertFunction()(L)
		}
		return 0
	}))
}

func SubModIndex(L *lua.LState, key string, objs ...LValueObject) lua.LValue {
	for _, obj := range objs {
		if v := obj.Index(L, key); v != lua.LNil {
			return v
		}
	}
	return lua.LNil
}

func NewUserData(L *lua.LState, obj LValueObject) *lua.LUserData {
	RegisterMetatable(L, obj.MetatableName())
	ud := L.NewUserData()
	ud.Value = obj
	L.SetMetatable(ud, L.GetTypeMetatable(obj.MetatableName()))
	return ud
}
