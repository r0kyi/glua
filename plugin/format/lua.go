package format

import (
	"fmt"

	"github.com/r0kyi/glua/core"
	lua "github.com/r0kyi/gopher-lua"
)

func (f *Format) String() string {
	return fmt.Sprintf(f.format, f.args...)
}

func (f *Format) Type() lua.LValueType {
	return lua.LTObject
}

func (f *Format) AssertFunction() (*lua.LFunction, bool) {
	return core.NewFunction(newFormatL), true
}

func (f *Format) Index(L *lua.LState, key string) lua.LValue {
	switch key {
	default:
		return lua.LNil
	}
}

func newFormatL(L *lua.LState) int {
	if L.GetTop() < 1 {
		L.Push(lua.LNil)
		return 1
	}

	format := L.CheckString(1)

	var args []any
	for i := 2; i <= L.GetTop(); i++ {
		args = append(args, L.CheckAny(i))
	}

	f := &Format{
		format: format,
		args:   args,
	}
	L.Push(f)

	return 1
}

func Preload() lua.LValue {
	f := &Format{}

	return f
}
