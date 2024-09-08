package models

type ShoppingList struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	OwnerID int64  `json:"owner_id"`
}
