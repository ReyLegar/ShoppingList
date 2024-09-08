package database

import (
	"time"

	"github.com/ReyLegar/ShoppingList/internal/models"
)

type UserRepository interface {
	Register(user *models.User) (int64, error)
	Login(user *models.Credentials) models.User
	CreateRefreshToken(user *models.User, refreshTokenString string, refreshExpirationTime time.Time)
}

type ShoppingListRepository interface {
	Create(list *models.ShoppingList, loginUser string) (int64, error)
}

type ItemRepository interface {
	Create(item *models.Item) (int64, error)
}
