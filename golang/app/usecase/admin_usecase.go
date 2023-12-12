package usecase

import (
	"app/model"
	"app/repository"
)

type AdminUsecaseInterface interface {
	GetAdminById(adminId uint) (model.AdminResponse, error)
}

type adminUsecase struct {
	adminRepo repository.AdminRepositoryInterface
}

func NewAdminUsecase(adminRepo repository.AdminRepositoryInterface) AdminUsecaseInterface {
	return &adminUsecase{adminRepo}
}

func (au *adminUsecase) GetAdminById(adminId uint) (model.AdminResponse, error) {
	admin := model.Admin{}
	if err := au.adminRepo.GetAdminById(&admin, adminId); err != nil {
		return model.AdminResponse{}, err
	}

	resAdmin := model.AdminResponse{
		Id:        admin.Id,
		Name:      admin.Name,
		LoginId:   admin.LoginId,
		Role:      admin.Role,
		Status:    admin.Status,
		CreatedAt: admin.CreatedAt,
	}
	return resAdmin, nil
}
