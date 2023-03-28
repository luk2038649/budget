/*
Copyright © 2023 Luke Schulz

*/
package cmd

import (
	"budget/internal/config"
	"github.com/spf13/cobra"
	"log"
)

// createCmd represents the create command
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
