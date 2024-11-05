package main

import (
	"bufio"
	"fmt"
	"os"
)

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

func main() {

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	defer out.Flush()

	n := readNum(in)
	adj := make([]string, n)
	for i := 0; i < n; i++ {
		adj[i] = readString(in)
	}

	m := readNum(in)
	p := readNNums(in, m)

	res := solve(adj, p)

	fmt.Fprintf(out, "%d\n", len(res))

	for _, x := range res {
		fmt.Fprintf(out, "%d ", x)
	}

	fmt.Fprintln(out)
}

func getDist(adj []string) [][]int32 {
	n := len(adj)

	que := make([]int, n)

	bfs := func(x int) []int32 {
		dist := make([]int32, n)
		for i := 0; i < n; i++ {
			dist[i] = -1
		}
		dist[x] = 0
		var head, tail int
		que[head] = x
		head++
		for tail < head {
			u := que[tail]
			tail++
			for i := 0; i < n; i++ {
				if adj[u][i] == '1' && dist[i] < 0 {
					dist[i] = dist[u] + 1
					que[head] = i
					head++
				}
			}
		}
		return dist
	}

	dist := make([][]int32, n)
	for i := 0; i < n; i++ {
		dist[i] = bfs(i)
	}
	return dist
}

func solve(adj []string, p []int) []int {
	dist := getDist(adj)
	m := len(p)

	last := p[0] - 1
	var arr []int
	arr = append(arr, p[0])
	var cnt int32
	for i := 1; i < m; i++ {
		cnt++
		u := p[i] - 1
		if dist[last][u] < cnt {
			arr = append(arr, p[i-1])
			last = p[i-1] - 1
			cnt = 1
		}
	}

	arr = append(arr, p[m-1])
	return arr
}
