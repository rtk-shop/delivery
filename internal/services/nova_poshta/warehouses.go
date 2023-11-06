package novaposhta

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type TempJson struct {
	Name string `json:"name"`
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

		fmt.Println(req)

		// res, err := http.DefaultClient.Do(req)
		// if err != nil {
		// 	log.Println(err)
		// 	return nil, fmt.Errorf("faild to DO request to Nova Poshta API")
		// }

		temp := []byte(`[{"name": "User"}, {"name": "Admin"}]`)

		var result []TempJson

		if err = json.Unmarshal(temp, &result); err != nil {
			log.Println(err)
			return nil, fmt.Errorf("faild to read json")
		}

		fmt.Println(result)

		file, err := os.Create("json/warehouses/nova/" + cityID + ".json")
		if err != nil {
			log.Println(err)
			return nil, fmt.Errorf("faild to create json file")
		}

		defer file.Close()

		encoder := json.NewEncoder(file)
		encoder.Encode(result)

		bt, err := json.Marshal(result)
		if err != nil {
			log.Println(err)
			return nil, fmt.Errorf("faild to marshal json for cache")
		}

		err = warehousesCache.Set(cityID, bt)
		if err != nil {
			log.Println(err)
			return nil, fmt.Errorf("faild to marshal json for cache")
		}

		w, err := warehousesCache.Get(cityID)
		if err != nil {
			log.Println(err)
			return nil, fmt.Errorf("warehouses are not in the cache")
		}

		return w, nil

		// return nil, fmt.Errorf("warehouses for city %q not found", cityID)

	}

	return data, nil
}
