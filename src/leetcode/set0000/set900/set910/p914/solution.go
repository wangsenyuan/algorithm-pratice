package p914

import "sort"

const MAX_NUM = 10000

func hasGroupsSizeX(deck []int) bool {
	cnt := make([]int, MAX_NUM)
	for i := 0; i < len(deck); i++ {
		cnt[deck[i]]++
	}

	var x int
	for i := 0; i < MAX_NUM; i++ {
		if cnt[i] == 0 {
			continue
		}
		if x == 0 {
			x = cnt[i]
		} else {
			x = gcd(x, cnt[i])
		}
	}
	return x > 1
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func hasGroupsSizeX1(deck []int) bool {
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
