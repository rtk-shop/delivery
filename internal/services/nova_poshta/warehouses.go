package novaposhta

import "fmt"

func (s *service) Warehouses(cityID string) ([]byte, error) {

	data, ok := novaposhtaCache[cityID]
	if !ok {
		return nil, fmt.Errorf("data for city %q not found", cityID)
	}

	return data, nil
}
