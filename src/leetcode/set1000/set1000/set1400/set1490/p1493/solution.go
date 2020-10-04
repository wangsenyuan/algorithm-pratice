package p1493

func longestSubarray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var prev int
	var best int
	var flag bool
	for i := -1; i < len(nums); {
		if i < 0 || nums[i] == 0 {
			if i >= 0 && nums[i] == 0 {
				flag = true
			}
			// try this one
			if i+1 == len(nums) || nums[i+1] == 0 {
				i++
				prev = 0
				continue
			}
			// nums[i+1] = 1
			i++
			j := i

			for i < len(nums) && nums[i] == 1 {
				i++
			}

			if prev+i-j > best {
				best = prev + i - j
			}
			prev = i - j
			continue
		}
		i++
	}

	if !flag {
		return len(nums) - 1
	}

	return best
}
