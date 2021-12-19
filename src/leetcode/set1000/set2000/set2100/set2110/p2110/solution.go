package p2110

func getDescentPeriods(prices []int) int64 {

	var res int64

	n := len(prices)

	for i := 0; i < n; {
		j := i
		i++
		for i < n && prices[i] == prices[i-1]-1 {
			i++
		}
		l := int64(i - j)
		res += l * (l + 1) / 2
	}

	return res
}
