package services

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"tobuy-app/server/api/models"
	"tobuy-app/server/api/repositories"
	"tobuy-app/server/api/utils/logic"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type ItemsService interface {
	GetAllItems(w http.ResponseWriter, userId int) ([]models.BaseItemResponse, error)
	GetItemById(w http.ResponseWriter, r *http.Request, userId int) (models.BaseItemResponse, error)
	SendAllItemsResponse(w http.ResponseWriter, items *[]models.BaseItemResponse)
	SendItemResponse(w http.ResponseWriter, item *models.BaseItemResponse)
}

type itemsService struct {
	ir repositories.ItemsRepository
	il logic.ItemsLogic
	rl logic.ResponseLogic
}

func NewItemsService(ir repositories.ItemsRepository, il logic.ItemsLogic, rl logic.ResponseLogic) ItemsService {
	return &itemsService{ir, il, rl}
}

func (is *itemsService) GetAllItems(w http.ResponseWriter, userId int) ([]models.BaseItemResponse, error) {
	var items []models.Item
	if err := is.ir.GetAllItems(&items, userId); err != nil {
		return nil, err
	}
	responseItems := is.il.CreateAllItemsResponse(&items)

	return responseItems, nil
}

func (is *itemsService) GetItemById(w http.ResponseWriter, r *http.Request, userId int) (models.BaseItemResponse, error) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		statusCode := http.StatusBadRequest
		errMessage := "invalid item_id"
		// エラーレスポンス送信
		is.rl.SendResponse(w, is.rl.CreateErrorStringResponse(errMessage), statusCode)
		return models.BaseItemResponse{}, err
	}

	var item models.Item
	if err := is.ir.GetItemById(&item, id, userId); err != nil {
		var errMessage string
		var statusCode int
		// https://gorm.io/ja_JP/docs/error_handling.html
		if errors.Is(err, gorm.ErrRecordNotFound) {
			statusCode = http.StatusBadRequest
			errMessage = "該当データは存在しません。"
		} else {
			statusCode = http.StatusInternalServerError
			errMessage = "データ取得に失敗しました。"
		}
		// エラーレスポンス送信
		is.rl.SendResponse(w, is.rl.CreateErrorStringResponse(errMessage), statusCode)
		return models.BaseItemResponse{}, err
	}

	responseItem := is.il.CreateItemResponse(&item)

	return responseItem, nil
}

func (is *itemsService) SendAllItemsResponse(w http.ResponseWriter, items *[]models.BaseItemResponse) {
	var response models.AllItemsResponse
	response.Items = *items
	responseBody, _ := json.Marshal(response)

	is.rl.SendResponse(w, responseBody, http.StatusOK)
}

func (is *itemsService) SendItemResponse(w http.ResponseWriter, item *models.BaseItemResponse) {
	var response models.ItemResponse
	response.Item = *item
	responseBody, _ := json.Marshal(response)
	is.rl.SendResponse(w, responseBody, http.StatusOK)
}
