package repository

import (
	"app/model"

	"gorm.io/gorm"
)

type AdminRepositoryInterface interface {
	GetAllAdmins(admins *[]model.Admin) error
	GetAdminById(admin *model.Admin, adminId uint) error
	CreateAdmin(admin *model.Admin) error
	UpdateAdmin(admin *model.Admin, adminId uint) error
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

func (ar *adminRepository) CreateAdmin(admin *model.Admin) error {
	if err := ar.db.Create(admin).Error; err != nil {
		return err
	}
	return nil
}

func (ar *adminRepository) UpdateAdmin(admin *model.Admin, adminId uint) error {
	db := ar.db.Model(admin).Where("id = ?", adminId).Omit("id").Updates(admin)
	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
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
