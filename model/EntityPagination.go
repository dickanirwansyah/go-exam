package model

import "gorm.io/gorm"

type EntityPagination interface {
	Count(db *gorm.DB) int64
	Grab(db *gorm.DB, limit int, offset int) interface{}
}
