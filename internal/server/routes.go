package server

import (
	"github.com/tellmeac/go-template/internal/server/handlers"
	"github.com/tellmeac/go-template/pkg/web"
	"net/http"
)

func (a *App) GetRoutes() []web.Route {
	return []web.Route{
		{
			Path:    "api/v1/users/:id",
			Method:  http.MethodGet,
			Handler: a.handlerAdapter(handlers.GetUser),
		},
	}
}
