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
		X := readNNums(reader, n)
		ladders := make([][]int, k)
		for i := 0; i < k; i++ {
			ladders[i] = readNNums(reader, 5)
		}
		ok, res := solve(n, m, ladders, X)

		if !ok {
			buf.WriteString("NO ESCAPE\n")
		} else {
			buf.WriteString(fmt.Sprintf("%d\n", res))
		}
	}
	fmt.Print(buf.String())
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

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

const INF = 1 << 62

func solve(n int, m int, ladders [][]int, X []int) (bool, int64) {
	// 到达某个梯子的出口时，可以计算这一层所有梯子入口的cost
	// 但是这里有个潜在的 m * m 的问题
	// 假设某个梯子出口的cost时y，同一层梯子入口的cost 可以更新为
	// y + dist * x - h
	// y + (a - b) * x - h
	// y - b * x + a * x - h
	// 假设b是第一个梯子的坐标
	type End struct {
		id  int
		pos int
		out bool
	}

	ends := make([][]End, n+1)

	for i, ld := range ladders {
		a, b, c, d := ld[0], ld[1], ld[2], ld[3]
		if ends[a] == nil {
			ends[a] = make([]End, 0, 1)
		}
		ends[a] = append(ends[a], End{i, b, false})
		if ends[c] == nil {
			ends[c] = make([]End, 0, 1)
		}
		ends[c] = append(ends[c], End{i, d, true})
	}

	for i := 1; i <= n; i++ {
		sort.Slice(ends[i], func(j, k int) bool {
			return ends[i][j].pos < ends[i][k].pos || ends[i][j].pos == ends[i][k].pos && ends[i][j].out
		})
	}
	dp := make([]int64, len(ladders))
	for i := 0; i < len(dp); i++ {
		dp[i] = INF
	}
	for _, end := range ends[1] {
		dp[end.id] = int64(end.pos-1)*int64(X[0]) - int64(ladders[end.id][4])
	}

	// y + dist * x - h
	// y + (a - b) * x - h
	// y - b * x + a * x - h
	// 假设b是第一个梯子的坐标

	for r := 2; r <= n; r++ {
		eds := ends[r]

		var best int64 = INF

		for i := 0; i < len(eds); i++ {
			if eds[i].out {
				// a out
				if dp[eds[i].id] == INF {
					continue
				}
				tmp := dp[eds[i].id] - int64(eds[i].pos)*int64(X[r-1])
				best = min(best, tmp)
			} else {
				if best == INF {
					continue
				}
				tmp := best + int64(eds[i].pos)*int64(X[r-1]) - int64(ladders[eds[i].id][4])
				dp[eds[i].id] = min(dp[eds[i].id], tmp)
			}
		}

		best = INF

		for i := len(eds) - 1; i >= 0; i-- {
			if eds[i].out {
				if dp[eds[i].id] == INF {
					continue
				}
				tmp := dp[eds[i].id] + int64(eds[i].pos)*int64(X[r-1])
				best = min(best, tmp)
			} else {
				if best == INF {
					continue
				}
				tmp := best - int64(eds[i].pos)*int64(X[r-1]) - int64(ladders[eds[i].id][4])
				dp[eds[i].id] = min(dp[eds[i].id], tmp)
			}
		}
	}

	var res int64 = INF

	for i := 0; i < len(ends[n]); i++ {
		if dp[ends[n][i].id] < INF {
			tmp := dp[ends[n][i].id] + int64(X[n-1])*int64(m-ends[n][i].pos)
			res = min(res, tmp)
		}
	}

	if res == INF {
		return false, 0
	}
	return true, res
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}

func search(n int, f func(int) bool) int {
	l, r := 0, n
	for l < r {
		mid := (l + r) / 2
		if f(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}
