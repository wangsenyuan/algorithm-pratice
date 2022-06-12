package p2300

import "sort"

func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)

	n := len(spells)

	P := make([]Pair, n)

	for i := 0; i < n; i++ {
		P[i] = Pair{spells[i], i}
	}

	sort.Slice(P, func(i int, j int) bool {
		return P[i].first < P[j].first
	})

	ans := make([]int, n)
	m := len(potions)
	for i, j := 0, m; i < n; i++ {
		cur := P[i]
		for j > 0 && int64(cur.first)*int64(potions[j-1]) >= success {
			j--
		}
		ans[cur.second] = m - j
	}
	return ans
}

type Pair struct {
	first  int
	second int
}
