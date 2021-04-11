package p1822

func arraySign(nums []int) int {
	var cnt int
	for _, num := range nums {
		if num == 0 {
			return 0
		}
		if num < 0 {
			cnt++
		}
	}
	if cnt%2 == 1 {
		return -1
	}
	return 1
}
