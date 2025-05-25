package repository

import (
	"github.com/ooooose/jira_go/model"
	"gorm.io/gorm"
)

func SaveIssue(db *gorm.DB, title, description string, status model.IssueStatus, priority model.IssuePriority, asigneeId, reporterId uint) (*model.Issue, error) {
	issue := &model.Issue{
		Title:       title,
		Description: description,
		Status:      status,
		Priority:    priority,
		AsigneeId:   asigneeId,
		ReporterId:  reporterId,
	}
	if err := db.Create(issue).Error; err != nil {
		return nil, err
	}
	return issue, nil
}
