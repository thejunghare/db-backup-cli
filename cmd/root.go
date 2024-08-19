package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "Command line database backup",
	Short: "A CLI tool for database backup and restoration",
}
