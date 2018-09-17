package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of cobra-demo",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cobra-demo version 0.1.0")
	},
}
