package database

import (
	"fmt"

	"github.com/r0kyi/glua/core"
	lua "github.com/r0kyi/gopher-lua"
)

func (s *Statement) String() string {
	return fmt.Sprintf("glua.database.statement: %p", s)
}

func (s *Statement) Type() lua.LValueType {
	return lua.LTObject
}

func (s *Statement) AssertFunction() (*lua.LFunction, bool) {
	return nil, false
}

func (s *Statement) queryL(L *lua.LState) int {
	var args []any
	for i := 1; i <= L.GetTop(); i++ {
		args = append(args, L.CheckAny(i))
	}

	rows, err := s.query(args)
	if err != nil {
		L.Push(lua.LFalse)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	tbl := core.MapToLTable(L, rows)
	L.Push(lua.LTrue)
	L.Push(tbl)

	return 2
}

func (s *Statement) closeL(L *lua.LState) int {
	err := s.close()
	if err != nil {
		L.Push(lua.LFalse)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	L.Push(lua.LTrue)

	return 1
}

func (s *Statement) Index(L *lua.LState, key string) lua.LValue {
	switch key {
	case "query":
		return L.NewFunction(s.queryL)
	case "close":
		return L.NewFunction(s.closeL)
	default:
		return lua.LNil
	}
}
