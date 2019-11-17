package p1262

func maxSumDivThree(nums []int) int {
	n := len(nums)
	// (a + b + c) % 3 == 0
	// a % 3 == 1, then b % 3 == 2 and c % 3 == 0
	const INF = 1 << 20
	var sum int
	var best int
	x := []int{INF, INF}
	y := []int{INF, INF}

	for i := 0; i < n; i++ {
		num := nums[i]
		sum += num

		if num%3 == 1 {
			update(x, num)
		} else if num%3 == 2 {
			update(y, num)
		}

		r := sum % 3
		if r == 0 {
			best = sum
		} else if r == 1 {
			// remove one x, or remove two y
			if x[0] < INF {
				best = max(best, sum-x[0])
			}
			if y[1] < INF {
				best = max(best, sum-y[0]-y[1])
			}
		} else {
			// remove one y, or remove two x
			if y[0] < INF {
				best = max(best, sum-y[0])
			}
			if x[1] < INF {
				best = max(best, sum-x[0]-x[1])
			}
		}
	}

	return best
}

func update(arr []int, num int) {
	if num < arr[0] {
		arr[1] = arr[0]
		arr[0] = num
	} else if num < arr[1] {
		arr[1] = num
	}
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
