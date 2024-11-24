package shared

func (s *service) PopularCities() ([]byte, error) {
	return s.populdarCities, nil
}

func (s *service) GetPopularCitiesHash() string {
	return s.populdarCitiesHash
}
