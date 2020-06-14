package p1478

const INF = 1 << 30

func minSumOfLengths(arr []int, target int) int {
	pos := make(map[int]int)
	var sum int
	pos[0] = -1

	n := len(arr)
	var flag bool
	left := make([]int, n)

	for i := 0; i < len(arr); i++ {
		if i > 0 {
			left[i] = left[i-1]
		} else {
			left[i] = INF
		}
		sum += arr[i]
		pos[sum] = i
		if j, found := pos[sum-target]; found {
			left[i] = min(left[i], i-j)
			flag = true
		}
	}

	if !flag {
		return -1
	}
	right := make([]int, n)

	pos2 := make(map[int]int)
	pos2[0] = n
	ans := INF
	sum = 0
	for i := n - 1; i >= 0; i-- {
		if i < n-1 {
			right[i] = right[i+1]
		} else {
			right[i] = INF
		}

		sum += arr[i]
		pos2[sum] = i

		if j, found := pos2[sum-target]; found {
			right[i] = min(j-i, right[i])
			if i > 0 && left[i-1] < INF {
				ans = min(ans, left[i-1]+right[i])
			}
		}
	}

	if ans < INF {
		return ans
	}
	return -1
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
