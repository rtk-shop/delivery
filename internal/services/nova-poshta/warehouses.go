package novaposhta

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"rtk/delivery/internal/entity"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/redis/go-redis/v9"
)

const (
	cityIdMaxLen    = 36
	cahceKey        = "nova_warehouses"
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

func (s *service) Warehouses(cityID string, warehouseType int) ([]byte, error) {

	// validate Warehouses(args)

	if c := utf8.RuneCountInString(cityID); c > cityIdMaxLen {
		return nil, fmt.Errorf("city_id length > %d symbols", cityIdMaxLen)
	}

	warehousesTypeRef, ok := warehouseTypesMap[warehouseType]
	if !ok {
		return nil, fmt.Errorf("unknown warehouse ID - %d", warehouseType)
	}

	key := cahceKey + ":" + cityID + ":" + warehousesTypeRef

	cacheResult, err := s.cache.Get(context.Background(), key).Bytes()
	if err == redis.Nil {

		fmt.Printf("warehouses for city %q not found\n", cityID)
		fmt.Printf("request to Nova Poshta API\n")

		data, err := s.fetchWarehouses(cityID, warehousesTypeRef)
		if err != nil {
			return nil, errors.New("NOVA POSHTA api not availdable")
		}

		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Println(err)
			return nil, fmt.Errorf("faild to marshal json for cache")
		}

		err = s.cache.Set(context.Background(), key, jsonData, 48*time.Hour).Err()
		if err != nil {
			fmt.Println(err)
			return nil, errors.New("failed to store data in cache")
		}

		return jsonData, nil

	} else if err != nil {

		fmt.Println(err)
		return nil, errors.New("failed to get data from cache")

	}

	fmt.Printf("cache: load warehouses for %q\n", cityID)

	return cacheResult, nil

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

	fmt.Printf("total count:%d warehouses len:%d\n", -1, len(warehousesDTO))
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
