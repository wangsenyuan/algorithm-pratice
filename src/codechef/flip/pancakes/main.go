package main

import "fmt"

func flipPancake(sides []rune) int {
	var doFlip func(int, int, rune) int

	changeExpect := func(current rune) rune {
		if current == '+' {
			return '-'
		}
		return '+'
	}

	doFlip = func(i int, flips int, expect rune) int {
		if i < 0 {
			return flips
		}

		if sides[i] == expect {
			return doFlip(i-1, flips, expect)
		}
		return doFlip(i-1, flips+1, changeExpect(expect))
	}

	return doFlip(len(sides)-1, 0, '+')
}

func main() {
	var t int
	fmt.Scanf("%d", &t)

	for i := 1; i <= t; i++ {
		var line string
		fmt.Scanf("%s", &line)
		y := flipPancake([]rune(line))
		fmt.Printf("Case #%d: %d\n", i, y)
	}
}
