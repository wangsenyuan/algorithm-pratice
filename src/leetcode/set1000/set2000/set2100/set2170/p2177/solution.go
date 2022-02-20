package p2177

func sumOfThree(num int64) []int64 {
	// num = a - 1 + a + a + 1
	// num =  3 * a
	if num%3 != 0 {
		return nil
	}
	a := num / 3
	return []int64{a - 1, a, a + 1}
}
