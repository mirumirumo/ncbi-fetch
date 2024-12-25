package search

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type SearchResult struct {
	IDs []string `xml:"IdList>Id"`
}

type Taxonid struct {
	Species string `json:"species"`
	Taxonid string `json:"taxon_id"`
}

func Org2Taxon(orgs []string) ([]byte, error) {
	baseURL := "https://eutils.ncbi.nlm.nih.gov/entrez/eutils/esearch.fcgi"
	var taxonIDs []Taxonid

	for _, org := range orgs {
		params := url.Values{}
		params.Add("db", "taxonomy")
		params.Add("term", org)
		params.Add("retmode", "xml")
		params.Add("retmax", "1")

		resp, err := http.Get(baseURL + "?" + params.Encode())
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
