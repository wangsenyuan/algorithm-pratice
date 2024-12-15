package p3387

func buttonWithLongestTime(events [][]int) int {
	ans := 1
	var prev int
	var record int
	for _, cur := range events {
		i, j := cur[0], cur[1]

		if j-prev > record || j-prev == record && i < ans {
			ans = i
			record = j - prev
		}

		prev = j
	}

	return ans
}
