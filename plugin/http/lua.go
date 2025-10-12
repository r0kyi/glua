package http

import (
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/r0kyi/glua/core"
	lua "github.com/yuin/gopher-lua"
)

func (h *Http) String() string { return "http" }
func (h *Http) AssertFunction() lua.LGFunction {
	return NewHttpL
}

func (h *Http) MetatableName() string {
	return "lua.table.http"
}

func (h *Http) getL(L *lua.LState) int {
	h.url = L.CheckString(1)
	err := h.get()

	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	L.Push(core.NewUserData(L, h.response))
	L.Push(lua.LNil)

	return 2
}

func (h *Http) postL(L *lua.LState) int {
	body := h.Body
	h.url = L.CheckString(1)

	if L.GetTop() == 2 {
		h.Body = L.CheckString(2)
	}

	err := h.post()
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	h.Body = body
	L.Push(core.NewUserData(L, h.response))
	L.Push(lua.LNil)

	return 2
}

func (h *Http) putL(L *lua.LState) int {
	body := h.Body
	h.url = L.CheckString(1)

	if L.GetTop() == 2 {
		h.Body = L.CheckString(2)
	}

	err := h.put()
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	h.Body = body
	L.Push(core.NewUserData(L, h.response))
	L.Push(lua.LNil)

	return 2
}

func (h *Http) deleteL(L *lua.LState) int {
	body := h.Body
	h.url = L.CheckString(1)

	if L.GetTop() == 2 {
		h.Body = L.CheckString(2)
	}

	err := h.delete()
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	h.Body = body
	L.Push(core.NewUserData(L, h.response))
	L.Push(lua.LNil)

	return 2
}

func (h *Http) patchL(L *lua.LState) int {
	body := h.Body
	h.url = L.CheckString(1)

	if L.GetTop() == 2 {
		h.Body = L.CheckString(2)
	}

	err := h.patch()
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	h.Body = body
	L.Push(core.NewUserData(L, h.response))
	L.Push(lua.LNil)

	return 2
}

func (h *Http) optionsL(L *lua.LState) int {
	h.url = L.CheckString(1)

	err := h.options()
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	L.Push(core.NewUserData(L, h.response))
	L.Push(lua.LNil)

	return 2
}

func (h *Http) headL(L *lua.LState) int {
	h.url = L.CheckString(1)

	err := h.head()
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	L.Push(core.NewUserData(L, h.response))
	L.Push(lua.LNil)

	return 2
}

func (h *Http) Index(L *lua.LState, key string) lua.LValue {
	switch key {
	case "get":
		return L.NewFunction(h.getL)
	case "post":
		return L.NewFunction(h.postL)
	case "put":
		return L.NewFunction(h.putL)
	case "delete":
		return L.NewFunction(h.deleteL)
	case "patch":
		return L.NewFunction(h.patchL)
	case "options":
		return L.NewFunction(h.optionsL)
	case "head":
		return L.NewFunction(h.headL)
	case "headers":
		return core.MapToLTable(L, h.Headers)
	case "args":
		return core.MapToLTable(L, h.Args)
	case "body":
		return lua.LString(h.Body)
	case "proxy":
		return lua.LString(h.Proxy)
	case "timeout":
		return lua.LNumber(h.Timeout)
	default:
		return lua.LNil
	}
}

func NewHttpL(L *lua.LState) int {
	ud := core.NewUserData(L, &Http{
		client: resty.New(),
		response: &Response{
			headers: make(map[string][]string),
		},
	})
	h := ud.Value.(*Http)

	if tbl, ok := L.Get(2).(*lua.LTable); ok {
		_ = core.LTableToStrut(tbl, h)
	}
	if h.Timeout > 0 {
		h.Timeout = h.Timeout * time.Second
	}
	for k, v := range h.Headers {
		for _, vv := range v {
			h.client.SetHeader(k, vv)
		}
	}
	if h.Proxy != "" {
		h.client.SetProxy(h.Proxy)
	}

	L.Push(ud)

	return 1
}

func Preload(L *lua.LState) lua.LValue {
	h := &Http{}
	ud := core.NewUserData(L, h)

	return ud
}
