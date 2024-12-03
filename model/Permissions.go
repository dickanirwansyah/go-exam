package model

import "time"

type Permissions struct {
	Id        uint      `gorm:"primaryKey;autoIncrement;column:id"`
	Name      string    `gorm:"type:varchar(255);not null;column:name"`
	Level     int64     `gorm:"not null;column:level"`
	Endpoint  string    `gorm:"type:varchar(255);not null;column:endpoint"`
	Icon      string    `gorm:"type:varchar(100);not null;column:icon"`
	ParentId  uint      `gorm:"column:parent_id"`
	IsDeleted int64     `gorm:"not null;column:is_deleted"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
