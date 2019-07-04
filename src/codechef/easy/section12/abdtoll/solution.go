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
	sign := int64(1)
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
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
		firstLine := readNInt64Nums(scanner, 3)
		n := int(firstLine[0])
		x := int(firstLine[1])
		k := firstLine[2]
		tl := readNNums(scanner, n)
		edges := make([][]int, n-1)

		for i := 0; i < n-1; i++ {
			edges[i] = readNNums(scanner, 2)
		}
		fmt.Println(solve(n, x, k, tl, edges))
	}
}

func solve(n, x int, k int64, tl []int, edges [][]int) int64 {
	outs := make([][]int, n)
	for i := 0; i < n; i++ {
		outs[i] = make([]int, 0, 3)
	}

	for _, edge := range edges {
		a, b := edge[0]-1, edge[1]-1
		outs[a] = append(outs[a], b)
		outs[b] = append(outs[b], a)
	}

	X := make([]int64, n)

	var dfs func(p int, u int) int64

	dfs = func(p int, u int) int64 {
		var b int64

		for _, v := range outs[u] {
			if p == v {
				continue
			}
			b = max(b, dfs(u, v))
		}

		b += int64(tl[u])

		X[u] = b

		return b
	}

	dfs(-1, x-1)

	V := make([]int64, n)

	var dfs2 func(p int, u int)

	dfs2 = func(p int, u int) {
		if X[u] > k {
			V[u] = min(int64(tl[u]), X[u]-k)
		}
		for _, v := range outs[u] {
			if p == v {
				continue
			}
			dfs2(u, v)
		}
	}

	dfs2(-1, x-1)

	var res int64

	for _, v := range V {
		res += v
	}

	return res * 2
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}
