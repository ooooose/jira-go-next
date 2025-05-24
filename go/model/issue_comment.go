package model

type IssueComment struct {
	ID          uint
	IssueId     uint
	Issue			 Issue `gorm:"foreignKey:IssueId"`
	AuthorId    uint
	Author      User `gorm:"foreignKey:AuthorId"`
	Content     string
	CreatedAt   string
	UpdatedAt   string
}
