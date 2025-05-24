package service

import "github.com/ooooose/jira_go/repository"

type IssueInput struct {
    Title       string `json:"title"`
    Description string `json:"description"`
}

func GetIssues() []repository.Issue {
    return repository.FindAllIssues()
}

func CreateIssue(input IssueInput) (repository.Issue, error) {
    return repository.SaveIssue(input.Title, input.Description)
}
