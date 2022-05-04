package services

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"tobuy-app/api/server/models"
	"tobuy-app/api/server/repositories"
	"tobuy-app/api/server/utils/logic"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type ItemsService interface {

	// アイテムに関する処理

	GetAllItems(w http.ResponseWriter, userId int) ([]models.BaseItemResponse, error)
	GetItemById(w http.ResponseWriter, r *http.Request, userId int) (models.BaseItemResponse, error)
	CreateItem(w http.ResponseWriter, r *http.Request, userId int) (models.BaseItemResponse, error)
	DeleteItem(w http.ResponseWriter, r *http.Request, userId int) error
	UpdateItem(w http.ResponseWriter, r *http.Request, userId int) (models.BaseItemResponse, error)

	// レスポンス送信

	SendAllItemsResponse(w http.ResponseWriter, items *[]models.BaseItemResponse)
	SendItemResponse(w http.ResponseWriter, item *models.BaseItemResponse)
	SendCreateItemResponse(w http.ResponseWriter, item *models.BaseItemResponse)
	SendDeleteItemResponse(w http.ResponseWriter)
}

type itemsService struct {
	ir repositories.ItemRepository
	il logic.ItemsLogic
	rl logic.ResponseLogic
}

func NewItemsService(ir repositories.ItemRepository, il logic.ItemsLogic, rl logic.ResponseLogic) ItemsService {
	return &itemsService{ir, il, rl}
}

func (is *itemsService) GetAllItems(w http.ResponseWriter, userId int) ([]models.BaseItemResponse, error) {
	var items []models.Item
	if err := is.ir.GetAllItems(&items, userId); err != nil {
		var errMessage string
		var statusCode int
		if errors.Is(err, gorm.ErrRecordNotFound) {
			statusCode = http.StatusBadRequest
			errMessage = "該当データは存在しません。"
		} else {
			statusCode = http.StatusInternalServerError
			errMessage = "データ取得に失敗しました。"
		}
		log.Error().Err(err).Int("status_code", statusCode).Msg(errMessage)
		is.rl.SendResponse(w, is.rl.CreateErrorStringResponse(errMessage), statusCode)
	}
	itemsResponse := is.il.CreateAllItemsResponse(&items)

	return itemsResponse, nil
}

func (is *itemsService) GetItemById(w http.ResponseWriter, r *http.Request, userId int) (models.BaseItemResponse, error) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		statusCode := http.StatusBadRequest
		errMessage := "アイテムIDの読み取りに失敗しました。"
		log.Error().Err(err).Int("status_code", statusCode).Msg(errMessage)
		is.rl.SendResponse(w, is.rl.CreateErrorStringResponse(errMessage), statusCode)
		return models.BaseItemResponse{}, err
	}

	var item models.Item
	if err := is.ir.GetItemById(&item, id, userId); err != nil {
		var errMessage string
		var statusCode int
		if errors.Is(err, gorm.ErrRecordNotFound) {
			statusCode = http.StatusBadRequest
			errMessage = "該当データは存在しません。"
		} else {
			statusCode = http.StatusInternalServerError
			errMessage = "データ取得に失敗しました。"
		}
		log.Error().Err(err).Int("status_code", statusCode).Msg(errMessage)
		is.rl.SendResponse(w, is.rl.CreateErrorStringResponse(errMessage), statusCode)
		return models.BaseItemResponse{}, err
	}

	itemResponse := is.il.CreateItemResponse(&item)

	return itemResponse, nil
}

func (is *itemsService) CreateItem(w http.ResponseWriter, r *http.Request, userId int) (models.BaseItemResponse, error) {
	resBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errMessage := "リクエストボディの読み取り処理でエラー発生"
		log.Error().Err(err).Msg(errMessage)
		is.rl.SendResponse(w, is.rl.CreateErrorStringResponse(errMessage), http.StatusInternalServerError)
		return models.BaseItemResponse{}, err
	}
	var mutationItemRequest models.MutationItemRequest
	if err := json.Unmarshal(resBody, &mutationItemRequest); err != nil {
		errMessage := "リクエストボディの構造体への変換処理でエラー発生"
		log.Error().Err(err).Msg(errMessage)
		is.rl.SendResponse(w, is.rl.CreateErrorStringResponse(errMessage), http.StatusInternalServerError)
		return models.BaseItemResponse{}, err
	}
	// TODO bodyのバリデーション

	var item models.Item
	item.Name = mutationItemRequest.Name
	item.Description = mutationItemRequest.Description
	item.Number = mutationItemRequest.Number
	item.Counter = mutationItemRequest.Counter
	item.Status = mutationItemRequest.Status
	item.CategoryId = mutationItemRequest.CategoryId
	item.UserId = &userId

	if err := is.ir.CreateItem(&item); err != nil {
		errMessage := "アイテム作成に失敗しました。"
		statusCode := http.StatusInternalServerError
		log.Error().Err(err).Int("status_code", statusCode).Msg(errMessage)
		is.rl.SendResponse(w, is.rl.CreateErrorStringResponse(errMessage), statusCode)
		return models.BaseItemResponse{}, err
	}

	if err := is.ir.GetLastItem(&item, userId); err != nil {
		var errMessage string
		var statusCode int
		if errors.Is(err, gorm.ErrRecordNotFound) {
			statusCode = http.StatusNotFound
			errMessage = "該当データは存在しません。"
		} else {
			statusCode = http.StatusInternalServerError
			errMessage = "アイテム取得に失敗しました。"
		}
		log.Error().Err(err).Int("status_code", statusCode).Msg(errMessage)
		is.rl.SendResponse(w, is.rl.CreateErrorStringResponse(errMessage), statusCode)
		return models.BaseItemResponse{}, err
	}

	itemResponse := is.il.CreateItemResponse(&item)

	return itemResponse, nil
}

