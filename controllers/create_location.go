package controllers

import (
	"encoding/json"
	"net/http"
	helper "superapps/helpers"
	"superapps/models"
	"superapps/services"
)

func CreateLocation(w http.ResponseWriter, r *http.Request) {

	data := &models.CreateLocation{}

	err := json.NewDecoder(r.Body).Decode(data)

	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		helper.Response(w, 400, true, "Internal server error ("+err.Error()+")", map[string]interface{}{})
		return
	}

	Name := data.Name
	ContactName := data.ContactName
	ContactPhone := data.ContactPhone
	Address := data.Address
	Note := data.Note
	PostalCode := data.PostalCode
	Latitude := data.Latitude
	Longitude := data.Longitude
	Type := data.Type
	UserId := data.UserId

	if Name == "" {
		helper.Logger("error", "In Server: name is required")
		helper.Response(w, 400, true, "name is required", map[string]any{})
		return
	}

	if ContactName == "" {
		helper.Logger("error", "In Server: contact_name is required")
		helper.Response(w, 400, true, "contact_name is required", map[string]any{})
		return
	}

	if ContactPhone == "" {
		helper.Logger("error", "In Server: contact_phone is required")
		helper.Response(w, 400, true, "contact_phone is required", map[string]any{})
		return
	}

	if Address == "" {
		helper.Logger("error", "In Server: address is required")
		helper.Response(w, 400, true, "address is required", map[string]any{})
		return
	}

	if Note == "" {
		helper.Logger("error", "In Server: note is required")
		helper.Response(w, 400, true, "note is required", map[string]any{})
		return
	}

	if PostalCode == 0 {
		helper.Logger("error", "In Server: postal_code is required")
		helper.Response(w, 400, true, "postal_code is required", map[string]any{})
		return
	}

	if Latitude == 0 {
		helper.Logger("error", "In Server: latitude is required")
		helper.Response(w, 400, true, "latitude is required", map[string]any{})
		return
	}

	if Longitude == 0 {
		helper.Logger("error", "In Server: latitude is required")
		helper.Response(w, 400, true, "latitude is required", map[string]any{})
		return
	}

	if Type == "" {
		helper.Logger("error", "In Server: type is required")
		helper.Response(w, 400, true, "type is required", map[string]any{})
		return
	}

	data.UserId = UserId

	result, err := services.CreateLocation(data)

	if err != nil {
		helper.Response(w, 400, true, err.Error(), map[string]any{})
		return
	}

	helper.Logger("info", "Form Region success")
	helper.Response(w, http.StatusOK, false, "Successfully", result["data"])
}
