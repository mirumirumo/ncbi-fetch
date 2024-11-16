package cmd

import (
	"fmt"
	"log"

	"github.com/mirumirumo/ncbi-cli/connect"
	"github.com/spf13/cobra"
)

// catCmd represents the cat command
var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "connect to the ftp server",

	Run: func(cmd *cobra.Command, args []string) {
		_, cancel, err := connect.Connect()
		defer cancel()
		if err != nil {
			log.Fatalf("Failed to connect: %v", err)
		}

		if err != nil {
			log.Fatalf("Failed to list directory: %v", err)
		}
		fmt.Printf("You can successfully connect to the NCBI ftp server: %s\n", connect.SERVER)

	},
}

func init() {
	rootCmd.AddCommand(connectCmd)
}
