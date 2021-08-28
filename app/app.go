package app

import "github.com/agnynureza/homework-rakamin-golang-sql/config"

type Application struct {
	Config *config.Config
}

func Init() *Application {
	application := &Application{
		Config: config.Init(),
	}

	return application
}
