package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	h, w, n := readThreeNums(reader)

	P := make([][]int, n)

	for i := 0; i < n; i++ {
		P[i] = readNNums(reader, 2)
	}

	fmt.Println(solve(h, w, P))
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
const MAX_N = 200010

var F [MAX_N]int64
var IF [MAX_N]int64

func init() {
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

func nCr(n, r int) int64 {
	if n < r {
		return 0
	}
	res := F[n]
	res = res * IF[r] % MOD
	res = res * IF[n-r] % MOD
	return res
}

func solve(h, w int, P [][]int) int {
	n := len(P)
	pp := make([]Point, n+1)
	for i := 0; i < n; i++ {
		pp[i] = Point{P[i][0], P[i][1]}
	}
	pp[n] = Point{h, w}
	sort.Slice(pp, func(i, j int) bool {
		return pp[i].Less(pp[j])
	})

	dp := make([]int64, n+1)

	for i := 0; i <= n; i++ {
		dp[i] = ways(1, 1, pp[i].r, pp[i].c)
		for j := i - 1; j >= 0; j-- {
			if pp[i].r >= pp[j].r && pp[i].c >= pp[j].c {
				dp[i] -= dp[j] * ways(pp[j].r, pp[j].c, pp[i].r, pp[i].c) % MOD
				dp[i] = (dp[i] + MOD) % MOD
			}
		}
	}
	return int(dp[n])
}

func ways(x1, y1, x2, y2 int) int64 {
	if x1 > x2 || y1 > y2 {
		return 0
	}
	dx := x2 - x1
	dy := y2 - y1
	return nCr(dx+dy, dx)
}

type Point struct {
	r, c int
}

func (this Point) Less(that Point) bool {
	return this.r < that.r || this.r == that.r && this.c < that.c
}
