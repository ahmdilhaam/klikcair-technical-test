package config

import (
	helper "digital-wallet/src/utils/helper"
	"strconv"
)

type Database struct {
	Host            string
	Port            string
	User            string
	Password        string
	Name            string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime int
}

func NewDatabase() *Database {
	maxIdle, _ := strconv.Atoi(helper.GetEnv("DB_IDLE_CONNECTIONS"))
	maxOpen, _ := strconv.Atoi(helper.GetEnv("DB_OPEN_CONNECTIONS"))
	maxLife, _ := strconv.Atoi(helper.GetEnv("DB_CONNECTION_LIFETIME"))

	return &Database{
		Host:            helper.GetEnv("DB_HOST"),
		Port:            helper.GetEnv("DB_PORT"),
		User:            helper.GetEnv("DB_USER"),
		Password:        helper.GetEnv("DB_PASSWORD"),
		Name:            helper.GetEnv("DB_DATABASE"),
		MaxIdleConns:    maxIdle,
		MaxOpenConns:    maxOpen,
		ConnMaxLifetime: maxLife,
	}
}
