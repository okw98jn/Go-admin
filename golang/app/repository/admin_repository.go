package repository

import (
	"app/model"

	"gorm.io/gorm"
)

type AdminRepositoryInterface interface {
	GetAdminById(admin *model.Admin, adminId uint) error
}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) AdminRepositoryInterface {
	return &adminRepository{db}
}

func (ar *adminRepository) GetAdminById(admin *model.Admin, adminId uint) error {
	if err := ar.db.Where("id = ?", adminId).First(admin).Error; err != nil {
		return err
	}
	return nil
}
