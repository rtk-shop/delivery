package novaposhta

import "github.com/allegro/bigcache/v3"

var citiesCache = make(map[string][]byte)
var warehousesCache *bigcache.BigCache

func init() {
	initNovaPoshtaCache()
}
