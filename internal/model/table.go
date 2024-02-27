// model/table.go

package model

type Table struct {
	TableIDUUID      string `json:"table_id_uuid"`      // Previously ID
	RestaurantIDUUID string `json:"restaurant_id_uuid"` // Previously RestaurantID
	QRCodeStatic     string `json:"qr_code_static"`     // As specified
	QRCodeDynamic    string `json:"qr_code_dynamic"`    // As specified
	IsActive         bool   `json:"is_active"`          // As specified
}
