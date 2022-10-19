package server

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/tellmeac/go-template/internal/config"
)

type Attacher interface {
	Attach(router gin.IRouter)
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

func (s Server) Attach(a Attacher) {
	if a == nil {
		return
	}
	a.Attach(s.Handler.(*gin.Engine))
}
