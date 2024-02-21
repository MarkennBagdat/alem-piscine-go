package piscine

func JumpOver(str string) string {
	if len(str) == 0 {
		return "\n"
	}

	var answer string
	for i, char := range str {
		if (i+1)%3 == 0 {
			answer += string(char)
		}
	}
	return answer + "\n"
}