func (is *itemsService) DeleteItem(w http.ResponseWriter, r *http.Request, userId int) error {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		statusCode := http.StatusBadRequest
		errMessage := "アイテムIDの読み取りに失敗しました。"
		log.Error().Err(err).Int("status_code", statusCode).Msg(errMessage)
		is.rl.SendResponse(w, is.rl.CreateErrorStringResponse(errMessage), statusCode)
		return err
	}
	if err := is.ir.DeleteItem(id, userId); err != nil {
		statusCode := http.StatusInternalServerError
		errMessage := "アイテムの削除に失敗しました。"
		log.Error().Err(err).Int("status_code", statusCode).Msg(errMessage)
		is.rl.SendResponse(w, is.rl.CreateErrorStringResponse(errMessage), statusCode)
		return err
	}
	return nil
}

func (is *itemsService) UpdateItem(w http.ResponseWriter, r *http.Request, userId int) (models.BaseItemResponse, error) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		statusCode := http.StatusBadRequest
		errMessage := "アイテムIDの読み取りに失敗しました。"
		log.Error().Err(err).Int("status_code", statusCode).Msg(errMessage)
		is.rl.SendResponse(w, is.rl.CreateErrorStringResponse(errMessage), statusCode)
		return models.BaseItemResponse{}, err
	}

	resBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		statusCode := http.StatusInternalServerError
		errMessage := "リクエストボディの読み取り処理でエラー発生"
		log.Error().Err(err).Int("status_code", statusCode).Msg(errMessage)
		is.rl.SendResponse(w, is.rl.CreateErrorStringResponse(errMessage), statusCode)
		return models.BaseItemResponse{}, err
	}
	var mutationItemRequest models.MutationItemRequest
	if err := json.Unmarshal(resBody, &mutationItemRequest); err != nil {
		statusCode := http.StatusInternalServerError
		errMessage := "リクエストボディの構造体への変換処理でエラー発生"
		log.Error().Err(err).Int("status_code", statusCode).Msg(errMessage)
		is.rl.SendResponse(w, is.rl.CreateErrorStringResponse(errMessage), statusCode)
		return models.BaseItemResponse{}, err
	}
	// TODO バリデーション

	var updateItem models.Item
	updateItem.Name = mutationItemRequest.Name
	updateItem.Description = mutationItemRequest.Description
	updateItem.Number = mutationItemRequest.Number
	updateItem.Counter = mutationItemRequest.Counter
	updateItem.Status = mutationItemRequest.Status
	updateItem.CategoryId = mutationItemRequest.CategoryId
	updateItem.UserId = &userId

	if err := is.ir.UpdateItem(&updateItem, id, userId); err != nil {
		statusCode := http.StatusInternalServerError
		errMessage := "アイテムの更新に失敗しました。"
		log.Error().Err(err).Int("status_code", statusCode).Msg(errMessage)
		is.rl.SendResponse(w, is.rl.CreateErrorStringResponse(errMessage), statusCode)
		return models.BaseItemResponse{}, err
	}

	var item models.Item
	if err := is.ir.GetItemById(&item, id, userId); err != nil {
		var errMessage string
		var statusCode int
		if errors.Is(err, gorm.ErrRecordNotFound) {
			statusCode = http.StatusNotFound
			errMessage = "該当データは存在しません。"
		} else {
			statusCode = http.StatusInternalServerError
			errMessage = "アイテム取得に失敗しました。"
		}
		log.Error().Err(err).Int("status_code", statusCode).Msg(errMessage)
		is.rl.SendResponse(w, is.rl.CreateErrorStringResponse(errMessage), statusCode)
		return models.BaseItemResponse{}, err
	}

	itemResponse := is.il.CreateItemResponse(&item)

	return itemResponse, nil
}

// レスポンス送信

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

func (is *itemsService) SendCreateItemResponse(w http.ResponseWriter, item *models.BaseItemResponse) {
	var response models.ItemResponse
	response.Item = *item
	responseBody, _ := json.Marshal(response)
	is.rl.SendResponse(w, responseBody, http.StatusCreated)
}

func (is *itemsService) SendDeleteItemResponse(w http.ResponseWriter) {
	is.rl.SendNotBodyResponse(w)
}
