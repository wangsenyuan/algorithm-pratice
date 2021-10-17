package lcp40

import "sort"

func maxmiumScore(cards []int, cnt int) int {
	sort.Ints(cards)
	reverse(cards)
	var sum int

	var i int
	j, k := -1, -1
	for i < cnt {
		sum += cards[i]
		if cards[i]%2 == 0 {
			j = i
		} else {
			k = i
		}
		i++
	}
	for i < len(cards) && sum%2 == 1 {
		num := cards[i]
		if num%2 == 1 {
			// a odd num, we need to remove the last min even number
			if j >= 0 {
				sum = sum - cards[j] + num
				k = i
			}
		} else {
			if k >= 0 {
				sum = sum - cards[k] + num
				j = i
			}
		}
		i++
	}
	if sum%2 == 1 {
		return 0
	}

	return sum
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
