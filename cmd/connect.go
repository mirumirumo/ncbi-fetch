package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/jlaffaye/ftp"
	"github.com/spf13/cobra"
)

const SERVER string = "ftp.ncbi.nlm.nih.gov:21"

// catCmd represents the cat command
var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "connect to the ftp server",

	Run: func(cmd *cobra.Command, args []string) {
		c, err := ftp.Dial(SERVER, ftp.DialWithTimeout(5*time.Second))
		if err != nil {
			log.Fatal(err)
		}
		defer func() {
			if err := c.Quit(); err != nil {
				log.Fatalf("Failed to close connection: %v", err)
			}
		}()

		err = c.Login("anonymous", "anonymous")
		if err != nil {
			log.Fatal(err)
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
