package search

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type EsearchResult struct {
	EsearchResult struct {
		Count  string   `json:"count"`
		IDList []string `json:"idlist"`
	} `json:"esearchresult"`
}

func fetchCount(taxonID string, extraQuery string) (int, error) {
	params := url.Values{}
	params.Add("db", "genome")
	params.Add("term", taxonID+"[Organism:exp]"+extraQuery)
	params.Add("retmode", "json")

	resp, err := http.Get(EsearchURL + "?" + params.Encode())
	if err != nil {
		return -1, fmt.Errorf("failed to get: %w", err)
	}
	defer resp.Body.Close()
	var result EsearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return -1, fmt.Errorf("failed to decode json: %w", err)
	}
	count, err := strconv.Atoi(result.EsearchResult.Count)
	if err != nil {
		return -1, fmt.Errorf("failed to convert count: %w", err)
	}
	return count, nil
}

func fetchIDs(taxonID string, extraQuery string) ([]string, error) {
	params := url.Values{}
	params.Add("db", "genome")
	params.Add("term", taxonID+"[Organism:exp]"+extraQuery)
	params.Add("retmode", "json")
	params.Add("retmax", Retmax)

	resp, err := http.Get((EsearchURL + "?" + params.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed to get: %w", err)
	}
	defer resp.Body.Close()

	// ここからIDsの取得をする処理を書こう
	return nil, nil
}
