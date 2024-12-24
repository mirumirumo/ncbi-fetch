package get

import (
	"fmt"

	"github.com/spf13/cobra"
)

var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "get the data from the NCBI database",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get the data from the NCBI database")
	},
}
