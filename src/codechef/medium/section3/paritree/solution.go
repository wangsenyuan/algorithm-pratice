package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		n, m := readTwoNums(scanner)
		for i := 0; i < n-1; i++ {
			scanner.Scan()
		}
		conditions := make([][]int, m)
		for i := 0; i < m; i++ {
			conditions[i] = readNNums(scanner, 3)
		}
		res := solve(n, nil, m, conditions)
		fmt.Println(res)
		tc--
	}
}

const MOD = 1000000007

func solve(n int, edges [][]int, m int, conditions [][]int) int {
	conn := make([]map[int]int, n)
	for i := 0; i < n; i++ {
		conn[i] = make(map[int]int)
	}

	for _, condition := range conditions {
		u, v, x := condition[0]-1, condition[1]-1, condition[2]
		conn[u][v] = x
		conn[v][u] = x
	}

	var dfs func(u, c int) bool

	vis := make([]bool, n)
	color := make([]int, n)
	dfs = func(u, c int) bool {
		if vis[u] {
			return color[u] == c
		}
		vis[u] = true
		color[u] = c
		for v, x := range conn[u] {
			var tmp bool
			if x == 0 {
				tmp = dfs(v, c)
			} else {
				tmp = dfs(v, 1-c)
			}
			if !tmp {
				return false
			}
		}

		return true
	}

	var c int
	for i := 0; i < n; i++ {
		if !vis[i] {
			c++
			if !dfs(i, 0) {
				return 0
			}
		}
	}

	return int(pow(2, c-1))
}

func solve1(n int, edges [][]int, m int, conditions [][]int) int {
	uf := CreateUF(n)

	for i := 0; i < m; i++ {
		u, v, x := conditions[i][0]-1, conditions[i][1]-1, conditions[i][2]
		if !uf.Union(u, v, x) {
			return 0
		}
	}

	return int(pow(2, uf.size-1))
}

func pow(a, n int) int64 {
	x := int64(a)
	y := int64(1)
	for n > 0 {
		if n&1 == 1 {
			y = (y * x) % MOD
		}
		x = (x * x) % MOD
		n >>= 1
	}

	return y
}

type UF struct {
	set   []int
	d     []int
	count []int
	size  int
}

func CreateUF(n int) UF {
	set := make([]int, n)
	d := make([]int, n)
	count := make([]int, n)
	size := n

	for i := 0; i < n; i++ {
		set[i] = i
		count[i] = 1
	}

	return UF{set, d, count, size}
}

func (uf *UF) Union(u, v, x int) bool {
	set := uf.set
	d := uf.d
	count := uf.count

	var find func(u int) int
	find = func(u int) int {
		if set[u] == u {
			return u
		}
		p := set[u]
		set[u] = find(p)
		d[u] = d[u] ^ d[p]
		return set[u]
	}

	pu := find(u)
	pv := find(v)
	if pu == pv {
		return x == d[u]^d[v]
	}

	if count[pu] > count[pv] {
		set[pv] = pu
		count[pu] += count[pv]
		d[pv] = d[u] ^ d[v] ^ x
	} else {
		set[pu] = pv
		count[pv] += count[pu]
		d[pu] = d[u] ^ d[v] ^ x
	}
	uf.size--
	return true
}
