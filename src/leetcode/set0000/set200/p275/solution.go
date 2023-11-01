package p275

func hIndex(citations []int) int {
	n := len(citations)
	ans := 0
	for i := n - 1; i >= 0; i-- {
		j := n - i
		if citations[i] >= j {
			ans = j
		}
	}
	return ans
}
