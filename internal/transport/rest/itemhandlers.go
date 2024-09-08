package rest

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"

	"github.com/ReyLegar/ShoppingList/internal/models"
	"github.com/ReyLegar/ShoppingList/internal/services"
)

type ItemHandlers struct {
	itemService services.ItemService
}

func NewItemHandlers(itemService services.ItemService) *ItemHandlers {
	return &ItemHandlers{itemService: itemService}
}

func (ih *ItemHandlers) HandleCreateItem(w http.ResponseWriter, r *http.Request) {

	_, ok := r.Context().Value(userContentKey).(string)

	re := regexp.MustCompile(`^/shopping_list/(\d+)/items$`)
	matches := re.FindStringSubmatch(r.URL.Path)

	if len(matches) != 2 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	shoppingListId, err := strconv.ParseInt(matches[1], 10, 64)
	if err != nil {
		http.Error(w, "Invalid shopping list ID", http.StatusBadRequest)
		return
	}

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

	var item models.Item
	err = json.Unmarshal(body, &item)
	fmt.Println(item)
	if err != nil {
		http.Error(w, "Ошибка при разборе JSON", http.StatusBadRequest)
		return
	}

	item.ShoppingListID = shoppingListId

	id, err := ih.itemService.CreateItem(&item)
	if err != nil {
		http.Error(w, "Ошибка при создании продукта", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]int64{"Список успешно создан": id}
	json.NewEncoder(w).Encode(response)
}
