package p2877

func minSizeSubarray(nums []int, target int) int {
	n := len(nums)

	arr := make([]int, 2*n)
	copy(arr, nums)
	copy(arr[n:], nums)

	pos := make(map[int]int)
	pos[0] = 0
	ans := -1
	var sum int
	for i := 1; i <= 2*n; i++ {
		sum += arr[i-1]
		if j, ok := pos[sum-target]; ok {
			if ans < 0 || ans > i-j {
				ans = i - j
			}
		}
		pos[sum] = i
	}

	sum /= 2

	if target >= sum {
		// some suf, whole nums, pref
		x := target / sum
		r := target % sum
		if r == 0 {
			ans = x * n
		}
		pos = make(map[int]int)
		pos[0] = 0
		var tmp int
		for i := 1; i <= 2*n; i++ {
			tmp += arr[i-1]
			if j, ok := pos[tmp-r]; ok {
				cur := x*n + i - j
				if ans < 0 || ans > cur {
					ans = cur
				}
			}
			pos[tmp] = i
		}
	}

	return ans
}
