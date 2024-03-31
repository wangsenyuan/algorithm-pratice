package p3097

const inf = 1 << 30

func minimumSubarrayLength(nums []int, k int) int {
	if k == 0 {
		return 1
	}
	pos := make([]int, 30)
	for i := 0; i < 30; i++ {
		pos[i] = -1
	}

	res := inf

	for i, num := range nums {
		for j := 0; j < 30; j++ {
			if (num>>j)&1 == 1 {
				pos[j] = i
			}
		}
		var arr []int
		for j := 29; j >= 0; j-- {
			if (k>>j)&1 == 1 {
				if pos[j] < 0 {
					arr = arr[:0]
					break
				}
				arr = append(arr, pos[j])
			} else if pos[j] >= 0 {
				x := min(min_of(arr), pos[j])
				if x < inf && i-x+1 < res {
					res = i - x + 1
				}
			}
		}

		x := min_of(arr)
		if x < inf && i-x+1 < res {
			res = i - x + 1
		}
	}

	if res >= inf {
		return -1
	}

	return res
}

func min_of(arr []int) int {
	res := inf
	for i := 0; i < len(arr); i++ {
		if res > arr[i] {
			res = arr[i]
		}
	}
	return res
}
