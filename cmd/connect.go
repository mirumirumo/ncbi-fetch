package cmd

import (
	"fmt"
	"log"

	"github.com/mirumirumo/ncbi-cli/connect"
	"github.com/spf13/cobra"
)

const SERVER string = "ftp.ncbi.nlm.nih.gov:21"

// catCmd represents the cat command
var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "connect to the ftp server",

	Run: func(cmd *cobra.Command, args []string) {
		c, cancel, err := connect.Connect()
		defer cancel()
		if err != nil {
			log.Fatalf("Failed to connect: %v", err)
		}

		entries, err := c.List("")
		if err != nil {
			log.Fatalf("Failed to list directory: %v", err)
		}

		for _, entry := range entries {
			fmt.Println(entry.Name)
		}

	},
}

func init() {
	rootCmd.AddCommand(connectCmd)
}
