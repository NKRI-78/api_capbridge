package controllers

import (
	"encoding/json"
	"net/http"
	"superapps/entities"
	helper "superapps/helpers"
	service "superapps/services"
)

func ProjectStore(w http.ResponseWriter, r *http.Request) {

	data := &entities.ProjectStore{}

	err := json.NewDecoder(r.Body).Decode(data)

	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		helper.Response(w, 400, true, "Internal server error ("+err.Error()+")", map[string]interface{}{})
		return
	}

	Title := data.Title
	Goal := data.Goal
	Capital := data.Capital
	Roi := data.Roi
	MinInvest := data.MinInvest
	UnitPrice := data.UnitPrice
	UnitTotal := data.UnitTotal
	NumberOfUnit := data.NumberOfUnit
	Periode := data.Periode

	if Title == "" {
		helper.Logger("error", "In Server: title is required")
		helper.Response(w, 400, true, "title is required", map[string]any{})
		return
	}

	if Goal == "" {
		helper.Logger("error", "In Server: goal is required")
		helper.Response(w, 400, true, "goal is required", map[string]any{})
		return
	}

	if Capital == "" {
		helper.Logger("error", "In Server: goal is required")
		helper.Response(w, 400, true, "goal is required", map[string]any{})
		return
	}

	if Roi == "" {
		helper.Logger("error", "In Server: roi is required")
		helper.Response(w, 400, true, "roi is required", map[string]any{})
		return
	}

	if MinInvest == "" {
		helper.Logger("error", "In Server: min_invest is required")
		helper.Response(w, 400, true, "min_invest is required", map[string]any{})
		return
	}

	if UnitPrice == "" {
		helper.Logger("error", "In Server: unit_price is required")
		helper.Response(w, 400, true, "unit_price is required", map[string]any{})
		return
	}

	if UnitTotal == "" {
		helper.Logger("error", "In Server: unit_total is required")
		helper.Response(w, 400, true, "unit_total is required", map[string]any{})
		return
	}

	if NumberOfUnit == "" {
		helper.Logger("error", "In Server: number_of_unit is required")
		helper.Response(w, 400, true, "number_of_unit is required", map[string]any{})
		return
	}

	if Periode == "" {
		helper.Logger("error", "In Server: periode is required")
		helper.Response(w, 400, true, "periode is required", map[string]any{})
		return
	}

	result, err := service.ProjectStore(data)

	if err != nil {
		helper.Response(w, 400, true, err.Error(), map[string]any{})
		return
	}

	helper.Logger("info", "Project Store success")
	helper.Response(w, http.StatusOK, false, "Successfully", result["data"])
}
