package piscine

func Join(strs []string, sep string) string {
	answer := ""
	for _, char := range strs[0 : len(strs)-1] {
		answer += char + sep
	}
	answer += strs[len(strs)-1]
	return answer
}
