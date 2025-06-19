package entities

import "time"

type Courier struct {
	Success  bool       `json:"success"`
	Object   string     `json:"object"`
	Couriers []Couriers `json:"couriers"`
}

type Couriers struct {
	AvailableCollectionMethod    []AvailableCollectionMethod `json:"available_collection_method"`
	AvailableForCashOnDelivery   bool                        `json:"available_for_cash_on_delivery"`
	AvailableForProofOfDelivery  bool                        `json:"available_for_proof_of_delivery"`
	AvailableForInstantWaybillId bool                        `json:"available_for_instant_waybill_id"`
	CourierName                  string                      `json:"courier_name"`
	CourierCode                  string                      `json:"courier_code"`
	CourierServiceName           string                      `json:"courier_service_name"`
	CourierServiceCode           string                      `json:"courier_service_code"`
	Tier                         string                      `json:"premium"`
	Description                  string                      `json:"description"`
	ServiceType                  string                      `json:"service_type"`
	ShippingType                 string                      `json:"shipping_type"`
	ShipmentDurationRange        string                      `json:"shipment_duration_range"`
	ShipmentDurationUnit         string                      `json:"shipment_duration_unit"`
}

type AvailableCollectionMethod string

type CourierRate struct {
	OriginLatitude       string `json:"origin_latitude"`
	OriginLangitude      string `json:"origin_longitude"`
	DestinationLatitude  string `json:"destination_latitude"`
	DestinationLongitude string `json:"destination_longitude"`
}

