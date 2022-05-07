package logic

import (
	"encoding/json"
	"net/http"
)

type IResponseLogic interface {
	SendResponse(w http.ResponseWriter, response []byte, code int)
	SendNotBodyResponse(w http.ResponseWriter)
	CreateErrorResponse(err error) []byte
	CreateErrorStringResponse(errMessage string) []byte
}

type ResponseLogic struct{}

func NewResponseLogic() *ResponseLogic {
	return &ResponseLogic{}
}

/*
APIレスポンス送信処理
*/
func (rl *ResponseLogic) SendResponse(w http.ResponseWriter, response []byte, code int) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

/*
APIレスポンス送信処理 (レスポンスBodyなし)
*/
func (rl *ResponseLogic) SendNotBodyResponse(w http.ResponseWriter) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

/*
エラーレスポンス作成
*/
func (rl *ResponseLogic) CreateErrorResponse(err error) []byte {
	response := map[string]interface{}{
		"error": err,
	}
	responseBody, _ := json.Marshal(response)

	return responseBody
}

/*
エラーレスポンス作成 (エラーメッセージはstring)
*/
func (rl *ResponseLogic) CreateErrorStringResponse(errMessage string) []byte {
	response := map[string]interface{}{
		"error": errMessage,
	}
	responseBody, _ := json.Marshal(response)

	return responseBody
}
