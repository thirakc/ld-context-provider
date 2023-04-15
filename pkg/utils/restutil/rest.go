package restutil

import "ld-context-provider/pkg/httpserver"

type HTTPHandler struct {
	path    string
	method  string
	handler httpserver.HTTPRequestHandler
}

func NewHTTPHandler(path, method string, handler httpserver.HTTPRequestHandler) *HTTPHandler {
	return &HTTPHandler{
		path:    path,
		method:  method,
		handler: handler,
	}
}

func (h *HTTPHandler) Path() string {
	return h.path
}

func (h *HTTPHandler) Method() string {
	return h.method
}

func (h *HTTPHandler) Handler() httpserver.HTTPRequestHandler {
	return h.handler
}