type CreateLocation struct {
	Name         string  `json:"name"`
	ContactName  string  `json:"contact_name"`
	ContactPhone string  `json:"contact_phone"`
	Address      string  `json:"address"`
	Note         string  `json:"note"`
	PostalCode   int     `json:"postal_code"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	Type         string  `json:"type"`
	UserId       string  `json:"user_id"`
}

type CreateLocationResponse struct {
	Success      bool      `json:"success"`
	Id           string    `json:"id"`
	Name         string    `json:"name"`
	Address      string    `json:"address"`
	Note         string    `json:"note"`
	PostalCode   int       `json:"postal_code"`
	ContactName  string    `json:"contact_name"`
	ContactPhone string    `json:"contact_phone"`
	Owned        bool      `json:"owned"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type RateByCoordinate struct {
	OriginLatitude       string                  `json:"origin_latitude"`
	OriginLongitude      string                  `json:"origin_longitude"`
	DestinationLatitude  string                  `json:"destination_latitude"`
	DestinationLongitude string                  `json:"destination_longitude"`
	Couriers             string                  `json:"couriers"`
	Items                []RateByCoordinateItems `json:"items"`
}

type RateByCoordinateItems struct {
	Name     string `json:"name"`
	Weight   int    `json:"weight"`
	Value    int    `json:"value"`
	Quantity int    `json:"quantity"`
}

type RateByCoordinateResponse struct {
	Success     bool                                `json:"success"`
	Object      string                              `json:"object"`
	Message     string                              `json:"message"`
	Code        int64                               `json:"code"`
	Origin      RateByCoordinateResponseOrigin      `json:"origin"`
	Stops       []any                               `json:"stops"`
	Destination RateByCoordinateResponseDestination `json:"destination"`
	Pricing     []RateByCoordinateResponsePricing   `json:"pricing"`
}

type RateByCoordinateResponseOrigin struct {
	LocationID                        any    `json:"location_id"`
	Latitude                          string `json:"latitude"`
	Longitude                         string `json:"longitude"`
	PostalCode                        any    `json:"postal_code"`
	AdministrativeDivisionLevel1_Name any    `json:"administrative_division_level_1_name"`
	AdministrativeDivisionLevel1_Type string `json:"administrative_division_level_1_type"`
	AdministrativeDivisionLevel2_Name any    `json:"administrative_division_level_2_name"`
	AdministrativeDivisionLevel2_Type string `json:"administrative_division_level_2_type"`
	AdministrativeDivisionLevel3_Name any    `json:"administrative_division_level_3_name"`
	AdministrativeDivisionLevel3_Type string `json:"administrative_division_level_3_type"`
	AdministrativeDivisionLevel4_Name any    `json:"administrative_division_level_4_name"`
	AdministrativeDivisionLevel4_Type string `json:"administrative_division_level_4_type"`
	Address                           any    `json:"address"`
}

type RateByCoordinateResponseDestination struct {
	LocationID                        any    `json:"location_id"`
	Latitude                          string `json:"latitude"`
	Longitude                         string `json:"longitude"`
	PostalCode                        any    `json:"postal_code"`
	AdministrativeDivisionLevel1_Name any    `json:"administrative_division_level_1_name"`
	AdministrativeDivisionLevel1_Type string `json:"administrative_division_level_1_type"`
	AdministrativeDivisionLevel2_Name any    `json:"administrative_division_level_2_name"`
	AdministrativeDivisionLevel2_Type string `json:"administrative_division_level_2_type"`
	AdministrativeDivisionLevel3_Name any    `json:"administrative_division_level_3_name"`
	AdministrativeDivisionLevel3_Type string `json:"administrative_division_level_3_type"`
	AdministrativeDivisionLevel4_Name any    `json:"administrative_division_level_4_name"`
	AdministrativeDivisionLevel4_Type string `json:"administrative_division_level_4_type"`
	Address                           any    `json:"address"`
}

type RateByCoordinateResponsePricing struct {
	AvailableCollectionMethod    []string `json:"available_collection_method"`
	AvailableForCashOnDelivery   bool     `json:"available_for_cash_on_delivery"`
	AvailableForProofOfDelivery  bool     `json:"available_for_proof_of_delivery"`
	AvailableForInstantWaybillID bool     `json:"available_for_instant_waybill_id"`
	AvailableForInsurance        bool     `json:"available_for_insurance"`
	Company                      string   `json:"company"`
	CourierName                  string   `json:"courier_name"`
	CourierCode                  string   `json:"courier_code"`
	CourierServiceName           string   `json:"courier_service_name"`
	CourierServiceCode           string   `json:"courier_service_code"`
	Currency                     string   `json:"currency"`
	Description                  string   `json:"description"`
	Duration                     string   `json:"duration"`
	ShipmentDurationRange        string   `json:"shipment_duration_range"`
	ShipmentDurationUnit         string   `json:"shipment_duration_unit"`
	ServiceType                  string   `json:"service_type"`
	ShippingType                 string   `json:"shipping_type"`
	Price                        int64    `json:"price"`
	TaxLines                     []any    `json:"tax_lines"`
	Type                         string   `json:"type"`
}

type OrderByCoordinate struct {
	Id                      string                   `json:"id"`
	ShipperContactName      string                   `json:"shipper_contact_name"`
	ShipperContactPhone     string                   `json:"shipper_contact_phone"`
	ShipperContactEmail     string                   `json:"shipper_contact_email"`
	ShipperOrganization     string                   `json:"shipper_organization"`
	OriginContactName       string                   `json:"origin_contact_name"`
	OriginContactPhone      string                   `json:"origin_contact_phone"`
	OriginAddress           string                   `json:"origin_address"`
	OriginNote              string                   `json:"origin_note"`
	OriginCoordinate        OriginCoordinate         `json:"origin_coordinate"`
	DestinationContactName  string                   `json:"destination_contact_name"`
	DesinationContactPhone  string                   `json:"destination_contact_phone"`
	DestinationContactEmail string                   `json:"destination_contact_email"`
	DestinationAddress      string                   `json:"destination_address"`
	DestinationNote         string                   `json:"destination_note"`
	DestinationCoordinate   DestinationCoordinate    `json:"destination_coordinate"`
	CourierCompany          string                   `json:"courier_company"`
	CourierType             string                   `json:"courier_type"`
	CourierInsurance        string                   `json:"courier_insurance"`
	DeliveryType            string                   `json:"delivery_type"`
	OrderNote               string                   `json:"order_note"`
	MetaData                any                      `json:"metadata"`
	Items                   []OrderByCoordinateItems `json:"items"`
}

type OrderInfo struct {
	Id string `json:"id"`
}

type OrderInfoResponse struct {
	Success      bool                  `json:"success"`
	Message      string                `json:"message"`
	Object       string                `json:"object"`
	ID           string                `json:"id"`
	ShortID      string                `json:"short_id"`
	Shipper      OrderInfoShipperInfo  `json:"shipper"`
	Origin       OrderInfoLocation     `json:"origin"`
	Destination  OrderInfoDestination  `json:"destination"`
	Delivery     OrderInfoDeliveryInfo `json:"delivery"`
	Voucher      OrderInfoVoucherInfo  `json:"voucher"`
	Courier      OrderInfoCourierInfo  `json:"courier"`
	ReferenceID  *string               `json:"reference_id"`
	InvoiceID    *string               `json:"invoice_id"`
	Items        []Item                `json:"items"`
	Extra        any                   `json:"extra"`
	Metadata     any                   `json:"metadata"`
	Tags         []string              `json:"tags"`
	Note         string                `json:"note"`
	Currency     string                `json:"currency"`
	TaxLines     []any                 `json:"tax_lines"`
	Price        int                   `json:"price"`
	Status       string                `json:"status"`
	TicketStatus *string               `json:"ticket_status"`
}

type OrderInfoShipperInfo struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Organization string `json:"organization"`
}

