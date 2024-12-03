package model

import "gorm.io/gorm"

type QuestionCategory struct {
	Id        uint   `gorm:"column:id;autoIncrement"`
	Name      string `gorm:"column:name;not null"`
	IsDeleted int64  `gorm:"column:is_deleted;not null"`
}

func (questionCategory *QuestionCategory) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&QuestionCategory{}).Count(&total)
	return total
}

func (questionCategory *QuestionCategory) Grab(db *gorm.DB, limit int, offset int) interface{} {
	var questionCategories []QuestionCategory
	db.Offset(offset).Limit(limit).Find(&questionCategories)
	return questionCategories
}
