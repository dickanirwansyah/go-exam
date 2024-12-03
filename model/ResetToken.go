package model

import "time"

type ResetToken struct {
	Id         uint      `gorm:"primaryKey;autoIncrement;column:id"`
	AccountId  uint      `gorm:"not null;column:account_id"`
	Email      string    `gorm:"type:varchar(255);not null;column:email"`
	IsExecuted string    `gorm:"type:varchar(20);not null;column:is_executed"`
	Token      string    `gorm:"type:varchar(100);not null;column:token"`
	Expires    time.Time `gorm:"not null;column:expires"`
}
