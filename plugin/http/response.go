package http

type Response struct {
	statusCode int                 `lua:"status_code"`
	headers    map[string][]string `lua:"headers"`
	body       string              `lua:"body"`
}
