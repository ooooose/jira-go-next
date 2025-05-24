package router

import (
	"github.com/labstack/echo/v4"
	"github.com/ooooose/jira_go/config"
	"github.com/ooooose/jira_go/handler"
)

func NewRouter(cfg *config.Config) *echo.Echo {
    e := echo.New()

    e.GET("/issues", handler.GetIssues)
    e.POST("/issues", handler.CreateIssue)

    return e
}