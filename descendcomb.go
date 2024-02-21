package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	for i := 99; i >= 1; i-- {
		first := i / 10
		second := i % 10
		for j := i - 1; j >= 0; j-- {
			third := j / 10
			fourth := j % 10
			z01.PrintRune(rune('0' + first))
			z01.PrintRune(rune('0' + second))
			z01.PrintRune(' ')
			z01.PrintRune(rune('0' + third))
			z01.PrintRune(rune('0' + fourth))
			if i-1 > 0 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
}
