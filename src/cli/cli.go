package cli

import (
	"github.com/spf13/cobra"
)

func Execute () {
	var rootCmd = &cobra.Command{Use: "akinet"}

	rootCmd.AddCommand(CMDInbound)
	
	rootCmd.Execute()
}