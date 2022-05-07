package logic

import "tobuy-app/api/server/models"

type IItemsLogic interface {
	CreateAllItemsResponse(items *[]models.Item) []models.BaseItemResponse
	CreateItemResponse(item *models.Item) models.BaseItemResponse
}

type ItemsLogic struct{}

func NewItemsLogic() *ItemsLogic {
	return &ItemsLogic{}
}

func (il *ItemsLogic) CreateAllItemsResponse(items *[]models.Item) []models.BaseItemResponse {
	var responseItems []models.BaseItemResponse
	for _, item := range *items {
		var newItem models.BaseItemResponse
		newItem.BaseModel.ID = item.BaseModel.ID
		newItem.BaseModel.CreatedAt = item.BaseModel.CreatedAt
		newItem.BaseModel.UpdatedAt = item.BaseModel.UpdatedAt
		newItem.BaseModel.DeletedAt = item.BaseModel.DeletedAt
		newItem.Name = item.Name
		newItem.Description = item.Description
		newItem.Number = item.Number
		newItem.Counter = item.Counter
		newItem.Status = item.Status
		newItem.CategoryId = item.CategoryId

		responseItems = append(responseItems, newItem)
	}

	return responseItems
}

func (il *ItemsLogic) CreateItemResponse(item *models.Item) models.BaseItemResponse {
	var responseItem models.BaseItemResponse
	responseItem.BaseModel.ID = item.BaseModel.ID
	responseItem.BaseModel.CreatedAt = item.BaseModel.CreatedAt
	responseItem.BaseModel.UpdatedAt = item.BaseModel.UpdatedAt
	responseItem.BaseModel.DeletedAt = item.BaseModel.DeletedAt
	responseItem.Name = item.Name
	responseItem.Description = item.Description
	responseItem.Number = item.Number
	responseItem.Counter = item.Counter
	responseItem.Status = item.Status
	responseItem.CategoryId = item.CategoryId

	return responseItem
}
