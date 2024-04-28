package p3130

func medianOfUniquenessArray(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 1
	}
	// 对于x，计算有多少个子数组的个数 >= x
	mx := nums[0]
	for i := 1; i < n; i++ {
		mx = max(mx, nums[i])
	}
	freq := make([]int, mx+1)

	add := func(v int, cnt *int) {
		freq[v]++
		if freq[v] == 1 {
			*cnt++
		}
	}

	rem := func(v int, cnt *int) {
		freq[v]--
		if freq[v] == 0 {
			*cnt--
		}
	}

	count := func(expect int) int {
		var cnt int
		for i := 0; i <= mx; i++ {
			freq[i] = 0
		}
		var res int
		var l, r int
		for ; r < n; r++ {
			// 以r结束的子树组，有多少 count distinct >= expect
			add(nums[r], &cnt)
			for l <= r && cnt >= expect {
				rem(nums[l], &cnt)
				if cnt < expect {
					add(nums[l], &cnt)
					break
				}
				l++
			}
			if cnt == expect {
				res += l + 1
			}
		}

		for l < r {
			rem(nums[l], &cnt)
			l++
		}

		return res
	}

	tot := n * (n + 1) / 2

	l, r := 1, n
	for l < r {
		mid := (l + r) / 2
		a := count(mid + 1)
		b := count(mid)
		x := b - a
		// 假设 x = distinct count of sub = mid 的计数
		//  y > mid的计数, z < mid 的计数
		// 那么如果mid是答案, 就有 x + y >= (tot + 1) / 2, (x + z) >= tot / 2
		// 如果 mid 处在靠左边的地方，那么就要增加l
		// 否则就减少r
		if tot-b+x >= (tot+1)/2 {
			r = mid
		} else {
			l = mid + 1
		}

	}

	return r
}
