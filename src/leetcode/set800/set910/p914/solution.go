package p914

import "sort"

func hasGroupsSizeX(deck []int) bool {
	if len(deck) == 0 {
		return false
	}

	sort.Ints(deck)

	var x int
	for x < len(deck) && deck[x] == deck[0] {
		x++
	}
	if x == 1 {
		return false
	}

	for i, j := x+1, x; i <= len(deck); i++ {
		if i < len(deck) && deck[i] == deck[i-1] {
			continue
		}
		y := i - j

		x = gcd(x, y)

		if x == 1 {
			return false
		}

		j = i
	}
	return true
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
