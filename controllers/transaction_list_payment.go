package controllers

import (
	"net/http"
	helper "superapps/helpers"
	service "superapps/services"
)

func TransactionListPayment(w http.ResponseWriter, r *http.Request) {

	orderId := r.URL.Query().Get("order_id")

	result, err := service.TransactionListPayment(orderId)

	if err != nil {
		helper.Response(w, 400, true, err.Error(), map[string]any{})
		return
	}

	helper.Logger("info", "Get All Transaction Payment success")
	helper.Response(w, http.StatusOK, false, "Successfully", result["data"])
}
