package model

import "time"

type Roles struct {
	Id        uint      `gorm:"primaryKey;autoIncrement;column:id"`
	Name      string    `gorm:"type:varchar(120);not null;column:name"`
	CreatedAt time.Time `gorm:"type:date;column:created_at"`
	UpdatedAt time.Time `gorm:"type:date;column:updated_at"`
	IsDeleted int       `gorm:"not null;column:is_deleted"`
}
