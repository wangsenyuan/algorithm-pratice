package p2717

func semiOrderedPermutation(nums []int) int {
	first := 0
	n := len(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			first = i
			break
		}
	}

	var res int
	for first != 0 {
		nums[first], nums[first-1] = nums[first-1], nums[first]
		first--
		res++
	}
	last := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == n {
			last = i
			break
		}
	}

	res += n - 1 - last
	return res
}
