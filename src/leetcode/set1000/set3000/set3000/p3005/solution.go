package p3005

func maxFrequencyElements(nums []int) int {
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}
	var x int
	for _, v := range freq {
		x = max(x, v)
	}

	var res int
	for _, v := range freq {
		if v == x {
			res += x
		}
	}
	return res
}
