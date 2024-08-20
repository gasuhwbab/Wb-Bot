package entities

type Warehouse struct {
	ID   int    `json:"ID"`
	Name string `json:"name"`
}

type WarehouseList struct {
	Warehouses []Warehouse
}
