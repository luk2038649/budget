/*
Copyright © 2023 Luke Schulz
*/
package cmd

import (
	"log"

	"budget/internal/config"
	"github.com/spf13/cobra"
)

// createCmd represents the create command.
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a new named budget config",
	Long: `Create a budget configuration which will be used to store budget
details to your file system. Name is first arg`,
	Run: func(cmd *cobra.Command, args []string) {
		err := config.Create(args[0])
		if err != nil {
			log.Println(err)
		}
	},
}

func init() {
	configCmd.AddCommand(createCmd)
}
