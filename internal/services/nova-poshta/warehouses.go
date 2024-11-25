package novaposhta

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
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

		s.logger.Warn("warehouses not found, request..", "city_id", cityID)

		data, err := s.fetchWarehouses(cityID, warehousesTypeRef)
		if err != nil {
			return nil, errors.New("NOVA POSHTA api not availdable")
		}

		jsonData, err := json.Marshal(data)
		if err != nil {
			s.logger.Error("warehouses marshal body", "error", err)
			return nil, errors.New("faild to marshal json for cache")
		}

		err = s.cache.Set(context.Background(), key, jsonData, 336*time.Hour).Err() // two weeks
		if err != nil {
			s.logger.Error("cache set warehouses", "error", err)
			return nil, errors.New("failed to store data")
		}

		return jsonData, nil

	} else if err != nil {

		s.logger.Error("cache get warehouses", "error", err)
		return nil, errors.New("failed to get data from cache")

	}

	s.logger.Info("cache: load warehouses for", "city_id", cityID)

	return cacheResult, nil

}

func (s *service) fetchWarehouses(cityID string, warehouseType string) ([]entity.NovaPoshtaWarehouse, error) {

	reqBodyString := fmt.Sprintf(warehousesBodyF, s.apiKey, cityID, warehouseType)

	req, err := http.NewRequest(http.MethodGet, s.config.NovaPoshtaURL, strings.NewReader(reqBodyString))
	if err != nil {
		s.logger.Error("warehouses build API request", "error", err)
		return nil, errors.New("faild to build request")
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		s.logger.Error("warehouses DO API request", "error", err)
		return nil, errors.New("faild to DO request to Nova Poshta API")
	}

	defer resp.Body.Close()

	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		s.logger.Error("warehouses read Body", "error", err)
		return nil, errors.New("faild to read body data")
	}

	var respDTO getWarehousesDTO

	if err = json.Unmarshal(respData, &respDTO); err != nil {
		s.logger.Error("warehouses unmarshal Body", "error", err)
		return nil, errors.New("faild to read json")
	}

	if !respDTO.Success {

		s.logger.Error("warehouses requst success:false", "error", string(respData))
		return nil, fmt.Errorf("api request %q was failed", "getWarehouses")

	}

	warehousesDTO := respDTO.Data

	s.logger.Warn("Nova Poshta API: loaded warehouses", "processed", len(warehousesDTO))

	warehouses := make([]entity.NovaPoshtaWarehouse, 0, len(warehousesDTO))

	for _, novaWarehouse := range warehousesDTO {

		warehouses = append(warehouses, entity.NovaPoshtaWarehouse{
			ID:   novaWarehouse.Ref,
			Name: novaWarehouse.Description,
		})

	}

	return warehouses, nil
}
