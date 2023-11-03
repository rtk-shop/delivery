package shared

import (
	"io"
	"log"
	"os"
)

var sharedMap = make(map[string][]byte)

const (
	POPULAR_CITIES_FILE = "popular_cities.json"
)

func init() {

	file, err := os.Open("json/" + POPULAR_CITIES_FILE)

	if err != nil {
		log.Println(err)
		panic("failed to open " + POPULAR_CITIES_FILE)
	}

	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		panic("fail to parse popular cities data")
	}

	sharedMap["popularCities"] = data
}
