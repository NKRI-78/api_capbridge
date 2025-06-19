package controllers

import (
	"encoding/json"
	"net/http"
	helper "superapps/helpers"
	"superapps/models"
	"superapps/services"
)

func RateByCoordinate(w http.ResponseWriter, r *http.Request) {

	data := &models.RateByCoordinate{}

	err := json.NewDecoder(r.Body).Decode(data)

	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		helper.Response(w, 400, true, "Internal server error ("+err.Error()+")", map[string]interface{}{})
		return
	}

	if data.OriginLatitude == "" {
		helper.Logger("error", "In Server: origin_latitude is required")
		helper.Response(w, 400, true, "origin_latitude is required", map[string]any{})
		return
	}

	if data.OriginLongitude == "" {
		helper.Logger("error", "In Server: origin_longitude is required")
		helper.Response(w, 400, true, "origin_longitude is required", map[string]any{})
		return
	}

	if data.Couriers == "" {
		helper.Logger("error", "In Server: couriers is required")
		helper.Response(w, 400, true, "couriers is required", map[string]any{})
		return
	}

	if data.DestinationLatitude == "" {
		helper.Logger("error", "In Server: destination_latitude is required")
		helper.Response(w, 400, true, "destination_latitude is required", map[string]any{})
		return
	}

	if data.DestinationLongitude == "" {
		helper.Logger("error", "In Server: destination_longitude is required")
		helper.Response(w, 400, true, "destination_longitude is required", map[string]any{})
		return
	}

	if len(data.Items) == 0 {
		helper.Logger("error", "In Server: items is required")
		helper.Response(w, 400, true, "items is required", map[string]any{})
		return
	}

	result, err := services.RateByCoordinate(data)

	if err != nil {
		helper.Response(w, 400, true, err.Error(), map[string]any{})
		return
	}

	helper.Logger("info", "Rate by Coordinate")
	helper.Response(w, http.StatusOK, false, "Successfully", result["data"])
}
