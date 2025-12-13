package http

import (
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/r0kyi/glua/core"
)

type Http struct {
	Headers map[string][]string `lua:"headers"`
	Proxy   string              `lua:"proxy"`
	Timeout time.Duration       `lua:"timeout"`

	client *resty.Client
}

func (h *Http) get(url string, args map[string]string) (*Response, error) {
	response, err := h.client.R().SetQueryParams(args).Get(url)
	if err != nil {
		return nil, err
	}

	r := &Response{
		headers: make(map[string][]string),
	}
	r.statusCode = response.StatusCode()
	r.headers = response.Header()
	r.body = core.B2S(response.Body())

	return r, nil
}

func (h *Http) post(url string, args map[string]string, body string) (*Response, error) {
	response, err := h.client.R().SetBody(body).SetQueryParams(args).Post(url)
	if err != nil {
		return nil, err
	}

	r := &Response{
		headers: make(map[string][]string),
	}
	r.statusCode = response.StatusCode()
	r.headers = response.Header()
	r.body = core.B2S(response.Body())

	return r, nil
}

func (h *Http) put(url string, args map[string]string, body string) (*Response, error) {
	response, err := h.client.R().SetBody(body).SetQueryParams(args).Put(url)
	if err != nil {
		return nil, err
	}

	r := &Response{
		headers: make(map[string][]string),
	}
	r.statusCode = response.StatusCode()
	r.headers = response.Header()
	r.body = core.B2S(response.Body())

	return r, nil
}

func (h *Http) delete(url string, args map[string]string, body string) (*Response, error) {
	response, err := h.client.R().SetBody(body).SetQueryParams(args).Delete(url)
	if err != nil {
		return nil, err
	}

	r := &Response{
		headers: make(map[string][]string),
	}
	r.statusCode = response.StatusCode()
	r.headers = response.Header()
	r.body = core.B2S(response.Body())

	return r, nil
}

func (h *Http) patch(url string, args map[string]string, body string) (*Response, error) {
	response, err := h.client.R().SetBody(body).SetQueryParams(args).Patch(url)
	if err != nil {
		return nil, err
	}

	r := &Response{
		headers: make(map[string][]string),
	}
	r.statusCode = response.StatusCode()
	r.headers = response.Header()
	r.body = core.B2S(response.Body())

	return r, nil
}

func (h *Http) options(url string, args map[string]string) (*Response, error) {
	response, err := h.client.R().SetQueryParams(args).Options(url)
	if err != nil {
		return nil, err
	}

	r := &Response{
		headers: make(map[string][]string),
	}
	r.statusCode = response.StatusCode()
	r.headers = response.Header()

	return r, nil
}

func (h *Http) head(url string, args map[string]string) (*Response, error) {
	response, err := h.client.R().SetQueryParams(args).Head(url)
	if err != nil {
		return nil, err
	}

	r := &Response{
		headers: make(map[string][]string),
	}
	r.statusCode = response.StatusCode()
	r.headers = response.Header()

	return r, nil
}
