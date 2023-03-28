package budget

type frequency int

const (
	Yearly  frequency = 1
	Monthly frequency = 12
	Weekly  frequency = 52
	Daily   frequency = 365
	Once    frequency = 0 // this might have to be one same as yearly? and rely on repeat flag for calculations.
	// TODO biweekly? more types?
)

func (f frequency) String() string {
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

// var (
//	frequencyMap = map[string]frequency{
//		"yearly":  Yearly,
//		"monthly": Monthly,
//		"weekly":  Weekly,
//		"daily":   Daily,
//		"once":    Once,
//		"y":       Yearly,
//		"m":       Monthly,
//		"w":       Weekly,
//		"d":       Daily,
//		"o":       Once,
//	}
//)

// func parseFrequencyStr(str string) (frequency, bool) {
//	f, ok := frequencyMap[strings.ToLower(str)]
//
//	return f, ok
//}
