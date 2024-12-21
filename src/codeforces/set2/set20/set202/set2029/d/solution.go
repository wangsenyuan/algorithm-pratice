package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	var buf bytes.Buffer
	reader := bufio.NewReader(os.Stdin)
	tc := readNum(reader)
	for tc > 0 {
		tc--
		res, _, _ := process(reader)
		buf.WriteString(fmt.Sprintf("%d\n", len(res)))
		for _, cur := range res {
			buf.WriteString(fmt.Sprintf("%d %d %d\n", cur[0], cur[1], cur[2]))
		}
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

func process(reader *bufio.Reader) (res [][]int, n int, edges [][]int) {
	n, m := readTwoNums(reader)
	edges = make([][]int, m)
	for i := range m {
		edges[i] = readNNums(reader, 2)
	}
	res = solve(n, edges)
	return
}

func solve(n int, edges [][]int) [][]int {

	g := make([]map[int]bool, n)

	for i := 0; i < n; i++ {
		g[i] = make(map[int]bool)
	}

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g[u][v] = true
		g[v][u] = true
	}

	var res [][]int

	remove := func(u int, v int) {
		delete(g[u], v)
		delete(g[v], u)
	}

	add := func(u int, v int) {
		g[u][v] = true
		g[v][u] = true
	}

	for u := 0; u < n; u++ {
		for len(g[u]) > 1 {
			v, w := -1, -1
			for k := range g[u] {
				if v == -1 {
					v = k
				} else {
					w = k
					break
				}
			}
			// v, w must exist
			res = append(res, []int{u + 1, v + 1, w + 1})
			remove(u, v)
			remove(u, w)
			if g[v][w] {
				remove(v, w)
			} else {
				add(v, w)
			}
		}
	}

	u, v := -1, -1

	for i := 0; i < n; i++ {
		if len(g[i]) == 1 {
			u = i
			v = keys(g[i])[0]
			break
		}
	}

	if u < 0 {
		return res
	}

	marked := make([]bool, n)
	marked[u] = true
	marked[v] = true

	for x := 0; x < n; x++ {
		if marked[x] {
			continue
		}
		if len(g[x]) == 0 {
			res = append(res, []int{u + 1, v + 1, x + 1})
			v = x
		} else {
			y := keys(g[x])[0]
			res = append(res, []int{u + 1, x + 1, y + 1})
			marked[y] = true
		}
	}

	return res
}

func keys(m map[int]bool) []int {
	var res []int
	for k := range m {
		res = append(res, k)
	}
	return res
}
