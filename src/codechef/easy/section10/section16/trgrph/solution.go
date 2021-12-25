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
		n, m := readTwoNums(scanner)
		edges := make([][]int, m)

		for i := 0; i < m; i++ {
			edges[i] = readNNums(scanner, 2)
		}

		res := solve(n, m, edges)

		if len(res) == 0 {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
			for i := 0; i < n; i++ {
				if i < n-1 {
					fmt.Printf("%d ", res[i])
				} else {
					fmt.Printf("%d\n", res[i])
				}
			}
		}

	}
}

func solve(n, m int, edges [][]int) []int {
	degree := make([]int, n)

	for _, edge := range edges {
		u, v := edge[0]-1, edge[1]-1
		degree[u]++
		degree[v]++
	}

	outs := make([][]int, n)

	for i := 0; i < n; i++ {
		outs[i] = make([]int, 0, degree[i])
	}

	for _, edge := range edges {
		u, v := edge[0]-1, edge[1]-1
		outs[u] = append(outs[u], v)
		outs[v] = append(outs[v], u)
	}

	var cmp = func(u, v int) bool {
		return degree[u] > degree[v]
	}

	for i := 0; i < n; i++ {
		conns := Conns{outs[i], cmp}
		sort.Sort(&conns)
	}

	cnt := make([]int, n)
	lvl := make([]int, n)
	vis := make([]bool, n)

	parent := make([]int, n)

	var dfs func(p, u int, level int)

	dfs = func(p, u, level int) {
		parent[u] = p + 1
		lvl[u] = level
		vis[u] = true

		for _, v := range outs[u] {
			if vis[v] {
				continue
			}
			dfs(u, v, level+1)
			cnt[u] += cnt[v]
		}
		cnt[u]++
	}

	r := -1

	for i := 0; i < n; i++ {
		if degree[i] == n-1 {
			r = i
			break
		}
	}

	if r < 0 {
		return nil
	}

	dfs(-1, r, 0)

	for i := 0; i < n; i++ {
		if cnt[i]-1 != degree[i]-lvl[i] || !vis[i] {
			return nil
		}
	}

	return parent
}

type Conns struct {
	conns []int
	fn    func(i, j int) bool
}

func (conns *Conns) Len() int {
	return len(conns.conns)
}

func (conns *Conns) Less(i, j int) bool {
	cc := conns.conns
	return conns.fn(cc[i], cc[j])
}

func (Conns *Conns) Swap(i, j int) {
	cc := Conns.conns
	cc[i], cc[j] = cc[j], cc[i]
}
