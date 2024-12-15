package p3387

func buttonWithLongestTime(events [][]int) int {
	record := make(map[int]int)
	var prev int
	for _, cur := range events {
		i, j := cur[0], cur[1]
		record[i] = max(record[i], j-prev)
		prev = j
	}

	ans := 1
	for k, v := range record {
		if v > record[ans] || v == record[ans] && k < ans {
			ans = k
		}
	}
	return ans
}
