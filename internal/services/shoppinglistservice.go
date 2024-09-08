package services

import (
	"github.com/ReyLegar/ShoppingList/internal/database"
	"github.com/ReyLegar/ShoppingList/internal/models"
)

type ShoppingListService interface {
	CreateShoppingList(spl *models.ShoppingList, loginUser string) (int64, error)
}

type shoppingListServiceImpl struct {
	db database.Repository
}

func NewShoppingListService(db database.Repository) ShoppingListService {
	return &shoppingListServiceImpl{db: db}
}

func (sls *shoppingListServiceImpl) CreateShoppingList(spl *models.ShoppingList, loginUser string) (int64, error) {
	id, err := sls.db.ShoppingList().Create(spl, loginUser)
	return id, err
}
