package p2342

func maximumSum(nums []int) int {
	mem := make([]int, 10*10)
	for i := 0; i < len(mem); i++ {
		mem[i] = -1
	}

	for i, num := range nums {
		dig := digitSum(num)
		if mem[dig] < 0 || nums[mem[dig]] < num {
			mem[dig] = i
		}
	}

	best := -1

	for i := 0; i < len(nums); i++ {
		dig := digitSum(nums[i])
		if mem[dig] != i {
			best = max(best, nums[mem[dig]]+nums[i])
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

func digitSum(num int) int {
	var res int
	for num > 0 {
		res += num % 10
		num /= 10
	}
	return res
}
