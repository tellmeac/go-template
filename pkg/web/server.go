package web

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"sync/atomic"
	"time"
)

type ServerConfig struct {
	Listen            string        `config:"listen"`
	DrainInterval     time.Duration `yaml:"drainInterval"`
	Profile           bool          `yaml:"profile"`
	ReadTimeout       time.Duration `yaml:"readTimeout"`
	ReadHeaderTimeout time.Duration `yaml:"readHeaderTimeout"`
	WriteTimeout      time.Duration `yaml:"writeTimeout"`
	IdleTimeout       time.Duration `yaml:"idleTimeout"`
	Env               string        `yaml:"env"`
	GZip              struct {
		CompLevel int `yaml:"compLevel"`
	} `yaml:"gzip,omitempty"`
}

type Server interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	Router() gin.IRouter
	Ready() bool
}

var _ Server = (*BaseServer)(nil)

type BaseServer struct {
	engine     *gin.Engine
	httpServer *http.Server

	config ServerConfig

	isNotReady int32
}

func NewServer(config ServerConfig, handler *gin.Engine) *BaseServer {
	s := &BaseServer{
		engine: handler,
		config: config,
	}

	listen := config.Listen
	// сервер может подниматься на рандомном порту для blackbox и нагрузочных тестов
	port := os.Getenv("RECIPE_PORT")
	if len(port) > 0 {
		listen = "127.0.0.1:" + port
	}

	s.httpServer = &http.Server{
		Addr:              listen,
		Handler:           s.engine.Handler(),
		ReadTimeout:       config.ReadTimeout,
		ReadHeaderTimeout: config.ReadHeaderTimeout,
		WriteTimeout:      config.WriteTimeout,
		IdleTimeout:       config.IdleTimeout,
	}

	return s
}

func (s *BaseServer) Start(ctx context.Context) error {
	// TODO: Profile routes with Profile options check

	s.Router().GET("/live", func(_ *gin.Context) {})
	s.Router().GET("/ping", s.getPing)

	go func() {
		for {
			<-ctx.Done()
			return
		}
	}()

	return s.httpServer.ListenAndServe()
}

func (s *BaseServer) Stop(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, s.config.DrainInterval)
	defer cancel()

	return s.httpServer.Shutdown(ctx)
}

func (s *BaseServer) Router() gin.IRouter {
	return s.engine
}

func (s *BaseServer) Ready() bool {
	return atomic.LoadInt32(&s.isNotReady) == 0
}

func (s *BaseServer) getPing(ctx *gin.Context) {
	if s.Ready() {
		_, _ = ctx.Writer.Write([]byte("pong"))
	} else {
		http.Error(ctx.Writer, "server cannot accept requests", http.StatusTeapot)
	}
}
