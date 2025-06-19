package controllers

import (
	"encoding/json"
	"net/http"
	helper "superapps/helpers"
	"superapps/models"
	"superapps/services"
)

func OrderByCoordinate(w http.ResponseWriter, r *http.Request) {

	data := &models.OrderByCoordinate{}

	err := json.NewDecoder(r.Body).Decode(data)

	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		helper.Response(w, 400, true, "Internal server error ("+err.Error()+")", map[string]any{})
		return
	}

	if data.ShipperContactName == "" {
		helper.Logger("error", "In Server: shipper_contact_name is required")
		helper.Response(w, 400, true, "shipper_contact_name is required", map[string]any{})
		return
	}

	if data.ShipperContactPhone == "" {
		helper.Logger("error", "In Server: shipper_contact_phone is required")
		helper.Response(w, 400, true, "shipper_contact_phone is required", map[string]any{})
		return
	}

	if data.ShipperContactEmail == "" {
		helper.Logger("error", "In Server: shipper_contact_email is required")
		helper.Response(w, 400, true, "shipper_contact_email is required", map[string]any{})
		return
	}

	if data.ShipperOrganization == "" {
		helper.Logger("error", "In Server: shipper_organization is required")
		helper.Response(w, 400, true, "shipper_organization is required", map[string]any{})
		return
	}

	if data.OriginContactName == "" {
		helper.Logger("error", "In Server: origin_contact_name is required")
		helper.Response(w, 400, true, "origin_contact_name is required", map[string]any{})
		return
	}

	if data.OriginContactPhone == "" {
		helper.Logger("error", "In Server: origin_contact_phone is required")
		helper.Response(w, 400, true, "origin_contact_phone is required", map[string]any{})
		return
	}

	if data.OriginAddress == "" {
		helper.Logger("error", "In Server: origin_address is required")
		helper.Response(w, 400, true, "origin_address is required", map[string]any{})
		return
	}

	if data.OriginNote == "" {
		helper.Logger("error", "In Server: origin_note is required")
		helper.Response(w, 400, true, "origin_note is required", map[string]any{})
		return
	}

	if data.OriginCoordinate.Latitude == 0 {
		helper.Logger("error", "In Server: origin_coordinate.latitude is required")
		helper.Response(w, 400, true, "origin_coordinate.latitude is required", map[string]any{})
		return
	}

	if data.OriginCoordinate.Longitude == 0 {
		helper.Logger("error", "In Server: origin_coordinate.longitude is required")
		helper.Response(w, 400, true, "origin_coordinate.longitude is required", map[string]any{})
		return
	}

	if data.DestinationContactName == "" {
		helper.Logger("error", "In Server: destination_contact_name is required")
		helper.Response(w, 400, true, "destination_contact_name is required", map[string]any{})
		return
	}

	if data.DesinationContactPhone == "" {
		helper.Logger("error", "In Server: destination_contact_phone is required")
		helper.Response(w, 400, true, "destination_contact_phone is required", map[string]any{})
		return
	}

	if data.DestinationContactEmail == "" {
		helper.Logger("error", "In Server: destination_contact_email is required")
		helper.Response(w, 400, true, "destination_contact_email is required", map[string]any{})
		return
	}

	if data.DestinationAddress == "" {
		helper.Logger("error", "In Server: destination_address is required")
		helper.Response(w, 400, true, "destination_address is required", map[string]any{})
		return
	}

	if data.DestinationNote == "" {
		helper.Logger("error", "In Server: destination_note is required")
		helper.Response(w, 400, true, "destination_note is required", map[string]any{})
		return
	}

	if data.DestinationCoordinate.Latitude == 0 {
		helper.Logger("error", "In Server: destination_coordinate.latitude is required")
		helper.Response(w, 400, true, "destination_coordinate.latitude is required", map[string]any{})
		return
	}

	if data.DestinationCoordinate.Longitude == 0 {
		helper.Logger("error", "In Server: destination_coordinate.longitude is required")
		helper.Response(w, 400, true, "destination_coordinate.longitude is required", map[string]any{})
		return
	}

	if data.CourierCompany == "" {
		helper.Logger("error", "In Server: courier_company is required")
		helper.Response(w, 400, true, "courier_company is required", map[string]any{})
		return
	}

	if data.CourierType == "" {
		helper.Logger("error", "In Server: courier_type is required")
		helper.Response(w, 400, true, "courier_type is required", map[string]any{})
		return
	}

	if data.DeliveryType == "" {
		helper.Logger("error", "In Server: delivery_type is required")
		helper.Response(w, 400, true, "delivery_type is required", map[string]any{})
		return
	}

	if data.OrderNote == "" {
		helper.Logger("error", "In Server: order_note is required")
		helper.Response(w, 400, true, "order_note is required", map[string]any{})
		return
	}

	if len(data.Items) == 0 {
		helper.Logger("error", "In Server: items is required")
		helper.Response(w, 400, true, "items is required", map[string]any{})
		return
	}

	result, err := services.OrderByCoordinate(data)

	if err != nil {
		helper.Response(w, 400, true, err.Error(), map[string]any{})
		return
	}

	helper.Logger("info", "Order by Coordinate")
	helper.Response(w, http.StatusOK, false, "Successfully", result["data"])
}
