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
		n, m := readTwoNums(reader)
		E := make([][]int, m)
		for i := 0; i < m; i++ {
			E[i] = readNNums(reader, 2)
		}
		ok, res := solve(n, E)
		if !ok {
			buf.WriteString("impossible\n")
			continue
		}
		buf.WriteString("possible\n")
		for i := 0; i < len(res); i++ {
			buf.WriteByte(byte('0' + res[i]))
		}
		buf.WriteByte('\n')
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

func solve(N int, E [][]int) (bool, []int) {
	// one more vertex to avoid check bound
	n := N + N&1
	degree := make([]int, n)
	rdegree := make([]int, n)
	for _, e := range E {
		u, v := e[0], e[1]
		u--
		v--
		degree[u^1]++
		degree[v^1]++

		rdegree[u]++
		rdegree[v]++
	}

	G := make([][]int, n)
	R := make([][]int, n)
	for i := 0; i < n; i++ {
		G[i] = make([]int, 0, degree[i])
		R[i] = make([]int, 0, rdegree[i])
	}

	for _, e := range E {
		u, v := e[0], e[1]
		u--
		v--
		G[u^1] = append(G[u^1], v)
		G[v^1] = append(G[v^1], u)
		R[u] = append(R[u], v^1)
		R[v] = append(R[v], u^1)
	}

	stack := make([]int, 0, n)
	vis := make([]int, n)
	var dfs1 func(p, u int)
	dfs1 = func(p, u int) {
		vis[u]++
		for _, v := range R[u] {
			if vis[v] == 0 {
				dfs1(u, v)
			}
		}
		stack = append(stack, u)
	}

	for i := 0; i < n; i++ {
		if vis[i] == 0 {
			dfs1(-1, i)
		}
	}

	comp := make([]int, n)
	for i := 0; i < n; i++ {
		comp[i] = -1
	}

	var dfs2 func(u int, c int)
	dfs2 = func(u int, c int) {
		comp[u] = c
		for _, v := range G[u] {
			if comp[v] < 0 {
				dfs2(v, c)
			}
		}
	}

	for i := n - 1; i >= 0; i-- {
		if comp[stack[i]] == -1 {
			dfs2(stack[i], stack[i])
		}
	}

	for i := 0; i < n; i += 2 {
		vis[i] = 0
		vis[i+1] = 0
		if comp[i] == comp[i^1] {
			return false, nil
		}
	}

	f := make([]int, n)
	for i := 0; i < n; i++ {
		f[i] = -1
	}

	var dfs3 func(u int, flag int)
	dfs3 = func(u int, flag int) {
		vis[u]++

		f[comp[u^1]] = 1 - flag

		for _, v := range G[u] {
			if vis[v] == 0 {
				dfs3(v, flag)
			}
		}
	}

	for i := n - 1; i >= 0; i-- {
		if vis[stack[i]] == 0 {
			u := stack[i]
			v := 1
			if f[comp[u]] >= 0 {
				v = f[comp[u]]
			}
			dfs3(u, v)
		}
	}
	ans := make([]int, N)

	for i := 0; i < N; i++ {
		ans[i] = f[comp[i]]
	}
	return true, ans
}
