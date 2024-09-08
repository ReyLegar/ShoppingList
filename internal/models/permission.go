package models

type Permission struct {
	ID             int64 `json:"id"`
	ShoppingListID int64 `json:"shopping_list_id"`
	UserID         int64 `json:"user_id"`
	CanEdit        bool  `json:"can_edit"`
}
