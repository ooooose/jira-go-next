package model

type IssueStatus int
type IssuePriority int

const (
	StatusOpen IssueStatus = iota
	StatusInProgress
	StatusClosed
)

const (
	PriorityLow IssuePriority = iota
	PriorityMedium
	PriorityHigh
)

type Issue struct {
	ID          uint
	Title       string
	Description string
	Status      IssueStatus
	Priority    IssuePriority
	AsigneeId   uint
	Asignee     User `gorm:"foreignKey:AsigneeId"`
	ReporterId  uint
	Reporter    User `gorm:"foreignKey:ReporterId"`
	CreatedAt   string
	UpdatedAt   string
}
