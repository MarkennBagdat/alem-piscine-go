package piscine

func Compact(ptr *[]string) int {
	count := 0
	var arr []string

	for _, char := range *ptr {
		if char != "" {
			arr = append(arr, char)
			count++
		}
	}
	*ptr = arr
	return count
}
