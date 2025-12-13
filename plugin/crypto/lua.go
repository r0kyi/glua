package crypto

import (
	"fmt"

	"github.com/r0kyi/glua/core"
	lua "github.com/r0kyi/gopher-lua"
)

func (c *Crypto) String() string {
	return fmt.Sprintf("glua.crypto: %p", c)
}

func (c *Crypto) Type() lua.LValueType {
	return lua.LTObject
}

func (c *Crypto) AssertFunction() (*lua.LFunction, bool) {
	return nil, false
}

func (c *Crypto) Index(L *lua.LState, key string) lua.LValue {
	switch key {
	case "aes":
		return c.aes
	default:
		return core.SubModIndex(L, key, c.hash)
	}
}

func Preload() lua.LValue {
	c := &Crypto{}

	return c
}
