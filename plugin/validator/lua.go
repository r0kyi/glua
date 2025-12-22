package validator

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	lua "github.com/r0kyi/gopher-lua"
)

func (v *Validator) String() string {
	return fmt.Sprintf("glua.validator: %p", v)
}

func (v *Validator) Type() lua.LValueType {
	return lua.LTObject
}

func (v *Validator) AssertFunction() (*lua.LFunction, bool) {
	return nil, false
}

func (v *Validator) validL(tag string) lua.LGFunction {
	return func(L *lua.LState) int {
		value := L.CheckString(1)

		L.Push(lua.LBool(v.valid(value, tag)))
		return 1
	}
}

func (v *Validator) Index(L *lua.LState, key string) lua.LValue {
	switch key {
	case "email", "uuid", "ip", "ipv4", "ipv6", "url", "hostname", "mac", "alpha", "alphanum", "cidr", "cidrv4", "cidrv6":
		return L.NewFunction(v.validL(key))
	default:
		return lua.LNil
	}
}

func Preload() lua.LValue {
	v := &Validator{
		validate: validator.New(),
	}

	return v
}
