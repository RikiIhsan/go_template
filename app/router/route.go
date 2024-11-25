package router

import "github.com/gofiber/fiber/v2"

// base router for your application
func Route(route fiber.Router) {
	route.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"message": "hello world",
		})
	})
	Auth(route.Group("/auth"))
	Api(route.Group("/api"))
}
