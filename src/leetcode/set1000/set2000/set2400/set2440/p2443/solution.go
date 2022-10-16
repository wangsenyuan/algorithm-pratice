package p2443

func countSubarrays(nums []int, minK int, maxK int) int64 {

	var res int64
	var left int
	pos1, pos2 := -1, -1

	for right, num := range nums {
		if num == minK {
			pos1 = right
		}
		if num == maxK {
			pos2 = right
		}
		if num < minK || num > maxK {
			left = right + 1
		}
		res += int64(max(0, min(pos1, pos2)-left+1))
	}
	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
