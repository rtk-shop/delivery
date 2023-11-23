package novaposhta

type APIResoonse struct {
	Success bool           `json:"success"`
	Data    []WarehouseDTO `json:"data"`
}

type WarehouseDTO struct {
	Ref         string `json:"Ref"`
	Description string `json:"Description"`
	CityRef     string `json:"CityRef"`
}
