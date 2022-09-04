package p2401

const H = 31

func longestNiceSubarray(nums []int) int {
	//在sub array中，从0到31位，每一位最多有一个数字
	pos := make([]int, H)
	for i := 0; i < H; i++ {
		pos[i] = -1
	}

	n := len(nums)

	var best int
	prev := -1
	for i := 0; i < n; i++ {
		cur := nums[i]
		for j := 0; j < H; j++ {
			if (cur>>j)&1 == 1 {
				prev = max(prev, pos[j])
			}
		}
		best = max(best, i-prev)

		for j := 0; j < H; j++ {
			if (cur>>j)&1 == 1 {
				pos[j] = i
			}
		}
	}

	return best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
