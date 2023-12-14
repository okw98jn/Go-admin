package controller

import (
	"app/model"
	"app/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type AdminControllerInterface interface {
	GetAllAdmins(c echo.Context) error
	GetAdminById(c echo.Context) error
	CreateAdmin(c echo.Context) error
	UpdateAdmin(c echo.Context) error
	DeleteAdmin(c echo.Context) error
}

type adminController struct {
	adminUsecase usecase.AdminUsecaseInterface
}

func NewAdminController(adminUsecase usecase.AdminUsecaseInterface) AdminControllerInterface {
	return &adminController{adminUsecase}
}

func (ac *adminController) GetAllAdmins(c echo.Context) error {
	admins, err := ac.adminUsecase.GetAllAdmins()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, admins)
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

func (ac *adminController) CreateAdmin(c echo.Context) error {
	admin := model.Admin{}
	if err := c.Bind(&admin); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	adminRes, err := ac.adminUsecase.CreateAdmin(admin)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, adminRes)
}

func (ac *adminController) UpdateAdmin(c echo.Context) error {
	admin := model.Admin{}
	id := c.Param("admin_id")
	adminId, _ := strconv.Atoi(id)
	if err := c.Bind(&admin); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	adminRes, err := ac.adminUsecase.UpdateAdmin(admin, uint(adminId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, adminRes)
}

func (ac *adminController) DeleteAdmin(c echo.Context) error {
	id := c.Param("admin_id")
	adminId, _ := strconv.Atoi(id)
	err := ac.adminUsecase.DeleteAdmin(uint(adminId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
