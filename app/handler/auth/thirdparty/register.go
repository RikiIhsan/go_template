package thirdparty

import (
	models "asrs/app/models/auth"
	request "asrs/app/request/auth"
	"asrs/app/utils"
	"encoding/base64"
	"fmt"
	"github.com/RikiIhsan/lib/env"
	"github.com/RikiIhsan/lib/validator"
	"github.com/gofiber/fiber/v2"
	"time"
)

// usage for register a key for a thirdparty application
func Register(ctx *fiber.Ctx) error {
	var req = new(request.RegisterKey)
	if err := ctx.BodyParser(&req); err != nil {
		return utils.ServerError(ctx)
	}
	if validate := validator.Validate(req); len(validate) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": validate,
		})
	}
	expired := time.Now().Add(time.Hour * 60 * 30)
	var key, _ = utils.Encrypt([]byte(fmt.Sprintf("%s:%s",
		req.Detail,
		expired.Unix())),
		[]byte(env.Get("SECRET_KEY")),
	)
	encodedString := base64.StdEncoding.EncodeToString(key)
	models := &models.Key{
		Details:    req.Detail,
		Token:      encodedString,
		Expired_at: expired,
	}
	if err := models.Create().Error; err != nil {
		return utils.ServerError(ctx)
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "register key successfully",
		"key":     encodedString,
	})
}
