package main

import (
	"github.com/ooooose/jira_go/config"
	"github.com/ooooose/jira_go/router"
	"github.com/ooooose/jira_go/util"
	"github.com/ooooose/jira_go/model"
)

func main() {
	cfg := config.Load()

	util.InitDB(cfg)

	// Automatically migrate the database schema
	util.DB.AutoMigrate(&model.Issue{}, &model.User{}, &model.Project{}, &model.IssueComment{}, &model.ProjectMember{})

	e := router.NewRouter(cfg)
	e.Logger.Fatal(e.Start(":" + cfg.Port))
}
