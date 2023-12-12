package model

import "time"

type Admin struct {
	Id        uint      `json:"id" gorm:"primary_key; autoIncrement"`
	Name      string    `json:"name" gorm:"not null; comment:名前"`
	LoginId   string    `json:"login_id" gorm:"unique; not null; comment:ログインID"`
	Role      bool      `json:"role" gorm:"not null; comment:権限"`
	Status    int8      `json:"status" gorm:"not null; default:1 ; comment:ステータス(0:無効,1:有効)"`
	CreatedAt time.Time `json:"created_at" gorm:"comment:作成日時"`
	UpdatedAt time.Time `json:"updated_at" gorm:"comment:更新日時"`
	DeletedAt time.Time `json:"deleted_at" gorm:"comment:削除日時"`
}

type AdminResponse struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	LoginId   string    `json:"login_id"`
	Role      bool      `json:"role"`
	Status    int8      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
