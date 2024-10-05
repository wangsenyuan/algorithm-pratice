package p3303

const B = 27

func minStartingIndex1(s string, pattern string) int {
	n := len(s)

	bases := make([]Hash, n+1)
	bases[0] = NewHash(1)

	for i := 1; i <= n; i++ {
		bases[i] = bases[i-1].MulInt(B)
	}
	pref := make([]Hash, n+1)
	pref[0] = NewHash(0)

	for i, x := range []byte(s) {
		pref[i+1] = pref[i].MulInt(B).AddInt(int(x - 'a'))
	}

	m := len(pattern)

	get := func(arr []Hash, l int, r int) Hash {
		res := arr[r+1]
		res = res.Sub(arr[l].Mul(bases[r-l+1]))
		return res
	}

	pos := make(map[Hash]int)

	for i := n - m; i >= 0; i-- {
		pos[get(pref, i, i+m-1)] = i
	}
	exp := make([]Hash, m+1)

	for i, x := range []byte(pattern) {
		exp[i+1] = exp[i].MulInt(B).AddInt(int(x - 'a'))
	}

	res := n
	if i, ok := pos[exp[m]]; ok {
		res = i
	}
	var first Hash
	for i := 0; i < m; i++ {
		first = first.MulInt(B)
		// 在改变s[i]的情况下
		second := get(exp, i+1, m-1)

		for x := 0; x < 26; x++ {
			tmp := first.AddInt(x).Mul(bases[m-i-1]).Add(second)
			if j, ok := pos[tmp]; ok {
				res = min(res, j)
			}
		}
		first = first.AddInt(int(pattern[i] - 'a'))
	}

	if res == n {
		return -1
	}

	return res
}

func minStartingIndex(s string, pattern string) int {
	n := len(s)

	bases := make([]Hash, n+1)
	bases[0] = NewHash(1)

	for i := 1; i <= n; i++ {
		bases[i] = bases[i-1].MulInt(B)
	}

	pref := make([]Hash, n+1)
	pref[0] = NewHash(0)

	for i, x := range []byte(s) {
		pref[i+1] = pref[i].MulInt(B).AddInt(int(x - 'a'))
	}

	m := len(pattern)
	exp := make([]Hash, m+1)

	for i, x := range []byte(pattern) {
		exp[i+1] = exp[i].MulInt(B).AddInt(int(x - 'a'))
	}

	get := func(arr []Hash, l int, r int) Hash {
		res := arr[r+1]
		res = res.Sub(arr[l].Mul(bases[r-l+1]))
		return res
	}

	var check func(l1 int, r1 int, l2 int, r2 int) bool

	check = func(l1 int, r1 int, l2 int, r2 int) bool {
		if l1 == r1 {
			// 直接修改
			return true
		}
		a := get(pref, l1, r1)
		b := get(exp, l2, r2)

		if a == b {
			// 完全相同
			return true
		}

		m1 := (l1 + r1) / 2
		m2 := (l2 + r2) / 2
		la := get(pref, l1, m1)
		lb := get(exp, l2, m2)
		if la == lb && check(m1+1, r1, m2+1, r2) {
			return true
		}
		ra := get(pref, m1+1, r1)
		rb := get(exp, m2+1, r2)
		if ra == rb && check(l1, m1, l2, m2) {
			return true
		}
		return false
	}

	for i := 0; i+m <= n; i++ {
		if check(i, i+m-1, 0, m-1) {
			return i
		}
	}
	return -1
}

var MOD = [...]uint{1e9 + 7, 1e9 + 9}

type Hash struct {
	h [2]uint
}

func NewHash(x uint) Hash {
	h := [2]uint{uint(x) % MOD[0], uint(x) % MOD[1]}
	return Hash{h}
}

func (this Hash) Sub(that Hash) Hash {
	h := [2]uint{0, 0}
	for i := 0; i < 2; i++ {
		h[i] = (this.h[i] + MOD[i] - that.h[i]) % MOD[i]
	}
	return Hash{h}
}

func (this Hash) Add(that Hash) Hash {
	h := [2]uint{0, 0}
	for i := 0; i < 2; i++ {
		h[i] = (this.h[i] + that.h[i]) % MOD[i]
	}
	return Hash{h}
}

func (this Hash) AddInt(x int) Hash {
	h := [2]uint{0, 0}
	for i := 0; i < 2; i++ {
		h[i] = (this.h[i] + uint(x)%MOD[i]) % MOD[i]
	}
	return Hash{h}
}

func (this Hash) Mul(that Hash) Hash {
	h := [2]uint{0, 0}
	for i := 0; i < 2; i++ {
		h[i] = (this.h[i] * that.h[i]) % MOD[i]
	}
	return Hash{h}
}

func (this Hash) MulInt(x int) Hash {
	h := [2]uint{0, 0}
	for i := 0; i < 2; i++ {
		h[i] = (this.h[i] * uint(x)) % MOD[i]
	}
	return Hash{h}
}
