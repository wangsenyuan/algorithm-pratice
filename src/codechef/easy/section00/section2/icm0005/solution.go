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
		X, Y := make([]int, n), make([]int, n)
		for i := 0; i < n; i++ {
			X[i], Y[i] = readTwoNums(reader)
		}
		res := solve(X, Y)
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

const MOD = 998244353

func pow(a, b int) int {
	A := int64(a)
	R := int64(1)
	for b > 0 {
		if b&1 == 1 {
			R = (R * A) % MOD
		}
		A = (A * A) % MOD
		b >>= 1
	}
	return int(R)
}

func solve(X, Y []int) int {
	P := make(Pairs, len(X))
	for i := 0; i < len(X); i++ {
		P[i] = Pair{X[i], Y[i]}
	}

	sort.Sort(P)

	n := len(P)

	g := make([]int64, n)
	g[n-1] = int64(P[n-1].ff)
	for i := n - 2; i >= 0; i-- {
		g[i] = gcd(int64(P[i].ff), g[i+1])
	}
	// ll is lcm so far, gg is gcd so far
	ll := int64(P[0].ff)
	var res int64
	for i := 0; i < n; i++ {
		ll = ll * int64(P[i].ff) / gcd(ll, int64(P[i].ff))
		if i+1 < n && g[i+1]%ll == 0 {
			res = (res + 1) % MOD
		}
		// ll == int64(P[i].ff)
		if g[i]%ll == 0 {
			res = (res + int64(pow(2, P[i].ss))) % MOD
			res = (res + MOD - 2) % MOD
		}
	}
	return int(res)
}

func gcd(a, b int64) int64 {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

type Pair struct {
	ff, ss int
}

func (this Pair) Less(that Pair) bool {
	return this.ff < that.ff
}

type Pairs []Pair

func (ps Pairs) Len() int {
	return len(ps)
}

func (ps Pairs) Less(i, j int) bool {
	return ps[i].Less(ps[j])
}

func (ps Pairs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
