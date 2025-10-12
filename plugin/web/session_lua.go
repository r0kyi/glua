package web

import (
	"github.com/r0kyi/glua/core"
	lua "github.com/yuin/gopher-lua"
)

func (s *Session) String() string {
	return "web.session"
}

func (s *Session) AssertFunction() lua.LGFunction {
	return NewSessionL
}

func (s *Session) MetatableName() string {
	return "lua.table.web.session"
}

func (s *Session) defaultL(L *lua.LState) int {
	context := L.CheckUserData(1)
	if context == nil {
		return 0
	}
	s.context = context.Value.(*Context).context
	s.default_()

	return 0
}

func (s *Session) getL(L *lua.LState) int {
	key := L.CheckString(1)
	s.sess.key = key
	s.get()
	L.Push(lua.LString(s.sess.value))

	return 1
}

func (s *Session) setL(L *lua.LState) int {
	key := L.CheckString(1)
	value := L.CheckString(2)
	s.sess.key = key
	s.sess.value = value
	s.set()

	return 0
}

func (s *Session) deleteL(L *lua.LState) int {
	key := L.CheckString(1)
	s.sess.key = key
	s.delete()

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

func NewSessionL(L *lua.LState) int {
	ud := core.NewUserData(L, &Session{})
	s := ud.Value.(*Session)

	if tbl, ok := L.Get(2).(*lua.LTable); ok {
		_ = core.LTableToStrut(tbl, s)
	}
	s.newStore()
	L.Push(ud)

	return 1
}
