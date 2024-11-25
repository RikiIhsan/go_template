package auth

import (
	models "asrs/app/models/auth"
	request "asrs/app/request/auth"
	"asrs/app/utils"
	"errors"
	argon2id "github.com/RikiIhsan/lib/argon2/id"
	"github.com/RikiIhsan/lib/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
	"time"
)

// sign up a user
func Signup(ctx *fiber.Ctx) error {
	req := new(request.UserSignUp)
	if err := ctx.BodyParser(&req); err != nil {
		return utils.ServerError(ctx)
	}
	if validate := validator.Validate(req); len(validate) > 0 {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{
				"error": validate,
			})
	}
	var haserr error
	req.Password, haserr = argon2id.CreateHash(req.Password, argon2id.DefaultParams)
	if haserr != nil {
		return utils.ServerError(ctx)
	}
	users := &models.User{
		Username: req.Username,
		Password: req.Password,
		Name:     req.Name,
		Email:    req.Email,
		Active:   false,
	}
	if err := users.Create().Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{
				"error": "username or email already exists",
			})
		}
		return utils.ServerError(ctx)
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User created successfully",
	})
}

// sign in user
func SignIn(ctx *fiber.Ctx) error {
	var req = new(request.UserSignIn)
	if err := ctx.BodyParser(&req); err != nil {
		return utils.ServerError(ctx)
	}
	if validate := validator.Validate(req); len(validate) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": validate,
		})
	}
	users := &models.User{
		Username: req.Username,
		Password: req.Password,
	}
	if err := users.FindByUsername().Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "user not found",
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "something went wrong, try again later",
		})
	}
	if match, err := argon2id.ComparePassAndHash(req.Password, users.Password); !match {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid password",
		})
	} else if err != nil {
		return utils.ServerError(ctx)
	}
	if !users.Active && !users.ActivateAt.Valid {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "you account not active,please contact to administrator for activation",
		})
	}
	atokenexp := time.Now().Add(time.Minute * 5)
	atokenclaims := jwt.MapClaims{
		"username": users.Username,
		"exp":      atokenexp.Unix(),
	}
	atoken := utils.CreateJwt(atokenclaims)
	rtokenexp := time.Now().Add(time.Hour * 1)
	rtokenclaims := jwt.MapClaims{
		"username": users.Username,
		"exp":      rtokenexp.Unix(),
	}
	rtoken := utils.CreateJwt(rtokenclaims)
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "login successfully",
		"accesstoken": fiber.Map{
			"token":   atoken,
			"expired": atokenexp,
		},
		"refreshtoken": fiber.Map{
			"token":   rtoken,
			"expired": rtokenexp,
		},
	})
}
