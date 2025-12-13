package crypto

import (
	"errors"
	"fmt"

	"github.com/r0kyi/glua/core"
	lua "github.com/r0kyi/gopher-lua"
)

func (a *Aes) String() string {
	return fmt.Sprintf("glua.crypto.aes: %p", a)
}

func (a *Aes) Type() lua.LValueType {
	return lua.LTObject
}

func (a *Aes) AssertFunction() (*lua.LFunction, bool) {
	return core.NewFunction(newAesL), true
}

func (a *Aes) encryptL(L *lua.LState) int {
	plaintext := L.CheckString(1)
	var ciphertext string
	var err error
	switch a.Mode {
	case "cbc":
		ciphertext, err = a.cbcEncrypt(plaintext)
	case "cfb":
		ciphertext, err = a.cfbEncrypt(plaintext)
	case "ofb":
		ciphertext, err = a.ofbEncrypt(plaintext)
	case "ctr":
		ciphertext, err = a.ctrEncrypt(plaintext)
	case "gcm":
		var tag string
		ciphertext, tag, err = a.gcmEncrypt(plaintext)
		if err != nil {
			L.Push(lua.LNil)
			L.Push(lua.LNil)
			L.Push(lua.LString(err.Error()))
			return 3
		}

		L.Push(lua.LString(ciphertext))
		L.Push(lua.LString(tag))
		L.Push(lua.LNil)

		return 3
	case "ecb":
		ciphertext, err = a.ecbEncrypt(plaintext)
	default:
		err = errors.New("mode: " + a.Mode + " not supported")
	}
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	L.Push(lua.LString(ciphertext))
	L.Push(lua.LNil)

	return 2
}

func (a *Aes) decryptL(L *lua.LState) int {
	ciphertext := L.CheckString(1)
	var plaintext string
	var err error
	switch a.Mode {
	case "cbc":
		plaintext, err = a.cbcDecrypt(ciphertext)
	case "cfb":
		plaintext, err = a.cfbDecrypt(ciphertext)
	case "ofb":
		plaintext, err = a.ofbDecrypt(ciphertext)
	case "ctr":
		plaintext, err = a.ctrDecrypt(ciphertext)
	case "gcm":
		tag := L.CheckString(2)
		plaintext, err = a.gcmDecrypt(ciphertext, tag)
	case "ecb":
		plaintext, err = a.ecbDecrypt(ciphertext)
	default:
		err = errors.New("mode: " + a.Mode + " not supported")
	}
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	L.Push(lua.LString(plaintext))
	L.Push(lua.LNil)

	return 2
}

func (a *Aes) Index(L *lua.LState, key string) lua.LValue {
	switch key {
	case "encrypt":
		return L.NewFunction(a.encryptL)
	case "decrypt":
		return L.NewFunction(a.decryptL)
	default:
		return lua.LNil
	}
}

func newAesL(L *lua.LState) int {
	a := &Aes{}

	tbl := L.CheckTable(1)
	_ = core.LTableToStrut(tbl, a)
	L.Push(a)

	return 1
}
