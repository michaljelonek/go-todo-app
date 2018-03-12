package main

import (
	"github.com/mjelonek92/go-todo-app/app"
	"github.com/mjelonek92/go-todo-app/config"
)

func main() {
	config := config.GetConf()
	app := &app.App{}
	app.Init(config)
	app.Run(":8080")
}
