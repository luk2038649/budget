/*
Copyright © 2023 Luke Schulz
*/
package cmd

import (
	"fmt"
	"log"

	"budget/internal/config"
	"github.com/spf13/cobra"
)

// useCmd represents the use command.
var useCmd = &cobra.Command{
	Use:   "use",
	Short: "use config",
	Long:  `Choose which budget config file to use. see config list for options.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			err := config.Use(args[0])
			if err != nil {
				log.Println(err)
			}
		} else {
			fmt.Println("Must pass in name argument")
		}

	},
}

func init() {
	configCmd.AddCommand(useCmd)
}
