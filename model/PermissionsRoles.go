package model

type PermissionsRoles struct {
	Id            uint `gorm:"column:id;primaryKey;autoIncrement"`
	RolesId       uint `gorm:"column:roles_id;not null"`
	PermissionsId uint `gorm:"column:permissions_id;not null"`
}
