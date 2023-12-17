package p2966

import "sort"

func minimumCost(nums []int) int64 {
	sort.Ints(nums)
	n := len(nums)
	sum := make([]int64, n+1)

	for i := 0; i < n; i++ {
		sum[i+1] = sum[i] + int64(nums[i])
	}

	// 回文数有多少个呢？
	// 1。。9， 11，22，33.。。
	// 1到9999 的mirror，就是回文数，再+中间为0.。。9
	// 所以一共有这么多的回文数
	// 所以就是找和中位数最接近的回文数吗？

	calc := func(x int) int64 {
		i := sort.SearchInts(nums, x)
		// nums[i] >= x
		a := int64(i)*int64(x) - sum[i]
		b := sum[n] - sum[i] - int64(n-i)*int64(x)
		return a + b
	}

	if nums[n/2] <= 10 {
		return calc(nums[n/2])
	}
	best := calc(9)

	for x := 1; x <= 9999; x++ {
		// 10, 100, 1000, 10000
		tmp := x
		y := x
		for y > 0 {
			tmp = tmp*10 + y%10
			y /= 10
		}
		best = min(best, calc(tmp))
		if tmp >= nums[n/2] {
			break
		}
	}

	for a := 0; a < 10; a++ {
		for x := 1; x <= 9999; x++ {
			tmp := x*10 + a
			y := x
			for y > 0 {
				tmp = tmp*10 + y%10
				y /= 10
			}
			best = min(best, calc(tmp))
			if tmp >= nums[n/2] {
				break
			}
		}
	}

	return best
}
