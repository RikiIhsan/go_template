package utils

import "github.com/gofiber/fiber/v2"

func ServerError(ctx *fiber.Ctx, err ...error) error {
	if len(err) == 0 {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "something went wrong,try again later",
		})
	} else {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err[0].Error(),
		})
	}
}
