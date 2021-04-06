package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	first := readNNums(reader, 4)
	p, q, c, m := first[0], first[1], first[2], first[3]
	M := make([][]int, m)
	for i := 0; i < m; i++ {
		M[i] = readNNums(reader, 2)
	}
	fmt.Println(solve(p, q, c, M))
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
const MAX_N = 4000010

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

func solve(p, q, c int, M [][]int) int {
	// p - c > q
	if p < q+c {
		return 0
	}

	n := len(M)
	pp := make([]Point, 0, n+1)

	for i := 0; i < n; i++ {
		if M[i][0] <= c && M[i][1] == 0 {
			return 0
		}
		if M[i][0] >= M[i][1]+c && M[i][0] <= p && M[i][1] <= q {
			pp = append(pp, Point{M[i][0], M[i][1]})
		}
	}

	pp = append(pp, Point{p, q})

	sort.Sort(Points(pp))

	dp := make([]int64, len(pp))

	for i := 0; i < len(pp); i++ {
		dp[i] = restricted(c, c, 0, pp[i].x, pp[i].y)
	}

	for i := 0; i < len(pp); i++ {
		for j := i - 1; j >= 0; j-- {
			if pp[i].x >= pp[j].x && pp[i].y >= pp[j].y {
				dp[i] -= dp[j] * restricted(c, pp[j].x, pp[j].y, pp[i].x, pp[i].y) % MOD
				dp[i] = (dp[i] + MOD) % MOD
			}
		}
	}
	return int(dp[len(pp)-1])
}

func restricted(c int, x1, y1, x2, y2 int) int64 {
	xref := y1 + (c - 1)
	yref := x1 - (c - 1)
	ans := unrestricted(x1, y1, x2, y2)
	ans -= unrestricted(xref, yref, x2, y2)
	if ans < MOD {
		ans += MOD
	}
	if ans >= MOD {
		ans -= MOD
	}
	return ans
}

func unrestricted(x1, y1, x2, y2 int) int64 {
	if x1 > x2 || y1 > y2 {
		return 0
	}
	return nCr(y2-y1+x2-x1, x2-x1)
}

type Point struct {
	x, y int
}

func (this Point) Less(that Point) bool {
	return this.x < that.x || this.x == that.x && this.y < that.y
}

type Points []Point

func (this Points) Len() int {
	return len(this)
}

func (this Points) Less(i, j int) bool {
	return this[i].Less(this[j])
}

func (this Points) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
