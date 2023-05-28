package p2707

func canTraverseAllPairs(nums []int) bool {
	n := len(nums)

	if n == 1 {
		return true
	}

	arr := make([]int, n)
	cnt := make([]int, n)
	sz := n
	for i := 0; i < n; i++ {
		arr[i] = i
		cnt[i]++
	}

	var find func(int) int
	find = func(u int) int {
		if arr[u] != u {
			arr[u] = find(arr[u])
		}
		return arr[u]
	}

	union := func(a, b int) bool {
		a = find(a)
		b = find(b)
		if a == b {
			return false
		}
		if cnt[a] < cnt[b] {
			a, b = b, a
		}
		cnt[a] += cnt[b]
		arr[b] = a
		sz--
		return true
	}

	x := max(nums...)

	lpf := make([]int, x+1)
	prev := make([]int, x+1)
	for i := 2; i <= x; i++ {
		prev[i] = -1
		if lpf[i] == 0 {
			lpf[i] = i
			if x/i < i {
				continue
			}
			for j := i * i; j <= x; j += i {
				if lpf[j] == 0 {
					lpf[j] = i
				}
			}
		}
	}

	for i, num := range nums {
		if num == 1 {
			return false
		}
		cur := num
		for cur > 1 {
			f := lpf[cur]
			if prev[f] != -1 {
				union(prev[f], i)
			}
			cur /= f
		}

		cur = num

		for cur > 1 {
			prev[lpf[cur]] = i
			cur /= lpf[cur]
		}
	}

	return sz == 1
}

func max(arr ...int) int {
	res := arr[0]
	for i := 1; i < len(arr); i++ {
		if res < arr[i] {
			res = arr[i]
		}
	}
	return res
}
