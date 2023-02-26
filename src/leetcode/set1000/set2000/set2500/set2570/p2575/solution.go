package p2575

func divisibilityArray(word string, m int) []int {
	n := len(word)

	res := make([]int, n)

	var sum int

	for i := 0; i < n; i++ {
		sum = sum*10 + int(word[i]-'0')
		sum %= m
		if sum == 0 {
			res[i] = 1
		}
	}

	return res
}
