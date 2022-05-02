package router

import (
	"net/http"
	"tobuy-app/server/api/controllers"

	"github.com/gorilla/mux"
)

type ItemsRouter interface {
	SetItemsRouting(router *mux.Router)
}

type itemsRouter struct {
	ic controllers.ItemsController
}

func NewItemsRouter(ic controllers.ItemsController) ItemsRouter {
	return &itemsRouter{ic}
}

func (ir *itemsRouter) SetItemsRouting(router *mux.Router) {
	// router.HandleFunc(basePath+"/items", ir.ic.ItemsPage).Methods("GET")
	router.HandleFunc(basePath+"/items", ir.ic.FetchAllItems).Methods(http.MethodGet)
	router.HandleFunc(basePath+"/items/{id}", ir.ic.FetchItemById).Methods(http.MethodGet)
}
