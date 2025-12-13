package web

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	lua "github.com/r0kyi/gopher-lua"
)

type Web struct {
	Addr           string            `lua:"addr"`
	Mode           string            `lua:"mode"`
	Pattern        string            `lua:"pattern"`
	Static         map[string]string `lua:"static"`
	TrustedProxies []string          `lua:"trusted_proxies"`

	engine  *gin.Engine
	session *Session
}

func (w *Web) get(path string, fn gin.HandlerFunc) {
	w.engine.GET(path, fn)
}

func (w *Web) post(path string, fn gin.HandlerFunc) {
	w.engine.POST(path, fn)
}

func (w *Web) put(path string, fn gin.HandlerFunc) {
	w.engine.PUT(path, fn)
}

func (w *Web) delete(path string, fn gin.HandlerFunc) {
	w.engine.DELETE(path, fn)
}

func (w *Web) patch(path string, fn gin.HandlerFunc) {
	w.engine.PATCH(path, fn)
}

func (w *Web) options(path string, fn gin.HandlerFunc) {
	w.engine.OPTIONS(path, fn)
}

func (w *Web) head(path string, fn gin.HandlerFunc) {
	w.engine.HEAD(path, fn)
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
		context := &Context{
			context: c,
		}
		if err := L.CallByParam(lua.P{
			Fn:      fn,
			NRet:    0,
			Protect: true,
		}, context); err != nil {
			c.String(500, "")
		}
	}
}
