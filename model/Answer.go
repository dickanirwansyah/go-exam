package model

import "time"

type Answer struct {
	Id         uint      `gorm:"primaryKey;column:id;autoIncrement"`
	QuestionId uint      `gorm:"column:question_id;not null"`
	Text       string    `gorm:"column:text;not null;column:text;type:text"`
	IsCorrect  bool      `gorm:"column:is_correct;not null"`
	CreatedAt  time.Time `gorm:"column:created_at;not null"`
	UpdatedAt  time.Time `gorm:"column:updated_at;not null"`
	IsDeleted  int64     `gorm:"column:is_deleted;not null"`
}
