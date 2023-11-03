package novaposhta

var novaposhtaCache = make(map[string][]byte)

func init() {
	novaposhtaCache["c1"] = []byte(`[{"id": "w1", "name": "Перше відділення"}]`)
}
