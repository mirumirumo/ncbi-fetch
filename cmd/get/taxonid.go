package get

import (
	"fmt"

	"github.com/mirumirumo/ncbi-cli/search"
	"github.com/spf13/cobra"
)

var organism string

var taxonidCmd = &cobra.Command{
	Use:   "taxonid",
	Short: "get the taxon id from the NCBI database",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if organism == "" {
			fmt.Println("Please specify the flag \"organism\" to get the taxon id")
		}
		taxonid, err := search.Org2Taxon(organism)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		fmt.Printf("Taxon ID for %s: %s\n", organism, taxonid)
	}}

func init() {
	GetCmd.AddCommand(taxonidCmd)
	taxonidCmd.Flags().StringVarP(&organism, "species", "s", "", "an organism name to get the taxon id")
}
