package p5611

import "sort"

func stoneGameVI(aliceValues []int, bobValues []int) int {
	n := len(aliceValues)
	sum := make([]Pair, n)
	for i := 0; i < n; i++ {
		sum[i] = Pair{aliceValues[i] + bobValues[i], i}
	}

	sort.Slice(sum, func(i, j int) bool {
		return sum[i].first > sum[j].first
	})

	var alice int
	var bob int

	for i := 0; i < n; i += 2 {
		alice += aliceValues[sum[i].second]
		if i+1 < n {
			bob += bobValues[sum[i+1].second]
		}
	}

	if alice > bob {
		return 1
	}

	if alice < bob {
		return -1
	}
	return 0
}

type Pair struct {
	first, second int
}
