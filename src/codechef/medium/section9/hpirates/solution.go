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

	n, m := readTwoNums(reader)

	E := make([][]int, m)

	for i := 0; i < m; i++ {
		E[i] = readNNums(reader, 2)
	}

	ok, res := solve(n, m, E)

	if !ok {
		fmt.Println("-1")
	} else {
		fmt.Printf("%d %d\n", len(res), len(res[0]))
		for i := 0; i < len(res); i++ {
			fmt.Println(res[i])
		}
	}
}

func solve(N, M int, E [][]int) (bool, []string) {
	if M != N-1 {
		return false, nil
	}
	conns := make([][]int, N)
	for i := 0; i < N; i++ {
		conns[i] = make([]int, 0, 10)
	}

	for _, e := range E {
		u, v := e[0], e[1]
		u--
		v--
		conns[u] = append(conns[u], v)
		conns[v] = append(conns[v], u)
	}

	vis := make([]bool, N)
	sz := make([]int, N)

	var dfs1 func(p int, u int) bool

	dfs1 = func(p int, u int) bool {
		if vis[u] {
			return false
		}
		vis[u] = true
		sz[u] = 1
		for _, v := range conns[u] {
			if p == v {
				continue
			}
			if !dfs1(u, v) {
				return false
			}
			sz[u] += sz[v] + 1
		}

		return true
	}

	if !dfs1(-1, 0) {
		return false, nil
	}

	for i := 0; i < N; i++ {
		if !vis[i] {
			return false, nil
		}
	}

	// a tree
	grid := make([][]byte, 0, N)

	C := sz[0]

	var dfs2 func(p, u, d, l int) int

	dfs2 = func(p, u, d, l int) int {
		if len(grid) == d {
			//it is the first grid
			grid = append(grid, make([]byte, C))
		}
		x := byte(d&1 + '0')
		var w = 1
		for _, v := range conns[u] {
			if p == v {
				continue
			}
			w += dfs2(u, v, d+1, l+w)

			for i := d + 1; i < len(grid); i++ {
				grid[i][l+w] = x
			}

			w++
		}

		for i := 0; i < w; i++ {
			grid[d][l+i] = x
		}

		for i := d + 1; i < len(grid); i++ {
			grid[i][l] = x
		}

		return w
	}

	dfs2(-1, 0, 0, 0)

	ans := make([]string, len(grid))
	for i := 0; i < len(grid); i++ {
		ans[i] = string(grid[i])
	}

	return true, ans
}
