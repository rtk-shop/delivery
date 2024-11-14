package novaposhta

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"rtk/delivery/internal/entity"
	"rtk/delivery/internal/services/nova-poshta/mock"
)

const cahceKey = "nova_warehouses:"

func (s *service) Warehouses(cityID string) ([]byte, error) {

	key := cahceKey + cityID
	ctx := context.Background()

	cacheResult, err := s.cache.Get(ctx, key).Result()
	if err != nil {
		fmt.Printf("warehouses for city %q not found\n", cityID)
		fmt.Printf("request Nova Poshta API\n")

		result, err := s.fetchFromAPI(cityID)
		if err != nil {
			fmt.Println(err)
			return nil, fmt.Errorf("faild to fetch API")
		}

		jsonData, err := json.Marshal(result)
		if err != nil {
			log.Println(err)
			return nil, fmt.Errorf("faild to marshal json for cache")
		}

		if ok := json.Valid(jsonData); !ok {
			return nil, fmt.Errorf("final json isn't valid")
		}

		_, err = s.cache.Set(ctx, key, jsonData, 0).Result()
		if err != nil {
			log.Println(err)
			return nil, fmt.Errorf("failed to set json for cache")
		}

		return jsonData, nil

	}

	fmt.Printf("cache: load warehouses for %q\n", cityID)

	return []byte(cacheResult), nil
}

func processWarehouses(warehouses []WarehouseDTO) []entity.Warehouse {

	dst := make([]entity.Warehouse, 0, len(warehouses))

	for _, w := range warehouses {
		dst = append(dst, entity.Warehouse{
			ID:          w.Ref,
			Description: w.Description,
		})
	}

	return dst
}

func (s *service) fetchFromAPI(cityID string) ([]entity.Warehouse, error) {

	/*
		"errors": [],
		"warnings": [],
		"info": {
			"totalCount": 850
		},
		"messageCodes": [],
		"errorCodes": [],
		"warningCodes": [],
		"infoCodes": []
	*/

	reqBodyString := fmt.Sprintf(`{
		"apiKey": "%s",
		"modelName": "Address",
		"calledMethod": "getWarehouses",
		"methodProperties": {
			"Language": "UA",
			"CityRef": "%s"
		}
	}`, s.apiKey, cityID)

	reqBodyJson := []byte(reqBodyString)

	// fmt.Println(reqBodyString)

	req, err := http.NewRequest(http.MethodGet, s.config.NovaPoshtaURL, bytes.NewReader(reqBodyJson))
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("faild to build request to Nova Poshta API")
	}

	// resp, err := http.DefaultClient.Do(req)
	resp, err := mock.MockHttpClient.Do(req)
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

	// fmt.Println(string(respData)[:100])

	var result APIResoonse

	if err = json.Unmarshal(respData, &result); err != nil {
		log.Println(err)
		return nil, fmt.Errorf("faild to read json")
	}

	if !result.Success {
		return nil, fmt.Errorf("request was failed because of logical error")
	}

	// fmt.Println(result)

	if result.Data[0].CityRef != cityID {
		return nil, fmt.Errorf("the data does not match the city ID")
	}

	return processWarehouses(result.Data), nil
}
