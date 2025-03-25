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
	for _, cur := range res {
		if len(cur) == 0 {
			buf.WriteString("Impossible\n")
		} else {
			buf.WriteString(fmt.Sprintf("%d %d\n", cur[0], cur[1]))
		}
	}
	buf.WriteTo(os.Stdout)
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
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

func process(reader *bufio.Reader) [][]int {
	n := readNum(reader)
	a := readNNums(reader, n)
	g := make([]string, n)
	for i := 0; i < n; i++ {
		g[i] = readString(reader)
	}
	m := readNum(reader)
	qs := make([][]int, m)
	for i := 0; i < m; i++ {
		qs[i] = readNNums(reader, 2)
	}
	return solve(a, g, qs)
}

func solve(a []int, g []string, queries [][]int) [][]int {
	n := len(a)
	que := make([]int, n)
	bfs := func(start int) [][]int {
		dist := make([][]int, n)
		dist[start] = []int{0, a[start]}
		var head, tail int
		que[head] = start
		head++
		for tail < head {
			u := que[tail]
			tail++
			for v := 0; v < n; v++ {
				if g[u][v] == 'Y' {
					if dist[v] == nil {
						dist[v] = []int{dist[u][0] + 1, dist[u][1] + a[v]}
						que[head] = v
						head++
					} else if dist[v][0] == dist[u][0]+1 && dist[v][1] < dist[u][1]+a[v] {
						dist[v][1] = dist[u][1] + a[v]
					}
				}
			}
		}
		return dist
	}

	dist := make([][][]int, n)
	for u := 0; u < n; u++ {
		dist[u] = bfs(u)
	}

	ans := make([][]int, len(queries))
	for i, q := range queries {
		u, v := q[0]-1, q[1]-1
		if dist[u][v] != nil {
			ans[i] = dist[u][v]
		}
	}

	return ans
}
