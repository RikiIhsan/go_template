package app

import (
	"asrs/app/config"
	"asrs/app/middleware"
	"asrs/app/models"
	"asrs/app/router"
	"fmt"
	"github.com/RikiIhsan/lib/database"
	"github.com/RikiIhsan/lib/validator"
	"github.com/gofiber/fiber/v2"
)

func Main() {
	//initialize a config
	appconfig := config.MainConfig()
	//initialize database
	if name, err := database.Init(appconfig.Database...); err != nil {
		panic(fmt.Sprintf("cannot connect to database : %s,with error : %s", name, err.Error()))
	}
	//migrate a models
	models.Migrate()
	//initialize validator
	validator.Init(appconfig.ValidatorTranslator...)
	skeleton := fiber.New(fiber.Config{
		AppName: appconfig.AppName,
	})
	//initialize CORS
	skeleton.Use(middleware.Cors())
	//initialize a ROUTE
	router.Route(skeleton)
	//run server
	if err := skeleton.Listen(fmt.Sprintf("%s:%s", appconfig.AppHost, appconfig.AppPort)); err != nil {
		panic(err)
	}
}
