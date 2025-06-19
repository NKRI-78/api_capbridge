package controllers

import (
	"net/http"
	helper "superapps/helpers"
	service "superapps/services"
)

func CourierRateList(w http.ResponseWriter, r *http.Request) {

	originLat := r.URL.Query().Get("origin_lat")
	originLng := r.URL.Query().Get("origin_lng")
	destLat := r.URL.Query().Get("dest_lat")
	destLng := r.URL.Query().Get("dest_lng")

	if originLat == "" {
		helper.Logger("error", "In Server: origin_lat is required")
		helper.Response(w, 400, true, "origin_lat is required", map[string]any{})
		return
	}

	if originLng == "" {
		helper.Logger("error", "In Server: origin_lng is required")
		helper.Response(w, 400, true, "origin_lng is required", map[string]any{})
		return
	}

	if destLat == "" {
		helper.Logger("error", "In Server: dest_lat is required")
		helper.Response(w, 400, true, "dest_lat is required", map[string]any{})
		return
	}

	if destLng == "" {
		helper.Logger("error", "In Server: dest_lng is required")
		helper.Response(w, 400, true, "dest_lng is required", map[string]any{})
		return
	}

	result, err := service.CourierRateList(originLat, originLng, destLat, destLng)

	if err != nil {
		helper.Response(w, 400, true, err.Error(), map[string]any{})
		return
	}

	helper.Logger("info", "Get All Courier success")
	helper.Response(w, http.StatusOK, false, "Successfully", result["data"])
}
