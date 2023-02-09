package server

import (
	"github.com/gin-gonic/gin"
	"github.com/tellmeac/go-template/internal/commands"
	"github.com/tellmeac/go-template/pkg/schema/json"
)

func (a *App) repoContextHandlerAdapter(handler func(ctx *gin.Context, repo *commands.Repository) error) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := handler(ctx, a.repo)
		handleError(ctx, err)
	}
}

func handleError(ctx *gin.Context, err error) {
	if err == nil {
		return
	}

	_ = json.APIError(ctx, 500, err)
}
