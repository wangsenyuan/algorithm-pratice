package jump2

func jump(nums []int) int {

	n := len(nums)
	var far, end, cnt int

	for i := 0; i < n-1; i++ {
		far = max(far, i+nums[i])
		if i == end {
			end = far
			cnt++
		}
	}

	return cnt
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func jump1(nums []int) int {
	n := len(nums)

	arr := make([]int, 2*n)

	for i := 0; i < len(arr); i++ {
		arr[i] = n
	}

	update := func(pos int, v int) {
		pos += n
		arr[pos] = v

		for pos > 1 {
			arr[pos>>1] = min(arr[pos], arr[pos^1])
			pos >>= 1
		}
	}

	get := func(left, right int) int {
		left += n
		right += n

		var res = n

		for left < right {
			if left&1 == 1 {
				res = min(res, arr[left])
				left++
			}
			if right&1 == 1 {
				right--
				res = min(res, arr[right])
			}
			left >>= 1
			right >>= 1
		}
		return res
	}
	update(n-1, 0)
	for i := n - 2; i >= 0; i-- {
		x := get(i+1, min(n, i+nums[i]+1)) + 1
		update(i, x)
	}

	return get(0, 1)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
