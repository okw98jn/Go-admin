package router

import (
	"app/controller"

	"github.com/labstack/echo/v4"
)

func NewRouter(adminController controller.AdminControllerInterface) *echo.Echo {
	e := echo.New()
	e.GET("/admin/:admin_id", adminController.GetAdminById)
	return e
}
