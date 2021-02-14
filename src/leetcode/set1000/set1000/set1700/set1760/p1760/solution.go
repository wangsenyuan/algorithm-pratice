package p1760

func minimumSize(nums []int, maxOperations int) int {

	check := func(x int) bool {
		var res int
		var i int
		for ; i < len(nums) && res <= maxOperations; i++ {
			num := nums[i]
			if num <= x {
				continue
			}
			res += (num+x-1)/x - 1
		}
		return i == len(nums) && res <= maxOperations
	}

	left, right := 1, max(nums)

	for left < right {
		mid := (left + right) / 2
		if check(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return right
}

func max(arr []int) int {
	res := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > res {
			res = arr[i]
		}
	}
	return res
}
