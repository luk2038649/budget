/*
Copyright © 2023 Luke Schulz
*/
package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

// addCmd represents the add command.
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add budget item",
	Long:  `Add budget items of various types to your budget.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			log.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.PersistentFlags().StringP("frequency", "f", "once", "Frequency, choose yearly, monthly, weekly, daily, or once. First initial also accepted as shorthand")
	addCmd.PersistentFlags().StringP("name", "n", "", "name for this budget item")
	addCmd.PersistentFlags().StringP("description", "d", "", "Optional longer description")
	err := addCmd.MarkPersistentFlagRequired("frequency")
	if err != nil {
		log.Println(err)
	}
	err = addCmd.MarkPersistentFlagRequired("name")
	if err != nil {
		log.Println(err)
	}
}
