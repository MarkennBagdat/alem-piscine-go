package piscine

func DescendAppendRange(max, min int) []int {
	if min >= max {
		return []int{}
	}
	var answer []int
	for i := max; i > min; i-- {
		answer = append(answer, i)
	}
	return answer
}
