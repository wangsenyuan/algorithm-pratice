package p1848

func getMinDistance(nums []int, target int, start int) int {
	res := len(nums)
	for i, num := range nums {
		if num == target {
			if abs(i-start) < res {
				res = abs(i - start)
			}
		}
	}
	return res
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