type OrderInfoLocation struct {
	ContactName  string              `json:"contact_name"`
	ContactPhone string              `json:"contact_phone"`
	Address      string              `json:"address"`
	Note         string              `json:"note"`
	PostalCode   int                 `json:"postal_code"`
	Coordinate   OrderInfoCoordinate `json:"coordinate"`
}

type OrderInfoCoordinate struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type OrderInfoDestination struct {
	ContactName     string                   `json:"contact_name"`
	ContactPhone    string                   `json:"contact_phone"`
	ContactEmail    string                   `json:"contact_email"`
	Address         string                   `json:"address"`
	Note            string                   `json:"note"`
	PostalCode      int                      `json:"postal_code"`
	Coordinate      OrderInfoCoordinate      `json:"coordinate"`
	ProofOfDelivery OrderInfoProofOfDelivery `json:"proof_of_delivery"`
	CashOnDelivery  OrderInfoCashOnDelivery  `json:"cash_on_delivery"`
}

type OrderInfoProofOfDelivery struct {
	Use  bool    `json:"use"`
	Fee  int     `json:"fee"`
	Note *string `json:"note"`
	Link *string `json:"link"`
}

type OrderInfoCashOnDelivery struct {
	ID             *string `json:"id"`
	Amount         int     `json:"amount"`
	AmountCurrency string  `json:"amount_currency"`
	Fee            int     `json:"fee"`
	FeeCurrency    string  `json:"fee_currency"`
	Note           *string `json:"note"`
	Type           *string `json:"type"`
}

type OrderInfoDeliveryInfo struct {
	Datetime     string  `json:"datetime"`
	Note         *string `json:"note"`
	Type         string  `json:"type"`
	Distance     float64 `json:"distance"`
	DistanceUnit string  `json:"distance_unit"`
}

type OrderInfoVoucherInfo struct {
	ID    *string `json:"id"`
	Name  *string `json:"name"`
	Value *string `json:"value"`
	Type  *string `json:"type"`
}

type OrderInfoCourierInfo struct {
	TrackingID        string           `json:"tracking_id"`
	WaybillID         string           `json:"waybill_id"`
	Company           string           `json:"company"`
	History           []CourierHistory `json:"history"`
	Link              string           `json:"link"`
	Name              string           `json:"name"`  // Deprecated
	Phone             string           `json:"phone"` // Deprecated
	DriverName        string           `json:"driver_name"`
	DriverPhone       string           `json:"driver_phone"`
	DriverPhotoURL    string           `json:"driver_photo_url"`
	DriverPlateNumber string           `json:"driver_plate_number"`
	Type              string           `json:"type"`
	ShipmentFee       int              `json:"shipment_fee"`
	Insurance         InsuranceInfo    `json:"insurance"`
}

type CourierHistory struct {
	ServiceType string `json:"service_type"`
	Status      string `json:"status"`
	Note        string `json:"note"`
	UpdatedAt   string `json:"updated_at"`
}

type InsuranceInfo struct {
	Amount         int     `json:"amount"`
	AmountCurrency string  `json:"amount_currency"`
	Fee            int     `json:"fee"`
	FeeCurrency    string  `json:"fee_currency"`
	Note           *string `json:"note"`
}

