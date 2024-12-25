package get

import (
	"fmt"

	"github.com/mirumirumo/ncbi-cli/search"
	"github.com/spf13/cobra"
)

var organisms []string

var taxonidCmd = &cobra.Command{
	Use:   "taxonid",
	Short: "get the taxon id from the NCBI database",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if len(organisms) == 0 {
			fmt.Println("Please specify the flag \"organism\" to get the taxon id")
		}
		taxonids, err := search.Org2Taxon(organisms)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}

		fmt.Printf("%s", string(taxonids))
	}}

func init() {
	GetCmd.AddCommand(taxonidCmd)
	taxonidCmd.Flags().StringSliceVarP(&organisms, "species", "s", nil, "an organism name to get the taxon id")
}
