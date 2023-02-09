package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tellmeac/go-template/internal/commands"
	"github.com/tellmeac/go-template/internal/core/actions"
	"github.com/tellmeac/go-template/pkg/schema/json"
	"net/http"
)

func GetUser(ctx *gin.Context, repo *commands.Repository) error {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return json.APIError(ctx, http.StatusBadRequest, fmt.Errorf("argument 'id': %w", err))
	}

	user, err := actions.GetUserAction(ctx, id, repo.Storage)
	if err != nil {
		// TODO: not always internal server error
		return json.APIError(ctx, http.StatusInternalServerError, err)
	}

	return json.APIOK(ctx, user)
}
