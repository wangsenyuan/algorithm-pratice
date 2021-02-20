package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	A := readNNums(reader, n)
	solver := NewSolver(n, A)

	q := readNum(reader)

	var buf bytes.Buffer

	for q > 0 {
		q--
		m := readNum(reader)
		buf.WriteString(fmt.Sprintf("%d\n", solver.Query(m)))
	}
	fmt.Print(buf.String())
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

const H = 30
const MOD = 998244353
const MAX_N = 200010

var F [MAX_N]int64
var IF [MAX_N]int64

const FFTN = 1 << 18

var W [FFTN + 5]int64

func pow(a, b int64) int64 {
	var res int64 = 1
	for b > 0 {
		if b&1 == 1 {
			res = (res * a) % MOD
		}
		a = (a * a) % MOD
		b >>= 1
	}
	return res
}

func init() {
	F[0] = 1
	for i := 1; i < MAX_N; i++ {
		F[i] = int64(i) * F[i-1] % MOD
	}
	IF[MAX_N-1] = pow(F[MAX_N-1], MOD-2)
	for i := MAX_N - 2; i >= 0; i-- {
		IF[i] = int64(i+1) * IF[i+1] % MOD
	}
	W[0] = 1
	W[1] = pow(3, (MOD-1)/FFTN)
	for i := 2; i <= FFTN; i++ {
		W[i] = W[i-1] * W[1] % MOD
	}
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

type Solver []int64

func NewSolver(n int, A []int) Solver {
	var N = 1
	for N <= n {
		N *= 2
	}
	ones := make([]int64, N+1)
	zeros := make([]int64, N+1)
	ans := make([]int64, n+1)
	for i := 0; i < H; i++ {
		var k int
		for j := 0; j < n; j++ {
			k += (A[j] >> uint(i)) & 1
		}

		for j := 1; j <= k; j += 2 {
			ones[j] = nCr(k, j)
		}
		for j := 0; j <= n-k; j++ {
			zeros[j] = nCr(n-k, j)
		}

		odd := fft(ones, k+1, zeros, n-k+1)
		for j := 1; j <= n; j++ {
			ans[j] = (ans[j] + odd[j]*(1<<uint(i))%MOD) % MOD
		}
	}

	for i := 1; i <= n; i++ {
		ans[i] += ans[i-1]
		ans[i] %= MOD
	}

	return Solver(ans)
}

func (solver Solver) Query(m int) int {
	return int(solver[m])
}

var res [FFTN + 5]int64
var R [FFTN + 5]int
var A [FFTN + 5]int64
var B [FFTN + 5]int64

var P [FFTN + 5]int64
var w [FFTN + 5]int64

func dft(a []int64, R []int, n int) {
	for i := 0; i < n; i++ {
		P[R[i]] = a[i]
	}

	for d := 1; d < n; d <<= 1 {
		ll := FFTN / (d << 1)
		for i, j := 0, 0; i < d; i, j = i+1, j+ll {
			w[i] = W[j]
		}
		for i := 0; i < n; i += d << 1 {
			for j := 0; j < d; j++ {
				y := P[i+j+d] * w[j] % MOD
				P[i+j+d] = P[i+j] + MOD - y
				P[i+j] += y
			}
		}
		if d == 1<<15 {
			for i := 0; i < n; i++ {
				P[i] %= MOD
			}
		}
	}
	for i := 0; i < n; i++ {
		a[i] = P[i] % MOD
	}
}

func idft(a []int64, R []int, n int) {
	for i := 0; i < n; i++ {
		P[R[i]] = a[i]
	}

	for d := 1; d < n; d <<= 1 {
		ll := FFTN / (d << 1)
		for i, j := 0, FFTN; i < d; i, j = i+1, j-ll {
			w[i] = W[j]
		}
		for i := 0; i < n; i += d << 1 {
			for j := 0; j < d; j++ {
				y := P[i+j+d] * w[j] % MOD
				P[i+j+d] = P[i+j] + MOD - y
				P[i+j] += y
			}
		}
		if d == 1<<15 {
			for i := 0; i < n; i++ {
				P[i] %= MOD
			}
		}
	}
	tmp := pow(int64(n), MOD-2)
	for i := 0; i < n; i++ {
		a[i] = P[i] * tmp % MOD
	}
}

func fft(a []int64, sa int, b []int64, sb int) []int64 {
	sa--
	sb--
	var sc = 1
	for sc <= sa+sb {
		sc <<= 1
	}

	for i := 0; i < sc; i++ {
		R[i] = R[i>>1] >> 1
		if i&1 == 1 {
			R[i] = R[i] | sc>>1
		}
	}
	for i := 0; i < sc; i++ {
		A[i] = 0
		if i <= sa {
			A[i] = a[i]
		}
		B[i] = 0
		if i <= sb {
			B[i] = b[i]
		}
	}
	dft(A[:], R[:], sc)
	dft(B[:], R[:], sc)

	for i := 0; i < sc; i++ {
		A[i] = A[i] * B[i] % MOD
	}

	idft(A[:], R[:], sc)

	for i := 0; i <= sa+sb; i++ {
		res[i] = A[i]
	}
	return res[:(sa + sb + 1)]
}
