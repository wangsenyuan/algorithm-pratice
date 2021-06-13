package p1897

import "sort"

func mergeTriplets(triplets [][]int, target []int) bool {
	sort.Slice(triplets, func(i, j int) bool {
		for k := 0; k < 3; k++ {
			if triplets[i][k] < triplets[j][k] {
				return true
			}
			if triplets[i][k] == triplets[j][k] {
				continue
			}
			return false
		}
		return false
	})

	var second, third bool

	for i := 0; i < len(triplets); i++ {
		cur := triplets[i]
		if cur[0] > target[0] {
			break
		}

		if cur[1] > target[1] || cur[2] > target[2] {
			continue
		}

		if cur[1] == target[1] && cur[2] <= target[2] {
			second = true
		}
		if cur[1] <= target[1] && cur[2] == target[2] {
			third = true
		}

		if cur[0] == target[0] {
			if second && third {
				return true
			}
		}
	}

	return false
}
