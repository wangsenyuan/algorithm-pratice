package p2455

func averageValue(nums []int) int {
	var sum int
	var cnt int

	for _, num := range nums {
		if num%6 == 0 {
			sum += num
			cnt++
		}
	}

	if cnt == 0 {
		return 0
	}

	return sum / cnt
}
