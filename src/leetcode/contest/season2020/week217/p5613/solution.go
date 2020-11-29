package p5613

func maximumWealth(accounts [][]int) int {
	var best int

	for i := 0; i < len(accounts); i++ {
		var sum int
		for j := 0; j < len(accounts[i]); j++ {
			sum += accounts[i][j]
		}
		if sum > best {
			best = sum
		}
	}
	return best
}
