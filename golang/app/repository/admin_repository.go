package repository

import (
	"app/model"

	"gorm.io/gorm"
)

type AdminRepositoryInterface interface {
	GetAllAdmins(admins *[]model.Admin) error
	GetAdminById(admin *model.Admin, adminId uint) error
	DeleteAdmin(adminId uint) error
}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) AdminRepositoryInterface {
	return &adminRepository{db}
}

func (ar *adminRepository) GetAllAdmins(admins *[]model.Admin) error {
	if err := ar.db.Find(admins).Error; err != nil {
		return err
	}
	return nil
}

func (ar *adminRepository) GetAdminById(admin *model.Admin, adminId uint) error {
	if err := ar.db.Where("id = ?", adminId).First(admin).Error; err != nil {
		return err
	}
	return nil
}

func (ar *adminRepository) DeleteAdmin(adminId uint) error {
	result := ar.db.Where("id = ?", adminId).Delete(&model.Admin{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
