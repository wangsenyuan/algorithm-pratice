package p3191

func minOperations(nums []int) int {

	var cnt int
	for _, num := range nums {
		if cnt%2 == num {
			cnt++
		}
	}

	return cnt
}
