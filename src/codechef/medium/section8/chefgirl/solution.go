package main

import (
	"bufio"
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

func fillNNums(bs []byte, n int, res []int) {
	x := 0
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	N, M := readTwoNums(reader)

	E := make([][]int, M)
	for i := 0; i < M; i++ {
		E[i] = readNNums(reader, 4)
	}

	fmt.Println(solve(N, M, E))
}

const INF = 1 << 28

func solve(N int, M int, E [][]int) int {

	conns := make([][]Pair, N)
	for i := 0; i < N; i++ {
		conns[i] = make([]Pair, 0, 2)
	}

	for i, e := range E {
		u, v := e[0], e[1]
		u--
		v--
		conns[u] = append(conns[u], Pair{v, i})
		conns[v] = append(conns[v], Pair{u, i})
	}

	dp := make([][][]int, N)

	for i := 0; i < N; i++ {
		dp[i] = make([][]int, 32)
		for j := 0; j < 32; j++ {
			dp[i][j] = make([]int, 32)
			for k := 0; k < 32; k++ {
				dp[i][j][k] = INF
			}
		}
	}

	for i := 0; i < 32; i++ {
		for j := i; j < 32; j++ {
			dp[0][i][j] = 0
		}
	}

	leaf := make([]bool, N)

	var dfs func(p, u int)

	dfs = func(p, u int) {
		leaf[u] = true

		for _, x := range conns[u] {
			if x.first == p {
				continue
			}
			leaf[u] = false
			v := x.first
			e := E[x.second]
			a, b := e[2]-1, e[3]-1
			for i := 0; i < 32; i++ {
				for j := i; j < 32; j++ {
					if dp[u][i][j] >= INF {
						continue
					}
					dp[v][i][j] = min(dp[v][i][j], dp[u][i][j]+max(0, a-i)+max(0, j-b))
				}
			}

			dfs(u, v)
		}
	}

	dfs(-1, 0)

	cost := make([][]int, 32)
	for i := 0; i < 32; i++ {
		cost[i] = make([]int, 32)
		for j := 0; j < 32; j++ {
			cost[i][j] = INF
		}
	}

	for v := 0; v < N; v++ {
		if !leaf[v] {
			continue
		}
		for i := 0; i < 32; i++ {
			for j := 0; j < 32; j++ {
				cost[i][j] = min(cost[i][j], dp[v][i][j])
			}
		}
	}

	best := make([]int, 33)
	for i := 0; i <= 32; i++ {
		best[i] = INF
	}
	best[0] = 0
	for k := 1; k <= 32; k++ {
		for j := 0; j < k; j++ {
			best[k] = min(best[k], best[j]+cost[j][k-1])
		}
	}
	return best[32]
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Pair struct {
	first  int
	second int
}
