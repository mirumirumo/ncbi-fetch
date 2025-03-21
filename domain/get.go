package domain

type GetInput struct {
	Entry TaxonID
	Args  map[string]string
}
type Getter interface {
	Get(input GetInput) error
}

func NewGetInput(entry TaxonID, args map[string]string) GetInput {
	return GetInput{Entry: entry, Args: args}
}
