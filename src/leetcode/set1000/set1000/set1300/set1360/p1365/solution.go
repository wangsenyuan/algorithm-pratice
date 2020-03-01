package p1365

func smallerNumbersThanCurrent(nums []int) []int {
	n := len(nums)

	res := make([]int, n)

	for i := 0; i < n; i++ {
		var cnt int

		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			if nums[j] < nums[i] {
				cnt++
			}
		}
		res[i] = cnt
	}

	return res
}
