package p2395

func maximumRobots(chargeTimes []int, runningCosts []int, budget int64) int {
	n := len(chargeTimes)

	arr := make([]int, 2*n)

	for i := n; i < 2*n; i++ {
		arr[i] = chargeTimes[i-n]
	}

	for i := n - 1; i > 0; i-- {
		arr[i] = max(arr[i*2], arr[i*2+1])
	}

	get := func(l, r int) int {
		l += n
		r += n
		var res int

		for l < r {
			if l&1 == 1 {
				res = max(res, arr[l])
				l++
			}
			if r&1 == 1 {
				r--
				res = max(res, arr[r])
			}
			l >>= 1
			r >>= 1
		}
		return res
	}

	sum := make([]int64, n)

	find := func(end int) int {
		lo, hi := 0, end+1

		for lo < hi {
			mid := (lo + hi) / 2
			cnt := (end - mid + 1)
			mx := get(mid, end+1)
			tmp := sum[end]
			if mid > 0 {
				tmp -= sum[mid-1]
			}
			tmp = int64(mx) + int64(cnt)*int64(tmp)
			if tmp <= budget {
				hi = mid
			} else {
				lo = mid + 1
			}
		}
		return end - hi + 1
	}

	var k int

	for i := 0; i < n; i++ {
		sum[i] = int64(runningCosts[i])
		if i > 0 {
			sum[i] += sum[i-1]
		}
		k = max(k, find(i))
	}

	return k
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
