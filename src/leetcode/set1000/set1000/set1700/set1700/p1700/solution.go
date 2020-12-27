package p1700

func averageWaitingTime(customers [][]int) float64 {
	var cur int
	var waiting int
	for _, cus := range customers {
		a, b := cus[0], cus[1]
		if cur < a {
			cur = a
		}
		cur += b
		waiting += cur - a
	}
	res := float64(waiting) / float64(len(customers))
	return res
}
