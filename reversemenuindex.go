package piscine

func ReverseMenuIndex(menu []string) []string {
	length := len(menu)
	reverseMenu := make([]string, length)
	for i, char := range menu {
		reverseMenu[length-i-1] = char
	}
	return reverseMenu
}
