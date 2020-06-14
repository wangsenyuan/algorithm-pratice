package p1475

func finalPrices(prices []int) []int {
	n := len(prices)
	st := make([]int, n)
	var p int
	res := make([]int, n)

	for i := 0; i < n; i++ {
		res[i] = prices[i]
		for p > 0 && prices[st[p-1]] >= prices[i] {
			res[st[p-1]] -= prices[i]
			p--
		}
		st[p] = i
		p++
	}

	return res
}
