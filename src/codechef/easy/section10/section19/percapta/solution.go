package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		A := readNNums(reader, n)
		B := readNNums(reader, n)
		E := make([][]int, m)
		for i := 0; i < m; i++ {
			E[i] = readNNums(reader, 2)
		}
		res := solve(n, A, B, E)
		fmt.Println(len(res))
		var buf bytes.Buffer

		for i := 0; i < len(res); i++ {
			buf.WriteString(strconv.Itoa(res[i]))
			buf.WriteByte(' ')
		}
		fmt.Println(buf.String())
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

func solve(n int, A []int, B []int, E [][]int) []int {
	a := int64(A[0])
	b := int64(B[0])

	for i := 1; i < n; i++ {
		x, y := int64(A[i]), int64(B[i])
		// x/y > a / b
		if x*b > a*y {
			a, b = x, y
		}
	}
	// we know the best per-capital value
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
	que := make([]int, n)
	color := make([]int, n)
	bfs := func(x int, flag int) int {
		var front, end int
		que[end] = x
		end++
		color[x] = flag

		for front < end {
			u := que[front]
			front++

			for _, v := range adj[u] {
				if color[v] == flag || int64(A[v])*b != int64(B[v])*a {
					continue
				}
				color[v] = flag
				que[end] = v
				end++
			}
		}
		return end
	}
	var best int
	var bestColor int
	cur := 1
	for i := 0; i < n; i++ {
		if color[i] == 0 && int64(A[i])*b == int64(B[i])*a {
			cnt := bfs(i, cur)
			if cnt > best {
				best = cnt
				bestColor = cur
			}
			cur++
		}
	}

	res := make([]int, 0, best)

	for i := 0; i < n; i++ {
		if color[i] == bestColor {
			res = append(res, i+1)
		}
	}
	return res
}
