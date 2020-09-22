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
		n := readNum(reader)
		V := readNNums(reader, n)
		E := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			E[i] = readNNums(reader, 2)
		}
		res := solve(n, V, E)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func solve(n int, V []int, E [][]int) int {
	G := getGraph(n, E)

	upto := make([]int, n)
	from := make([]int, n)

	var dfs1 func(p, u int)

	dfs1 = func(p, u int) {
		if p >= 0 && V[u] >= V[p] {
			upto[u] = upto[p] + 1
		} else {
			upto[u] = 1
		}
		from[u] = 1
		for _, v := range G[u] {
			if p == v {
				continue
			}
			dfs1(u, v)
			if V[v] >= V[u] {
				from[u] = max(from[u], from[v]+1)
			}
		}
	}

	dfs1(-1, 0)

	ans := make([]int, n+1)
	var best int
	var dfs2 func(p, u int, ii int)

	dfs2 = func(p, u int, ii int) {
		i := search(ii, func(i int) bool {
			return ans[i] > V[u]
		})
		i--
		// i == 0 || ans[i-1] <= V[u]
		if i > 0 {
			best = max(best, i+from[u])
		}
		tmp := ans[upto[u]]

		ans[upto[u]] = min(ans[upto[u]], V[u])

		for _, v := range G[u] {
			if v != p {
				dfs2(u, v, ii+1)
			}
		}
		ans[upto[u]] = tmp
	}
	ans[0] = 0

	for i := 1; i <= n; i++ {
		ans[i] = 1 << 30
	}

	dfs2(-1, 0, 1)

	return best
}

func search(n int, fn func(int) bool) int {
	var left = 0
	var right = n
	for left < right {
		mid := (left + right) / 2
		if fn(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
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
