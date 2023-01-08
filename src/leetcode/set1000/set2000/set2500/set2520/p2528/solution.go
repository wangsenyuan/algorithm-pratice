package p2528

func maxPower(stations []int, rng int, k int) int64 {
	n := len(stations)

	sum := make([]int64, n+1)

	check := func(power int64) bool {
		cnt := int64(k)

		for i := 0; i <= n; i++ {
			sum[i] = 0
		}

		for i := 0; i < n; i++ {
			a := max(0, i-rng)
			b := min(n-1, i+rng)
			sum[a] += int64(stations[i])
			sum[b+1] -= int64(stations[i])
		}

		for i := 0; i < n; i++ {
			if i > 0 {
				sum[i] += sum[i-1]
			}
			if sum[i] < power {
				diff := power - sum[i]
				if diff > cnt {
					return false
				}
				cnt -= diff
				sum[i] += diff
				x := min(n-1, i+rng)
				x = min(n-1, x+rng)
				sum[x+1] -= diff
			}
		}
		return true
	}

	var l, r int64 = 0, 1 << 40

	for l < r {
		mid := (l + r) / 2
		if !check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r - 1

}

func maxPower1(stations []int, rng int, k int) int64 {
	n := len(stations)

	arr := make([]int64, 2*n)

	set := func(p int, v int) {
		p += n
		arr[p] += int64(v)
		for p > 1 {
			arr[p>>1] = arr[p] + arr[p^1]
			p >>= 1
		}
	}

	get := func(l, r int) int64 {
		var res int64
		l += n
		r += n

		for l < r {
			if l&1 == 1 {
				res += arr[l]
				l++
			}
			if r&1 == 1 {
				r--
				res += arr[r]
			}
			l >>= 1
			r >>= 1
		}
		return res
	}

	check := func(power int64) bool {
		cnt := int64(k)

		for i := 0; i < 2*n; i++ {
			arr[i] = 0
		}

		for i := 0; i < n; i++ {
			set(i, stations[i])
		}

		for i := 0; i < n; i++ {
			a := max(0, i-rng)
			b := min(n-1, i+rng)
			tmp := get(a, b+1)
			// tmp += int64(stations[i])
			if tmp < power {
				diff := power - tmp
				if diff > cnt {
					return false
				}
				cnt -= diff
				set(b, int(diff))
			}
		}

		return true
	}

	var l, r int64 = 0, 1 << 40

	for l < r {
		mid := (l + r) / 2
		if !check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r - 1
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
