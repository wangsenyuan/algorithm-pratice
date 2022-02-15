package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--

		n := readNum(reader)

		P := make([][]int, n)
		for i := 0; i < n; i++ {
			P[i] = readNNums(reader, 2)
		}

		res := solve(n, P)

		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
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

const MOD = 1000000007

const N = 200010

var F []int64
var I []int64

func init() {
	F = make([]int64, N)
	I = make([]int64, N)
	F[0] = 1
	for i := 1; i < N; i++ {
		F[i] = int64(i) * F[i-1] % MOD
	}
	I[N-1] = pow(F[N-1], MOD-2)
	for i := N - 2; i >= 0; i-- {
		I[i] = int64(i+1) * I[i+1] % MOD
	}
}

func nCr(n, r int) int64 {
	if n < r {
		return 0
	}
	res := F[n]
	res = mul(res, I[r])
	res = mul(res, I[n-r])
	return res
}

func pow(a, b int64) int64 {
	var res int64 = 1
	for b > 0 {
		if b&1 == 1 {
			res = mul(res, a)
		}
		a = mul(a, a)
		b >>= 1
	}
	return res
}

func mul(a, b int64) int64 {
	a *= b
	return a % MOD
}

func solve(n int, P [][]int) int {
	L := make([][]int, n+1)
	R := make([][]int, n+1)

	for i := 1; i <= n; i++ {
		L[i] = make([]int, 0, 2)
		L[i] = append(L[i], i-1)
		R[i] = make([]int, 0, 2)
		R[i] = append(R[i], i+1)
	}

	set := make(map[Pair]bool)
	for _, p := range P {
		l, r := p[0], p[1]

		set[Pair{l, r}] = true
		L[l] = append(L[l], r)
		R[r] = append(R[r], l)
	}

	for i := 1; i <= n; i++ {
		sort.Ints(L[i])
		sort.Ints(R[i])
	}
	var loop func(l, r int) int64

	loop = func(l, r int) int64 {
		if r < l {
			return 1
		}
		if !set[Pair{l, r}] {
			return 0
		}

		if l == r {
			return 1
		}
		i := sort.Search(len(L[l]), func(i int) bool {
			return L[l][i] >= r
		})
		mx := L[l][i-1] + 1

		j := sort.Search(len(R[r]), func(j int) bool {
			return R[r][j] > l
		})
		if mx != R[r][j]-1 {
			return 0
		}

		// find pos, which is x + 1 (also y - 1)
		a := mul(loop(l, mx-1), loop(mx+1, r))
		b := nCr(r-l, mx-l)
		return mul(a, b)
	}

	return int(loop(1, n))
}

type Pair struct {
	first, second int
}
