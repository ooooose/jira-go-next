package model

type ProjectMember struct {
	ID        uint
	ProjectID uint
	Project   Project `gorm:"foreignKey:ProjectID"`
	UserID    uint
	User      User `gorm:"foreignKey:UserID"`
}
