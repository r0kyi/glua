package format

import (
	"fmt"

	. "github.com/r0kyi/glua/core"
	lua "github.com/yuin/gopher-lua"
)

func (f *Format) String() string {
	return fmt.Sprintf(f.format, f.args...)
}

func (f *Format) AssertFunction() lua.LGFunction {
	return NewFormatL
}

func (f *Format) MetatableName() string {
	return "lua.table.format"
}

func (f *Format) Index(L *lua.LState, key string) lua.LValue {
	switch key {
	default:
		return lua.LNil
	}
}

func NewFormatL(L *lua.LState) int {
	if L.GetTop() < 3 {
		L.Push(lua.LNil)
		return 1
	}

	format := L.CheckString(2)

	var args []any
	for i := 3; i <= L.GetTop(); i++ {
		args = append(args, L.CheckAny(i))
	}

	f := &Format{format: format, args: args}
	ud := NewUserData(L, f)
	L.Push(ud)

	return 1
}

func Preload(L *lua.LState) lua.LValue {
	f := &Format{}
	ud := NewUserData(L, f)

	return ud
}
