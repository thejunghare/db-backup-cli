package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listDBCmd = &cobra.Command{
	Use:   "list",
	Short: "See supported databases",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Mysql, Postgress")
	},
}

func init() {
	RootCmd.AddCommand(listDBCmd)
}
