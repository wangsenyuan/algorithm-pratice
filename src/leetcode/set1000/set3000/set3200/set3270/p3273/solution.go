package p3273

import "slices"

func minDamage(power int, damage []int, health []int) int64 {
	type pair struct {
		first  int
		second int
	}

	n := len(damage)
	monsters := make([]pair, n)
	for i := 0; i < n; i++ {
		monsters[i] = pair{damage[i], (health[i] + power - 1) / power}
	}

	slices.SortFunc(monsters, func(a, b pair) int {
		return a.second*b.first - a.first*b.second
	})

	var res int
	var cnt int

	for i := 0; i < n; i++ {
		cur := monsters[i]
		cnt += cur.second
		res += cur.first * cnt
	}

	return int64(res)
}
