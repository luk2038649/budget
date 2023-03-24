package budget

import "fmt"

// expenseReport basic function for setting up go test ci.
func expenseReport(n int) string {
	return fmt.Sprintf("- %d money.", n)
}
