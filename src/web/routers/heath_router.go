package routers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HealthRouter struct {}

func InitHealthRouter() *HealthRouter {
	return &HealthRouter{}
}

func (router *HealthRouter) OnGetHealth(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "ok")
}

func (router *HealthRouter) Bind(server *echo.Echo) {
	group := server.Group("/api/v1/health")
	group.GET("", router.OnGetHealth)
}