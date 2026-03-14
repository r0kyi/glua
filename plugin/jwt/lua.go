package jwt

import (
	"fmt"

	"github.com/r0kyi/glua/core"
	lua "github.com/r0kyi/gopher-lua"
)

func (j *Jwt) String() string {
	return fmt.Sprintf("glua.jwt: %p", j)
}

func (j *Jwt) Type() lua.LValueType {
	return lua.LTObject
}

func (j *Jwt) AssertFunction() (*lua.LFunction, bool) {
	return nil, false
}

func (j *Jwt) signL(L *lua.LState) int {
	key := L.CheckString(1)
	alg := L.CheckString(2)
	jwt := L.CheckTable(3)

	jwt_, err := core.LTableToMap[any](jwt)
	if err != nil {
		L.Push(lua.LFalse)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	raw, err := j.sign(key, alg, jwt_)
	if err != nil {
		L.Push(lua.LFalse)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	L.Push(lua.LTrue)
	L.Push(lua.LString(raw))

	return 2
}

func (j *Jwt) verifyL(L *lua.LState) int {
	key := L.CheckString(1)
	raw := L.CheckString(2)

	jwt, err := j.verify(key, raw)
	if err != nil {
		L.Push(lua.LFalse)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	L.Push(lua.LTrue)
	L.Push(core.MapToLTable(L, jwt))

	return 2
}

func (j *Jwt) Index(L *lua.LState, key string) lua.LValue {
	switch key {
	case "sign":
		return L.NewFunction(j.signL)
	case "verify":
		return L.NewFunction(j.verifyL)
	default:
		return lua.LNil
	}
}

func Preload() lua.LValue {
	j := &Jwt{}

	return j
}
