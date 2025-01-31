package cmd

import (
	"fmt"
	"log"

	"github.com/mirumirumo/ncbi-cli/client/connect"
	"github.com/spf13/cobra"
)

// catCmd represents the cat command
var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "connect to the ftp server",

	Run: func(cmd *cobra.Command, args []string) {
		//connect the server via ftp
		_, cancel, err := connect.ConnectGoFtp()
		defer cancel()
		if err != nil {
			log.Fatalf("Failed to connect: %v", err)
		}

		if err != nil {
			log.Fatalf("Failed to list directory: %v", err)
		}
		fmt.Printf("You can successfully connect to the NCBI ftp server: %s\n", connect.HOST)

		//----------attempt to connect the server via sftp. However, this doesn't work now-----------
		// _, sshCancel, sftpCancel, err := connect.ConnectSftp()
		// defer sshCancel()
		// defer sftpCancel()
		// if err != nil {
		// 	log.Fatalf("Failed to connect: %v", err)
		// }
		// fmt.Printf("You can successfully connect to the NCBI sftp server: %s\n", connect.HOST)

	},
}

func init() {
	rootCmd.AddCommand(connectCmd)
}
