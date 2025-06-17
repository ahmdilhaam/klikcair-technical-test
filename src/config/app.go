package config

import (
	helper "digital-wallet/src/utils/helper"
)

type App struct {
	AppName string
	Port    string
}

func NewApp() *App {
	return &App{
		AppName: helper.GetEnv("APP_NAME"),
		Port:    helper.GetEnv("PORT"),
	}
}
