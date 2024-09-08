package postgres

import (
	"fmt"

	"github.com/ReyLegar/ShoppingList/internal/models"
)

type ItemDatabase struct {
	database *Database
}

func (i *ItemDatabase) Create(item *models.Item) (int64, error) {
	sqlStatemant := `INSERT INTO items (shopping_list_id, name, quantity, is_purchased) VALUES ($1, $2, $3, $4) RETURNING id`

	var id int64
	if err := i.database.db.Ping(); err != nil {
		fmt.Println("Не удалось подключиться к базе данных:", err)
		return -1, err
	}
	err := i.database.db.QueryRow(sqlStatemant, item.ShoppingListID, item.Name, item.Quantity, item.IsPurchased).Scan(&id)

	if err != nil {
		fmt.Println("Ошибка", err)
		return -1, err
	}
	return id, nil
}
