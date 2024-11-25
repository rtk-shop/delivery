package novaposhta

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"rtk/delivery/internal/entity"
	"strings"
	"unicode/utf8"
)

const (
	cityNameMaxLen   = 36
	settlementsBodyF = `{
	   "apiKey": "%s",
	   "modelName": "AddressGeneral",
	   "calledMethod": "searchSettlements",
	   "methodProperties": {
			"CityName" : "%s",
			"Limit" : "50",
			"Page" : "1"
		}
	}
	`
)

// DOC: https://developers.novaposhta.ua/view/model/a0cf0f5f-8512-11ec-8ced-005056b2dbe1/method/a0eb83ab-8512-11ec-8ced-005056b2dbe1

func (s *service) Settlements(cityName string) ([]entity.NovaPoshtaSettlement, error) {

	if cityName == "" {
		return nil, errors.New("city_name empty")
	}

	if c := utf8.RuneCountInString(cityName); c > cityNameMaxLen {
		return nil, fmt.Errorf("city_name length > %d symbols", cityNameMaxLen)
	}

	re := regexp.MustCompile(`[a-zA-Z]`)
	if ok := re.MatchString(cityName); ok {
		return nil, errors.New("city_name Cyrillic only")
	}

	s.logger.Info("search settlements for", "city_name", cityName)

	reqBodyString := fmt.Sprintf(settlementsBodyF, s.apiKey, cityName)

	req, err := http.NewRequest(http.MethodGet, s.config.NovaPoshtaURL, strings.NewReader(reqBodyString))
	if err != nil {
		s.logger.Error("search-settlements build request", "error", err)
		return nil, errors.New("faild to build request to Nova Poshta API")
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		s.logger.Error("search-settlements DO request", "error", err)
		return nil, errors.New("faild to DO request to Nova Poshta API")
	}

	defer resp.Body.Close()

	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		s.logger.Error("search-settlements read Body", "error", err)
		return nil, errors.New("faild to read body data")
	}

	// fmt.Println(string(respData)[:100])

	var respDTO searchSettlementsApiResponse

	if err = json.Unmarshal(respData, &respDTO); err != nil {
		s.logger.Error("search-settlements unmarshal body", "error", err)
		return nil, errors.New("faild to read json")
	}

	if !respDTO.Success {
		return nil, fmt.Errorf("api request %q was failed", "searchSettlements")
	}

	target := respDTO.Data[0]

	settlements := make([]entity.NovaPoshtaSettlement, 0, len(target.Addresses))

	for _, novaSettlement := range target.Addresses {

		if novaSettlement.AddressDeliveryAllowed {

			settlements = append(settlements, entity.NovaPoshtaSettlement{
				ID:   novaSettlement.DeliveryCity,
				Name: novaSettlement.Present,
			})

		}

	}

	return settlements, nil

}
