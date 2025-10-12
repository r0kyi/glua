package web

import (
	"io"

	"github.com/gin-gonic/gin"
	"github.com/r0kyi/glua/core"
)

type Context struct {
	request struct {
		method string
		body   string
		path   string
		uri    string
	}

	response struct {
		statusCode int
		name       string
		obj        map[string]any
		format     string
		values     []any
	}

	cookie struct {
		name     string
		value    string
		maxAge   int
		path     string
		domain   string
		secure   bool
		httpOnly bool
	}

	header struct {
		key   string
		value string
	}

	query struct {
		key   string
		value string
	}

	form struct {
		key   string
		value string
	}

	param struct {
		key   string
		value string
	}

	context *gin.Context
}

func (c *Context) json() {
	c.context.JSON(c.response.statusCode, gin.H(c.response.obj))
}

func (c *Context) asciiJson() {
	c.context.AsciiJSON(c.response.statusCode, gin.H(c.response.obj))
}

func (c *Context) string() {
	c.context.String(c.response.statusCode, c.response.format, c.response.values...)
}

func (c *Context) html() {
	c.context.HTML(c.response.statusCode, c.response.name, gin.H(c.response.obj))
}

func (c *Context) getCookie() {
	cookie, _ := c.context.Cookie(c.response.name)
	c.cookie.value = cookie
}

func (c *Context) setCookie() {
	c.context.SetCookie(c.cookie.name, c.cookie.value, c.cookie.maxAge, c.cookie.path, c.cookie.domain, c.cookie.secure, c.cookie.httpOnly)
}

func (c *Context) getHeader() {
	c.header.value = c.context.GetHeader(c.header.key)
}

func (c *Context) setHeader() {
	c.context.Header(c.header.key, c.header.value)
}

func (c *Context) getQuery() {
	value := c.context.Query(c.query.key)
	c.query.value = value
}

func (c *Context) getForm() {
	value := c.context.PostForm(c.form.key)
	c.form.value = value
}

func (c *Context) getParam() {
	value := c.context.Param(c.param.key)
	c.param.value = value
}

func (c *Context) body() {
	body, _ := io.ReadAll(c.context.Request.Body)
	c.request.body = core.B2S(body)
}

func (c *Context) method() {
	method := c.context.Request.Method
	c.request.method = method
}

func (c *Context) path() {
	path := c.context.Request.URL.Path
	c.request.path = path
}

func (c *Context) uri() {
	uri := c.context.Request.RequestURI
	c.request.uri = uri
}
