package budget

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"text/tabwriter"
	"time"

	"github.com/luk2038649/budget/internal/config"
	"github.com/luk2038649/budget/internal/file"
)

const tabWidth = 8
const budgetHeaders = "KIND\tAMOUNT\tNAME\tFREQUENCY\tDESCRIPTION"

type Budget struct {
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
	Amount      int
	Frequency   Frequency
	Name        string
	Description string
	Kind        ItemKind
}

func Show() error {
	b, err := getCurrent()
	if err != nil {
		return fmt.Errorf("show: %w", err)
	}
	err = b.print()
	if err != nil {
		return fmt.Errorf("show: %w", err)
	}

	return nil
}
func Add(amount int, freq Frequency, name, description string, kind ItemKind) error {
	i := newItem(amount, freq, name, description, kind)
	b, err := getCurrent()
	if err != nil {
		return fmt.Errorf("add: %w", err)
	}
	b.addItem(i)
	err = b.save()
	if err != nil {
		return fmt.Errorf("add: %w", err)
	}
	fmt.Println("Item added.")

	return nil
}
func (b *Budget) addItem(i Item) {
	b.Items = append(b.Items, i)
}
func newItem(amount int, freq Frequency, name, description string, kind ItemKind) Item {
	i := Item{
		Amount:      amount,
		Frequency:   freq,
		Name:        name,
		Description: description,
		Kind:        kind,
	}

	return i
}

func getCurrent() (Budget, error) {
	var b = New()
	path, err := config.GetCurrentDataFilePath()
	if err != nil {
		return b, fmt.Errorf("save: %w", err)
	}
	ok, err := file.Exists(path)
	if err != nil {
		return b, fmt.Errorf("get current: %w", err)
	}
	if !ok {
		return b, nil
	}
	bBytes, err := file.Load(path)
	if err != nil {
		return b, fmt.Errorf("getCurrent: %w", err)
	}
	if len(bBytes) > 0 {
		err = json.Unmarshal(bBytes, &b)
		if err != nil {
			return b, fmt.Errorf("load config: %w", err)
		}
	}

	return b, nil
}

func (b *Budget) save() error {
	path, err := config.GetCurrentDataFilePath()
	if err != nil {
		return fmt.Errorf("save: %w", err)
	}
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("save create file: %w", err)
	}

	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	bBytes, err := json.Marshal(&b)
	if err != nil {
		return fmt.Errorf("save: %w", err)
	}

	_, err = f.Write(bBytes)
	if err != nil {
		return fmt.Errorf("save write: %w", err)
	}

	return nil
}

// New creates a new named budget struct.
func New() Budget {
	b := Budget{
		CreationDate: time.Now(),
		LastEdited:   time.Now(),
		Items:        []Item{},
	}

	return b
}

func (i Item) String() string {
	return fmt.Sprintf("%s\t%d\t%s\t%s\t%s", i.Kind, i.Amount, i.Name, i.Frequency.String(), i.Description)
}

func (b *Budget) print() error {
	tw := tabwriter.NewWriter(os.Stdout, 0, tabWidth, 1, '\t', tabwriter.AlignRight)
	_, err := fmt.Fprintln(tw, budgetHeaders)
	if err != nil {
		return fmt.Errorf("budget print: %w", err)
	}
	err = b.printItemsByKind(tw, Income)
	if err != nil {
		return fmt.Errorf("print: %w ", err)
	}
	err = b.printItemsByKind(tw, Expense)
	if err != nil {
		return fmt.Errorf("print: %w ", err)
	}
	err = tw.Flush()
	if err != nil {
		return fmt.Errorf("budget print: %w", err)
	}

	return nil
}

func (b *Budget) printItemsByKind(w io.Writer, k ItemKind) error {
	for _, i := range b.Items {
		if i.Kind == k {
			_, err := fmt.Fprintln(w, i.String())
			if err != nil {
				return fmt.Errorf("print expenses: %w", err)
			}
		}
	}

	return nil
}
