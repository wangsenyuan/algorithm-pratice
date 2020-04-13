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

func readThreeNums(scanner *bufio.Scanner) (a int, b int, c int) {
	res := readNNums(scanner, 3)
	a, b, c = res[0], res[1], res[2]
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	n, m := readTwoNums(scanner)
	segs := make([][]int, m)

	for i := 0; i < m; i++ {
		segs[i] = readNNums(scanner, 2)
	}

	fmt.Println(solve(n, m, segs))
}

const MOD = 1000000007

func solve(n, m int, segs [][]int) int {
	edges := make([][]int, n+1)

	for i := 0; i <= n; i++ {
		edges[i] = make([]int, 0, 3)
	}

	for _, seg := range segs {
		a, b := seg[0], seg[1]
		a--
		edges[a] = append(edges[a], b)
		edges[b] = append(edges[b], a)
	}

	vis := make([]bool, n+1)

	var dfs func(u int) int

	dfs = func(u int) int {
		var res = 0
		vis[u] = true

		for _, v := range edges[u] {
			if vis[v] {
				continue
			}
			res++
			res += dfs(v)
		}

		return res
	}

	var count int
	for i := 0; i < n; i++ {
		count += dfs(i)
	}

	return pow(2, count)
}

func pow(a, b int) int {
	A := int64(a)
	R := int64(1)

	for b > 0 {
		if b&1 == 1 {
			R *= A
			R %= MOD
		}
		A *= A
		A %= MOD
		b >>= 1
	}

	return int(R)
}
