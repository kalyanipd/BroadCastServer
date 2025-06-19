package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "broadcast-server",
	Short: "Broadcast server using WebSockets",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
