package novaposhta

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"rtk/delivery/internal/entity"
	"strings"
)

const (
	cahceKey        = "nova_warehouses:"
	warehousesBodyF = `{
		"apiKey": "%s",
		"modelName": "AddressGeneral",
		"calledMethod": "getWarehouses",
		"methodProperties": {
			"CityRef": "%s",
			"Page": "",
			"Limit": "",
			"Language": "UA",
			"TypeOfWarehouseRef": "%s"
		}
	}`
)

// [36]
// "TypeOfWarehouseRef": "841339c7-591a-42e2-8233-7a0a00f0ed6f"

func (s *service) Warehouses(cityID string, warehouseType int) ([]entity.NovaPoshtaWarehouse, error) {

	// validate Warehouses(args)

	// key := cahceKey + cityID
	// ctx := context.Background()

	// cacheResult, err := s.cache.Get(ctx, key).Result()
	// if err != nil {
	// 	fmt.Printf("warehouses for city %q not found\n", cityID)
	// 	fmt.Printf("request Nova Poshta API\n")

	// 	result, err := s.fetchFromAPI(cityID)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return nil, fmt.Errorf("faild to fetch API")
	// 	}

	// 	jsonData, err := json.Marshal(result)
	// 	if err != nil {
	// 		log.Println(err)
	// 		return nil, fmt.Errorf("faild to marshal json for cache")
	// 	}

	// 	if ok := json.Valid(jsonData); !ok {
	// 		return nil, fmt.Errorf("final json isn't valid")
	// 	}

	// 	_, err = s.cache.Set(ctx, key, jsonData, 0).Result()
	// 	if err != nil {
	// 		log.Println(err)
	// 		return nil, fmt.Errorf("failed to set json for cache")
	// 	}

	// 	return jsonData, nil

	// }

	// fmt.Printf("cache: load warehouses for %q\n", cityID)

	warehousesTypeRef, ok := warehouseTypesMap[warehouseType]
	if !ok {
		return nil, fmt.Errorf("unknown warehouse ID - %d", warehouseType)
	}

	data, err := s.fetchWarehouses(cityID, warehousesTypeRef)

	if err != nil {
		return nil, errors.New("NOVA POSHTA api not availdable")
	}

	return data, nil

}

func (s *service) fetchWarehouses(cityID string, warehouseType string) ([]entity.NovaPoshtaWarehouse, error) {

	reqBodyString := fmt.Sprintf(warehousesBodyF, s.apiKey, cityID, warehouseType)

	req, err := http.NewRequest(http.MethodGet, s.config.NovaPoshtaURL, strings.NewReader(reqBodyString))
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("faild to build request to Nova Poshta API")
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("faild to DO request to Nova Poshta API")
	}

	defer resp.Body.Close()

	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("faild to read body data")
	}

	var respDTO getWarehousesDTO

	if err = json.Unmarshal(respData, &respDTO); err != nil {
		log.Println(err)
		return nil, fmt.Errorf("faild to read json")
	}

	if !respDTO.Success {

		fmt.Println(reqBodyString)
		fmt.Println(string(respData))

		return nil, fmt.Errorf("request %q was failed", "getWarehouses")

	}

	warehousesDTO := respDTO.Data

	fmt.Printf("total count:%d warehouses len:%d", -1, len(warehousesDTO))
	// fmt.Printf("total count:%d warehouses len:%d", respDTO.Info.TotalCount, len(warehousesDTO))

	warehouses := make([]entity.NovaPoshtaWarehouse, 0, len(warehousesDTO))

	for _, novaWarehouse := range warehousesDTO {

		warehouses = append(warehouses, entity.NovaPoshtaWarehouse{
			ID:   novaWarehouse.Ref,
			Name: novaWarehouse.Description,
		})

	}

	return warehouses, nil
}
