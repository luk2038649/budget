package budget

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestExpenseReport calls expenseReport function to validate our testing CI works.
func TestExpenseReport(t *testing.T) {
	t.Parallel()
	expense := 30
	result := expenseReport(expense)
	assert.Contains(t, result, strconv.Itoa(expense))
}
