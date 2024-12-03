package model

import "time"

type Questions struct {
	Id                 uint      `gorm:"primaryKey;autoIncrement;column:id"`
	Text               string    `gorm:"column:text;type:text;not null"`
	QuestionCategoryId uint      `gorm:"not null;column:question_category_id"`
	IsDeleted          int64     `gorm:"not null;column:is_deleted"`
	CreatedAt          time.Time `gorm:"column:created_at;not null"`
	UpdatedAt          time.Time `gorm:"column:updated_at;not null"`
}
