package handlers

import (
	"go-blog/api/payloads/response"
	"net/http"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	response.ResponseMessage(w, 200, "Connected to REST Blog", nil)
}
