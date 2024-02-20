package middlewares

import (
	"errors"
	"go-blog/api/helpers"
	"go-blog/api/payloads/response"
	"net/http"
)

func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		_, err := helpers.Authorization(r)

		if err != nil {
			response.ResponseError(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}

		next(w, r)
	}

}
