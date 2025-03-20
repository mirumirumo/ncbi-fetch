package domain

type GetInput struct {
	Entry TaxonID
	Args  map[string]interface{}
}
type Getter interface {
	FetchFromAPI(input GetInput) error
}

func NewGetInput(entry TaxonID, args map[string]interface{}) GetInput {
	return GetInput{Entry: entry, Args: args}
}
