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

	for tc > 0 {
		tc--
		n := readNum(reader)
		A := readNNums(reader, n)
		B := solve(n, A)

		var buf bytes.Buffer
		for i := 0; i < n; i++ {
			buf.WriteString(fmt.Sprintf("%d ", B[i]))
		}
		fmt.Println(buf.String())

	}
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

const MAX_N = 100000

var fact []int64

var ifact []int64

func init() {
	fact = make([]int64, MAX_N+1)
	ifact = make([]int64, MAX_N+1)

	fact[0] = 1
	for i := 1; i <= MAX_N; i++ {
		fact[i] = (fact[i-1] * int64(i)) % MOD
	}

	ifact[MAX_N] = inverse(fact[MAX_N])

	for i := MAX_N; i > 0; i-- {
		ifact[i-1] = (int64(i) * ifact[i]) % MOD
	}
}

func inverse(q int64) int64 {
	return pow(q, MOD-2)
}

func pow(a, b int64) int64 {
	r := int64(1)

	for b > 0 {
		if b&1 == 1 {
			r = (r * a) % MOD
		}
		a = (a * a) % MOD
		b >>= 1
	}
	return r
}

func nCr(n int, r int) int64 {
	if n < r {
		return 0
	}
	res := (ifact[r] * ifact[n-r]) % MOD
	res *= fact[n]
	return res % MOD
}

func solve(n int, A []int) []int {
	cnt := make(map[int]int)

	for _, a := range A {
		cnt[a]++
	}

	ps := make([]Pair, 0, len(cnt))

	for k, v := range cnt {
		ps = append(ps, Pair{k, v})
	}

	sort.Sort(Pairs(ps))

	ans := make([]int, n)
	dp := make([][]int64, len(ps))
	fp := make([][]int64, len(ps))
	for i := 0; i < len(ps); i++ {
		dp[i] = make([]int64, n+1)

		for j := 0; j <= n; j++ {
			dp[i][j] = nCr(ps[i].second, j)
			if j > 0 {
				dp[i][j] += dp[i][j-1]
				dp[i][j] %= MOD
			}
		}

		fp[i] = make([]int64, n+1)
		for j := 0; j <= n; j++ {
			fp[i][j] = dp[i][j]
		}
		if i == 0 {
			continue
		}
		for j := 0; j <= n; j++ {
			fp[i][j] *= fp[i-1][j]
			fp[i][j] %= MOD
		}
	}

	gp := make([][]int64, len(ps))
	for i := len(ps) - 1; i >= 0; i-- {
		gp[i] = make([]int64, n+1)
		for j := 0; j <= n; j++ {
			gp[i][j] = dp[i][j]
		}

		if i == len(ps)-1 {
			continue
		}
		for j := 0; j <= n; j++ {
			gp[i][j] *= gp[i+1][j]
			gp[i][j] %= MOD
		}
	}

	for i := 0; i < len(ps); i++ {
		cur := ps[i]

		var res int64

		for x := 1; x <= cur.second; x++ {
			var a int64 = 1
			if i > 0 {
				a = fp[i-1][x-1]
			}
			var b int64 = 1
			if i < len(ps)-1 {
				b = gp[i+1][x]
			}
			c := (a * b) % MOD
			d := nCr(cur.second, x)
			res += (c * d) % MOD
			res %= MOD
		}

		ans[cur.first-1] = int(res)
	}

	return ans
}

type Pair struct {
	first, second int
}

type Pairs []Pair

func (ps Pairs) Len() int {
	return len(ps)
}

func (ps Pairs) Less(i, j int) bool {
	return ps[i].first < ps[j].first
}

func (ps Pairs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
