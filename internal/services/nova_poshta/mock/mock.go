package mock

import (
	"bytes"
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
}

func (m *Client) Do(req *http.Request) (*http.Response, error) {

	fileData, err := os.ReadFile("test-data/kharkiv_warehouses.json")
	if err != nil {
		return nil, err
	}

	r := io.NopCloser(bytes.NewReader(fileData))

	// time.Sleep(1 * time.Second)

	return &http.Response{
		StatusCode: 200,
		Body:       r,
	}, nil

}
