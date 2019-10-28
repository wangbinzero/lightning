package v1

import (
	"github.com/labstack/echo"
	"net/http"
)

func V1Hello(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Hello Lightning exchange")
}
