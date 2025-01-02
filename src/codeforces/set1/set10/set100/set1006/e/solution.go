package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}
	os.Stdout.Write(buf.Bytes())
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

func process(reader *bufio.Reader) []int {
	n, m := readTwoNums(reader)
	p := readNNums(reader, n-1)
	queries := make([][]int, m)
	for i := 0; i < m; i++ {
		queries[i] = readNNums(reader, 2)
	}
	return solve(n, p, queries)
}

func solve(n int, p []int, queries [][]int) []int {
	g := make([][]int, n+1)

	for i := 0; i < n-1; i++ {
		g[p[i]] = append(g[p[i]], i+2)
	}
	sz := make([]int, n+1)
	var dfs func(u int)

	var order []int
	pos := make([]int, n+1)
	dfs = func(u int) {
		pos[u] = len(order)
		order = append(order, u)
		sz[u]++
		for _, v := range g[u] {
			dfs(v)
			sz[u] += sz[v]
		}
	}

	dfs(1)

	find := func(u int, k int) int {
		if sz[u] < k {
			return -1
		}
		// sz[u] >= k
		return order[pos[u]+k-1]
	}

	ans := make([]int, len(queries))

	for i, cur := range queries {
		ans[i] = find(cur[0], cur[1])
	}
	return ans
}
