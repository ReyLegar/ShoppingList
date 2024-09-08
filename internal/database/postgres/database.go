package postgres

import (
	"database/sql"

	"github.com/ReyLegar/ShoppingList/internal/database"
)

type Database struct {
	db *sql.DB

	userDatabase         *UserDatabase
	shoppingListDatabase *ShoppingListDatabase
	itemDatabase         *ItemDatabase
}

func New(db *sql.DB) *Database {
	return &Database{
		db: db,
	}
}

func (db *Database) User() database.UserRepository {
	if db.userDatabase != nil {
		return db.userDatabase
	}

	db.userDatabase = &UserDatabase{
		database: db,
	}

	return db.userDatabase
}

func (db *Database) ShoppingList() database.ShoppingListRepository {
	if db.shoppingListDatabase != nil {
		return db.shoppingListDatabase
	}

	db.shoppingListDatabase = &ShoppingListDatabase{
		database: db,
	}

	return db.shoppingListDatabase
}

func (db *Database) Item() database.ItemRepository {
	if db.itemDatabase != nil {
		return db.itemDatabase
	}

	db.itemDatabase = &ItemDatabase{
		database: db,
	}

	return db.itemDatabase
}
