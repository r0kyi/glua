package web

import (
	"fmt"

	"github.com/r0kyi/glua/core"
	lua "github.com/r0kyi/gopher-lua"
)

func (s *Session) String() string {
	return fmt.Sprintf("glua.web.session: %p", s)
}

func (s *Session) Type() lua.LValueType {
	return lua.LTObject
}

func (s *Session) AssertFunction() (*lua.LFunction, bool) {
	return core.NewFunction(newSessionL), true
}

func (s *Session) defaultL(L *lua.LState) int {
	context := L.CheckAny(1)
	if context == nil {
		return 0
	}

	context_ := context.(*Context).context
	s.default_(context_)

	return 0
}

func (s *Session) getL(L *lua.LState) int {
	key := L.CheckString(1)

	L.Push(lua.LString(s.get(key)))

	return 1
}

func (s *Session) setL(L *lua.LState) int {
	key := L.CheckString(1)
	value := L.CheckString(2)

	s.set(key, value)

	return 0
}

func (s *Session) deleteL(L *lua.LState) int {
	key := L.CheckString(1)

	s.delete(key)

	return 0
}

func (s *Session) clearL(L *lua.LState) int {
	s.clear()

	return 0
}

func (s *Session) saveL(L *lua.LState) int {
	err := s.save()
	if err != nil {
		L.Push(lua.LString(err.Error()))
		return 1
	}

	L.Push(lua.LNil)

	return 1
}

func (s *Session) Index(L *lua.LState, key string) lua.LValue {
	switch key {
	case "default":
		return L.NewFunction(s.defaultL)
	case "get":
		return L.NewFunction(s.getL)
	case "set":
		return L.NewFunction(s.setL)
	case "delete":
		return L.NewFunction(s.deleteL)
	case "clear":
		return L.NewFunction(s.clearL)
	case "save":
		return L.NewFunction(s.saveL)
	default:
		return lua.LNil
	}
}

func newSessionL(L *lua.LState) int {
	s := &Session{}

	tbl := L.CheckTable(1)
	_ = core.LTableToStrut(tbl, s)
	s.newStore()
	L.Push(s)

	return 1
}
