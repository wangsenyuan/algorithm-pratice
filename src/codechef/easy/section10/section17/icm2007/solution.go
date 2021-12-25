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
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	res := readNNums(scanner, 2)
	a, b = res[0], res[1]
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
	return res
}

func fillNNums(scanner *bufio.Scanner, n int, res []int) {
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var tmp int64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readNInt64Nums(scanner *bufio.Scanner, n int) []int64 {
	res := make([]int64, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt64(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		n, m := readTwoNums(scanner)
		E := make([][]int, m)
		for i := 0; i < m; i++ {
			E[i] = readNNums(scanner, 3)
		}
		res := solve(n, m, E)

		if res {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func solve(n int, m int, E [][]int) bool {
	conns := make([][]int, n)

	for i := 0; i < n; i++ {
		conns[i] = make([]int, 0, 10)
	}

	connect := func(u, v, w int) {
		conns[u] = append(conns[u], v)
		conns[u] = append(conns[u], w)
	}

	for _, e := range E {
		u, v, w := e[0]-1, e[1]-1, e[2]
		connect(u, v, w)
		connect(v, u, -w)
	}
	vis := make([]bool, n)

	dist := make([]int, n)

	var dfs func(u int, d int) bool
	dfs = func(u, d int) bool {
		if vis[u] {
			return dist[u] == d
		}
		dist[u] = d
		vis[u] = true

		for i := 0; i < len(conns[u]); i += 2 {
			v := conns[u][i]
			w := conns[u][i+1]
			if !dfs(v, d+w) {
				return false
			}
		}
		return true
	}

	for i := 0; i < n; i++ {
		if !vis[i] && !dfs(i, 0) {
			return false
		}
	}
	return true
}
