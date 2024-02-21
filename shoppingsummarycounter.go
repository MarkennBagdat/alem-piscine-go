package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	summary := make(map[string]int)

	for _, i := range SplitWhiteSpaces(str) {
		summary[i]++
	}
	return summary
}

func SplitWhiteSpaces(s string) []string {
	wout := make([]string, 0, len(s))
	word := ""

	for _, i := range s {
		if i == ' ' {
			wout = append(wout, word)
			word = ""
		} else {
			word += string(i)
		}
	}
	wout = append(wout, word)
	return wout
}
