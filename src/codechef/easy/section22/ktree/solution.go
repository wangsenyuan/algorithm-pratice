package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		var n int
		var K uint64
		s, _ := reader.ReadBytes('\n')

		pos := readInt(s, 0, &n)
		readUint64(s, pos+1, &K)
		E := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			E[i] = readNNums(reader, 2)
		}
		A := readNNums(reader, n)
		res := solve(n, K, E, A)
		if res {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
		}
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

func solve(n int, K uint64, E [][]int, A []int) bool {
	g := getGraph(n, E)
	cnt := make([]int, 64)

	var M int

	for i := 0; i < 64; i++ {
		if K&(1<<uint(i)) > 0 {
			M = i
		}
	}

	var dfs func(p, u int, big int, small int) bool

	dfs = func(p, u int, big int, small int) bool {
		big = max(big, A[u]-1)
		small = min(small, A[u]-1)
		cnt[A[u]-1]++
		if len(g[u]) == 1 && p >= 0 {
			// u is leaf
			if big > M && cnt[big]&1 == 0 {
				return true
			}
			if big == M && cnt[big]&1 == 1 {
				return true
			}
			if small > M && big > small {
				return true
			}
			if small == M {
				return true
			}
		} else {
			for _, v := range g[u] {
				if v == p {
					continue
				}
				if dfs(u, v, big, small) {
					return true
				}
			}
		}
		cnt[A[u]-1]--
		return false
	}

	return dfs(-1, 0, 0, 64)
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
