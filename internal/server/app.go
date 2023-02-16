package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/tellmeac/go-template/internal/commands"
	"github.com/tellmeac/go-template/pkg/web"
)

func New(repo *commands.Repository) *App {
	app := App{
		repo: repo,
	}

	app.InitRoutes(app.GetRoutes())
	app.server = web.NewServer(repo.Config.ServerConfig, app.router)

	return &app
}

type App struct {
	server web.Server
	router *gin.Engine
	repo   *commands.Repository
}

func (a *App) Start(ctx context.Context) error {
	return a.server.Start(ctx)
}

func (a *App) Stop(ctx context.Context) error {
	return a.server.Stop(ctx)
}

func (a *App) InitRoutes(routes []web.Route) {
	for _, r := range routes {
		a.router.Handle(r.Method, r.Path, r.Handler)
	}
}
