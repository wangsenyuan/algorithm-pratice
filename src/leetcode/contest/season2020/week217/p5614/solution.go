package p5614

const INF = 1 << 30

func mostCompetitive(nums []int, k int) []int {
	// if we place x at position p, then we need to make sure, there are k - p after position xi
	n := len(nums)

	arr := make([]Pair, 2*n)

	for i := 0; i < len(arr); i++ {
		arr[i] = Pair{INF, i}
	}

	update := func(p int, v Pair) {
		p += n
		arr[p] = v
		for p > 0 {
			arr[p>>1] = min(arr[p], arr[p^1])
			p >>= 1
		}
	}

	get := func(l, r int) Pair {
		l += n
		r += n
		res := Pair{INF, -1}
		for l < r {
			if l&1 == 1 {
				res = min(res, arr[l])
				l++
			}
			if r&1 == 1 {
				r--
				res = min(res, arr[r])
			}
			l >>= 1
			r >>= 1
		}
		return res
	}

	res := make([]int, k)
	var p int
	for i, j := 0, 0; i < n; i++ {
		update(i, Pair{nums[i], i})
		if i >= n-k {
			x := get(j, i+1)
			res[p] = x.first
			p++
			for j <= x.second {
				update(j, Pair{INF, j})
				j++
			}
		}
	}

	return res
}

type Pair struct {
	first, second int
}

func min(a, b Pair) Pair {
	if a.first < b.first || a.first == b.first && a.second < b.second {
		return a
	}
	return b
}
