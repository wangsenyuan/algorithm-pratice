package p2419

const H = 22

func longestSubarray(nums []int) int {
	n := len(nums)

	// 计算以nums[i] 为最大值的长度

	pos := make([]int, H)

	for i := 0; i < H; i++ {
		pos[i] = -1
	}
	var best int
	var max_value int

	for i := 0; i < n; i++ {
		x := i + 1
		for j := 0; j < H; j++ {
			if (nums[i]>>j)&1 == 1 {
				if pos[j] < 0 {
					x = 1
					break
				}
				x = min(x, i-pos[j]+1)
			}
		}

		if nums[i] > max_value {
			best = x
			max_value = nums[i]
		} else if nums[i] == max_value {
			best = max(best, x)
		}
		for j := 0; j < H; j++ {
			if (nums[i]>>j)&1 == 1 {
				if pos[j] < 0 {
					pos[j] = i
				}
			} else {
				pos[j] = -1
			}
		}
	}

	return best
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
