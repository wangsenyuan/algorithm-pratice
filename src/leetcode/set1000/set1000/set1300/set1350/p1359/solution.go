package p1359

const MOD = 1000000007
const MAX_N = 1002

var F []int64
var P []int64

func init() {
	F = make([]int64, MAX_N)
	F[0] = 1
	for i := 1; i < MAX_N; i++ {
		F[i] = F[i-1] * int64(i)
		F[i] %= MOD
	}

	P = make([]int64, MAX_N)
	P[0] = 1
	for i := 1; i < MAX_N; i++ {
		P[i] = 2 * P[i-1]
		P[i] %= MOD
	}
}

func pow(a int64, b int) int64 {
	r := int64(1)

	for b > 0 {
		if b&1 == 1 {
			r *= a
			r %= MOD
		}
		a *= a
		a %= MOD
		b >>= 1
	}
	return r
}

func inverse(num int64) int64 {
	return pow(num, MOD-2)
}

func countOrders(n int) int {
	// if n == 1, then 1
	// pi, di, there are total 2 * n positions, if not considering order, then there are (2 * n)!
	// 2 ^^ n
	// answer seems to be (2 * n) | / 2 ^^ n
	res := F[2*n]
	p := P[n]
	r := inverse(p)
	// x := (p * r) % MOD
	// fmt.Fprintf(os.Stderr, "%d\n", x)
	res *= r
	res %= MOD
	return int(res)
}
