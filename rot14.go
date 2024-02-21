package piscine

func Rot14(s string) string {
	answer := ""

	for _, res := range s {
		answer += string(rot14(res))
	}
	return answer
}

func rot14(a rune) rune {
	if a >= 'A' && a <= 'L' || a >= 'a' && a <= 'l' {
		return a + 14
	}

	if a >= 'M' && a <= 'Z' || a >= 'm' && a <= 'z' {
		return a - 12
	}
	return a
}
