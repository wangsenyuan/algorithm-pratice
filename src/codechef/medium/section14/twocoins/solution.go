package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

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

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(reader)
		E := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			E[i] = readNNums(reader, 2)
		}
		res := solve(n, E)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
}

const INF = 1 << 30

func solve(n int, E [][]int) int {
	if n == 1 {
		return -1
	}
	G := getGraph(n, E)

	dp := make([][]int, n)

	for i := 0; i < n; i++ {
		//0 for current, 1 for parent, 2 for parent of parent
		dp[i] = make([]int, 8)
		for j := 0; j < 8; j++ {
			dp[i][j] = INF
		}
	}

	var dfs func(p, u, flag int) int

	dfs = func(p, u, flag int) int {
		flag &= 7
		if dp[u][flag] < INF {
			return dp[u][flag]
		}

		if p >= 0 && len(G[u]) == 1 {
			//has to be black
			cnt := (flag & 1) + max((flag>>1)&1, (flag>>2)&1)
			if cnt == 2 {
				dp[u][flag] = 1
			}
			return dp[u][flag]
		}

		if flag&1 > 0 {
			//black
			res := 1
			for _, v := range G[u] {
				if v == p {
					continue
				}
				res += min(dfs(u, v, flag<<1), dfs(u, v, (flag<<1)|1))
			}
			dp[u][flag] = res
			return res
		}

		// white
		fp := make([][]int, 2)
		for i := 0; i < 2; i++ {
			fp[i] = make([]int, 3)
			for j := 0; j < 3; j++ {
				fp[i][j] = INF
			}
		}
		fp[0][0] = 0

		var cur int

		for _, v := range G[u] {
			if v == p {
				continue
			}
			nxt := cur ^ 1
			for j := 0; j < 3; j++ {
				fp[nxt][j] = INF
			}
			for j := 0; j < 3; j++ {
				for take := 0; take < 2; take++ {
					k := min(2, j+take)
					fp[nxt][k] = min(fp[nxt][k], fp[cur][j]+dfs(u, v, (flag<<1)|take))
				}
			}
			cur = nxt
		}

		dp[u][flag] = fp[cur][2]
		if flag&2 > 0 {
			//parent is black
			dp[u][flag] = min(dp[u][flag], fp[cur][1])
		}
		return dp[u][flag]
	}
	res := min(dfs(-1, 0, 0), dfs(-1, 0, 1))

	return res
}

func getGraph(n int, E [][]int) [][]int {
	degree := make([]int, n)
	for _, e := range E {
		u, v := e[0]-1, e[1]-1
		degree[u]++
		degree[v]++
	}
	adj := make([][]int, n)
	for i := 0; i < n; i++ {
		adj[i] = make([]int, 0, degree[i])
	}

	for _, e := range E {
		u, v := e[0]-1, e[1]-1
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	return adj
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
