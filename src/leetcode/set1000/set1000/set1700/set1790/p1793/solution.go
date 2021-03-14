package p1793

const INF = 1 << 30

func maximumScore(nums []int, k int) int {
	// i <= k <= j
	// 如果向左延长，确定以[i..k]之间的最小值为码，可以向右延长到j
	// 所以用segtree, 找到最右的j，且k...j >= x
	n := len(nums)
	tree1 := make([]int, 4*n)
	tree2 := make([]int, 4*n)

	var build func(i int, ll, rr int)

	build = func(i int, ll, rr int) {
		if ll == rr {
			if ll >= k {
				tree1[i] = nums[ll]
				tree2[i] = INF
			} else {
				tree1[i] = INF
				tree2[i] = nums[ll]
			}
			return
		}
		mid := (ll + rr) / 2
		build(i<<1, ll, mid)
		build((i<<1)|1, mid+1, rr)
		tree1[i] = min(tree1[i<<1], tree1[(i<<1)|1])
		tree2[i] = min(tree2[i<<1], tree2[(i<<1)|1])
	}

	build(1, 0, n-1)

	findRight := func(x int) int {
		i := 1
		ll, rr := 0, n-1
		for ll < rr {
			if tree1[i<<1] >= x {
				// go right
				i = (i << 1) | 1
				ll = (ll+rr)/2 + 1
			} else {
				i = i << 1
				rr = (ll + rr) / 2
			}
		}

		return ll
	}

	findLeft := func(x int) int {
		i := 1
		ll, rr := 0, n-1
		for ll < rr {
			if tree2[(i<<1)|1] >= x {
				// go left
				i = i << 1
				rr = (ll + rr) / 2
			} else {
				i = (i << 1) | 1
				ll = (ll+rr)/2 + 1
			}
		}

		return ll
	}

	var best int
	x := INF
	for i := k; i >= 0; i-- {
		x = min(x, nums[i])
		j := findRight(x)
		if j < n && nums[j] < x {
			j--
		}
		best = max(best, x*(j-i+1))
	}
	x = INF
	for j := k; j < n; j++ {
		x = min(x, nums[j])
		i := findLeft(x)
		if i < n && nums[i] < x {
			i++
		}
		best = max(best, x*(j-i+1))
	}

	return best
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
