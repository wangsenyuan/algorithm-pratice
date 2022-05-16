package p2275

func largestCombination(candidates []int) int {

	var best int

	for i := 0; i < 24; i++ {
		var cnt int
		for _, num := range candidates {
			if num&(1<<i) > 0 {
				cnt++
			}
		}
		if cnt > best {
			best = cnt
		}
	}
	return best
}
