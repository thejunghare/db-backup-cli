package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "See app version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("---v0.0.1")
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
