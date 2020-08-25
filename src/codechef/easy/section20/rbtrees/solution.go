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
		C, _ := reader.ReadString('\n')
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

const INF = 1 << 30

func solve(n int, E [][]int, C string) int {
	G := getGraph(n, E)

	cnt := make([]int, 2)

	for i := 0; i < n; i++ {
		cnt[int(C[i]-'0')]++
	}

	D := make([]int, n)

	var dfs func(p, u, d int)
	dfs = func(p, u, d int) {
		D[u] = d
		for _, v := range G[u] {
			if v == p {
				continue
			}
			dfs(u, v, 1-d)
		}
	}

	dfs(-1, 0, 0)

	cnt2 := make([]int, 2)

	for i := 0; i < n; i++ {
		cnt2[D[i]]++
	}

	var res int

	var dfs2 func(p, u int, red, blue int) (excess int, req int)

	dfs2 = func(p, u int, red, blue int) (excess int, req int) {
		// try assign red to level 0, and blue to level 1

		for _, v := range G[u] {
			if p == v {
				continue
			}
			a, b := dfs2(u, v, red, blue)
			excess += a
			req += b
		}

		if D[u] == 0 && int(C[u]-'0') == blue {
			excess++
		}
		if D[u] == 1 && int(C[u]-'0') == red {
			req++
		}
		if excess > req {
			res += excess - req
			excess -= req
			req = 0
		} else if excess < req {
			res += req - excess
			req -= excess
			excess = 0
		}
		return
	}

	var ans = INF

	if cnt[0] == cnt2[0] {
		//red is 0, blue is 1
		res = 0
		dfs2(-1, 0, 0, 1)
		ans = min(ans, res)
	}

	if cnt[0] == cnt2[1] {
		res = 0
		dfs2(-1, 0, 1, 0)
		ans = min(ans, res)
	}

	if ans == INF {
		return -1
	}
	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
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
