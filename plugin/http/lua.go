package http

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/r0kyi/glua/core"
	lua "github.com/r0kyi/gopher-lua"
)

func (h *Http) String() string {
	return fmt.Sprintf("glua.http: %p", h)
}

func (h *Http) Type() lua.LValueType {
	return lua.LTObject
}

func (h *Http) AssertFunction() (*lua.LFunction, bool) {
	return core.NewFunction(newHttpL), true
}

func (h *Http) getL(L *lua.LState) int {
	url, args, _, err := getParams(L)
	if err != nil {
		L.Push(lua.LFalse)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	r, err := h.get(url, args)

	if err != nil {
		L.Push(lua.LFalse)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	L.Push(lua.LTrue)
	L.Push(r)

	return 2
}

func (h *Http) postL(L *lua.LState) int {
	url, args, body, err := getParams(L)
	if err != nil {
		L.Push(lua.LFalse)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	r, err := h.post(url, args, body)
	if err != nil {
		L.Push(lua.LFalse)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	L.Push(lua.LTrue)
	L.Push(r)

	return 2
}

func (h *Http) putL(L *lua.LState) int {
	url, args, body, err := getParams(L)
	if err != nil {
		L.Push(lua.LFalse)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	r, err := h.put(url, args, body)
	if err != nil {
		L.Push(lua.LFalse)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	L.Push(lua.LTrue)
	L.Push(r)

	return 2
}

func (h *Http) deleteL(L *lua.LState) int {
	url, args, body, err := getParams(L)
	if err != nil {
		L.Push(lua.LFalse)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	r, err := h.delete(url, args, body)
	if err != nil {
		L.Push(lua.LFalse)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	L.Push(lua.LTrue)
	L.Push(r)

	return 2
}

func (h *Http) patchL(L *lua.LState) int {
	url, args, body, err := getParams(L)
	if err != nil {
		L.Push(lua.LFalse)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	r, err := h.patch(url, args, body)
	if err != nil {
		L.Push(lua.LFalse)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	L.Push(lua.LTrue)
	L.Push(r)

	return 2
}

func (h *Http) optionsL(L *lua.LState) int {
	url, args, _, err := getParams(L)
	if err != nil {
		L.Push(lua.LFalse)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	r, err := h.options(url, args)
	if err != nil {
		L.Push(lua.LFalse)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	L.Push(lua.LTrue)
	L.Push(r)

	return 2
}

func (h *Http) headL(L *lua.LState) int {
	url, args, _, err := getParams(L)
	if err != nil {
		L.Push(lua.LFalse)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	r, err := h.head(url, args)
	if err != nil {
		L.Push(lua.LFalse)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	L.Push(lua.LTrue)
	L.Push(r)

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
	default:
		return lua.LNil
	}
}

func getParams(L *lua.LState) (string, map[string]string, string, error) {
	var url string
	args := make(map[string]string)
	var body string
	var err error

	if L.GetTop() == 1 {
		url = L.CheckString(1)
	} else if L.GetTop() == 2 {
		url = L.CheckString(1)
		tbl := L.CheckTable(2)

		args, err = core.LTableToMap[string](tbl)
		if err != nil {
			return "", nil, "", err
		}
	} else if L.GetTop() == 3 {
		url = L.CheckString(1)
		tbl := L.CheckTable(2)
		body = L.CheckString(3)

		args, err = core.LTableToMap[string](tbl)
		if err != nil {
			return "", nil, "", err
		}
	}

	return url, args, body, nil
}

func newHttpL(L *lua.LState) int {
	h := &Http{
		client: resty.New(),
	}

	if L.GetTop() == 1 {
		tbl := L.CheckTable(1)
		_ = core.LTableToStrut(tbl, h)

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
	}

	L.Push(h)

	return 1
}

func Preload() lua.LValue {
	h := &Http{}

	return h
}
