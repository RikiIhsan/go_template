package middleware

import (
	"asrs/app/models/auth"
	"github.com/RikiIhsan/lib/env"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"time"
)

func Jwt() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(env.Get("SECRET_KEY"))},
	})
}

func ThirdParty(ctx *fiber.Ctx) error {
	key := ctx.Get("asrs_key")
	if key != "" {
		keys := &auth.Key{
			Token: key,
		}
		if err := keys.Find().Error; err == nil && time.Now().Unix() < keys.Expired_at.Unix() {
			return ctx.Next()
		}
	}
	return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": "you are not authorized!",
	})
}
