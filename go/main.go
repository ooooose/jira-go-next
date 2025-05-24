package main

import (
    "github.com/ooooose/jira_go/config"
    "github.com/ooooose/jira_go/router"
    "github.com/ooooose/jira_go/util"
)

func main() {
    cfg := config.Load()

    util.InitDB(cfg)

    e := router.NewRouter(cfg)
    e.Logger.Fatal(e.Start(":" + cfg.Port))
}
