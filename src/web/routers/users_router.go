package routers

import (
	"Weapons-API/src/users"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UsersRouter struct {
	usersService users.IUsersService
}

func InitUsersRouter() *UsersRouter {
	return &UsersRouter{
		usersService: users.InitUsersService(),
	}
}

func (router *UsersRouter) OnGetUser(ctx echo.Context) error {
	userId := ctx.Param("userId")
	user, err := router.usersService.Get(userId)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, user)
}

func (router *UsersRouter) Bind(server *echo.Echo) {
	group := server.Group("/api/v1/users")
	group.GET("/:userId", router.OnGetUser)
}