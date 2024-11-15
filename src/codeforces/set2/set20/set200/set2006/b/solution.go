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
		res := process(reader)
		s := fmt.Sprintf("%v", res)
		buf.WriteString(s[1 : len(s)-1])
		buf.WriteByte('\n')
	}

	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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
	n, w := readTwoNums(reader)
	p := readNNums(reader, n-1)
	events := make([][]int, n-1)
	for i := range n - 1 {
		events[i] = readNNums(reader, 2)
	}
	return solve(n, w, p, events)
}

func solve(n int, w int, p []int, events [][]int) []int {
	g := make([][]int, n)

	for i := range n - 1 {
		g[p[i]-1] = append(g[p[i]-1], i+1)
	}

	next := make([]int, n)
	deg := make([]int, n)

	var dfs func(u int, x int)

	dfs = func(u int, x int) {
		if u > 0 {
			next[u] = x
			deg[x]++
			deg[u]++
		}

		for i := 0; i < len(g[u]); i++ {
			v := g[u][i]
			if i+1 < len(g[u]) {
				dfs(v, g[u][i+1])
			} else {
				dfs(v, x)
			}
		}
	}

	dfs(0, 0)

	var sum int

	ans := make([]int, n-1)

	cnt := n
	// cnt表示，还剩多少个节点，在它的前面可以放置tot，可以计算dist[i, i+1]的

	for i, cur := range events {
		x, y := cur[0], cur[1]
		x--
		w -= y
		sum += y

		u := next[x]
		deg[u]--
		if deg[u] == 0 {
			cnt--
		}

		deg[x]--

		if deg[x] == 0 {
			cnt--
		}

		ans[i] = sum*2 + cnt*w
	}

	return ans
}
