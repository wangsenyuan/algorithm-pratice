package p2303

func calculateTax(brackets [][]int, income int) float64 {
	var res float64
	var prev int
	var i int
	for income > brackets[i][0] {
		res += float64(brackets[i][0]-prev) * float64(brackets[i][1]) / 100
		prev = brackets[i][0]
		i++
	}

	res += float64(income-prev) * float64(brackets[i][1]) / 100

	return res
}
