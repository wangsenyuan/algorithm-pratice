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
	ask := func(nodes []int) (node int, dist int) {
		buf.WriteString(fmt.Sprintf("? %d", len(nodes)))
		for i := 0; i < len(nodes); i++ {
			buf.WriteString(fmt.Sprintf(" %d", nodes[i]))
		}
		fmt.Println(buf.String())
		buf.Reset()
		node, dist = readTwoNums(reader)
		return
	}

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		E := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			E[i] = readNNums(reader, 2)
		}
		first, second := solve(n, E, ask)
		fmt.Printf("! %d %d\n", first, second)
		reader.ReadString('\n')
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

func solve(n int, E [][]int, ask func(arr []int) (int, int)) (first int, second int) {
	G := getGraph(n, E)
	height := make([]int, n)
	maxHeight := make([]int, n)

	var dfs func(p, u int)

	dfs = func(p, u int) {
		maxHeight[u] = height[u]
		for _, v := range G[u] {
			if p == v {
				continue
			}
			height[v] = height[u] + 1
			dfs(u, v)
			maxHeight[u] = max(maxHeight[u], maxHeight[v])
		}
	}

	nodes := make([]int, n)
	var findNodeAt func(p, u int, need int, idx *int)
	findNodeAt = func(p, u int, need int, idx *int) {
		if height[u] == need {
			nodes[*idx] = u + 1
			*idx++
			return
		}

		for _, v := range G[u] {
			if p != v && maxHeight[v] <= n/2 {
				findNodeAt(u, v, need, idx)
			}
		}
	}
	var p int
	for i := 0; i < n; i++ {
		nodes[p] = i + 1
		p++
	}
	root, dist := ask(nodes)
	first = root

	dfs(-1, root)

	st, ed := 0, n/2

	for st <= ed {
		mid := st + ed>>1
		p = 0
		findNodeAt(-1, root, mid, &p)
		if p == 0 {
			ed = mid - 1
		} else {
			node, tmpDist := ask(nodes[:p])
			if tmpDist == dist {
				st = mid + 1
				first = node
			} else {
				ed = mid - 1
			}
		}
	}
	// first is known
	vis := getDistanceFrom(n, G, first-1)

	p = 0

	for i := 0; i < n; i++ {
		if vis[i] == dist {
			nodes[p] = i + 1
			p++
		}
	}

	second, _ = ask(nodes[:p])
	return
}

func getDistanceFrom(n int, G [][]int, first int) []int {
	vis := make([]int, n)
	for i := 0; i < n; i++ {
		vis[i] = -1
	}
	vis[first] = 0

	que := make([]int, n)
	var end int
	que[end] = first
	end++
	var front int
	for front < end {
		u := que[front]
		front++
		for _, v := range G[u] {
			if vis[v] < 0 {
				vis[v] = vis[u] + 1
				que[end] = v
				end++
			}
		}
	}

	return vis
}

func getGraph(n int, E [][]int) [][]int {
	degree := make([]int, n)
	for _, e := range E {
		u, v := e[0], e[1]
		degree[u-1]++
		degree[v-1]++
	}

	adj := make([][]int, n)
	for i := 0; i < n; i++ {
		adj[i] = make([]int, 0, degree[i])
	}

	for _, e := range E {
		u, v := e[0], e[1]
		u--
		v--
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
