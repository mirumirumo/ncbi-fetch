package search

import (
	"encoding/xml"
	"fmt"
	"io"

	"github.com/mirumirumo/ncbi-fetch/domain"
)

type searchResult struct {
	IDs []string `xml:"IdList>Id"`
}

type Taxonid struct {
	Species string `json:"species"`
	Taxonid string `json:"taxon_id"`
}

func (t *Taxonid) Get(client domain.ApiFetcher, species string) error {
	// esClient := api.EsearchClient{}

	client.SetParams("db", "taxonomy")
	client.SetParams("term", species)
	client.SetParams("retmode", "xml")
	client.SetParams("retmax", "1")
	resp, err := client.GetResponse()
	if err != nil {
		return fmt.Errorf("failed to get: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read body: %w", err)
	}
	var result searchResult
	if err := xml.Unmarshal(body, &result); err != nil {
		return fmt.Errorf("failed to unmarshal xml: %w", err)
	}
	t.Taxonid = result.IDs[0]
	t.Species = species

	return nil
}
