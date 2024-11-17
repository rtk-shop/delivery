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

type getWarehousesDTO struct {
	Success bool `json:"success"`
	Data    []struct {
		Ref             string `json:"Ref"`
		Description     string `json:"Description"`
		TypeOfWarehouse string `json:"TypeOfWarehouse"`
	} `json:"data"`
	Errors   []any `json:"errors"`
	Warnings []any `json:"warnings"`
	// Info     struct {
	// 	TotalCount int `json:"totalCount"`
	// } `json:"info"`
	// MessageCodes []any `json:"messageCodes"`
	// ErrorCodes   []any `json:"errorCodes"`
	// WarningCodes []any `json:"warningCodes"`
	// InfoCodes    []any `json:"infoCodes"`
}

type searchSettlementsDataDTO struct {
	TotalCount int `json:"TotalCount"`
	Addresses  []struct {
		Present                string `json:"Present"`
		DeliveryCity           string `json:"DeliveryCity"`
		AddressDeliveryAllowed bool   `json:"AddressDeliveryAllowed"`
		// Warehouses             int    `json:"Warehouses"`
		// MainDescription        string `json:"MainDescription"`
		// Area                   string `json:"Area"`
		// Region                 string `json:"Region"`
		// SettlementTypeCode     string `json:"SettlementTypeCode"`
		// Ref                    string `json:"Ref"`
		// StreetsAvailability    bool   `json:"StreetsAvailability"`
		// ParentRegionTypes      string `json:"ParentRegionTypes"`
		// ParentRegionCode       string `json:"ParentRegionCode"`
		// RegionTypes            string `json:"RegionTypes"`
		// RegionTypesCode        string `json:"RegionTypesCode"`
	} `json:"Addresses"`
}

type searchSettlementsApiResponse struct {
	Success bool                       `json:"success"`
	Data    []searchSettlementsDataDTO `json:"data"`
}
