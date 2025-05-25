package model

type UserRole int

const (
	RoleGeneral UserRole = iota // 0: 一般
	RoleAdmin                   // 1: 管理者
)

type User struct {
	ID           uint
	Name         string
	Email        string
	PasswordHash string
	Role         UserRole
	CreatedAt    string
	UpdatedAt    string
}
