package p2326

const MOD = 1000000007

func peopleAwareOfSecret(n int, delay int, forget int) int {

	arr := make([]int, n+1)

	set := func(p int, v int) {
		p++
		for p <= n {
			arr[p] += v
			if arr[p] >= MOD {
				arr[p] -= MOD
			}
			p += p & -p
		}
	}

	get := func(p int) int {
		p++
		var res int
		for p > 0 {
			res += arr[p]
			if res >= MOD {
				res -= MOD
			}
			p -= p & -p
		}
		return res
	}

	setRange := func(l int, r int, v int) {
		if l >= n {
			return
		}
		if v >= MOD {
			v -= MOD
		}
		set(l, v)
		if r < n {
			set(r, MOD-v)
		}
	}

	setRange(0, 1, 1)
	var res int
	for i := 0; i < n; i++ {
		x := get(i)
		res += x
		if res >= MOD {
			res -= MOD
		}
		if i-forget >= 0 {
			y := get(i - forget)
			res += MOD - y
			if res >= MOD {
				res -= MOD
			}
		}
		setRange(i+delay, i+forget, x)
	}

	return res
}
