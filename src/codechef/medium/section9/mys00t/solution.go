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

	ask := func(x int) int {
		fmt.Printf("1 %d\n", x+1)

		y := readNum(reader)

		if y == -1 {
			fmt.Printf("2 %d\n", x+1)
			reader.ReadLine()
			return -1
		}

		return y - 1
	}

	tc := readNum(reader)
	for tc > 0 {
		tc--
		N := readNum(reader)
		E := make([][]int, N-1)
		for i := 0; i < N-1; i++ {
			E[i] = readNNums(reader, 2)
		}
		solve(N, E, ask)
	}
}

func solve(N int, E [][]int, ask func(int) int) {

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

	sz := make([]int, N)
	block := make([]bool, N)

	var dfs func(p, u int)

	dfs = func(p, u int) {
		sz[u]++
		for _, v := range conns[u] {
			if v == p || block[v] {
				continue
			}
			dfs(u, v)
			sz[u] += sz[v]
		}
	}

	var get func(p, u, n int) int

	get = func(p, u, n int) int {
		for _, v := range conns[u] {
			if p == v || block[v] {
				continue
			}
			if 2*sz[v] > n {
				return get(u, v, n)
			}
		}
		return u
	}

	findCentroid := func(p, u int) int {
		for i := 0; i < N; i++ {
			sz[i] = 0
		}

		dfs(p, u)

		return get(p, u, sz[u])
	}

	x := findCentroid(-1, 0)

	for {
		y := ask(x)

		if y < 0 {
			break
		}
		block[x] = true
		x = findCentroid(x, y)
	}
}
