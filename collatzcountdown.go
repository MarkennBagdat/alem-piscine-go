package piscine

func CollatzCountdown(start int) int {
	count := 0

	if start == 0 || start < 0 {
		return -1
	} else {
		for start != 1 {
			if start%2 == 0 {
				start = start / 2
			} else if start%2 == 1 {
				start = start*3 + 1
			}
			count++
		}
	}
	return count
}
