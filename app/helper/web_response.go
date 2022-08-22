package helper

import (
	"encoding/json"
	"go-crud/app/model/web"
	"net/http"
)

func WebResponse(w http.ResponseWriter, code int, status string, data interface{}) {

	webResponse := web.WebResponse{
		Code:   code,
		Status: status,
		Data:   data,
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	encoder := json.NewEncoder(w)
	encoder.Encode(webResponse)
}
