package web

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/r0kyi/glua/core"
	lua "github.com/r0kyi/gopher-lua"
)

func (w *Web) String() string {
	return fmt.Sprintf("glua.web: %p", w)
}

func (w *Web) Type() lua.LValueType {
	return lua.LTObject
}

func (w *Web) AssertFunction() (*lua.LFunction, bool) {
	return core.NewFunction(newWebL), true
}

func (w *Web) getL(L *lua.LState) int {
	path := L.CheckString(1)
	fn := L.CheckFunction(2)

	w.get(path, w.toHandler(L, fn))

	return 0
}

func (w *Web) postL(L *lua.LState) int {
	path := L.CheckString(1)
	fn := L.CheckFunction(2)

	w.post(path, w.toHandler(L, fn))

	return 0
}

func (w *Web) putL(L *lua.LState) int {
	path := L.CheckString(1)
	fn := L.CheckFunction(2)

	w.put(path, w.toHandler(L, fn))

	return 0
}

func (w *Web) deleteL(L *lua.LState) int {
	path := L.CheckString(1)
	fn := L.CheckFunction(2)

	w.delete(path, w.toHandler(L, fn))

	return 0
}

func (w *Web) patchL(L *lua.LState) int {
	path := L.CheckString(1)
	fn := L.CheckFunction(2)

	w.patch(path, w.toHandler(L, fn))

	return 0
}

func (w *Web) optionsL(L *lua.LState) int {
	path := L.CheckString(1)
	fn := L.CheckFunction(2)

	w.options(path, w.toHandler(L, fn))

	return 0
}

func (w *Web) headL(L *lua.LState) int {
	path := L.CheckString(1)
	fn := L.CheckFunction(2)

	w.head(path, w.toHandler(L, fn))

	return 0
}

func (w *Web) useL(L *lua.LState) int {
	session := L.CheckAny(1)
	if session == nil {
		return 0
	}

	w.session = session.(*Session)
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
		return w.session
	default:
		return lua.LNil
	}
}

func newWebL(L *lua.LState) int {
	w := &Web{}

	tbl := L.CheckTable(1)
	_ = core.LTableToStrut(tbl, w)

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

	L.Push(w)

	return 1
}

func Preload() lua.LValue {
	w := &Web{}

	return w
}
