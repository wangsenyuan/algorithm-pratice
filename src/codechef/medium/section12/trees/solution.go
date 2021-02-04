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
		n, k := readTwoNums(reader)
		E := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			E[i] = readNNums(reader, 2)
		}
		res := solve(n, k, E)
		buf.WriteString(fmt.Sprintf("%d\n", res))
		reader.ReadBytes('\n')
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

func solve(n, k int, E [][]int) int64 {
	dp := make([][]int64, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int64, k+1)
	}

	next := make([]int, 2*n)
	to := make([]int, 2*n)
	var cur int
	nodes := make([]int, n)

	addEdge := func(u, v int) {
		cur++
		next[cur] = nodes[u]
		nodes[u] = cur
		to[cur] = v
	}

	for _, e := range E {
		u, v := e[0], e[1]
		addEdge(u, v)
		addEdge(v, u)
	}

	pref := make([]int64, k+1)

	var dfs func(p, u int)
	dfs = func(p, u int) {
		dp[u][0] = 1

		for i := nodes[u]; i > 0; i = next[i] {
			v := to[i]
			if p == v {
				continue
			}
			dfs(u, v)

			for j := 0; j <= k; j++ {
				pref[j] = 0
			}
			for j := 0; j <= k; j++ {
				pref[j] += dp[u][j]
				for jj := 0; jj+j < k; jj++ {
					pref[max(j, jj+1)] += dp[u][j] * dp[v][jj]
				}
			}

			for j := 0; j <= k; j++ {
				dp[u][j] = pref[j]
			}
		}
	}

	dfs(-1, 0)

	var ans int64

	for i := 0; i < n; i++ {
		for j := 0; j <= k; j++ {
			ans += dp[i][j]
		}
	}
	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
