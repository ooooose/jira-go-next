package handler

import (
    "net/http"
    "github.com/ooooose/jira_go/service"

    "github.com/labstack/echo/v4"
)

func GetIssues(c echo.Context) error {
    issues := service.GetIssues()
    return c.JSON(http.StatusOK, issues)
}

func CreateIssue(c echo.Context) error {
    var input service.IssueInput
    if err := c.Bind(&input); err != nil {
        return c.JSON(http.StatusBadRequest, err)
    }

    issue, err := service.CreateIssue(input)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err)
    }

    return c.JSON(http.StatusCreated, issue)
}
