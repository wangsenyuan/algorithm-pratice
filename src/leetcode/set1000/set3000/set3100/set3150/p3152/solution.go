package p3152

func sumDigitDifferences(nums []int) int64 {
	freq := make([][]int, 11)
	for i := 0; i < 11; i++ {
		freq[i] = make([]int, 10)
	}
	var res int
	for i, num := range nums {
		for j := 0; j < 11; j++ {
			x := num % 10
			res += i - freq[j][x]
			freq[j][x]++
			num /= 10
		}
	}

	return int64(res)
}
