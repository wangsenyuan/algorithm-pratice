package p1542

func maxNonOverlapping(nums []int, target int) int {
	pos := make(map[int]int)
	pos[0] = 0
	var sum int

	var cnt int
	var prev int = -1

	for i := 1; i <= len(nums); i++ {
		sum += nums[i-1]
		if j, found := pos[sum-target]; found {
			if j >= prev {
				cnt++
				prev = i
			}
		}
		pos[sum] = i
	}

	return cnt
}
