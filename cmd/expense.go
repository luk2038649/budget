/*
Copyright © 2023 Luke Schulz
*/
package cmd

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/luk2038649/budget/internal/budget"
	"github.com/spf13/cobra"
)

// incomeCmd represents the income command.
var expenseCmd = &cobra.Command{
	Use:   "expense",
	Short: "add expense",
	Long:  `Add an expense item to your budget`,
	Run: func(cmd *cobra.Command, args []string) {
		err := parseAddItem(cmd, args, budget.Expense)
		if err != nil {
			log.Println(err)
		}
	},
}

func parseAddItem(cmd *cobra.Command, args []string, k budget.ItemKind) error {
	if len(args) < 1 {
		return errors.New("arg 1 must be expense amount")
	}

	amount, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("argument failed to convert to int: %w", err)
	}
	name, err := cmd.Flags().GetString("name")
	if err != nil {
		return fmt.Errorf("add income name flag: %w", err)
	}
	description, err := cmd.Flags().GetString("description")
	if err != nil {
		return fmt.Errorf("add income description flag: %w", err)
	}
	freqStr, err := cmd.Flags().GetString("frequency")
	if err != nil {
		return fmt.Errorf("add income frequency flag: %w", err)
	}
	freq, ok := budget.ParseFrequencyStr(freqStr)
	if !ok {
		// TODO print list of acceptable frequencies
		return errors.New("bad frequency value")
	}
	err = budget.Add(amount, freq, name, description, k)
	if err != nil {
		return fmt.Errorf("add income: %w", err)
	}

	return nil
}
func init() {
	addCmd.AddCommand(expenseCmd)
}
