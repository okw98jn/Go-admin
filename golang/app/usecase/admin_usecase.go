package usecase

import (
	"app/model"
	"app/repository"
)

type AdminUsecaseInterface interface {
	GetAdminById(adminId uint) (model.Admin, error)
}

type adminUsecase struct {
	adminRepo repository.AdminRepositoryInterface
}

func NewAdminUsecase(adminRepo repository.AdminRepositoryInterface) AdminUsecaseInterface {
	return &adminUsecase{adminRepo}
}

func (au *adminUsecase) GetAdminById(adminId uint) (model.Admin, error) {
	admin := model.Admin{}
	if err := au.adminRepo.GetAdminById(&admin, adminId); err != nil {
		return model.Admin{}, err
	}

	resAdmin := model.Admin{
		Id:        admin.Id,
		Name:      admin.Name,
		LoginId:   admin.LoginId,
		Role:      admin.Role,
		Status:    admin.Status,
		CreatedAt: admin.CreatedAt,
		UpdatedAt: admin.UpdatedAt,
		DeletedAt: admin.DeletedAt,
	}
	return resAdmin, nil
}
