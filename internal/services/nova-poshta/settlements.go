package novaposhta

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
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
		return nil, fmt.Errorf("city_name empty")
	}

	if c := utf8.RuneCountInString(cityName); c > cityNameMaxLen {
		return nil, fmt.Errorf("city_name length > %d symbols", cityNameMaxLen)
	}

	re := regexp.MustCompile(`[a-zA-Z]`)
	if ok := re.MatchString(cityName); ok {
		return nil, fmt.Errorf("city_name Cyrillic only")
	}

	log.Println("get settlements by", cityName)

	reqBodyString := fmt.Sprintf(settlementsBodyF, s.apiKey, cityName)

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

	// fmt.Println(string(respData)[:100])

	var respDTO searchSettlementsApiResponse

	if err = json.Unmarshal(respData, &respDTO); err != nil {
		log.Println(err)
		return nil, fmt.Errorf("faild to read json")
	}

	if !respDTO.Success {
		return nil, fmt.Errorf("request %q was failed", "searchSettlements")
	}

	target := respDTO.Data[0]

	fmt.Printf("total count:%d addresses len:%d", target.TotalCount, len(target.Addresses))

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
