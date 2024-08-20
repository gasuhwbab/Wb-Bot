package entities

import "time"

type Coefficient struct {
	Date        time.Time `json:"date"`
	Coefficient int       `json:"coefficient"`
	WarehouseID int       `json:"warehouseID"`
	BoxTypeName string    `json:"boxTypeName"`
	BoxTypeID   int       `json:"boxTypeID"`
}

type CoefficientList struct {
	Coefficients []Coefficient
}
