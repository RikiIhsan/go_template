package main

import (
	"asrs/app"
	"github.com/RikiIhsan/lib/env"
)

func main() {
	//load environtment file
	if err := env.Load(); err != nil {
		panic(err)
	}
	//run application
	app.Main()
}
