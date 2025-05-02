package adapter

import (
	"github.com/mirumirumo/ncbi-fetch/domain"
	"github.com/mirumirumo/ncbi-fetch/service/search"
)

type GetTaxons struct {
	client domain.ApiFetcher
}

func NewGetTaxons(client domain.ApiFetcher) *GetTaxons {
	return &GetTaxons{
		client: client,
	}
}

func (g *GetTaxons) Get(species []string) ([]domain.TaxonID, error) {
	var taxonIds []domain.TaxonID
	for _, s := range species {
		taxonIdClient := &search.Taxonid{Species: s}
		err := taxonIdClient.Get(g.client, s)
		if err != nil {
			return nil, err
		}
		taxonIds = append(taxonIds, domain.TaxonID(taxonIdClient.Taxonid))
		g.client.RefreshParams()
	}
	return taxonIds, nil
}
