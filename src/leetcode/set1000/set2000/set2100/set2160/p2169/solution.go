package p2169

func countOperations(num1 int, num2 int) int {
	// a - b
	var cnt int

	for num1 > 0 && num2 > 0 {
		cnt++
		if num1 >= num2 {
			num1 -= num2
		} else {
			num2 -= num1
		}
	}
	return cnt
}
