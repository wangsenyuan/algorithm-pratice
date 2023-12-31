package p2980

func hasTrailingZeros(nums []int) bool {
	var cnt int
	for _, num := range nums {
		if num%2 == 0 {
			cnt++
		}
		if cnt == 2 {
			return true
		}
	}
	return false
}
