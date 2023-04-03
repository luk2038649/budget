/*
Copyright © 2023 Luke Schulz
*/
package cmd

import (
	"log"

	"github.com/luk2038649/budget/internal/budget"
	"github.com/spf13/cobra"
)

// incomeCmd represents the income command.
var incomeCmd = &cobra.Command{
	Use:   "income",
	Short: "add income",
	Long:  `Add an income item to your budget`,
	Run: func(cmd *cobra.Command, args []string) {
		err := parseAddItem(cmd, args, budget.Income)
		if err != nil {
			log.Println(err)
		}
	},
}

func init() {
	addCmd.AddCommand(incomeCmd)
}
