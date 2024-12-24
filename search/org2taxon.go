package search

import (
	"encoding/xml"
	"io"
	"net/http"
	"net/url"
)

type SearchResult struct {
	IDs []string `xml:"IdList>Id"`
}

func Org2Taxon(org string) (string, error) {
	baseURL := "https://eutils.ncbi.nlm.nih.gov/entrez/eutils/esearch.fcgi"
	params := url.Values{}
	params.Add("db", "taxonomy")
	params.Add("term", org)
	params.Add("retmode", "xml")
	params.Add("retmax", "1")

	resp, err := http.Get(baseURL + "?" + params.Encode())
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var result SearchResult
	if err := xml.Unmarshal(body, &result); err != nil {
		return "", err
	}

	if len(result.IDs) > 0. {
		return result.IDs[0], nil
	} else {
		return "", nil
	}

}
