package api

import (
	"fmt"
	"net/http"
	"net/url"
)

type EsearchClient struct {
	params url.Values
}

func (e *EsearchClient) SetParams(params string, value string) {
	e.params.Add(params, value)
}

func (e *EsearchClient) GetResponse() (*http.Response, error) {
	resp, err := http.Get(EsearchURL + "?" + e.params.Encode())
	if err != nil {
		return nil, fmt.Errorf("failed to get: %w", err)
	}
	return resp, nil
}
