package handler

import (
	"net/http"

	"github.com/ooooose/jira_go/service"

	"github.com/labstack/echo/v4"
	"github.com/ooooose/jira_go/util"
)

func GetIssues(c echo.Context) error {
	issues, err := service.GetIssues(util.DB)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, issues)
}

func CreateIssue(c echo.Context) error {
	var input service.IssueInput
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	issue, err := service.CreateIssue(util.DB, input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, issue)
}
