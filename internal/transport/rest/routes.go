package rest

import (
	"net/http"

	"github.com/ReyLegar/ShoppingList/internal/services"
)

func NewRouter(userService services.UserService, shoppingListService services.ShoppingListService, itemService services.ItemService) http.Handler {
	handlerUser := NewUserHandlers(userService)
	handlerShoppingList := NewShoppingListHandlers(shoppingListService)
	handlerItem := NewItemHandlers(itemService)

	mux := http.NewServeMux()

	mux.HandleFunc("/register", handlerUser.HandleRegisterUser)
	mux.HandleFunc("/login", handlerUser.HandleLoginUser)

	mux.HandleFunc("/create_shopping_list", JWTMiddleware(handlerShoppingList.HandleCreateShoppingList))

	mux.HandleFunc("/shopping_list/", JWTMiddleware(handlerItem.HandleCreateItem))

	return mux
}
