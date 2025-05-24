package repository

type Issue struct {
    ID          int    `json:"id"`
    Title       string `json:"title"`
    Description string `json:"description"`
}

var issues = []Issue{
    {ID: 1, Title: "Initial Issue", Description: "This is the first issue."},
}

func FindAllIssues() []Issue {
    return issues
}

func SaveIssue(title, description string) (Issue, error) {
    issue := Issue{
        ID:          len(issues) + 1,
        Title:       title,
        Description: description,
    }
    issues = append(issues, issue)
    return issue, nil
}
