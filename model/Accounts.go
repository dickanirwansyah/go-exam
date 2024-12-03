package model

import "time"

type Accounts struct {
	Id            uint      `gorm:"primaryKey;autoIncrement;column:id"`
	Email         string    `gorm:"type:varchar(150);not null;column:email"`
	FullName      string    `gorm:"type:varchar(200);not null;column:full_name"`
	PhoneNumber   string    `gorm:"type:varchar(30);not null;column:phone_number"`
	RolesId       uint      `gorm:"not null;column:roles_id"`
	RolesName     string    `gorm:"type:varchar(120);column:roles_name"`
	ImageProfile  string    `gorm:"type:varchar(255);column:image_profile"`
	AddressDetail string    `gorm:"type:varchar(255);column:address_detail"`
	Password      string    `gorm:"type:varchar(255);column:password"`
	CreatedAt     time.Time `gorm:"type:date;column:created_at"`
	UpdatedAt     time.Time `gorm:"type:date;column:updated_at"`
	IsDeleted     int64     `gorm:"column:is_deleted"`
}
