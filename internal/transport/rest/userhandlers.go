package rest

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/ReyLegar/ShoppingList/internal/models"
	"github.com/ReyLegar/ShoppingList/internal/services"
)

type UserHandlers struct {
	userService services.UserService
}

func NewUserHandlers(userService services.UserService) *UserHandlers {
	return &UserHandlers{userService: userService}
}

func (h *UserHandlers) HandleRegisterUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Не тот метод", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Не удалось прочитать тело запроса", http.StatusBadRequest)
		return
	}

	var user models.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, "Ошибка при разборе JSON", http.StatusBadRequest)
		return
	}

	id, err := h.userService.RegisterUser(&user)
	if err != nil {
		http.Error(w, "Ошибка при создании пользователя", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]int64{"Пользователь успешно зарегистрирован": id}
	json.NewEncoder(w).Encode(response)
}

func (h *UserHandlers) HandleLoginUser(w http.ResponseWriter, r *http.Request) {
	var creds models.Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Неверный запрос", http.StatusBadRequest)
		return
	}

	tokens := h.userService.LoginUser(&creds)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tokens)

}
