package p1342

func numberOfSteps(num int) int {
	// if num == 0 {
	// 	return 0
	// }
	// if num&1 == 1 {
	// 	return 1 + numberOfSteps(num-1)
	// }
	// return 1 + numberOfSteps(num>>1)

	var res int
	for num > 0 {
		res++
		if num&1 == 1 {
			num--
		} else {
			num >>= 1
		}
	}
	return res
}
