package p2093

import "sort"

func findEvenNumbers(digits []int) []int {
	sort.Ints(digits)
	var res []int
	for i := 0; i < len(digits); i++ {
		if digits[i] == 0 || i > 0 && digits[i] == digits[i-1] {
			continue
		}
		for j := 0; j < len(digits); j++ {
			if i == j || digits[j]%2 == 1 {
				continue
			}
			for k := 0; k < len(digits); k++ {
				if i == k || j == k {
					continue
				}
				res = append(res, digits[i]*100+digits[k]*10+digits[j])
			}
		}
	}
	sort.Ints(res)
	var p int
	for i := 1; i <= len(res); i++ {
		if i == len(res) || res[i] > res[i-1] {
			res[p] = res[i-1]
			p++
		}
	}
	return res[:p]
}
