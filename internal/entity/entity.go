package entity

type PopolarCity struct {
	ID          string `json:"city_id"`
	CityName    string `json:"city_name"`
	NovaPohtaID string `json:"nova_poshta_id"`
	UkrPohtaID  string `json:"ukrposhta_id"`
}

type Warehouse struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}

type NovaPoshtaSettlement struct {
	ID   string `json:"settlement_id"`
	Name string `json:"name"`
}

type NovaPoshtaWarehouse struct {
	ID   string `json:"warehouse_id"`
	Name string `json:"name"`
}
