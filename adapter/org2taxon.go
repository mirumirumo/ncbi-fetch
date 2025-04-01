package adapter

import (
	"encoding/xml"
	"fmt"
	"io"

	"github.com/mirumirumo/ncbi-fetch/domain"
)

// type TaxonInfoGetter interface {
// 	Get(client domain.ApiFetcher, species string) (domain.TaxonID, error)
// }
// type FetchTaxonInfo struct {
// 	Client domain.ApiFetcher
// 	Getter TaxonInfoGetter
// }

// type FetchTaxonOutput struct {
// 	Species []string         `json:"species"`
// 	Taxonid []domain.TaxonID `json:"taxon_id"`
// }

// func (f *FetchTaxonInfo) Get(species []string) (FetchTaxonOutput, error) {
// 	output := FetchTaxonOutput{}
// 	for _, s := range species {
// 		taxonId, err := f.Getter.Get(f.Client, s)
// 		if err != nil {
// 			return output, fmt.Errorf("failed to get taxon info: %w", err)
// 		}
// 		output.Species = append(output.Species, s)
// 		output.Taxonid = append(output.Taxonid, taxonId)

// 	}
// 	return output, nil
// }

type SearchResult struct {
	IDs []string `xml:"IdList>Id"`
}

type TaxonGet struct {
	Client domain.ApiFetcher
}

func (t *TaxonGet) Get(species string) (domain.TaxonID, error) {
	var taxonId domain.TaxonID
	t.Client.SetParams("db", "taxonomy")
	t.Client.SetParams("term", species)
	t.Client.SetParams("retmode", "xml")
	t.Client.SetParams("retmax", "1")
	resp, err := t.Client.GetResponse()
	if err != nil {
		return taxonId, fmt.Errorf("failed to get: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return taxonId, fmt.Errorf("failed to read body: %w", err)
	}
	var result SearchResult
	if err := xml.Unmarshal(body, &result); err != nil {
		return taxonId, fmt.Errorf("failed to unmarshal xml: %w", err)
	}
	taxonId = domain.TaxonID(result.IDs[0])

	return taxonId, nil
}

// func Org2Taxon(orgs []string) ([]byte, error) {
// 	var taxonIDs []Taxonid

// 	for _, org := range orgs {
// 		esClient := api.EsearchClient{}
// 		esClient.SetParams("db", "taxonomy")
// 		esClient.SetParams("term", org)
// 		esClient.SetParams("retmode", "xml")
// 		esClient.SetParams("retmax", "1")
// 		resp, err := esClient.GetResponse()
// 		if err != nil {
// 			return nil, fmt.Errorf("failed to get: %w", err)
// 		}
// 		defer resp.Body.Close()

// 		body, err := io.ReadAll(resp.Body)
// 		if err != nil {
// 			return nil, fmt.Errorf("failed to read body: %w", err)
// 		}
// 		var result SearchResult
// 		if err := xml.Unmarshal(body, &result); err != nil {
// 			return nil, fmt.Errorf("failed to unmarshal xml: %w", err)
// 		}

// 		if len(result.IDs) > 0 {
// 			taxonIDs = append(taxonIDs, Taxonid{Species: org, Taxonid: result.IDs[0]})
// 		} else {
// 			return nil, nil
// 		}
// 	}
// 	out, err := json.Marshal(taxonIDs)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to marshal json: %w", err)
// 	}
// 	return out, nil
// }
