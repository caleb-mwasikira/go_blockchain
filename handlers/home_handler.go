package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func GetHomePage(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "Go Blockchain",
	})
}

func ErrorHandler(c *fiber.Ctx, err error) error {
	errStatusCode := fiber.StatusNotFound
	errMsg := err.Error()

	// check if error is a fiber Error
	fiberErr, ok := err.(*fiber.Error)
	if ok {
		errStatusCode = fiberErr.Code
		errMsg = fiberErr.Message
	}

	return fiber.NewError(errStatusCode, errMsg)
}
