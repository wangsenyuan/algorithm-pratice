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

	line := readNNums(scanner, 3)
	n, m, k := line[0], line[1], line[2]
	A := readNNums(scanner, k)
	E := make([][]int, m)

	for i := 0; i < m; i++ {
		E[i] = readNNums(scanner, 2)
	}
	fmt.Println(solve(n, m, k, A, E))
}

func solve(n, m, k int, A []int, E [][]int) int {
	conns := make([][]int, n)

	for i := 0; i < n; i++ {
		conns[i] = make([]int, 0, 5)
	}

	for _, e := range E {
		u, v := e[0]-1, e[1]-1
		conns[u] = append(conns[u], v)
		conns[v] = append(conns[v], u)
	}

	var bfs func(dist []int, x int)

	que := make([]int, n)

	bfs = func(dist []int, x int) {
		for i := 0; i < n; i++ {
			dist[i] = -1
		}
		dist[x] = 0
		var front, end int
		que[end] = x
		end++
		for front < end {
			u := que[front]
			front++

			for _, v := range conns[u] {
				if dist[v] < 0 {
					dist[v] = dist[u] + 1
					que[end] = v
					end++
				}
			}
		}
	}
	dist := make([][]int, 2)
	dist[0] = make([]int, n)
	dist[1] = make([]int, n)

	bfs(dist[0], 0)
	bfs(dist[1], n-1)

	ps := make([]Pair, k)

	for i := 0; i < k; i++ {
		a := A[i] - 1
		d := dist[0][a] - dist[1][a]
		ps[i] = Pair{d, a}
	}

	sort.Sort(Pairs(ps))

	var best int
	var res = -1 << 30

	for i := 0; i < k; i++ {
		a := ps[i].second
		best = max(best, res+dist[1][a])
		res = max(res, dist[0][a])
	}

	return min(best+1, dist[0][n-1])
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Pair struct {
	first, second int
}

type Pairs []Pair

func (ps Pairs) Len() int {
	return len(ps)
}

func (ps Pairs) Less(i, j int) bool {
	return ps[i].first < ps[j].first
}

func (ps Pairs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
