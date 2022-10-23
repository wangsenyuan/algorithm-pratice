package p2449

import "sort"

func makeSimilar(nums []int, target []int) int64 {
	sort.Ints(nums)
	sort.Ints(target)

	a, b := partition(nums, isEven)
	c, d := partition(target, isEven)
	a = append(a, b...)
	c = append(c, d...)

	return solve(a, c)
}

func partition(arr []int, f func(int) bool) ([]int, []int) {
	var a []int
	var b []int
	for _, num := range arr {
		if f(num) {
			a = append(a, num)
		} else {
			b = append(b, num)
		}
	}
	return a, b
}

func isEven(x int) bool {
	return x%2 == 0
}

func solve(nums []int, target []int) int64 {

	// parity same, can add 2/-2
	var res int64
	var cnt int
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] == target[i] {
			continue
		}
		delta := (nums[i] - target[i]) / 2
		if delta > 0 {
			// need to -2
			if cnt < 0 {
				if -cnt >= delta {
					cnt += delta
				} else {
					x := delta + cnt
					res += int64(x)
					cnt = x
				}
			} else {
				// 现在减去-2， 需要在后续步骤增加2
				cnt += delta
				res += int64(delta)
			}
		} else {
			delta *= -1
			// nums[i] < target[i] 需要增加2, 那么后续要减去2
			if cnt > 0 {
				if cnt >= delta {
					cnt -= delta
				} else {
					// cnt
					x := delta - cnt
					res += int64(x)
					cnt = -x
				}
			} else {
				cnt -= delta
				res += int64(delta)
			}
		}
	}
	// cnt == 0
	return res
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
