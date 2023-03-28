package budget

import (
	"fmt"
	"io"
	"os"
	"text/tabwriter"
	"time"
)

type Budget struct {
	Name         string
	Items        []Item
	CreationDate time.Time
	LastEdited   time.Time
}

type ItemKind string

const (
	Expense ItemKind = "expense"
	Income  ItemKind = "income"
)

type Item struct {
	amount      int
	frequency   frequency
	name        string
	description string
	key         int
	kind        ItemKind
}

// New creates a new named budget struct
func New(name string) Budget {
	b := Budget{
		Name:         name,
		CreationDate: time.Now(),
		LastEdited:   time.Now(),
		Items:        []Item{},
	}
	return b
}

func (b Item) String() string {
	return fmt.Sprintf("%s\t%d:\t$d\t%s\n", b.name, b.amount, b.frequency.String())
}

func (b Budget) print() error {
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

func (b Budget) printItemsByKind(w io.Writer, k ItemKind) error {
	for _, i := range b.Items {
		if i.kind == k {
			_, err := fmt.Fprintln(w, i.String())
			if err != nil {
				return fmt.Errorf("print expenses: %w", err)
			}
		}
	}
	return nil
}
