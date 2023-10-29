package p2913

const mod = 1e9 + 7

type lazySeg []struct{ sum, todo int }

func (t lazySeg) do(o, l, r, add int) {
	t[o].sum += add * (r - l + 1)
	t[o].todo += add
}

// o=1  [l,r] 1<=l<=r<=n
// 把 [L,R] 加一，同时返回加一之前的区间和
func (t lazySeg) addValue(o, l, r, L, R int) (res int) {
	if L <= l && r <= R {
		res = t[o].sum
		t.do(o, l, r, 1)
		return
	}
	m := (l + r) >> 1
	if add := t[o].todo; add != 0 {
		t.do(o<<1, l, m, add)
		t.do(o<<1|1, m+1, r, add)
		t[o].todo = 0
	}
	if L <= m {
		res = t.addValue(o<<1, l, m, L, R)
	}
	if m < R {
		res += t.addValue(o<<1|1, m+1, r, L, R)
	}
	t[o].sum = t[o<<1].sum + t[o<<1|1].sum
	return
}

func sumCounts(nums []int) (ans int) {
	last := map[int]int{}
	n := len(nums)
	t := make(lazySeg, n*4)
	s := 0
	for i, x := range nums {
		i++
		j := last[x]
		s += t.addValue(1, 1, n, j+1, i)*2 + i - j
		ans = (ans + s) % mod
		last[x] = i
	}
	return
}
