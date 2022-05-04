package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"tobuy-app/api/server/controllers"
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
	// FetchAllItems
	router.HandleFunc(basePath+"/items", ir.ic.FetchAllItems).Methods(http.MethodGet)
	// CreateItem
	router.HandleFunc(basePath+"/items", ir.ic.CreateItem).Methods(http.MethodPost)
	// FetchItemById
	router.HandleFunc(basePath+"/items/{id}", ir.ic.FetchItemById).Methods(http.MethodGet)
	// DeleteItem
	router.HandleFunc(basePath+"/items/{id}/delete", ir.ic.DeleteItem).Methods(http.MethodPost)
	// UpdateItem
	router.HandleFunc(basePath+"/items/{id}/update", ir.ic.UpdateItem).Methods(http.MethodPost)
}
