package p1521

func closestToTarget(arr []int, target int) int {
	if len(arr) == 0 {
		return target
	}
	best := abs(target - arr[0])
	nums := []int{arr[0]}

	for _, x := range arr {
		mem := make(map[int]bool)
		for _, num := range nums {
			mem[x&num] = true
		}
		mem[x] = true
		nums = nums[:0]
		for x := range mem {
			d := abs(x - target)
			best = min(best, d)
			nums = append(nums, x)
		}
	}

	return best
}

func closestToTarget1(arr []int, target int) int {

	merge := func(l, m, r int) int {
		ll, rr := m, m+1
		y := arr[ll] & arr[rr]

		if y < target {
			// and more nums, get less, best is y
			return target - y
		}

		for ll >= l || rr <= r {
			if ll > l && y&arr[ll-1] >= target {
				ll--
				y &= arr[ll]
				continue
			}
			if rr < r && y&arr[rr+1] >= target {
				rr++
				y &= arr[rr]
				continue
			}
			break
		}

		d := y - target

		if ll > l && abs(y&arr[ll-1]-target) < d {
			d = abs(y&arr[ll-1] - target)
		}
		if rr < r && abs(y&arr[rr+1]-target) < d {
			d = abs(y&arr[rr+1] - target)
		}
		return d
	}

	var dfs func(l, r int) int

	dfs = func(l, r int) int {
		if l == r {
			return abs(arr[l] - target)
		}
		mid := (l + r) / 2
		a := dfs(l, mid)
		b := dfs(mid+1, r)
		c := merge(l, mid, r)
		return min3(a, b, c)
	}

	return dfs(0, len(arr)-1)
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func min3(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= a && b <= c {
		return b
	}
	return c
}
