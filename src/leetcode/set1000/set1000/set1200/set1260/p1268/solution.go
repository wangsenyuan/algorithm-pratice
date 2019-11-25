package p1268

import "sort"

func suggestedProducts(products []string, searchWord string) [][]string {
	sort.Strings(products)

	n := len(searchWord)
	res := make([][]string, n)
	start := 0
	end := len(products)

	for i := 0; i < n; i++ {
		res[i] = make([]string, 0, 3)
		if start >= end {
			continue
		}

		c := searchWord[i]

		y := search(start, end, func(j int) bool {
			if len(products[j]) <= i {
				return false
			}
			return products[j][i] > c
		})

		x := search(start, end, func(j int) bool {
			if len(products[j]) <= i {
				return false
			}
			return products[j][i] >= c
		})

		for j := x; j < y && j < x+3; j++ {
			res[i] = append(res[i], products[j])
		}

		start, end = x, y
	}

	return res
}

func search(i, n int, fn func(int) bool) int {
	left := i
	right := n

	for left < right {
		mid := (left + right) / 2
		if fn(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return right
}
