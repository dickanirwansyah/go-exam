package model

type QuestionCategory struct {
	Id        uint   `gorm:"column:id;autoIncrement"`
	Name      string `gorm:"column:name;not null"`
	IsDeleted int64  `gorm:"column:is_deleted;not null"`
}
