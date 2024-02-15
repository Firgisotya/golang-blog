package response

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseAuth struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Token   string      `json:"token"`
}

func ResponseToken(w http.ResponseWriter, status int, message string, data interface{}, token string) {
	res := ResponseAuth{
		Status:  status,
		Message: message,
		Data:    data,
		Token:   token,
	}

	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Fatal(err)
	}
}

func ResponseMessage(w http.ResponseWriter, status int, message string, data interface{}) {
	res := Response{
		Status:  status,
		Message: message,
		Data:    data,
	}

	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Fatal(err)
	}
}

func ResponseError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)

	res := Response{
		Message: err.Error(),
		Data: nil,
	}

	err = json.NewEncoder(w).Encode(res)

	if err != nil {
		log.Fatal(err)
	}
}