type OrderInfoItem struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	SKU         *string `json:"sku"`
	Value       int     `json:"value"`
	Quantity    int     `json:"quantity"`
	Length      int     `json:"length"`
	Width       int     `json:"width"`
	Height      int     `json:"height"`
	Weight      int     `json:"weight"`
}

type OriginCoordinate struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type DestinationCoordinate struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type OrderByCoordinateItems struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Value       int    `json:"value"`
	Quantity    int    `json:"quantity"`
	Weight      int    `json:"weight"`
}

type OrderByCoordinateResponse struct {
	Success     bool                        `json:"success"`
	Message     string                      `json:"message"`
	Object      string                      `json:"object"`
	ID          string                      `json:"id"`
	Shipper     ShipperOrderByCoordinate    `json:"shipper"`
	Origin      LocationOrderByCoordinate   `json:"origin"`
	Destination DestinationOrdeByCoordinate `json:"destination"`
	Courier     Courier                     `json:"courier"`
	Delivery    Delivery                    `json:"delivery"`
	ReferenceID *string                     `json:"reference_id"`
	Items       []Item                      `json:"items"`
	Extra       []any                       `json:"extra"`
	Currency    string                      `json:"currency"`
	TaxLines    []any                       `json:"tax_lines"`
	Price       int                         `json:"price"`
	Metadata    map[string]any              `json:"metadata"`
	Note        string                      `json:"note"`
	Status      string                      `json:"status"`
}

type ShipperOrderByCoordinate struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Organization string `json:"organization"`
}

type LocationOrderByCoordinate struct {
	ContactName  string     `json:"contact_name"`
	ContactPhone string     `json:"contact_phone"`
	Coordinate   Coordinate `json:"coordinate"`
	Address      string     `json:"address"`
	Note         string     `json:"note"`
	PostalCode   int        `json:"postal_code"`
}

type DestinationOrdeByCoordinate struct {
	ContactName     string          `json:"contact_name"`
	ContactPhone    string          `json:"contact_phone"`
	ContactEmail    string          `json:"contact_email"`
	Address         string          `json:"address"`
	Note            string          `json:"note"`
	ProofOfDelivery ProofOfDelivery `json:"proof_of_delivery"`
	CashOnDelivery  CashOnDelivery  `json:"cash_on_delivery"`
	Coordinate      Coordinate      `json:"coordinate"`
	PostalCode      int             `json:"postal_code"`
}

type Coordinate struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type ProofOfDelivery struct {
	Use  bool    `json:"use"`
	Fee  int     `json:"fee"`
	Note *string `json:"note"`
	Link *string `json:"link"`
}

type CashOnDelivery struct {
	ID             string  `json:"id"`
	Amount         int     `json:"amount"`
	AmountCurrency string  `json:"amount_currency"`
	Fee            int     `json:"fee"`
	FeeCurrency    string  `json:"fee_currency"`
	Note           *string `json:"note"`
	Type           string  `json:"type"`
}

type CourierOrderByCoordinate struct {
	TrackingID        string    `json:"tracking_id"`
	WaybillID         *string   `json:"waybill_id"`
	Company           string    `json:"company"`
	Name              *string   `json:"name"`
	Phone             *string   `json:"phone"`
	DriverName        *string   `json:"driver_name"`
	DriverPhone       *string   `json:"driver_phone"`
	DriverPhotoURL    *string   `json:"driver_photo_url"`
	DriverPlateNumber *string   `json:"driver_plate_number"`
	Type              string    `json:"type"`
	Link              *string   `json:"link"`
	Insurance         Insurance `json:"insurance"`
	RoutingCode       *string   `json:"routing_code"`
}

type Insurance struct {
	Amount         int    `json:"amount"`
	AmountCurrency string `json:"amount_currency"`
	Fee            int    `json:"fee"`
	FeeCurrency    string `json:"fee_currency"`
	Note           string `json:"note"`
}

type Delivery struct {
	Datetime     string  `json:"datetime"`
	Note         *string `json:"note"`
	Type         string  `json:"type"`
	Distance     float64 `json:"distance"`
	DistanceUnit string  `json:"distance_unit"`
}

