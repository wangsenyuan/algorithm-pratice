package p2334

func bruteForce(arr []int, threshold int) int {
	for sz := 1; sz <= len(arr); sz++ {
		for i := 0; i+sz <= len(arr); i++ {
			ok := true
			for j := i; j < i+sz; j++ {
				ok = arr[j] > threshold/sz
				if !ok {
					break
				}
			}
			if ok {
				return sz
			}
		}
	}
	return -1
}

func validSubarraySize(nums []int, threshold int) int {
	n := len(nums)
	stack := make([]int, n)
	L := make([]int, n)
	var p int
	for i := 0; i < n; i++ {
		for p > 0 && nums[stack[p-1]] >= nums[i] {
			p--
		}
		if p > 0 {
			L[i] = stack[p-1]
		} else {
			L[i] = -1
		}
		stack[p] = i
		p++
	}

	p = 0

	for i := n - 1; i >= 0; i-- {
		for p > 0 && nums[stack[p-1]] >= nums[i] {
			p--
		}
		R := n
		if p > 0 {
			R = stack[p-1]
		}
		sz := R - L[i] - 1
		if nums[i] > threshold/sz {
			return sz
		}
		stack[p] = i
		p++
	}
	return -1
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
