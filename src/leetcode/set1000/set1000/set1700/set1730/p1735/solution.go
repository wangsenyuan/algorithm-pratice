package p1735

const MOD = 1000000007

const MAX_K = 10010

var P []int
var pf [MAX_K]int

const MAX_N = 13000

var F [MAX_N]int64
var IF [MAX_N]int64

func init() {
	P = make([]int, 0, 1231)
	for x := 2; x < MAX_K; x++ {
		if pf[x] == 0 {
			P = append(P, x)
			for y := x; y < MAX_K; y += x {
				if pf[y] == 0 {
					pf[y] = x
				}
			}
		}
	}

	F[0] = 1
	for i := 1; i < MAX_N; i++ {
		F[i] = int64(i) * F[i-1] % MOD
	}

	IF[MAX_N-1] = pow(F[MAX_N-1], MOD-2)

	for i := MAX_N - 2; i >= 0; i-- {
		IF[i] = int64(i+1) * IF[i+1] % MOD
	}
}

func pow(a, b int64) int64 {
	var r int64 = 1
	for b > 0 {
		if b&1 == 1 {
			r = r * a % MOD
		}
		a = a * a % MOD
		b >>= 1
	}
	return r
}

func nCr(n, r int) int64 {
	if n < r {
		return 0
	}
	res := F[n]
	res = res * IF[r] % MOD
	res = res * IF[n-r] % MOD
	return res
}

func waysToFillArray(queries [][]int) []int {

	ans := make([]int, len(queries))
	//cnt := make([]int, len(P))
	for i := 0; i < len(queries); i++ {
		cur := queries[i]
		n, k := cur[0], cur[1]
		ans[i] = solve(n, k)
	}

	return ans
}

func solve(n, k int) int {
	// k = product(p ^^ q)
	// then F[k] = C(n, q1) * C(n - q1, q2) ...
	var res int64 = 1

	for i := 0; i < len(P) && P[i] <= k; i++ {
		x := P[i]
		if k%x != 0 {
			continue
		}
		var cnt int
		for k%x == 0 {
			cnt++
			k /= x
		}

		tmp := nCr(n+cnt-1, cnt)
		res = res * tmp % MOD
	}

	return int(res)
}
