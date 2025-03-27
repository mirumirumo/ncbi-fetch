package search

import "github.com/mirumirumo/ncbi-fetch/domain"

type Getter interface {
	Get(input domain.GetInput) error
}
