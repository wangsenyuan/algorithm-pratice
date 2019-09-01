package p1176

func dietPlanPerformance(calories []int, k int, lower int, upper int) int {
	var sum int

	var i int

	for i < k {
		sum += calories[i]
		i++
	}
	var res int
	if sum < lower {
		res--
	} else if sum > upper {
		res++
	}

	for i < len(calories) {
		sum += calories[i]
		sum -= calories[i-k]
		if sum < lower {
			res--
		} else if sum > upper {
			res++
		}
		i++
	}

	return res
}
