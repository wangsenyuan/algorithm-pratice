package chapter4

func BottomUpMinimalChange(nums []int) [][]int {
	var process func(nums []int) [][]int

	process = func(nums []int) [][]int {
		if len(nums) == 0 {
			return [][]int{{}}
		}
		n := len(nums)
		last := nums[n-1]
		subs := process(nums[:n-1])
		res := make([][]int, 0, 2*len(subs))
		dir := -1
		for _, sub := range subs {
			start, end := n-1, -1
			if dir == 1 {
				start, end = 0, n
			}
			for i := start; i != end; i += dir {
				tmp := make([]int, n)
				if i < n-1 {
					copy(tmp, sub[:i])
					tmp[i] = last
					copy(tmp[i+1:], sub[i:])
				} else {
					copy(tmp, sub)
					tmp[n-1] = last
				}
				res = append(res, tmp)
			}
			dir = -dir
		}

		return res
	}

	return process(nums)
}

func JohnsonTrotter(n int) [][]int {
	nums := make([]int, n)
	arrow := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i + 1
		arrow[i] = -1
	}
	res := make([][]int, 0, 100)
	res = append(res, makeCopy(nums))
	for {
		k, j := -1, -1
		for i := 0; i < n; i++ {
			if arrow[i] == -1 {
				if i-1 >= 0 && nums[i-1] < nums[i] && (k == -1 || nums[k] < nums[i]) {
					k = i
					j = i - 1
				}
			} else {
				if i+1 < n && nums[i+1] < nums[i] && (k == -1 || nums[k] < nums[i]) {
					k = i
					j = i + 1
				}
			}
		}
		if k < 0 {
			break
		}
		nums[k], nums[j] = nums[j], nums[k]
		arrow[k], arrow[j] = arrow[j], arrow[k]
		for i := 0; i < n; i++ {
			if nums[i] > nums[j] {
				arrow[i] = -arrow[i]
			}
		}

		res = append(res, makeCopy(nums))
	}
	return res
}

func makeCopy(nums []int) []int {
	tmp := make([]int, len(nums))
	copy(tmp, nums)
	return tmp
}

func LexicographicPermute(n int) [][]int {
	res := make([][]int, 0, 100)

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}
	res = append(res, makeCopy(nums))
	for {
		i := n - 2
		for i >= 0 {
			if nums[i+1] > nums[i] {
				break
			}
			i--
		}
		if i < 0 {
			break
		}

		j := n - 1
		for j > i {
			if nums[j] > nums[i] {
				break
			}
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]

		i++
		j = n - 1
		for i < j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
		res = append(res, makeCopy(nums))
	}
	return res
}
