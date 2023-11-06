package novaposhta

import (
	"bags2on/delivery/internal/utils"
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/allegro/bigcache/v3"
)

func initNovaPoshtaCache() {
	log.Println("init of nova-cache")

	config := bigcache.Config{
		CleanWindow: 0,
		Shards:      1024,
	}

	warehouses, err := bigcache.New(context.Background(), config)
	if err != nil {
		panic(err)
	}

	warehousesCache = loadWarehouses(warehouses)
	log.Println("✅ warehouses cahce -", utils.PrettyByteSize(warehouses.Capacity()))

}

func loadWarehouses(cache *bigcache.BigCache) *bigcache.BigCache {
	dirEntries, err := os.ReadDir("json/warehouses/nova")
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

		data, err := os.ReadFile("json/warehouses/nova/" + entryName)
		if err != nil {
			log.Println(err)
			panic("fail to proccess " + entryName)
		}

		err = cache.Set(utils.FileNameWithoutExt(entryName), data)
		if err != nil {
			log.Println(err)
			panic("fail to set data of " + entryName)
		}
	}

	return cache
}
