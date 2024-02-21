package piscine

func Unmatch(a []int) int {
	for _, i := range a {
		res := 0
		for _, j := range a {
			if i == j {
				res++
			}
		}
		if res == 1 || res%2 == 1 {
			return i
		}
	}
	return -1
}
