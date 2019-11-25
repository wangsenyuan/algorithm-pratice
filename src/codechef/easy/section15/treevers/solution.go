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
		W := readNNums(scanner, n)
		edges := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			edges[i] = readNNums(scanner, 2)
		}
		fmt.Println(solve(n, W, edges))
	}
}

func solve(n int, W []int, edges [][]int) int {

	outs := make([][]int, n)

	for i := 0; i < n; i++ {
		outs[i] = make([]int, 0, 3)
	}

	for _, edge := range edges {
		u, v := edge[0]-1, edge[1]-1
		outs[u] = append(outs[u], v)
		outs[v] = append(outs[v], u)
	}

	ones := make([]int, n)
	zeros := make([]int, n)

	var dfs func(p int, u int)

	dfs = func(p, u int) {
		// we count 0
		if W[u] == 0 {
			zeros[u]++
		} else {
			ones[u]++
		}

		for _, v := range outs[u] {
			if p == v {
				continue
			}
			dfs(u, v)
			zeros[u] += zeros[v]
			ones[u] += ones[v]
		}
	}

	dfs(-1, 0)

	tour := make([]int, n)
	var index int

	var dfs2 func(p, u int)

	dfs2 = func(p, u int) {
		tour[index] = W[u]
		index++

		ps := make([]PP, 0, len(outs[u]))

		for _, v := range outs[u] {
			if p == v {
				continue
			}
			ps = append(ps, PP{ones[v], zeros[v], v})
		}

		sort.Sort(PS(ps))

		for i := 0; i < len(ps); i++ {
			dfs2(u, ps[i].third)
		}
	}

	dfs2(-1, 0)

	var res int
	var sum int
	for i := n - 1; i >= 0; i-- {
		if tour[i] == 0 {
			res += sum
		}
		sum += tour[i]
	}

	return res
}

type PP struct {
	first  int
	second int
	third  int
}

type PS []PP

func (ps PS) Len() int {
	return len(ps)
}

func (ps PS) Less(i, j int) bool {
	p := ps[i]
	q := ps[j]
	return p.first*q.second < p.second*q.first
}

func (ps PS) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
