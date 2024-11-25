package config

import (
	"fmt"
	"github.com/RikiIhsan/lib/database"
	"github.com/RikiIhsan/lib/env"
	"github.com/RikiIhsan/lib/validator"
	"gorm.io/gorm"
)

type Config struct {
	//Application Name
	AppName string
	//Host for running application
	AppHost string
	//Port for running application
	AppPort string
	//Database configuration
	//reference to gorm.io for full docomentation
	Database []database.Config
	//Validator translator usage for custom message from validation error
	//eg. "Field {x} must have a value!" for a rules required
	ValidatorTranslator []validator.Translator
}

func MainConfig() *Config {
	config := &Config{
		AppName: env.Get("APP_NAME"),
		AppHost: env.Get("APP_HOST"),
		AppPort: env.Get("APP_PORT"),
		Database: []database.Config{
			{
				Name: "test",
				Dsn: fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
					env.Get("database_user"),
					env.Get("database_password"),
					env.Get("database_host"),
					env.Get("database_port"),
					env.Get("database_name"),
				),
				Driver: "", //support mysql and sqlsrv
				Config: &gorm.Config{},
			},
		},
		ValidatorTranslator: []validator.Translator{
			//{
			//	Tag:     "required",
			//	Message: "{0} must have a value!",
			//},
		},
	}
	return config
}
