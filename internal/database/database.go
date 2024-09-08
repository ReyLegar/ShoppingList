package database

type Repository interface {
	User() UserRepository
	ShoppingList() ShoppingListRepository
	Item() ItemRepository
}
