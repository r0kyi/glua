package http

import (
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/r0kyi/glua/core"
)

type Http struct {
	url      string
	Headers  map[string][]string `lua:"headers"`
	Args     map[string]string   `lua:"args"`
	Body     string              `lua:"body"`
	Proxy    string              `lua:"proxy"`
	Timeout  time.Duration       `lua:"timeout"`
	client   *resty.Client
	response *Response
}

func (h *Http) get() error {
	response, err := h.client.R().SetQueryParams(h.Args).Get(h.url)
	if err != nil {
		return err
	}

	h.response.statusCode = response.StatusCode()
	h.response.headers = response.Header()
	h.response.body = core.B2S(response.Body())

	return nil
}

func (h *Http) post() error {
	response, err := h.client.R().SetBody(h.Body).SetQueryParams(h.Args).Post(h.url)
	if err != nil {
		return err
	}

	h.response.statusCode = response.StatusCode()
	h.response.headers = response.Header()
	h.response.body = core.B2S(response.Body())

	return nil
}

func (h *Http) put() error {
	response, err := h.client.R().SetBody(h.Body).SetQueryParams(h.Args).Put(h.url)
	if err != nil {
		return err
	}

	h.response.statusCode = response.StatusCode()
	h.response.headers = response.Header()
	h.response.body = core.B2S(response.Body())

	return nil
}

func (h *Http) delete() error {
	response, err := h.client.R().SetBody(h.Body).SetQueryParams(h.Args).Delete(h.url)
	if err != nil {
		return err
	}

	h.response.statusCode = response.StatusCode()
	h.response.headers = response.Header()
	h.response.body = core.B2S(response.Body())

	return nil
}

func (h *Http) patch() error {
	response, err := h.client.R().SetBody(h.Body).SetQueryParams(h.Args).Patch(h.url)
	if err != nil {
		return err
	}

	h.response.statusCode = response.StatusCode()
	h.response.headers = response.Header()
	h.response.body = core.B2S(response.Body())

	return nil
}

func (h *Http) options() error {
	response, err := h.client.R().SetQueryParams(h.Args).Options(h.url)
	if err != nil {
		return err
	}

	h.response.statusCode = response.StatusCode()
	h.response.headers = response.Header()
	h.response.body = ""

	return nil
}

func (h *Http) head() error {
	response, err := h.client.R().SetQueryParams(h.Args).Head(h.url)
	if err != nil {
		return err
	}

	h.response.statusCode = response.StatusCode()
	h.response.headers = response.Header()
	h.response.body = ""

	return nil
}
