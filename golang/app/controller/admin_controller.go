package controller

import (
	"app/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type AdminControllerInterface interface {
	GetAdminById(c echo.Context) error
}

type adminController struct {
	adminUsecase usecase.AdminUsecaseInterface
}

func NewAdminController(adminUsecase usecase.AdminUsecaseInterface) AdminControllerInterface {
	return &adminController{adminUsecase}
}

func (ac *adminController) GetAdminById(c echo.Context) error {
	id := c.Param("admin_id")
	adminId, _ := strconv.Atoi(id)
	admin, err := ac.adminUsecase.GetAdminById(uint(adminId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, admin)
}
