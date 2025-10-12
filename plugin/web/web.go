package web

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/r0kyi/glua/core"
	lua "github.com/yuin/gopher-lua"
)

type Web struct {
	Addr           string            `lua:"addr"`
	Mode           string            `lua:"mode"`
	Pattern        string            `lua:"pattern"`
	Static         map[string]string `lua:"static"`
	TrustedProxies []string          `lua:"trusted_proxies"`

	route struct {
		path string
		fn   gin.HandlerFunc
	}

	engine  *gin.Engine
	context *Context
	session *Session
}

func (w *Web) get() {
	w.engine.GET(w.route.path, w.route.fn)
}

func (w *Web) post() {
	w.engine.POST(w.route.path, w.route.fn)
}

func (w *Web) put() {
	w.engine.PUT(w.route.path, w.route.fn)
}

func (w *Web) delete() {
	w.engine.DELETE(w.route.path, w.route.fn)
}

func (w *Web) patch() {
	w.engine.PATCH(w.route.path, w.route.fn)
}

func (w *Web) options() {
	w.engine.OPTIONS(w.route.path, w.route.fn)
}

func (w *Web) head() {
	w.engine.HEAD(w.route.path, w.route.fn)
}

func (w *Web) use() {
	w.engine.Use(sessions.Sessions(w.session.Name, *w.session.store))
}

func (w *Web) run() error {
	err := w.engine.Run(w.Addr)
	if err != nil {
		return err
	}
	return nil
}

func (w *Web) toHandler(L *lua.LState, fn *lua.LFunction) gin.HandlerFunc {
	return func(c *gin.Context) {
		w.context.context = c
		if err := L.CallByParam(lua.P{
			Fn:      fn,
			NRet:    0,
			Protect: true,
		}, core.NewUserData(L, w.context)); err != nil {
			c.String(500, "")
		}
	}
}
