package p2862

func maxNumberOfAlloys(n int, k int, budget int, composition [][]int, stock []int, cost []int) int {
	var best int
	for _, comp := range composition {
		tmp := produce(n, k, budget, comp, stock, cost)

		if tmp > best {
			best = tmp
		}
	}
	return best
}

const inf = 1 << 30

func produce(n int, k int, budget int, comp []int, stock []int, cost []int) int {

	check := func(x int) bool {
		var tot int64
		for i := 0; i < n; i++ {
			need := int64(comp[i])*int64(x) - int64(stock[i])
			if need < 0 {
				need = 0
			}
			tot += need * int64(cost[i])
		}
		return tot <= int64(budget)
	}

	l, r := 1, inf

	for l < r {
		mid := (l + r) / 2
		if !check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r - 1
}
