package p3357

const inf = 1 << 30

func minDifference(nums []int) int {
	n := len(nums)
	p1 := -1
	p2 := -1

	var diff int
	for i := range n {
		if nums[i] > 0 {
			if i-1 >= 0 && nums[i-1] < 0 || i+1 < n && nums[i+1] < 0 {
				if p1 < 0 || nums[i] < nums[p1] {
					p1 = i
				}
				if p2 < 0 || nums[i] > nums[p2] {
					p2 = i
				}
			}

			if i+1 < n && nums[i+1] > 0 {
				diff = max(diff, abs(nums[i]-nums[i+1]))
			}
		}
	}

	if p1 == -1 {
		return diff
	}

	check := func(expect int) bool {
		x := nums[p1] + expect
		y := nums[p2] - expect
		for i := 0; i < n; {
			if nums[i] > 0 {
				i++
				continue
			}
			j := i
			for i < n && nums[i] < 0 {
				i++
			}
			if j == 0 && i == n {
				// expect >= 0
				return true
			}
			if j == 0 {
				if abs(x-nums[i]) <= expect {
					// put x
					continue
				} else if abs(y-nums[i]) <= expect {
					continue
				} else {
					return false
				}
			}
			if i == n {
				if abs(x-nums[j-1]) <= expect {
					continue
				} else if abs(y-nums[j-1]) <= expect {
					continue
				} else {
					return false
				}
			}
			a, b := min(nums[j-1], nums[i]), max(nums[j-1], nums[i])
			if abs(x-a) <= expect && abs(x-b) <= expect {
				continue
			} else if abs(y-a) <= expect && abs(y-b) <= expect {
				continue
			} else if i-j > 1 && abs(y-x) <= expect && abs(x-a) <= expect && abs(y-b) <= expect {
				continue
			} else {
				return false
			}
		}

		return true
	}

	l, r := diff, inf

	for l < r {
		mid := (l + r) / 2
		if check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return r
}

func abs(num int) int {
	return max(num, -num)
}
