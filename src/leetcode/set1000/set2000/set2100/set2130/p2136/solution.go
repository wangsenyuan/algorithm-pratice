package p2136

import "sort"

func earliestFullBloom(plantTime []int, growTime []int) int {
	n := len(plantTime)
	seeds := make([]Seed, n)
	for i := 0; i < n; i++ {
		seeds[i] = Seed{plantTime[i], growTime[i]}
	}
	sort.Slice(seeds, func(i, j int) bool {
		return seeds[i].Less(seeds[j])
	})
	var ans int
	var sum int
	for i := 0; i < n; i++ {
		sum += seeds[i].plant
		ans = max(ans, sum+seeds[i].grow)
	}

	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Seed struct {
	plant int
	grow  int
}

func (this Seed) Less(that Seed) bool {
	return this.grow > that.grow
}
