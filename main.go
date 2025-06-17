package main

import (
	config "digital-wallet/src/config"
	database "digital-wallet/src/utils/database"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	var db database.Connection

	app := fiber.New(fiber.Config{
		AppName:       config.NewApp().AppName,
		CaseSensitive: true,
		StrictRouting: true,
	})

	app.Use(recover.New())
	app.Use(logger.New())

	app.Get("/balance-inquiry/:phoneNumber", func(c *fiber.Ctx) error {
		phoneNumber, _ := c.ParamsInt("phoneNumber", 0)

		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"code":    fiber.StatusOK,
			"status":  "ok",
			"message": "success to get current balance!",
			"data":    db.CheckBalance(phoneNumber),
		})
	})

	app.Post("/withdrawal", func(c *fiber.Ctx) error {
		type request struct {
			PhoneNumber int     `json:"phone_number" form:"phone_number"`
			Amount      float64 `json:"amount" form:"amount"`
		}

		data := new(request)

		if err := c.BodyParser(data); err != nil {
			fmt.Printf("error when get user detail = %s\n", err.Error())
		}

		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"code":    fiber.StatusOK,
			"status":  "ok",
			"message": db.Withdrawal(data.PhoneNumber, data.Amount),
			"data":    nil,
		})
	})

	// 404 handling
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"code":    fiber.StatusNotFound,
			"status":  "not_found",
			"message": "page not found",
			"data":    fiber.Map{},
		})
	})

	app.Listen(":" + config.NewApp().Port)
}
