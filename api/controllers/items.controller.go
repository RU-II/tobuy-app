package controllers

import (
	"fmt"
	"net/http"
	"tobuy-app/api/services"
)

type ItemsController interface {
	ItemsPage(w http.ResponseWriter, r *http.Request)
	FetchAllItems(w http.ResponseWriter, r *http.Request)
	FetchItemById(w http.ResponseWriter, r *http.Request)
	// CreateItem(w http.ResponseWriter, r *http.Request)
	// DeleteItem(w http.ResponseWriter, r *http.Request)
	// UpdateItem(w http.ResponseWriter, r *http.Request)
}

type itemsController struct {
	is services.ItemsService
}

func NewItemsController(is services.ItemsService) ItemsController {
	return &itemsController{is}
}

func (itc *itemsController) ItemsPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Items")
	fmt.Println("Items endpoint is hooked!")
}

// FetchAllItems godoc
// @Summary      Fetch all items
// @Description  Fetch all items
// @Tags         fetchAllItems
// @Produce      json
// @Success      200
// @Router       /items [get]
func (ic *itemsController) FetchAllItems(w http.ResponseWriter, r *http.Request) {
	userId := 1

	allItems, err := ic.is.GetAllItems(w, userId)
	if err != nil {
		return
	}

	ic.is.SendAllItemsResponse(w, &allItems)
}

// FetchItemById godoc
// @Summary      Fetch item by id
// @Description  Fetch user's item by id
// @Tags         fetchItemBuyId
// @Produce      json
// @Param        id   path      int  true  "Item Id"
// @Success      200  {object}  models.ItemResponse
// @Failure      400
// @Router       /items/{id} [get]
func (ic *itemsController) FetchItemById(w http.ResponseWriter, r *http.Request) {
	userId := 1

	item, err := ic.is.GetItemById(w, r, userId)
	if err != nil {
		return
	}

	ic.is.SendItemResponse(w, &item)
}
