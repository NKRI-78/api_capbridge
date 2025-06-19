package controllers

import (
	"net/http"
	helper "superapps/helpers"
	service "superapps/services"
)

func TransactionListPPOB(w http.ResponseWriter, r *http.Request) {

	result, err := service.TransactionListPPOB()

	if err != nil {
		helper.Response(w, 400, true, err.Error(), map[string]any{})
		return
	}

	helper.Logger("info", "Get All Transaction PPOB success")
	helper.Response(w, http.StatusOK, false, "Successfully", result["data"])
}
