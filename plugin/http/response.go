package http

type Response struct {
	statusCode int
	headers    map[string][]string
	body       string
}
