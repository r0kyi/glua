package crypto

import (
	"github.com/r0kyi/glua/core"
	lua "github.com/yuin/gopher-lua"
)

func (c *Crypto) String() string {
	return "glua.crypto"
}

func (c *Crypto) AssertFunction() lua.LGFunction {
	return nil
}

func (c *Crypto) MetatableName() string {
	return "lua.table.crypto"
}

func (c *Crypto) Index(L *lua.LState, key string) lua.LValue {
	switch key {
	case "aes":
		return core.NewUserData(L, c.aes)
	}

	return core.SubModIndex(L, key, c.hash)
}

func Preload(L *lua.LState) lua.LValue {
	c := &Crypto{}
	ud := core.NewUserData(L, c)
	return ud
}
