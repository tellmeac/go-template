package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/tellmeac/go-template/internal/commands"
	"github.com/tellmeac/go-template/pkg/web"
	"net/http"
)

func New(repo *commands.Repository) *App {
	app := App{
		server: gin.Default(),
		repo:   repo,
	}

	app.InitRoutes(app.GetRoutes())

	return &app
}

type App struct {
	// TODO: custom WebApp with all useful routes and configurations
	server *gin.Engine
	repo   *commands.Repository
}

func (a *App) Start(_ context.Context) error {
	// TODO: use pkg for WebApp with pprof routes, pingRouter, ctx and etc
	return http.ListenAndServe(a.repo.Config.Server.Listen, a.server)
}

func (a *App) InitRoutes(routes []web.Route) {
	for _, r := range routes {
		a.server.Handle(r.Method, r.Path, r.Handler)
	}
}
