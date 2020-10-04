package p1578

const INF = 1 << 29

func minCost(s string, cost []int) int {
	n := len(s)
	var res int

	for i := 0; i < n; i++ {
		j := i
		var sum int
		var max int
		for j < n && s[j] == s[i] {
			sum += cost[j]
			if cost[j] > max {
				max = cost[j]
			}
			j++
		}
		res += sum - max
		i = j - 1
	}
	return res
}
