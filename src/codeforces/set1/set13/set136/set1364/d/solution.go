package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
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

	firstLine := readNNums(reader, 3)
	n, m, k := firstLine[0], firstLine[1], firstLine[2]
	E := make([][]int, m)
	for i := 0; i < m; i++ {
		E[i] = readNNums(reader, 2)
	}
	t, res := solve(n, k, E)

	fmt.Println(t)

	if t == 2 {
		fmt.Println(len(res))
	}

	var buf bytes.Buffer

	for i := 0; i < len(res); i++ {
		buf.WriteString(strconv.Itoa(res[i]))
		if i < len(res)-1 {
			buf.WriteByte(' ')
		}
	}
	fmt.Println(buf.String())
}

func solve(n, k int, E [][]int) (int, []int) {
	G := getGraph(n, E)
	pos := make([]int, n)
	for i := 0; i < n; i++ {
		pos[i] = -1
	}
	cnt := make([]int, 2)
	vis := make([]int, 0, n)
	cyc := make([]int, 0, n)
	var dfs func(u int) bool

	dfs = func(u int) bool {
		pos[u] = len(vis)
		cnt[pos[u]&1]++
		vis = append(vis, u)
		var low = -1
		for _, v := range G[u] {
			// visited before and not parent
			if pos[v] >= 0 && pos[u] > pos[v]+1 {
				low = max(low, pos[v])
			}
		}

		if low >= 0 {
			for i := low; i <= pos[u]; i++ {
				cyc = append(cyc, vis[i])
			}
			return true
		}

		for _, v := range G[u] {
			if pos[v] < 0 {
				if dfs(v) {
					return true
				}
			}
		}
		vis = vis[:pos[u]]
		return false
	}

	hasCycle := dfs(0)

	res := make([]int, 0, n)
	h := (k + 1) / 2

	if !hasCycle {
		var x int
		if cnt[1] > cnt[0] {
			x = 1
		}
		for i := 0; i < n; i++ {
			if pos[i]&1 == x {
				res = append(res, i+1)
			}
		}
		return 1, res[:h]
	}

	if len(cyc) <= k {
		for i := 0; i < len(cyc); i++ {
			res = append(res, cyc[i]+1)
		}
		return 2, res
	}

	for i := 0; i < h; i++ {
		res = append(res, cyc[2*i]+1)
	}

	return 1, res
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
