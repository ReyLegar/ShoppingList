package models

type SharedList struct {
	ID             int64 `json:"id"`
	ShoppingListID int64 `json:"shopping_list_id"`
	SharedUserID   int64 `json:"shared_user_id"`
}
