package novaposhta

import (
	"bags2on/delivery/internal/entity"
	"bags2on/delivery/internal/services/nova_poshta/mock"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

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

type APIResoonse struct {
	Success bool           `json:"success"`
	Data    []WarehouseDTO `json:"data"`
}

func (s *service) Warehouses(cityID string) ([]byte, error) {

	data, err := warehousesCache.Get(cityID)
	if err != nil {
		fmt.Printf("warehouses for city %q not found\n", cityID)
		fmt.Printf("request Nova Poshta API\n")

		reqBodyString := fmt.Sprintf(`{
			"apiKey": "%s",
			"modelName": "Address",
			"calledMethod": "getWarehouses",
			"methodProperties": {
				"Language": "UA",
				"CityRef": "%s"
			}
		}`, s.config.NovaPoshtaKey, cityID)

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

		wReady := processWarehouses(result.Data)

		file, err := os.Create("json/warehouses/nova/" + cityID + ".json")
		if err != nil {
			log.Println(err)
			return nil, fmt.Errorf("faild to create json file")
		}

		defer file.Close()

		encoder := json.NewEncoder(file)
		encoder.Encode(wReady)

		bt, err := json.Marshal(wReady)
		if err != nil {
			log.Println(err)
			return nil, fmt.Errorf("faild to marshal json for cache")
		}

		err = warehousesCache.Set(cityID, bt)
		if err != nil {
			log.Println(err)
			return nil, fmt.Errorf("faild to set json for cache")
		}

		w, err := warehousesCache.Get(cityID)
		if err != nil {
			log.Println(err)
			return nil, fmt.Errorf("warehouses are not in the cache after insertion")
		}

		return w, nil

	}

	fmt.Printf("cache: load warehouses for %q\n", cityID)

	return data, nil
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
