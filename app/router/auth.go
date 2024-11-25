package router

import (
	"asrs/app/handler/auth"
	"asrs/app/handler/auth/thirdparty"
	"asrs/app/middleware"
	"github.com/gofiber/fiber/v2"
)

// route for request with a path /auth
func Auth(route fiber.Router) {
	route.Post("/users/signup", auth.Signup)
	route.Post("/users/signin", auth.SignIn)
	route.Post("/key/create", middleware.Jwt(), thirdparty.Register)
}
