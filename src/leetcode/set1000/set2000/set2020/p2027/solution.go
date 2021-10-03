package p2027

func missingRolls(rolls []int, mean int, n int) []int {
	// sum + sum2 = (n + m) * mean
	// sum2 = (n + m) * mean - sum
	// sum2 >= n && sum2 <= 6 * n
	var sum int
	for _, roll := range rolls {
		sum += roll
	}
	sum2 := (n+len(rolls))*mean - sum
	if sum2 < n || sum2 > 6*n {
		return nil
	}
	arr := make([]int, n)
	for sum2 > 0 {
		for i := 0; i < n && sum2 > 0; i++ {
			arr[i]++
			sum2--
		}
	}
	return arr
}
