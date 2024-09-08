package postgres

import (
	"fmt"

	"github.com/ReyLegar/ShoppingList/internal/models"
)

type ShoppingListDatabase struct {
	database *Database
}

func (u *ShoppingListDatabase) Create(spl *models.ShoppingList, loginUser string) (int64, error) {

	sqlStatemant := `
	INSERT INTO shoppinglists (name, owner_id) 
    VALUES ($1, (SELECT id FROM users WHERE login = $2)) 
    RETURNING id`

	var id int64
	if err := u.database.db.Ping(); err != nil {
		fmt.Println("Не удалось подключиться к базе данных:", err)
		return -1, err
	}
	err := u.database.db.QueryRow(sqlStatemant, spl.Name, loginUser).Scan(&id)

	if err != nil {
		fmt.Println("Ошибка", err)
		return -1, err
	}
	return id, nil
}
