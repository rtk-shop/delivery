package shared

import (
	"context"
	"fmt"
	"log"
)

func (s *service) PopularCities() ([]byte, error) {

	res, err := s.cache.JSONGet(context.Background(), "popular_cities", ".").Result()
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("failed to get popular cities")
	}

	return []byte(res), nil
}
