package handlers

import (
	database "digital-wallet/src/utils/database"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func CheckBalance(c *fiber.Ctx) error {
	var db database.Connection

	phoneNumber, _ := c.ParamsInt("phoneNumber", 0)
	balance := db.CheckBalance(phoneNumber)

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"code":    fiber.StatusOK,
		"status":  "ok",
		"message": "success to get current balance!",
		"data":    balance,
	})
}

func Withdrawal(c *fiber.Ctx) error {
	var db database.Connection

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
}
