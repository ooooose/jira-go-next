package service

import (
	"github.com/ooooose/jira_go/model"
	"github.com/ooooose/jira_go/repository"
	"gorm.io/gorm"
)

type IssueInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      model.IssueStatus
	Priority    model.IssuePriority
	AsigneeId   uint
	ReporterId  uint
}

func GetIssues(db *gorm.DB) ([]model.Issue, error) {
	var issues []model.Issue
	if err := db.Find(&issues).Error; err != nil {
		return nil, err
	}
	return issues, nil
}

func CreateIssue(db *gorm.DB, input IssueInput) (*model.Issue, error) {
	return repository.SaveIssue(
		db,
		input.Title,
		input.Description,
		input.Status,
		input.Priority,
		input.AsigneeId,
		input.ReporterId,
	)
}
