package web

import (
	"github.com/gin-gonic/gin"
	"github.com/r0kyi/glua/core"
	lua "github.com/yuin/gopher-lua"
)

func (w *Web) String() string {
	return "web"
}

func (w *Web) AssertFunction() lua.LGFunction {
	return NewWebL
}

func (w *Web) MetatableName() string {
	return "lua.table.web"
}

func (w *Web) getL(L *lua.LState) int {
	path := L.CheckString(1)
	fn := L.CheckFunction(2)
	w.route.path = path
	w.route.fn = w.toHandler(L, fn)
	w.get()

	return 0
}

func (w *Web) postL(L *lua.LState) int {
	path := L.CheckString(1)
	fn := L.CheckFunction(2)
	w.route.path = path
	w.route.fn = w.toHandler(L, fn)
	w.post()

	return 0
}

func (w *Web) putL(L *lua.LState) int {
	path := L.CheckString(1)
	fn := L.CheckFunction(2)
	w.route.path = path
	w.route.fn = w.toHandler(L, fn)
	w.put()

	return 0
}

func (w *Web) deleteL(L *lua.LState) int {
	path := L.CheckString(1)
	fn := L.CheckFunction(2)
	w.route.path = path
	w.route.fn = w.toHandler(L, fn)
	w.delete()

	return 0
}

func (w *Web) patchL(L *lua.LState) int {
	path := L.CheckString(1)
	fn := L.CheckFunction(2)
	w.route.path = path
	w.route.fn = w.toHandler(L, fn)
	w.patch()

	return 0
}

func (w *Web) optionsL(L *lua.LState) int {
	path := L.CheckString(1)
	fn := L.CheckFunction(2)
	w.route.path = path
	w.route.fn = w.toHandler(L, fn)
	w.options()

	return 0
}

func (w *Web) headL(L *lua.LState) int {
	path := L.CheckString(1)
	fn := L.CheckFunction(2)
	w.route.path = path
	w.route.fn = w.toHandler(L, fn)
	w.head()

	return 0
}

func (w *Web) useL(L *lua.LState) int {
	session := L.CheckUserData(1)
	if session == nil {
		return 0
	}
	w.session = session.Value.(*Session)
	w.use()

	return 0
}

func (w *Web) runL(L *lua.LState) int {
	err := w.run()
	if err != nil {
		L.Push(lua.LString(err.Error()))
		return 1
	}

	L.Push(lua.LNil)
	return 1
}

func (w *Web) Index(L *lua.LState, key string) lua.LValue {
	switch key {
	case "get":
		return L.NewFunction(w.getL)
	case "post":
		return L.NewFunction(w.postL)
	case "put":
		return L.NewFunction(w.putL)
	case "delete":
		return L.NewFunction(w.deleteL)
	case "patch":
		return L.NewFunction(w.patchL)
	case "options":
		return L.NewFunction(w.optionsL)
	case "head":
		return L.NewFunction(w.headL)
	case "use":
		return L.NewFunction(w.useL)
	case "run":
		return L.NewFunction(w.runL)
	case "session":
		return core.NewUserData(L, w.session)
	default:
		return lua.LNil
	}
}

func NewWebL(L *lua.LState) int {
	ud := core.NewUserData(L, &Web{
		context: &Context{},
	})
	w := ud.Value.(*Web)

	if tbl, ok := L.Get(2).(*lua.LTable); ok {
		_ = core.LTableToStrut(tbl, w)
	}
	switch w.Mode {
	case "debug":
		gin.SetMode(gin.DebugMode)
	case "release":
		gin.SetMode(gin.ReleaseMode)
	case "test":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.DebugMode)
	}
	w.engine = gin.Default()
	if w.Pattern != "" {
		w.engine.LoadHTMLGlob(w.Pattern)
	}
	for k, v := range w.Static {
		w.engine.Static(k, v)
	}
	if len(w.TrustedProxies) > 0 {
		w.engine.SetTrustedProxies(w.TrustedProxies)
	}
	L.Push(ud)

	return 1
}

func Preload(L *lua.LState) lua.LValue {
	w := &Web{}
	ud := core.NewUserData(L, w)

	return ud
}
