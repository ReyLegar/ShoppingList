package app

import (
	"database/sql"
	"fmt"

	"github.com/ReyLegar/ShoppingList/config"
	"github.com/ReyLegar/ShoppingList/internal/database/postgres"
	"github.com/ReyLegar/ShoppingList/internal/services"
	restTransport "github.com/ReyLegar/ShoppingList/internal/transport/rest"

	_ "github.com/lib/pq"
)

func Run(cfg *config.Config) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.DB.Host, cfg.DB.DBPort, cfg.DB.User, cfg.DB.Password, cfg.DB.DBname)
	db, err := sql.Open("postgres", connStr)
	fmt.Println(connStr)
	if err != nil {
		panic(err)
	}

	postgresDB := postgres.New(db)

	userService := services.NewUserService(postgresDB)
	shoppingListService := services.NewShoppingListService(postgresDB)
	itemService := services.NewItemService(postgresDB)

	router := restTransport.NewRouter(userService, shoppingListService, itemService)

	server := restTransport.NewServer(cfg.Port, router)
	server.Start()
}
