package p1774

import "sort"

func closestCost(baseCosts []int, toppingCosts []int, target int) int {
	//sort.Ints(baseCosts)
	//sort.Ints(toppingCosts)

	// cost + a * m + b * n
	var best = 1 << 30
	var bestDiff = 1 << 20

	update := func(cur int) {
		if abs(cur-target) < bestDiff {
			bestDiff = abs(cur - target)
			best = cur
		} else if abs(cur-target) == bestDiff && cur < best {
			best = cur
		}
	}

	m := len(toppingCosts)

	M := 1 << m

	sum := make([]int, M)
	for cur := 0; cur < M; cur++ {
		var tmp int
		for i := 0; i < m; i++ {
			if (cur & (1 << i)) > 0 {
				tmp += toppingCosts[i]
			}
		}
		sum[cur] = tmp
	}

	sort.Ints(sum)

	for _, cost := range baseCosts {
		update(cost)
		for _, a := range sum {
			update(cost + a)
			j := sort.Search(len(sum), func(j int) bool {
				return sum[j]+cost+a > target
			})
			if j < len(sum) {
				update(sum[j] + cost + a)
			}
			if j-1 >= 0 {
				update(sum[j-1] + cost + a)
			}
		}
	}
	return best
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
