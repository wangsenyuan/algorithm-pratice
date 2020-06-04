package main

import "fmt"

func main() {
	var s string
	fmt.Scanf("%s", &s)

	solver := NewSolver(s)

	fmt.Println(solver.GetResult())

	var q, a, b int

	fmt.Scanf("%d", &q)

	for q > 0 {
		q--
		fmt.Scanf("%d %d", &a, &b)
		solver.Flip(a, b)
		fmt.Println(solver.GetResult())
	}
}

const MOD = 1000000007
const R = 500000004

type Solver struct {
	deg      []int64
	s        []int64
	f        []int64
	reversed []bool
	n        int
}

func NewSolver(S string) Solver {
	n := len(S)
	deg := make([]int64, n+1)
	deg[0] = 1
	for i := 1; i <= n; i++ {
		deg[i] = 2 * deg[i-1]
		if deg[i] >= MOD {
			deg[i] -= MOD
		}
	}

	s := make([]int64, n*4)
	f := make([]int64, n*4)
	r := make([]bool, n*4)
	var loop func(i int, l int, r int)

	loop = func(i int, l int, r int) {
		if l == r {
			s[i] = int64(S[l] - '0')
			f[i] = s[i]
			return
		}
		mid := (l + r) / 2
		loop(i<<1, l, mid)
		loop(i<<1|1, mid+1, r)
		s[i] = ((s[i<<1]*deg[r-mid])%MOD + s[i<<1|1]) % MOD
		f[i] = (f[i<<1] + f[i<<1|1])
	}

	loop(1, 0, n-1)

	return Solver{deg, s, f, r, n}
}

func (solver *Solver) Push(i int, ll int, rr int) {
	if solver.reversed[i] {
		if ll < rr {
			solver.reversed[i<<1] = !solver.reversed[i<<1]
			solver.reversed[i<<1|1] = !solver.reversed[i<<1|1]
		}
		sum := solver.deg[rr-ll+1] - 1
		solver.s[i] = sum - solver.s[i]
		if solver.s[i] < 0 {
			solver.s[i] += MOD
		}
		solver.f[i] = int64(rr-ll+1) - solver.f[i]
		solver.reversed[i] = false
	}
}

func (solver *Solver) Flip(l, r int) {
	l--
	r--
	deg := solver.deg
	s := solver.s
	f := solver.f
	reversed := solver.reversed

	var loop func(i int, ll, rr int)

	loop = func(i int, ll, rr int) {
		if l > rr || r < ll {
			return
		}
		solver.Push(i, ll, rr)
		if l <= ll && rr <= r {
			reversed[i] = !reversed[i]
			solver.Push(i, ll, rr)
			return
		}

		mid := (ll + rr) >> 1
		loop(i<<1, ll, mid)
		loop(i<<1|1, mid+1, rr)
		f[i] = f[i<<1] + f[i<<1|1]
		s[i] = ((s[i<<1]*deg[r-mid])%MOD + s[i<<1|1]) % MOD
	}

	loop(1, 0, solver.n-1)
}

func (solver *Solver) GetResult() int {

	deg := solver.deg
	s := solver.s
	f := solver.f

	var loop func(i int, ll int, rr int, l, r int) int64

	loop = func(i int, ll int, rr int, l, r int) int64 {
		solver.Push(i, ll, rr)
		if l > rr || r < ll {
			return 0
		}
		if l <= ll && rr <= r {
			return f[i]
		}
		mid := (ll + rr) >> 1
		res := loop(i<<1, ll, mid, l, r) + loop(i<<1|1, mid+1, rr, l, r)
		res %= MOD
		f[i] = f[i<<1] + f[i<<1|1]
		s[i] = ((s[i<<1]*deg[r-mid])%MOD + s[i<<1|1]) % MOD
		return res
	}

	n := solver.n

	lastDigit := loop(1, 0, n-1, n-1, n-1)

	if lastDigit == 1 {
		//N is odd
		ans := ((s[1] + 1) * R) % MOD
		return int(ans)
	}
	cnt := 1 - loop(1, 0, n-1, 0, n-1)&1
	ans := (s[1] * R) % MOD
	ans += cnt

	return int(ans % MOD)
}
