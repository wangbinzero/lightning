package routers

import (
	"github.com/labstack/echo"
	v1 "lightning/api/v1"
)

func SetV1Interfaces(e *echo.Echo) {
	e.GET("/api/:platform/v1/hello", v1.V1Hello)
}
