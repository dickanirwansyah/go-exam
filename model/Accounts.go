package model

import (
	"time"

	"gorm.io/gorm"
)

type Accounts struct {
	Id            uint      `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	Email         string    `gorm:"type:varchar(150);not null;column:email" json:"email"`
	FullName      string    `gorm:"type:varchar(200);not null;column:full_name" json:"full_name"`
	PhoneNumber   string    `gorm:"type:varchar(30);not null;column:phone_number" json:"phone_number"`
	RolesId       uint      `gorm:"not null;column:roles_id" json:"roles_id"`
	RolesName     string    `gorm:"type:varchar(120);column:roles_name" json:"roles_name"`
	ImageProfile  string    `gorm:"type:varchar(255);column:image_profile" json:"image_profile"`
	AddressDetail string    `gorm:"type:varchar(255);column:address_detail" json:"address_detail"`
	Password      string    `gorm:"type:varchar(255);column:password" json:"-"`
	CreatedAt     time.Time `gorm:"type:date;column:created_at" json:"-"`
	UpdatedAt     time.Time `gorm:"type:date;column:updated_at" json:"-"`
	IsDeleted     int64     `gorm:"column:is_deleted" json:"-"`
}

func (accounts *Accounts) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&Accounts{}).Count(&total)
	return total
}

func (accounts *Accounts) Grab(db *gorm.DB, limit int, offset int) interface{} {
	var listAccount []Accounts
	db.Offset(offset).Limit(limit).Find(&listAccount)
	return listAccount
}
