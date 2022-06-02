package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		A := readNNums(reader, n)
		Q := make([]int, m)
		for i := 0; i < m; i++ {
			Q[i] = readNum(reader)
		}
		res := solve(A, Q)
		for i := 0; i < m; i++ {
			buf.WriteString(fmt.Sprintf("%d\n", res[i]))
		}
	}

	fmt.Print(buf.String())
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

const MOD = 1000000007
const MAX_N = 100010

var F []int64
var I []int64
var X [][]Pair

func init() {
	F = make([]int64, MAX_N)
	I = make([]int64, MAX_N)

	F[0] = 1
	for i := 1; i < MAX_N; i++ {
		F[i] = int64(i) * F[i-1] % MOD
	}
	I[MAX_N-1] = pow(F[MAX_N-1], MOD-2)
	for i := MAX_N - 2; i >= 0; i-- {
		I[i] = int64(i+1) * I[i+1] % MOD
	}

	createLovelyCandiates()
}

func createLovelyCandiates() {

	arr := make([]int, 7)

	check := func(n int) bool {
		// the whole seqence not check
		N := 1 << uint(n)
		for i := 1; i < N-1; i++ {
			var sum int
			for j := 0; j < n; j++ {
				if (i>>uint(j))&1 == 1 {
					sum += arr[j]
				}
			}
			if sum%7 == 0 {
				return false
			}
		}
		return true
	}

	var create func(sz int, sum int, x int)

	create = func(sz int, sum int, x int) {
		if sz == 7 {
			if sum%7 == 0 && check(sz) {
				X = append(X, convert(arr))
			}
			return
		}
		if sum%7 > 0 && check(sz) {
			X = append(X, convert(arr[:sz]))
		}
		for i := x; i < 7; i++ {
			arr[sz] = i
			create(sz+1, sum+i, i)
		}
	}

	create(0, 0, 1)
}

func convert(arr []int) []Pair {
	var res []Pair
	cnt := make([]int, 7)
	for i := 0; i < len(arr); i++ {
		cnt[arr[i]]++
	}
	for i := 1; i < 7; i++ {
		if cnt[i] > 0 {
			res = append(res, Pair{cnt[i], i})
		}
	}
	return res
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

func nCr(n, r int) int {
	if n < r {
		return 0
	}
	res := F[n]
	res *= I[r]
	res %= MOD
	res *= I[n-r]
	res %= MOD
	return int(res)
}

func solve(A []int, Q []int) []int {
	cnt := make([]int, 7)

	for _, num := range A {
		cnt[num%7]++
	}
	// 1 ~ 7
	dp := make([]int, 8)

	for k := 1; k <= 7; k++ {
		for _, cur := range X {
			l := sumFirst(cur)
			if l != k {
				continue
			}
			prod := 1
			for _, p := range cur {
				i := p.second
				prod = modMul(prod, nCr(cnt[i], p.first))
			}
			dp[k] = modAdd(dp[k], prod)
		}
	}

	ans := make([]int, len(Q))

	for i := 0; i < len(Q); i++ {
		k := Q[i]
		if k <= 7 {
			ans[i] = dp[k]
		} else {
			for j := 1; j < 7; j++ {
				ans[i] = modAdd(ans[i], nCr(cnt[j], k))
			}
		}
	}

	return ans
}

func modAdd(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func modMul(a, b int) int {
	A := int64(a) * int64(b)
	A %= MOD
	return int(A)
}

type Pair struct {
	first  int
	second int
}

type Pairs []Pair

func (this Pairs) Len() int {
	return len(this)
}

func (this Pairs) Less(i, j int) bool {
	return this[i].first < this[j].first
}

func (this Pairs) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func sumFirst(p []Pair) int {
	var res int
	for _, cur := range p {
		res += cur.first
	}
	return res
}
