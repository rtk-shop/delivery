package novaposhta

import (
	"context"
	"log"

	"github.com/allegro/bigcache/v3"
)

func initNovaPoshtaCache() {
	log.Println("init of nova-cache")

	config := bigcache.Config{
		CleanWindow: 0,
		Shards:      1024,
	}

	citiesCache["c1"] = []byte(`[{"id": "w1", "name": "Перше відділення"}]`)

	cache, err := bigcache.New(context.Background(), config)
	if err != nil {
		panic(err)
	}

	warehousesCache = cache
}
