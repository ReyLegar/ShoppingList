package models

type Item struct {
	ID             int64  `json:"id"`
	ShoppingListID int64  `json:"shopping_list_id"`
	Name           string `json:"name"`
	Quantity       int64  `json:"quantity"`
	IsPurchased    bool   `json:"is_purchased"`
}
