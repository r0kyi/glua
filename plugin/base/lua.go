package base

import (
	"encoding/base32"
	"encoding/base64"
	"fmt"

	lua "github.com/r0kyi/gopher-lua"
)

func (b *Base) String() string {
	return fmt.Sprintf("glua.base: %p", b)
}

func (b *Base) Type() lua.LValueType {
	return lua.LTObject
}

func (b *Base) AssertFunction() (*lua.LFunction, bool) {
	return nil, false
}

func (b *Base) base32EncodeL(L *lua.LState) int {
	raw := L.CheckString(1)
	encoded := b.base32Encode(raw)

	L.Push(lua.LString(encoded))

	return 1
}

func (b *Base) base32DecodeL(L *lua.LState) int {
	encoded := L.CheckString(1)
	raw := b.base32Decode(encoded)

	L.Push(lua.LString(raw))

	return 1
}

func (b *Base) base64EncodeL(L *lua.LState) int {
	raw := L.CheckString(1)
	encoded := b.base64Encode(raw)

	L.Push(lua.LString(encoded))

	return 1
}

func (b *Base) base64DecodeL(L *lua.LState) int {
	encoded := L.CheckString(1)
	raw := b.base64Decode(encoded)

	L.Push(lua.LString(raw))

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

func Preload() lua.LValue {
	b := &Base{
		b32: base32.StdEncoding,
		b64: base64.StdEncoding,
	}

	return b
}
