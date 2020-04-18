package a

func minCount(coins []int) int {

	var res int

	for i := 0; i < len(coins); i++ {
		res += (coins[i] + 1) / 2
	}

	return res
}
