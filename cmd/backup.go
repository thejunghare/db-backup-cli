package cmd

import (
	"fmt"
	"log"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/thejunghare/db-backup-cli/pkg/db"
)

var backCommand = &cobra.Command{
	Use:   "backup",
	Short: "Start backup",
	Run: func(cmd *cobra.Command, args []string) {
		// get username
		prompt_username := promptui.Prompt{
			Label: "Database Username",
		}
		username, err := prompt_username.Run()
		if err != nil {
			log.Fatalf("Failed to get username: %v", err)
		}

		// get password
		prompt_password := promptui.Prompt{
			Label: "Database password",
			Mask:  '*',
		}
		password, err := prompt_password.Run()
		if err != nil {
			log.Fatalf("Failed to get password: %v", err)
		}

		// get dbname
		prompt_dbname := promptui.Prompt{
			Label: "Database dbname",
		}
		dbname, err := prompt_dbname.Run()
		if err != nil {
			log.Fatalf("Failed to get password: %v", err)
		}

		// try to connect to database
		backup, err := db.Backup(dbname, username, password)
		if err != nil {
			log.Fatalf("Failed to perform backup: %v", err)
		}
		fmt.Print("Done!", backup)
	},
}

func init() {
	RootCmd.AddCommand(backCommand)
}

//  'Droid Sans Mono', 'monospace', monospace
