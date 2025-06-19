package cmd

import (
	"BroadCastServer/server"
	"fmt"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the broadcast server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting broadcast server on :8080")
		server.StartServer()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
