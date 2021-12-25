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

func readNUint64s(scanner *bufio.Scanner, n int) []uint64 {
	res := make([]uint64, n)
	var pos int
	scanner.Scan()
	bs := scanner.Bytes()
	for i := 0; i < n; i++ {
		pos = readUint64(bs, pos, &res[i]) + 1
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	tc := readNum(scanner)

	for tc > 0 {
		tc--
		n := readNum(scanner)
		edges := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			edges[i] = readNNums(scanner, 2)
		}
		V := readNUint64s(scanner, n)
		M := readNUint64s(scanner, n)
		res := solve(n, edges, V, M)
		for i := 0; i < len(res); i++ {
			if i == len(res)-1 {
				fmt.Printf("%d\n", res[i])
			} else {
				fmt.Printf("%d ", res[i])
			}
		}
	}
}

func solve(n int, edges [][]int, V []uint64, M []uint64) []uint64 {
	out := make([][]int, n+1)

	for i := 1; i <= n; i++ {
		out[i] = make([]int, 0, 10)
	}

	for _, edge := range edges {
		a, b := edge[0], edge[1]
		out[a] = append(out[a], b)
		out[b] = append(out[b], a)
	}
	parent := make([]int, n+1)
	leaf := make([]bool, n+1)

	var dfs func(p, u int)
	dfs = func(p, u int) {
		leaf[u] = true
		parent[u] = p
		for _, v := range out[u] {
			if p == v {
				continue
			}
			leaf[u] = false
			dfs(u, v)
		}
	}

	dfs(0, 1)

	process := func(x int) uint64 {
		m := M[x-1]
		g := V[x-1]
		var cur = parent[x]
		for cur > 0 {
			g = gcd(g, V[cur-1])
			cur = parent[cur]
		}
		return m - gcd(g, m)
	}

	var res []uint64

	for i := 1; i <= n; i++ {
		if leaf[i] {
			r := process(i)
			res = append(res, r)
		}
	}

	return res
}

func gcd(a, b uint64) uint64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
