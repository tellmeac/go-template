package v1

import (
	"github.com/tellmeac/go-template/internal/app"
	"github.com/tellmeac/go-template/internal/pkg/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func New(app *app.App) *Endpoints {
	return &Endpoints{app: app}
}

type Endpoints struct {
	app *app.App
}

func (e Endpoints) Bind(router gin.IRouter) {
	router.GET("/api/v1/greet", e.GetGreeting)
	router.PATCH("/api/v1/greet", e.ModifyGreeting)
}

func (e Endpoints) GetGreeting(ctx *gin.Context) {
	var request GreetingRequest
	if err := ctx.BindQuery(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, Error{Msg: err.Error()})
		return
	}

	greeting := e.app.GetGreeting(ctx, request.Name)
	ctx.JSON(http.StatusOK, GreetingResponse{Text: greeting})
}

func (e Endpoints) ModifyGreeting(ctx *gin.Context) {
	var requestBody app.ModifyRequest
	if err := ctx.BindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, Error{Msg: err.Error()})
		return
	}

	err := e.app.ModifyGreeting(ctx, requestBody)
	if err != nil {
		switch {
		case errors.Is(err, errors.ErrInvalidTemplate):
			ctx.JSON(http.StatusBadRequest, Error{Msg: err.Error()})
		default:
			ctx.Status(http.StatusInternalServerError)
		}
		return
	}

	ctx.Status(http.StatusNoContent)
}
