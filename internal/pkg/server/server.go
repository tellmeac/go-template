package server

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/tellmeac/go-template/internal/config"
)

type Binder interface {
	Bind(router gin.IRouter)
}

func New(cfg config.Config, middlewares ...gin.HandlerFunc) Server {
	if !cfg.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	handler := gin.Default()
	for _, mw := range middlewares {
		handler.Use(mw)
	}

	return Server{
		Server: &http.Server{
			Addr:    cfg.ListenAddress,
			Handler: handler,
		},
	}
}

type Server struct {
	*http.Server
}

func (s Server) Bind(a Binder) {
	if a == nil {
		return
	}
	a.Bind(s.Handler.(*gin.Engine))
}
