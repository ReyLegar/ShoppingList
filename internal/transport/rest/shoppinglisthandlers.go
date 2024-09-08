package rest

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/ReyLegar/ShoppingList/internal/models"
	"github.com/ReyLegar/ShoppingList/internal/services"
)

type ShoppingListHandlers struct {
	shoppingListService services.ShoppingListService
}

func NewShoppingListHandlers(shoppingListService services.ShoppingListService) *ShoppingListHandlers {
	return &ShoppingListHandlers{shoppingListService: shoppingListService}
}

func (h *ShoppingListHandlers) HandleCreateShoppingList(w http.ResponseWriter, r *http.Request) {

	login, ok := r.Context().Value(userContentKey).(string)

	if !ok {
		http.Error(w, "Не удалось получить пользователя", http.StatusInternalServerError)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Не тот метод", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Не удалось прочитать тело запроса", http.StatusBadRequest)
		return
	}

	var list models.ShoppingList
	err = json.Unmarshal(body, &list)
	if err != nil {
		http.Error(w, "Ошибка при разборе JSON", http.StatusBadRequest)
		return
	}

	id, err := h.shoppingListService.CreateShoppingList(&list, login)
	if err != nil {
		http.Error(w, "Ошибка при создании пользователя", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]int64{"Список успешно создан": id}
	json.NewEncoder(w).Encode(response)
}
