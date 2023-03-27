package budget

import (
	"fmt"
	"io"
	"os"
	"text/tabwriter"
)

type budget struct {
	items []budgetItem
}

type budgetItemKind string

const (
	Expense budgetItemKind = "expense"
	Income  budgetItemKind = "income"
)

type budgetItem struct {
	amount      int
	frequency   frequency
	name        string
	description string
	key         int
	kind        budgetItemKind
}

func (b budgetItem) String() string {
	return fmt.Sprintf("%s\t%d:\t$d\t%s\n", b.name, b.amount, b.frequency.String())
}

func (b budget) print() error {
	tw := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)
	err := b.printItemsByKind(tw, Income)
	if err != nil {
		return fmt.Errorf("print: %w ", err)
	}
	err = b.printItemsByKind(tw, Expense)
	if err != nil {
		return fmt.Errorf("print: %w ", err)
	}
	return nil
}

func (b budget) printItemsByKind(w io.Writer, k budgetItemKind) error {
	for _, i := range b.items {
		if i.kind == k {
			_, err := fmt.Fprintln(w, i.String())
			if err != nil {
				return fmt.Errorf("print expenses: %w", err)
			}
		}
	}
	return nil
}
