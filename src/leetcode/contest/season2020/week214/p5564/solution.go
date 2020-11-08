package p5564

const N = 100001
const MOD = 1000000007

func createSortedArray(instructions []int) int {
	arr := make([]int, 2*N)

	get := func(l, r int) int {
		l += N
		r += N
		var res int
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

	update := func(p int) {
		p += N
		arr[p]++
		for p > 0 {
			arr[p>>1] = arr[p] + arr[p^1]
			p >>= 1
		}
	}

	var res int

	for i := 0; i < len(instructions); i++ {
		cur := instructions[i]
		a := get(0, cur)
		b := get(cur+1, N)
		res += min(a, b)
		if res >= MOD {
			res -= MOD
		}
		update(cur)
	}

	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
