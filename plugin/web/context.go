package web

import (
	"io"

	"github.com/gin-gonic/gin"
	"github.com/r0kyi/glua/core"
)

type Context struct {
	context *gin.Context
}

func (c *Context) json(statusCode int, obj map[string]any) {
	c.context.JSON(statusCode, gin.H(obj))
}

func (c *Context) asciiJson(statusCode int, obj map[string]any) {
	c.context.AsciiJSON(statusCode, gin.H(obj))
}

func (c *Context) string(statusCode int, format string, values ...any) {
	c.context.String(statusCode, format, values...)
}

func (c *Context) html(statusCode int, name string, obj map[string]any) {
	c.context.HTML(statusCode, name, gin.H(obj))
}

func (c *Context) getCookie(name string) string {
	cookie, err := c.context.Cookie(name)
	if err != nil {
		return ""
	}

	return cookie
}

func (c *Context) setCookie(name string, value string, maxAge int, path string, domain string, secure bool, httpOnly bool) {
	c.context.SetCookie(name, value, maxAge, path, domain, secure, httpOnly)
}

func (c *Context) getHeader(key string) string {
	return c.context.GetHeader(key)
}

func (c *Context) setHeader(key string, value string) {
	c.context.Header(key, value)
}

func (c *Context) getQuery(key string) string {
	return c.context.Query(key)
}

func (c *Context) getForm(key string) string {
	return c.context.PostForm(key)
}

func (c *Context) getParam(key string) string {
	return c.context.Param(key)
}

func (c *Context) body() string {
	body, err := io.ReadAll(c.context.Request.Body)
	if err != nil {
		return ""
	}

	return core.B2S(body)
}

func (c *Context) method() string {
	return c.context.Request.Method
}

func (c *Context) path() string {
	return c.context.Request.URL.Path
}

func (c *Context) uri() string {
	return c.context.Request.RequestURI
}
