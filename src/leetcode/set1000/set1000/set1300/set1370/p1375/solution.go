package p1375

func numTimesAllBlue(light []int) int {
	var res int

	n := len(light)

	max := -1

	for i := 1; i <= n; i++ {
		j := light[i-1]
		if j > max {
			max = j
		}
		if max == i {
			res++
		}
	}

	return res
}
