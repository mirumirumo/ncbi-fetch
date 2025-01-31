package search

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"

	"github.com/mirumirumo/ncbi-fetch/client/api"
)

type SearchResult struct {
	IDs []string `xml:"IdList>Id"`
}

type Taxonid struct {
	Species string `json:"species"`
	Taxonid string `json:"taxon_id"`
}

func Org2Taxon(orgs []string) ([]byte, error) {
	var taxonIDs []Taxonid

	for _, org := range orgs {
		esClient := api.EsearchClient{}
		esClient.SetParams("db", "taxonomy")
		esClient.SetParams("term", org)
		esClient.SetParams("retmode", "xml")
		esClient.SetParams("retmax", "1")
		resp, err := esClient.GetResponse()
		if err != nil {
			return nil, fmt.Errorf("failed to get: %w", err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read body: %w", err)
		}
		var result SearchResult
		if err := xml.Unmarshal(body, &result); err != nil {
			return nil, fmt.Errorf("failed to unmarshal xml: %w", err)
		}

		if len(result.IDs) > 0 {
			taxonIDs = append(taxonIDs, Taxonid{Species: org, Taxonid: result.IDs[0]})
		} else {
			return nil, nil
		}
	}
	out, err := json.Marshal(taxonIDs)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal json: %w", err)
	}
	return out, nil
}
