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
		d := make([]int, m)
		for i := 0; i < m; i++ {
			d[i] = readNum(reader)
		}
		edges, res := solve(n, d)
		for _, e := range edges {
			buf.WriteString(fmt.Sprintf("%d %d\n", e[0], e[1]))
		}
		for _, x := range res {
			buf.WriteString(fmt.Sprintf("%d %d %d\n", x[0], x[1], x[2]))
		}
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

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
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

func solve(n int, d []int) (edges [][]int, ops [][]int) {
	// 一开始生成一条直线
	parent := make([]int, n)
	dist := make([]int, n+1)
	for i := 1; i < n; i++ {
		edges = append(edges, []int{i, i + 1})
		parent[i] = i - 1
		dist[i] = i
	}
	// 3个端点
	// 分割点在1
	ends := []int{n - 1, 1}

	ops = make([][]int, len(d))
	arr := make([]int, n)

	process1 := func(b int, c int, dx int) []int {
		u := b
		var top int
		for dist[parent[u]] > dx {
			arr[top] = u
			top++
			u = parent[u]
		}
		arr[top] = u
		top++
		v1 := parent[u]

		ends[0] = v1
		ends[1] = b

		v2 := c
		// dist[parent[u]] == dx
		parent[u] = c
		for i := top - 1; i >= 0; i-- {
			x := arr[i]
			dist[x] = dist[parent[x]] + 1
		}
		return []int{u + 1, v1 + 1, v2 + 1}
	}

	process2 := func(b int, c int, dx int) []int {
		// 从c那边移动一部分到b这里
		u := c
		var top int
		for dist[b]+top < dx {
			arr[top] = u
			top++
			u = parent[u]
		}

		u = arr[top-1]
		// dist[b] + top == dx
		v1 := parent[u]
		v2 := b
		parent[u] = b
		for i := top - 1; i >= 0; i-- {
			x := arr[i]
			dist[x] = dist[parent[x]] + 1
		}

		ends[0] = c
		ends[1] = v1

		return []int{u + 1, v1 + 1, v2 + 1}
	}

	process := func(it int, dx int) {
		b := max(ends[0], ends[1])
		c := min(ends[0], ends[1])

		if dist[b] == dx || dist[c] == dx || c != 1 && dist[b]+dist[c]-2 == dx {
			ops[it] = []int{-1, -1, -1}
			return
		}
		if dist[b] > dx {
			ops[it] = process1(b, c, dx)
			return
		}
		if dist[c] > dx {
			ops[it] = process1(c, b, dx)
			return
		}

		// dist[b] < dx && dist[c] < dx
		if dist[b] < dist[c] {
			b, c = c, b
		}
		ops[it] = process2(b, c, dx)
	}

	for i, x := range d {
		process(i, x)
	}

	return
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
