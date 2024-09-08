package main

import (
	"github.com/ReyLegar/ShoppingList/config"
	"github.com/ReyLegar/ShoppingList/internal/app"
)

func main() {
	cfg, err := config.NewConfig()

	if err != nil {
		panic(err)
	}

	app.Run(cfg)
}
