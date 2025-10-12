package crypto

import (
	"errors"

	"github.com/r0kyi/glua/core"
	lua "github.com/yuin/gopher-lua"
)

func (a *Aes) String() string {
	return "glua.crypto.aes"
}

func (a *Aes) AssertFunction() lua.LGFunction {
	return NewAesL
}

func (a *Aes) MetatableName() string {
	return "lua.table.crypto.aes"
}

func (a *Aes) EncryptL(L *lua.LState) int {
	a.plaintext = L.CheckString(1)
	var err error
	switch a.Mode {
	case "cbc":
		err = a.cbcEncrypt()
	case "cfb":
		err = a.cfbEncrypt()
	case "ofb":
		err = a.ofbEncrypt()
	case "ctr":
		err = a.ctrEncrypt()
	case "gcm":
		err = a.gcmEncrypt()
		if err != nil {
			L.Push(lua.LNil)
			L.Push(lua.LNil)
			L.Push(lua.LString(err.Error()))
			return 3
		}

		L.Push(lua.LString(a.ciphertext))
		L.Push(lua.LString(a.tag))
		L.Push(lua.LNil)

		return 3
	case "ecb":
		err = a.ecbEncrypt()
	default:
		err = errors.New("mode: " + a.Mode + " not supported")
	}
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	L.Push(lua.LString(a.ciphertext))
	L.Push(lua.LNil)

	return 2
}

func (a *Aes) DecryptL(L *lua.LState) int {
	a.ciphertext = L.CheckString(1)
	var err error
	switch a.Mode {
	case "cbc":
		err = a.cbcDecrypt()
	case "cfb":
		err = a.cfbDecrypt()
	case "ofb":
		err = a.ofbDecrypt()
	case "ctr":
		err = a.ctrDecrypt()
	case "gcm":
		a.tag = L.CheckString(2)
		err = a.gcmDecrypt()
	case "ecb":
		err = a.ecbDecrypt()
	default:
		err = errors.New("mode: " + a.Mode + " not supported")
	}
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	L.Push(lua.LString(a.plaintext))
	L.Push(lua.LNil)

	return 2
}

func (a *Aes) Index(L *lua.LState, key string) lua.LValue {
	switch key {
	case "encrypt":
		return L.NewFunction(a.EncryptL)
	case "decrypt":
		return L.NewFunction(a.DecryptL)
	default:
		return lua.LNil
	}
}

func NewAesL(L *lua.LState) int {
	ud := core.NewUserData(L, &Aes{})
	a := ud.Value.(*Aes)

	if tbl, ok := L.Get(2).(*lua.LTable); ok {
		_ = core.LTableToStrut(tbl, a)
	}

	L.Push(ud)

	return 1
}
