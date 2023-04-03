/*
Copyright © 2023 Luke Schulz
*/
package cmd

import (
	"log"

	"github.com/luk2038649/budget/internal/config"
	"github.com/spf13/cobra"
)

// listCmd represents the list command.
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list available configs",
	Long:  `list available config data files which can be used for viewing and editing budgets`,
	Run: func(cmd *cobra.Command, args []string) {
		err := config.Show()
		if err != nil {
			log.Println(err)
		}
	},
}

func init() {
	configCmd.AddCommand(listCmd)
}
