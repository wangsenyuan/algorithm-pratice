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
		E := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			E[i] = readNNums(reader, 2)
		}

		k, color := solve(n, E)
		buf.WriteString(fmt.Sprintf("%d\n", k))
		for i := 0; i < n; i++ {
			if color[i] == 0 {
				buf.WriteString(fmt.Sprintf("%d ", i+1))
			}
		}
		buf.WriteByte('\n')
		for i := 0; i < n; i++ {
			if color[i] == 1 {
				buf.WriteString(fmt.Sprintf("%d ", i+1))
			}
		}
		buf.WriteByte('\n')
		fmt.Print(buf.String())

		buf.Reset()
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

func solve(n int, E [][]int) (k int, color []int) {
	G := getGraph(n, E)
	SZ := make([]int, n)
	color = make([]int, n)

	var dfs func(p, u int)

	dfs = func(p, u int) {
		SZ[u]++
		for _, v := range G[u] {
			if v != p {
				color[v] = 1 ^ color[u]
				dfs(u, v)
				SZ[u] += SZ[v]
			}
		}
	}

	dfs(-1, 0)

	var cc int

	for i := 0; i < n; i++ {
		cc += color[i]
	}

	if cc*2 == n {
		k = 1
		return
	}

	F := make([]bool, n)

	var dfs2 func(p, u int)
	dfs2 = func(p, u int) {
		if p >= 0 && len(G[u]) == 1 {
			//leaf
			color[u] = 0
			return
		}
		if len(G[u]) == 1 && p < 0 || len(G[u]) == 2 && p >= 0 {
			// one child
			var v = G[u][0]
			if p >= 0 {
				v ^= G[u][0] ^ p
			}
			dfs2(u, v)
			color[u] = 1 ^ color[v]
			return
		}

		for _, v := range G[u] {
			if p != v {
				dfs2(u, v)
			}
		}

		c := make([]int, 2)
		for _, v := range G[u] {
			if p != v && SZ[v]&1 == 1 {
				c[color[v]]++
			}
		}

		for _, v := range G[u] {
			if p != v {
				if SZ[v]&1 == 1 && c[color[v]] > c[1^color[v]] {
					F[v] = true
					c[color[v]]--
					color[v] ^= 1
					c[color[v]]++
				}
			}
		}

		if SZ[u]&1 == 0 {
			//c[0] == c[1] +/- 1
			if c[0] < c[1] {
				color[u] = 0
			} else {
				color[u] = 1
			}
		} else {
			color[u] = 0
		}

		var ff bool
		var e = -1
		for _, v := range G[u] {
			if p != v {
				if color[v] != color[u] {
					ff = true
				}
				if SZ[v]&1 == 0 {
					e = v
				}
			}
		}

		if !ff {
			F[e] = true
			color[e] ^= 1
		}
	}

	dfs2(-1, 0)

	var flip func(p, u, e int)

	flip = func(p, u, e int) {
		color[u] ^= e
		if F[u] {
			e ^= 1
		}
		for _, v := range G[u] {
			if p != v {
				flip(u, v, e)
			}
		}
	}

	flip(-1, 0, 0)
	k = 2

	return
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
