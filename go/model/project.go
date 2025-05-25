package model

type Project struct {
	ID          uint
	Name        string
	Description string
	OwnerId     uint
	Owner       User `gorm:"foreignKey:OwnerId"`
	CreatedAt   string
	UpdatedAt   string
}
