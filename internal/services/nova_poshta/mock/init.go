package mock

import (
	"bags2on/delivery/internal/utils"
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/allegro/bigcache/v3"
)

var warehousesCache *bigcache.BigCache

var citiesCache = map[string]string{
	"kharkiv_warehouses": "db5c88e0-391c-11dd-90d9-001a92567626",
	"kuiv_warehouses":    "8d5a980d-391c-11dd-90d9-001a92567626",
}

func initNovaPoshtaCache() {

	warehouses, err := bigcache.New(context.Background(), bigcache.Config{
		CleanWindow: 0,
		Shards:      2048,
	})
	if err != nil {
		log.Fatal(err)
	}

	warehousesCache = loadWarehouses(warehouses, "test-data/nova")
	log.Println("✅ warehouses cahce -", utils.PrettyByteSize(warehouses.Capacity()))

}

func loadWarehouses(cache *bigcache.BigCache, path string) *bigcache.BigCache {

	dirEntries, err := os.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	for _, entry := range dirEntries {
		if entry.IsDir() {
			continue
		}

		entryName := entry.Name()

		if filepath.Ext(entryName) != ".json" {
			continue
		}

		data, err := os.ReadFile(path + "/" + entryName)
		if err != nil {
			log.Println(err)
			panic("fail to proccess " + entryName)
		}

		keyName := utils.FileNameWithoutExt(entryName)

		cityID, ok := citiesCache[keyName]
		if !ok {
			panic(fmt.Sprintf("unexpected city key - %q", keyName))
		}

		// fmt.Printf("city id: %s for %s\n", cityID, keyName)

		err = cache.Set(cityID, data)
		if err != nil {
			log.Println(err)
			panic("fail to set data of " + entryName)
		}

	}

	return cache
}
