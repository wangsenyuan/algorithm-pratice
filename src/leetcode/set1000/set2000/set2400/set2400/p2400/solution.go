package p2400

const MOD = 1000000007

func modMul(a, b int) int {
	r := int64(a) * int64(b)
	return int(r % MOD)
}

func modAdd(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func pow(a, b int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r = modMul(r, a)
		}
		a = modMul(a, a)
		b >>= 1
	}
	return r
}

const N = 1010

var F [N]int
var I [N]int

func init() {
	F[0] = 1
	for i := 1; i < N; i++ {
		F[i] = modMul(F[i-1], i)
	}

	I[N-1] = pow(F[N-1], MOD-2)
	for i := N - 2; i >= 0; i-- {
		I[i] = modMul(I[i+1], i+1)
	}
}

func nCr(n, r int) int {
	res := F[n]
	res = modMul(res, I[r])
	res = modMul(res, I[n-r])
	return res
}

func numberOfWays(startPos int, endPos int, k int) int {
	if startPos > endPos {
		startPos, endPos = endPos, startPos
	}
	right := endPos - startPos
	if right > k || k&1 != right&1 {
		return 0
	}
	// 2 * x + right = k
	x := (k - right) / 2
	return nCr(k, x)
}
