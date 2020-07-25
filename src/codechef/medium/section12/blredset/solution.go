package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		E := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			E[i] = readNNums(reader, 2)
		}
		C := readNNums(reader, n)
		fmt.Println(solve(n, E, C))
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

const MOD = 1000000007

func solve(n int, E [][]int, C []int) int {
	adj := getAdjGraph(n, E)

	total := make([]int, 3)
	for i := 0; i < n; i++ {
		total[C[i]]++
	}

	count := make([][]int, n)
	for i := 0; i < n; i++ {
		count[i] = make([]int, 3)
	}

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}

	var dfs func(p, u int)

	dfs = func(p, u int) {
		dp[u][0] = 0
		dp[u][1] = 0
		if C[u] == 0 {
			dp[u][1] = 1
		}
		for k := 0; k < 3; k++ {
			count[u][k] = 0
		}

		count[u][C[u]]++

		for _, v := range adj[u] {
			if p == v {
				continue
			}
			dfs(u, v)
			modAdd(&dp[u][0], dp[v][0])
			if C[u] == 0 {
				modAdd(&dp[u][1], modMul(dp[u][1], dp[v][1]))
			}
			for k := 0; k < 3; k++ {
				count[u][k] += count[v][k]
			}
		}
		modAdd(&dp[u][0], dp[u][1])
	}

	var markGood func(p, u int)

	markGood = func(p, u int) {
		if C[u] == 0 {
			var c int
			for _, v := range adj[u] {
				if p == v {
					continue
				}
				if count[v][1]+count[v][2] > 0 {
					c++
				}
			}
			if total[1]+total[2]-count[u][1]-count[u][2] > 0 {
				c++
			}
			if c > 1 {
				C[u] = 2
			}
		}

		for _, v := range adj[u] {
			if p == v {
				continue
			}
			markGood(u, v)
		}
	}

	dfs(-1, 0)

	all := dp[0][0]

	markGood(-1, 0)

	dfs(-1, 0)
	bad := dp[0][0]

	return (all - bad + MOD) % MOD
}

func modMul(a, b int) int {
	c := int64(a) * int64(b)
	return int(c % MOD)
}

func modAdd(a *int, b int) {
	*a += b
	if *a >= MOD {
		*a -= MOD
	}
}

func getAdjGraph(n int, E [][]int) [][]int {
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
