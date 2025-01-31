package api

import "net/http"

type ApiFetcher interface {
	SetParams(params string, value string)
	GetResponse() (*http.Response, error)
}
