package p3432

func countPartitions(nums []int) int {
	var sum int
	for _, x := range nums {
		sum += x
	}
	var pref int
	var res int
	for _, x := range nums {
		sum -= x
		pref += x
		if abs(sum-pref)%2 == 0 {
			res++
		}
	}
	if pref%2 == 0 {
		res--
	}
	return res
}

func abs(num int) int {
	return max(num, -num)
}
