package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
		n := readNum(scanner)
		A := readNNums(scanner, n)
		edges := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			edges[i] = readNNums(scanner, 2)
		}
		res := solve(n, A, edges)
		fmt.Println(res)
	}
}

const MOD = 1000000007

func solve(n int, A []int, edges [][]int) int {
	leaf := n

	degree := make([]int, n)
	var root int
	for _, edge := range edges {
		u, v := edge[0]-1, edge[1]-1
		degree[u]++

		if degree[u] == 2 {
			root = u
			leaf--
		}

		degree[v]++

		if degree[v] == 2 {
			root = v
			leaf--
		}
	}

	conns := make([][]int, n)

	for i := 0; i < n; i++ {
		conns[i] = make([]int, 0, degree[i])
	}

	for _, edge := range edges {
		u, v := edge[0]-1, edge[1]-1
		conns[u] = append(conns[u], v)
		conns[v] = append(conns[v], u)
	}

	L := make([]int, n)

	B := make([]int, n)

	var dfs func(p, u int)

	dfs = func(p, u int) {
		if degree[u] == 1 {
			L[u] = 1
		}
		var s int
		for _, v := range conns[u] {
			if p == v {
				continue
			}
			dfs(u, v)
			L[u] += L[v]
			s += L[v]
		}
		var val int

		for _, v := range conns[u] {
			if p == v {
				continue
			}
			val += (s - L[v]) * L[v]
		}
		val /= 2
		val += L[u] * (leaf - L[u])
		val %= MOD
		B[u] = val
	}

	dfs(-1, root)

	sort.Ints(A)
	sort.Ints(B)

	var res int64

	for i := 0; i < n; i++ {
		tmp := int64(A[i]) * int64(B[i])
		tmp %= MOD
		res += tmp
		res %= MOD
	}
	return int(res)
}

func solve1(n int, A []int, edges [][]int) int {
	sort.Ints(A)

	leaf := n

	degree := make([]int, n)

	for _, edge := range edges {
		u, v := edge[0]-1, edge[1]-1
		degree[u]++

		if degree[u] == 2 {
			leaf--
		}

		degree[v]++

		if degree[v] == 2 {
			leaf--
		}
	}

	conns := make([][]int, n)

	for i := 0; i < n; i++ {
		conns[i] = make([]int, 0, degree[i])
	}

	for _, edge := range edges {
		u, v := edge[0]-1, edge[1]-1
		conns[u] = append(conns[u], v)
		conns[v] = append(conns[v], u)
	}

	que := make([]int, n)

	var front, end int
	var p int
	B := make([]int, n)

	for i := 0; i < n; i++ {
		if degree[i] == 1 {
			que[end] = i
			end++
			B[i] = A[p]
			p++
		}
	}

	for front < end {
		u := que[front]
		front++

		for _, v := range conns[u] {
			degree[v]--
			// p++
			if degree[v] == 1 {
				B[v] = A[p]
				p++
				que[end] = v
				end++
			}
		}
	}

	var dfs func(p, u int)
	var res int64
	child := make([]int, n)

	dfs = func(p, u int) {
		cnt := leaf

		for _, v := range conns[u] {
			if p == v {
				continue
			}
			dfs(u, v)
			child[u] += child[v]
			x := int64(child[v]) * int64(cnt-child[v])
			x %= MOD
			x *= int64(B[u])
			x %= MOD
			res += x
			res %= MOD
			cnt -= child[v]
		}

		if cnt == leaf {
			// a real leaf
			child[u] = 1
			x := int64(cnt - 1)
			x %= MOD
			x *= int64(B[u])
			x %= MOD
			res += x
			res %= MOD
		}
	}

	dfs(-1, 0)

	return int(res)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
