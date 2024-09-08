package services

import (
	"github.com/ReyLegar/ShoppingList/internal/database"
	"github.com/ReyLegar/ShoppingList/internal/models"
)

type ItemService interface {
	CreateItem(item *models.Item) (int64, error)
}

type ItemServiceImpl struct {
	db database.Repository
}

func NewItemService(db database.Repository) ItemService {
	return &ItemServiceImpl{db: db}
}

func (isi *ItemServiceImpl) CreateItem(item *models.Item) (int64, error) {
	id, err := isi.db.Item().Create(item)
	return id, err
}
