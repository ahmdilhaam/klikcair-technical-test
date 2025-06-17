package main

import (
	config "digital-wallet/src/config"
	handlers "digital-wallet/src/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:       config.NewApp().AppName,
		CaseSensitive: true,
		StrictRouting: true,
	})

	app.Use(recover.New())
	app.Use(logger.New())

	// Create a /transaction endpoint group
	transcation := app.Group("/transaction")

	transcation.Get("/balance-inquiry/:phoneNumber", handlers.CheckBalance)
	transcation.Post("/withdrawal", handlers.Withdrawal)

	// 404 handling
	app.Use(handlers.NotFound)

	app.Listen(":" + config.NewApp().Port)
}
