package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	_, _, _, r, res := process(reader)

	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", r))
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d ", x))
	}
	buf.WriteByte('\n')
	buf.WriteTo(os.Stdout)
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

func process(reader *bufio.Reader) (n int, k int, edges [][]int, r int, res []int) {
	n, k = readTwoNums(reader)
	edges = make([][]int, n-1)
	for i := range n - 1 {
		edges[i] = readNNums(reader, 2)
	}
	r, res = solve(n, edges, k)
	return
}

func solve(n int, edges [][]int, k int) (r int, res []int) {
	g := make([][]int, n)
	for i, e := range edges {
		u, v := e[0]-1, e[1]-1
		g[u] = append(g[u], i)
		g[v] = append(g[v], i)
	}

	res = make([]int, n-1)
	var cnt int
	var dfs func(p int, u int, x int, mid int)
	dfs = func(p int, u int, x int, mid int) {
		if len(g[u]) <= mid {
			c := 1
			for _, i := range g[u] {
				e := edges[i]
				v := u ^ (e[0] - 1) ^ (e[1] - 1)
				if p == v {
					continue
				}
				if c == x {
					c++
				}
				res[i] = c
				dfs(u, v, c, mid)

				c++
			}
		} else {
			cnt++
			c := x
			if c < 0 {
				c = 1
			}
			for _, i := range g[u] {
				e := edges[i]
				v := u ^ (e[0] - 1) ^ (e[1] - 1)
				if p != v {
					res[i] = c
					dfs(u, v, c, mid)
				}
			}
		}
	}

	check := func(mid int) bool {
		cnt = 0

		dfs(-1, 0, -1, mid)

		return cnt <= k
	}
	var l int

	for u := range n {
		r = max(r, len(g[u]))
	}

	for l < r {
		mid := (l + r) / 2
		if check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	check(r)

	return
}
