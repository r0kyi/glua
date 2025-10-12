package base

import (
	"github.com/r0kyi/glua/core"
	lua "github.com/yuin/gopher-lua"
)

func (b *Base) String() string {
	return "glua.base"
}

func (b *Base) AssertFunction() lua.LGFunction {
	return nil
}

func (b *Base) MetatableName() string {
	return "lua.table.base"
}

func (b *Base) base32EncodeL(L *lua.LState) int {
	b.raw = L.CheckString(1)
	b.base32Encode()
	L.Push(lua.LString(b.encoded))

	return 1
}

func (b *Base) base32DecodeL(L *lua.LState) int {
	b.encoded = L.CheckString(1)
	b.base32Decode()
	L.Push(lua.LString(b.raw))

	return 1
}

func (b *Base) base64EncodeL(L *lua.LState) int {
	b.raw = L.CheckString(1)
	b.base64Encode()
	L.Push(lua.LString(b.encoded))

	return 1
}

func (b *Base) base64DecodeL(L *lua.LState) int {
	b.encoded = L.CheckString(1)
	b.base64Decode()
	L.Push(lua.LString(b.raw))

	return 1
}

func (b *Base) Index(L *lua.LState, key string) lua.LValue {
	switch key {
	case "b32encode":
		return L.NewFunction(b.base32EncodeL)
	case "b32decode":
		return L.NewFunction(b.base32DecodeL)
	case "b64encode":
		return L.NewFunction(b.base64EncodeL)
	case "b64decode":
		return L.NewFunction(b.base64DecodeL)
	default:
		return lua.LNil
	}
}

func Preload(L *lua.LState) lua.LValue {
	b := &Base{}
	ud := core.NewUserData(L, b)

	return ud
}
