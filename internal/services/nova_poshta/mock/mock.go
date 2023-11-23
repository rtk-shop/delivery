package mock

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

var (
	MockHttpClient HttpClient
)

func init() {
	MockHttpClient = &Client{}
	initNovaPoshtaCache()
}

func (m *Client) Do(req *http.Request) (*http.Response, error) {

	bodyData, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}

	var data struct {
		MethodProperties struct {
			CityRef string `json:"CityRef"`
		} `json:"methodProperties"`
	}

	err = json.Unmarshal(bodyData, &data)
	if err != nil {
		return &http.Response{
			StatusCode: http.StatusBadRequest,
			Body:       nil,
		}, nil
	}

	warehouses, err := warehousesCache.Get(data.MethodProperties.CityRef)
	if err != nil {
		return &http.Response{
			StatusCode: http.StatusNotFound,
			Body:       nil,
		}, nil
	}

	r := io.NopCloser(bytes.NewReader(warehouses))

	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       r,
	}, nil

}
