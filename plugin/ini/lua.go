package ini

import (
	"fmt"

	"github.com/r0kyi/glua/core"
	lua "github.com/r0kyi/gopher-lua"
)

func (i *Ini) String() string {
	return fmt.Sprintf("glua.ini: %p", i)
}

func (i *Ini) Type() lua.LValueType {
	return lua.LTObject
}

func (i *Ini) AssertFunction() (*lua.LFunction, bool) {
	return nil, false
}

func (i *Ini) loadL(L *lua.LState) int {
	filename := L.CheckString(1)

	cfg, err := i.load(filename)
	if err != nil {
		L.Push(lua.LFalse)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	L.Push(lua.LTrue)
	L.Push(core.MapToLTable(L, cfg))

	return 2
}

func (i *Ini) saveL(L *lua.LState) int {
	filename := L.CheckString(1)
	cfg := L.CheckTable(2)

	cfg_, err := core.LTableToMap[any](cfg)
	if err != nil {
		L.Push(lua.LFalse)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	err = i.save(filename, cfg_)
	if err != nil {
		L.Push(lua.LFalse)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	L.Push(lua.LTrue)

	return 1
}

func (i *Ini) Index(L *lua.LState, key string) lua.LValue {
	switch key {
	case "load":
		return L.NewFunction(i.loadL)
	case "save":
		return L.NewFunction(i.saveL)
	default:
		return lua.LNil
	}
}

func Preload() lua.LValue {
	i := &Ini{}

	return i
}