type Item struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	SKU         *string `json:"sku"`
	Value       int     `json:"value"`
	Quantity    int     `json:"quantity"`
	Length      int     `json:"length"`
	Width       int     `json:"width"`
	Height      int     `json:"height"`
	Weight      int     `json:"weight"`
}

type CourierPricingResponse struct {
	Success     bool            `json:"success"`
	Object      string          `json:"object"`
	Message     string          `json:"message"`
	Code        int             `json:"code"`
	Origin      LocationPricing `json:"origin"`
	Destination LocationPricing `json:"destination"`
	Pricing     []Pricing       `json:"pricing"`
}

type LocationPricing struct {
	LocationID                       string `json:"location_id"`
	Latitude                         string `json:"latitude"`
	Longitude                        string `json:"longitude"`
	PostalCode                       int    `json:"postal_code"`
	CountryName                      string `json:"country_name"`
	CountryCode                      string `json:"country_code"`
	AdministrativeDivisionLevel1Name string `json:"administrative_division_level_1_name"`
	AdministrativeDivisionLevel1Type string `json:"administrative_division_level_1_type"`
	AdministrativeDivisionLevel2Name string `json:"administrative_division_level_2_name"`
	AdministrativeDivisionLevel2Type string `json:"administrative_division_level_2_type"`
	AdministrativeDivisionLevel3Name string `json:"administrative_division_level_3_name"`
	AdministrativeDivisionLevel3Type string `json:"administrative_division_level_3_type"`
	AdministrativeDivisionLevel4Name string `json:"administrative_division_level_4_name"`
	AdministrativeDivisionLevel4Type string `json:"administrative_division_level_4_type"`
	Address                          string `json:"address"`
}

type Pricing struct {
	AvailableForCashOnDelivery   bool     `json:"available_for_cash_on_delivery"`
	AvailableForProofOfDelivery  bool     `json:"available_for_proof_of_delivery"`
	AvailableForInstantWaybillID bool     `json:"available_for_instant_waybill_id"`
	AvailableForInsurance        bool     `json:"available_for_insurance"`
	AvailableCollectionMethod    []string `json:"available_collection_method"`
	Company                      string   `json:"company"`
	CourierName                  string   `json:"courier_name"`
	CourierCode                  string   `json:"courier_code"`
	CourierServiceName           string   `json:"courier_service_name"`
	CourierServiceCode           string   `json:"courier_service_code"`
	Currency                     string   `json:"currency"`
	Description                  string   `json:"description"`
	Duration                     string   `json:"duration"`
	ShipmentDurationRange        string   `json:"shipment_duration_range"`
	ShipmentDurationUnit         string   `json:"shipment_duration_unit"`
	ServiceType                  string   `json:"service_type"`
	ShippingType                 string   `json:"shipping_type"`
	Price                        int      `json:"price"`
	TaxLines                     []any    `json:"tax_lines"`
	Type                         string   `json:"type"`
}

type GetTracking struct {
	Id string `json:"id"`
}

type Tracking struct {
	Success     bool                   `json:"success"`
	Message     string                 `json:"messsage"`
	Object      string                 `json:"object"`
	ID          string                 `json:"id"`
	WaybillID   string                 `json:"waybill_id"`
	Courier     TrackingCourierInfo    `json:"courier"`
	Origin      TrackingLocationInfo   `json:"origin"`
	Destination TrackingLocationInfo   `json:"destination"`
	History     []TrackingHistoryEntry `json:"history"`
	Link        string                 `json:"link"`
	OrderID     string                 `json:"order_id"`
	Status      string                 `json:"status"`
}

type TrackingCourierInfo struct {
	Company           string `json:"company"`
	Name              string `json:"name"`  // Deprecated
	Phone             string `json:"phone"` // Deprecated
	DriverName        string `json:"driver_name"`
	DriverPhone       string `json:"driver_phone"`
	DriverPhotoURL    string `json:"driver_photo_url"`
	DriverPlateNumber string `json:"driver_plate_number"`
}

type TrackingLocationInfo struct {
	ContactName string `json:"contact_name"`
	Address     string `json:"address"`
}

type TrackingHistoryEntry struct {
	Note        string `json:"note"`
	ServiceType string `json:"service_type"`
	UpdatedAt   string `json:"updated_at"`
	Status      string `json:"status"`
}
