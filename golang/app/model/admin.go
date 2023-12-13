package model

import (
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	Id        uint           `json:"id" gorm:"primary_key; autoIncrement"`
	Name      string         `json:"name" gorm:"not null; comment:名前"`
	LoginId   string         `json:"login_id" gorm:"unique; not null; comment:ログインID"`
	Role      int8           `json:"role" gorm:"not null; comment:権限"`
	Status    bool           `json:"status" gorm:"not null; default:1 ; comment:ステータス(0:無効,1:有効)"`
	CreatedAt time.Time      `json:"created_at" gorm:"comment:作成日時"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"comment:更新日時"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"comment:論理削除日時"`
}

type AdminResponse struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	LoginId   string    `json:"login_id"`
	Role      int8      `json:"role"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
