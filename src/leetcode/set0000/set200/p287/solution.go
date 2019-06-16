package p287

func findDuplicate1(nums []int) int {
	fast, slow := nums[0], nums[0]

	for {
		fast = nums[nums[fast]]
		slow = nums[slow]
		if fast == slow {
			break
		}
	}

	slow = nums[0]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	return slow
}

func findDuplicate(nums []int) int {
	n := len(nums)
	left, right := 1, n-1
	for left < right {
		mid := (left + right) / 2
		var cnt int
		for _, num := range nums {
			if num <= mid {
				cnt++
			}
		}
		if cnt > mid {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
