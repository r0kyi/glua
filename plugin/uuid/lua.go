package uuid

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/r0kyi/glua/core"
	lua "github.com/r0kyi/gopher-lua"
)

func (u *UUID) String() string {
	return fmt.Sprintf("glua.uuid: %p", u)
}

func (u *UUID) Type() lua.LValueType {
	return lua.LTObject
}

func (u *UUID) AssertFunction() (*lua.LFunction, bool) {
	return core.NewFunction(newUUIDL), true
}

func (u *UUID) Index(L *lua.LState, key string) lua.LValue {
	switch key {
	default:
		return lua.LNil
	}
}

func newUUIDL(L *lua.LState) int {
	var uuid_ uuid.UUID
	var err error
	var version string
	if L.GetTop() != 0 {
		version = L.CheckString(1)
	}

	switch version {
	case "v1":
		uuid_, err = uuid.NewUUID()
	case "v3":
		namespace := L.CheckString(2)
		data := L.CheckString(3)
		uuid_, err = uuid.Parse(namespace)
		uuid_ = uuid.NewMD5(uuid_, core.S2B(data))
	case "v4":
		uuid_ = uuid.New()
	case "v5":
		namespace := L.CheckString(2)
		data := L.CheckString(3)
		uuid_, err = uuid.Parse(namespace)
		uuid_ = uuid.NewSHA1(uuid_, core.S2B(data))
	case "v6":
		uuid_, err = uuid.NewV6()
	case "v7":
		uuid_, err = uuid.NewV7()
	default:
		uuid_ = uuid.New()
	}

	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	L.Push(lua.LString(uuid_.String()))
	L.Push(lua.LNil)

	return 2
}

func Preload() lua.LValue {
	u := &UUID{}

	return u
}
