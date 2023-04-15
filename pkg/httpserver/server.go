package httpserver

import (
	"errors"
	"fmt"
	"ld-context-provider/pkg/logz"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

var logger = logz.NewLogger()

type HTTPRequestHandler func(c *gin.Context)

type HTTPHandler interface {
	Path() string
	Method() string
	Handler() HTTPRequestHandler
}

type Server struct {
	httpServer *http.Server
}

func New(mode, addr string, handlers ...HTTPHandler) *Server {
	gin.SetMode(mode)
	r := gin.New()

	for _, handler := range handlers {
		r.Handle(handler.Method(), handler.Path(), gin.HandlerFunc(handler.Handler()))
	}

	h := cors.New(
		cors.Options{
			AllowedMethods: []string{
				http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions,
			},
			AllowedHeaders: []string{"*"},
		},
	).Handler(r)

	return &Server{
		httpServer: &http.Server{
			Addr:    addr,
			Handler: h,
		},
	}
}

func (s *Server) Start() error {
	go func() {
		logger.Info(fmt.Sprintf("Listening for requests on [%s]", s.httpServer.Addr))

		err := s.httpServer.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic(fmt.Sprintf("Failed to start server on [%s]: %s", s.httpServer.Addr, err))
		}
		logger.Info("Server has stopped")
	}()
	return nil
}
