package controllers

import (
	"encoding/json"
	"net/http"
	"superapps/entities"
	helper "superapps/helpers"
	"superapps/services"
)

func Login(w http.ResponseWriter, r *http.Request) {

	data := &entities.Login{}

	err := json.NewDecoder(r.Body).Decode(data)

	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		helper.Response(w, 400, true, "Internal server error ("+err.Error()+")", map[string]interface{}{})
		return
	}

	Email := data.Email
	Password := data.Password

	if Email == "" {
		helper.Logger("error", "In Server: email is required")
		helper.Response(w, 400, true, "email is required", map[string]any{})
		return
	}

	if Password == "" {
		helper.Logger("error", "In Server: password is required")
		helper.Response(w, 400, true, "password is required", map[string]any{})
		return
	}

	result, err := services.Login(data)

	if err != nil {
		helper.Response(w, 400, true, err.Error(), map[string]any{})
		return
	}

	helper.Logger("info", "Login success")
	helper.Response(w, http.StatusOK, false, "Successfully", result["data"])
}
