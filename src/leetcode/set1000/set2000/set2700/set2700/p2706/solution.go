package p2706

func buyChoco(prices []int, money int) int {
	first := -1
	second := -1

	for _, p := range prices {
		if first < 0 || p <= first {
			second = first
			first = p
		} else if second < 0 || p <= second {
			second = p
		}
	}

	if first+second <= money {
		return money - first - second
	}
	return money
}
