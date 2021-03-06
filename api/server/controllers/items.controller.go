package controllers

import (
	"fmt"
	"net/http"

	"tobuy-app/api/server/services"
)

type ItemsController interface {
	ItemsPage(w http.ResponseWriter, r *http.Request)
	FetchAllItems(w http.ResponseWriter, r *http.Request)
	FetchItemById(w http.ResponseWriter, r *http.Request)
	CreateItem(w http.ResponseWriter, r *http.Request)
	DeleteItem(w http.ResponseWriter, r *http.Request)
	UpdateItem(w http.ResponseWriter, r *http.Request)
}

type itemsController struct {
	is services.ItemsService
	as services.AuthService
}

func NewItemsController(is services.ItemsService, as services.AuthService) ItemsController {
	return &itemsController{is, as}
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
// @Security     JWT
// @Success      200  {object}  models.AllItemsResponse
// @Failure      404
// @Failure      500
// @Router       /items [get]
func (ic *itemsController) FetchAllItems(w http.ResponseWriter, r *http.Request) {
	userId, err := ic.as.GetUserIdFromToken(w, r)
	if err != nil {
		return
	}

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
// @Security     JWT
// @Param        id   path      int  true  "Item ID"
// @Success      200  {object}  models.ItemResponse
// @Failure      404
// @Failure      500
// @Router       /items/{id} [get]
func (ic *itemsController) FetchItemById(w http.ResponseWriter, r *http.Request) {
	userId, err := ic.as.GetUserIdFromToken(w, r)
	if err != nil {
		return
	}

	item, err := ic.is.GetItemById(w, r, userId)
	if err != nil {
		return
	}

	ic.is.SendItemResponse(w, &item)
}

// CreateItem godoc
// @Summary      Create item
// @Description  Create user's item
// @Tags         createItem
// @Accept       json
// @Produce      json
// @Security     JWT
// @Param        requestBody  body      models.MutationItemRequest  true  "Create Item Request"
// @Success      201          {object}  models.ItemResponse
// @Failure      400
// @Failure      500
// @Router       /items [post]
func (ic *itemsController) CreateItem(w http.ResponseWriter, r *http.Request) {
	userId, err := ic.as.GetUserIdFromToken(w, r)
	if err != nil {
		return
	}

	itemResponse, err := ic.is.CreateItem(w, r, userId)
	if err != nil {
		return
	}

	ic.is.SendCreateItemResponse(w, &itemResponse)
}

// DeleteItem godoc
// @Summary      Delete item
// @Description  Delete user's item
// @Tags         deleteItem
// @Accept       json
// @Produce      json
// @Security     JWT
// @Param        id  path  int  true  "Item ID"
// @Success      204
// @Failure      400
// @Failure      500
// @Router       /items/{id}/delete [post]
func (ic *itemsController) DeleteItem(w http.ResponseWriter, r *http.Request) {
	userId, err := ic.as.GetUserIdFromToken(w, r)
	if err != nil {
		return
	}

	if err := ic.is.DeleteItem(w, r, userId); err != nil {
		return
	}

	ic.is.SendDeleteItemResponse(w)
}

// UpdateItem godoc
// @Summary      Update item
// @Description  Update user's item
// @Tags         updateItem
// @Accept       json
// @Produce      json
// @Security     JWT
// @Param        id           path      int                         true  "Item ID"
// @Param        requestBody  body      models.MutationItemRequest  true  "Update Item Request"
// @Success      200          {object}  models.ItemResponse
// @Failure      400
// @Failure      500
// @Router       /items/{id}/update [post]
func (ic *itemsController) UpdateItem(w http.ResponseWriter, r *http.Request) {
	userId, err := ic.as.GetUserIdFromToken(w, r)
	if err != nil {
		return
	}

	itemResponse, err := ic.is.UpdateItem(w, r, userId)
	if err != nil {
		return
	}

	ic.is.SendItemResponse(w, &itemResponse)
}
