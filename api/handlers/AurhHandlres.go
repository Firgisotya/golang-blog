package handlers

import (
	"encoding/json"
	"go-blog/api/configs"
	"go-blog/api/helpers"
	"go-blog/api/payloads/request"
	"go-blog/api/payloads/response"
	"go-blog/api/services"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	db, err := configs.SetupConnection()

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	defer configs.CloseConnection(db)

	var loginRequest request.LoginRequest
	err = json.NewDecoder(r.Body).Decode(&loginRequest)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	err = loginRequest.Validate()

	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	user, err := services.LoginService(db, loginRequest)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	if user.ID == 0 {
		hashPass := user.Password
		pwd := loginRequest.Password

		hash := helpers.VerifyPassword(hashPass, pwd)

		if hash == nil {
			token, err := helpers.CreateToken(user.ID, user.Email)

			if err != nil {
				response.ResponseError(w, http.StatusInternalServerError, err)
				return
			}

			response.ResponseToken(w, http.StatusOK, "Login Success", user, token)
		} else {
			response.ResponseError(w, http.StatusInternalServerError, err)
			return
		}
	}

}
