package p3439

func maxDifference(s string, k int) int {
	// 只有0.。。4， 也就是5个字符
	// n <= 3 * 1e4,不能有 n * n 的复杂性
	// 固定r，也就是要知道在区间 [0...r - k]范围内的某个值
	// 如果固定最多的是x，那么就是找这个区间内 pref[x] - pref[y]的最大值
	// pref[x][r]  - pref[y][r] - (pref[x][l] - pref[y][l])的最大值
	// 那么正好就是 5 * 4 个segment tree
	// a出现奇数次，b出现偶数次，没有考虑进去
	n := len(s)
	trs := make([][][]*SegTree, 5)
	for i := 0; i < 5; i++ {
		trs[i] = make([][]*SegTree, 5)
		for j := 0; j < 5; j++ {
			trs[i][j] = make([]*SegTree, 4)
			for p := 0; p < 4; p++ {
				trs[i][j][p] = NewSegTree(n + 1)
			}
			trs[i][j][0].Update(0, 0)
		}
	}

	// 有可能是负值
	res := -inf

	pref := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		pref[i] = make([]int, 5)
	}

	for i := 0; i < n; i++ {
		x := int(s[i] - '0')
		copy(pref[i+1], pref[i])
		pref[i+1][x]++

		for u := 0; u < 5; u++ {
			for v := 0; v < 5; v++ {
				if u == v {
					continue
				}
				p := pref[i+1][u] & 1
				p2 := pref[i+1][v] & 1
				trs[u][v][(p<<1)|p2].Update(i+1, pref[i+1][u]-pref[i+1][v])

				if i >= k-1 {
					cur := pref[i+1][u] - pref[i+1][v]
					// 如果x这里出现了奇数次，
					// 如果 y要求出现偶数次，
					// 奇数 - 奇数 = 偶数
					// 偶数 - 偶数 = 偶数
					// 那么必须是前面一个偶数次的组合
					// 如果x这里出现了偶数次
					//  偶数 - 奇数 = 奇数
					p ^= 1
					p = (p << 1) | p2
					tmp := trs[u][v][p].Get(0, i-k+2)
					if tmp.first == inf {
						continue
					}
					cur -= tmp.first
					j := tmp.second
					if pref[i+1][u]-pref[j][u] > 0 && pref[i+1][v]-pref[j][v] > 0 {
						res = max(res, cur)
					}
				}
			}
		}
	}

	return res
}

const inf = 1 << 60

type pair struct {
	first  int
	second int
}

func min_of(a, b pair) pair {
	if a.first < b.first || a.first == b.first && a.second < b.second {
		return a
	}
	return b
}

type SegTree struct {
	arr []pair
	sz  int
}

func NewSegTree(n int) *SegTree {
	arr := make([]pair, 2*n)
	for i := n; i < 2*n; i++ {
		arr[i] = pair{inf, i - n}
	}
	for i := n - 1; i > 0; i-- {
		arr[i] = arr[2*i]
	}
	return &SegTree{arr, n}
}

func (t *SegTree) Update(p int, v int) {
	p += t.sz
	t.arr[p].first = v
	for p > 1 {
		t.arr[p>>1] = min_of(t.arr[p], t.arr[p^1])
		p >>= 1
	}
}

func (t *SegTree) Get(l int, r int) pair {
	l += t.sz
	r += t.sz

	res := pair{inf, -1}

	for l < r {
		if l&1 == 1 {
			res = min_of(res, t.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = min_of(res, t.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
