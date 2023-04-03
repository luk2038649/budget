package budget

import "strings"

type Frequency int

const (
	Yearly  Frequency = 1
	Monthly Frequency = 12
	Weekly  Frequency = 52
	Daily   Frequency = 365
	Once    Frequency = 0 // this might have to be one same as yearly? and rely on repeat flag for calculations.
	// TODO biweekly? more types?
)

func (f Frequency) String() string {
	switch f {
	case Yearly:
		return "yearly"
	case Monthly:
		return "monthly"
	case Weekly:
		return "weekly"
	case Daily:
		return "daily"
	case Once:
		return "once"
	default:
		return "unknown"
	}
}

func frequencyMap() map[string]Frequency {
	return map[string]Frequency{
		"yearly":  Yearly,
		"monthly": Monthly,
		"weekly":  Weekly,
		"daily":   Daily,
		"once":    Once,
		"y":       Yearly,
		"m":       Monthly,
		"w":       Weekly,
		"d":       Daily,
		"o":       Once,
	}
}

func ParseFrequencyStr(str string) (Frequency, bool) {
	f, ok := frequencyMap()[strings.ToLower(str)]

	return f, ok
}
