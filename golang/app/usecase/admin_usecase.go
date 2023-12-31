package usecase

import (
	"app/model"
	"app/repository"
)

type AdminUsecaseInterface interface {
	GetAllAdmins() ([]model.AdminResponse, error)
	GetAdminById(adminId uint) (model.AdminResponse, error)
	CreateAdmin(admin model.Admin) (model.AdminResponse, error)
	UpdateAdmin(admin model.Admin, adminId uint) (model.AdminResponse, error)
	DeleteAdmin(adminId uint) error
}

type adminUsecase struct {
	adminRepo repository.AdminRepositoryInterface
}

func NewAdminUsecase(adminRepo repository.AdminRepositoryInterface) AdminUsecaseInterface {
	return &adminUsecase{adminRepo}
}

func (au *adminUsecase) GetAllAdmins() ([]model.AdminResponse, error) {
	admins := []model.Admin{}
	if err := au.adminRepo.GetAllAdmins(&admins); err != nil {
		return nil, err
	}

	resAdmins := []model.AdminResponse{}
	for _, admin := range admins {
		resAdmins = append(resAdmins, model.AdminResponse{
			Id:        admin.Id,
			Name:      admin.Name,
			LoginId:   admin.LoginId,
			Role:      admin.Role,
			Status:    admin.Status,
			CreatedAt: admin.CreatedAt,
		})
	}
	return resAdmins, nil
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

func (au *adminUsecase) CreateAdmin(admin model.Admin) (model.AdminResponse, error) {
	if err := au.adminRepo.CreateAdmin(&admin); err != nil {
		return model.AdminResponse{}, err
	}

	resAdmin := model.AdminResponse{
		Id:        admin.Id,
		Name:      admin.Name,
		LoginId:   admin.LoginId,
		Role:      admin.Role,
		Status:    admin.Status,
		CreatedAt: admin.CreatedAt,
		UpdatedAt: admin.UpdatedAt,
	}
	return resAdmin, nil
}

func (au *adminUsecase) UpdateAdmin(admin model.Admin, adminId uint) (model.AdminResponse, error) {
	if err := au.adminRepo.UpdateAdmin(&admin, adminId); err != nil {
		return model.AdminResponse{}, err
	}

	resAdmin := model.AdminResponse{
		Id:        admin.Id,
		Name:      admin.Name,
		LoginId:   admin.LoginId,
		Role:      admin.Role,
		Status:    admin.Status,
		CreatedAt: admin.CreatedAt,
		UpdatedAt: admin.UpdatedAt,
	}
	return resAdmin, nil
}

func (au *adminUsecase) DeleteAdmin(adminId uint) error {
	if err := au.adminRepo.DeleteAdmin(adminId); err != nil {
		return err
	}
	return nil
}
