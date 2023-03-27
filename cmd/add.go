/*
Copyright © 2023 Luke Schulz
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// addCmd represents the add command.
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add budget item",
	Long:  `Add budget items of various types to your budget.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.PersistentFlags().StringP("frequency", "f", "once", "Frequency, choose yearly, monthly, weekly, daily, or once. First initial also accepted as shorthand")
	addCmd.PersistentFlags().StringP("name", "n", "", "name for this budget item")
	addCmd.PersistentFlags().StringP("description", "d", "", "Optional longer description")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
