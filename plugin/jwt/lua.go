package jwt

import (
	"github.com/r0kyi/glua/core"
	lua "github.com/yuin/gopher-lua"
)

func (j *JWT) String() string {
	return "jwt"
}

func (j *JWT) AssertFunction() lua.LGFunction {
	return nil
}

func (j *JWT) MetatableName() string {
	return "lua.table.jwt"
}

func (j *JWT) signL(L *lua.LState) int {
	key := L.CheckString(1)
	alg := L.CheckString(2)
	jwt := L.CheckTable(3)
	j.key = key
	j.alg = alg
	j.jwt = core.LTableToMap(jwt)

	err := j.sign()
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	L.Push(lua.LString(j.raw))
	L.Push(lua.LNil)

	return 2
}

func (j *JWT) verifyL(L *lua.LState) int {
	key := L.CheckString(1)
	raw := L.CheckString(2)
	j.key = key
	j.raw = raw

	err := j.verify()
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	L.Push(core.MapToLTable(L, j.jwt))
	L.Push(lua.LNil)

	return 2
}

func (j *JWT) Index(L *lua.LState, key string) lua.LValue {
	switch key {
	case "sign":
		return L.NewFunction(j.signL)
	case "verify":
		return L.NewFunction(j.verifyL)
	default:
		return lua.LNil
	}
}

func Preload(L *lua.LState) lua.LValue {
	j := &JWT{}
	ud := core.NewUserData(L, j)

	return ud
}
