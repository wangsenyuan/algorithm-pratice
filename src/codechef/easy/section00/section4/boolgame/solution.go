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

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, m, k := readThreeNums(reader)
		G := readNNums(reader, n)
		L := make([][]int, m)
		for i := 0; i < m; i++ {
			L[i] = readNNums(reader, 3)
		}
		res := solve(k, G, L)
		for i := 0; i < len(res); i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
		buf.WriteByte('\n')
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

func solve(K int, G []int, L [][]int) []int64 {
	n := len(G)
	cc := make([]int, n+1)
	for _, l := range L {
		cc[l[1]]++
	}
	reward := make([][]Pair, n+1)
	for i := 0; i <= n; i++ {
		reward[i] = make([]Pair, 0, cc[i])
	}

	for _, l := range L {
		u, v, d := l[0], l[1], l[2]
		reward[v] = append(reward[v], Pair{u, d})
	}

	for i := 0; i <= n; i++ {
		sort.Sort(Pairs(reward[i]))
	}
	dp := make([][][]int64, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([][]int64, n+1)
		for j := 0; j <= n; j++ {
			dp[i][j] = make([]int64, 0, K)
		}
	}
	push := func(i, j int, v int64) {
		dp[i][j] = append(dp[i][j], v)
	}

	push(0, 0, 0)
	// dp[i][j] = j < i, 且x[j] = 0, x[j+1...i] 为1时的最大集合
	for i := 1; i <= n; i++ {
		tmp := int64(G[i-1])
		for _, r := range reward[i] {
			// 所有以i结尾的线段的收益
			tmp += int64(r.se)
		}
		var p int
		for j := 0; j < i; j++ {
			for p < len(reward[i]) && reward[i][p].fi <= j {
				tmp -= int64(reward[i][p].se)
				p++
			}
			for _, r := range dp[i-1][j] {
				push(i, j, tmp+r)
			}
		}
		// if set i is zero, dp[i][j] = dp[i-1][j]
		for j := 0; j < i; j++ {
			for _, r := range dp[i-1][j] {
				push(i, i, r)
			}
		}
		sort.Sort(Int64Gt(dp[i][i]))
		dp[i][i] = dp[i][i][:min(len(dp[i][i]), K)]
	}
	res := make([]int64, 0, K)

	for j := 0; j <= n; j++ {
		for _, r := range dp[n][j] {
			res = append(res, r)
		}
	}
	sort.Sort(Int64Gt(res))
	return res[:min(len(res), K)]
}

type Pair struct {
	fi, se int
}

type Pairs []Pair

func (ps Pairs) Len() int {
	return len(ps)
}

func (ps Pairs) Less(i, j int) bool {
	return ps[i].fi < ps[j].fi
}

func (ps Pairs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

type Int64Gt []int64

func (this Int64Gt) Len() int {
	return len(this)
}

func (this Int64Gt) Less(i, j int) bool {
	return this[i] > this[j]
}

func (this Int64Gt) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
