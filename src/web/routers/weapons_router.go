package routers

import (
	"Weapons-API/src/weapons"
	"net/http"

	"github.com/labstack/echo/v4"
)

type WeaponsRouter struct {
	weaponsService weapons.IWeaponsService
}

func InitWeaponsRouter() *WeaponsRouter {
	return &WeaponsRouter{
		weaponsService: weapons.InitWeaponsService(),
	}
}

func (router *WeaponsRouter) OnGetWeapon(ctx echo.Context) error {
	weaponId := ctx.Param("weaponId")
	weapon, err := router.weaponsService.Get(weaponId)
	if err != nil {
		return err
	} 
	return ctx.JSON(http.StatusOK, weapon)
}

func (router *WeaponsRouter) Bind(server *echo.Echo) {
	group := server.Group("/api/v1/weapons")
	group.GET("/:weaponId", router.OnGetWeapon)
}