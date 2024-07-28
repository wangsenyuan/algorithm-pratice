package p3233

func canAliceWin(nums []int) bool {
	var sum int
	res := []int{0, 0}
	for _, num := range nums {
		if num < 10 {
			res[0] += num
		} else if num < 100 {
			res[1] += num
		}
		sum += num
	}
	return 2*res[0] > sum || 2*res[1] > sum
}
