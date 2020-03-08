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

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	sign := int64(1)
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := int64(0)
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		N, M := readTwoNums(scanner)

		E := make([][]int, M)

		for i := 0; i < M; i++ {
			E[i] = readNNums(scanner, 2)
		}

		fmt.Println(solve(N, M, E))
	}
}

func solve(N int, M int, E [][]int) int {
	conns := make([][]int, N)
	for i := 0; i < N; i++ {
		conns[i] = make([]int, 0, 5)
	}

	for _, e := range E {
		u, v := e[0]-1, e[1]-1
		conns[u] = append(conns[u], v)
		conns[v] = append(conns[v], u)
	}

	vis := make([]int, N)
	depth := make([]int, N)
	hasPotential := make([]bool, N)
	sum := make([]int, N)
	var dfs func(v, p, d int) Pair

	var total int
	cnt := make([]int, N)

	dfs = func(v, p, d int) Pair {
		ret := NewPair(INF, INF)
		hasPotential[v] = true
		depth[v] = d
		vis[v] = 1
		for _, u := range conns[v] {
			if p == u {
				continue
			}
			if vis[u] == 0 {
				x := dfs(u, v, d+1)
				sum[v] += sum[u]
				ret.Update(x)
				if x[1] < d {
					hasPotential[v] = false
				}
			} else if vis[u] == 1 {
				// a back edge
				total++
				cnt[u]++
				cnt[v]++

				sum[p]++
				sum[u]--

				ret.UpdateNum(depth[u])
			}
		}

		vis[v] = 2
		return ret
	}

	for i := 0; i < N; i++ {
		if vis[i] == 0 {
			dfs(i, -1, 0)
		}
	}

	if total == 0 {
		return -1
	}

	for i := 0; i < N; i++ {
		if hasPotential[i] && cnt[i]+sum[i] == total {
			return i + 1
		}
	}
	return -1
}

type Pair []int

const INF = 1 << 30

func NewPair(a, b int) Pair {
	return Pair([]int{a, b})
}
func (this Pair) UpdateNum(num int) {
	if this[1] > num {
		this[1] = num
	}
	if this[0] > this[1] {
		this[1], this[0] = this[0], this[1]
	}
}

func (this Pair) Update(that Pair) {
	this.UpdateNum(that[1])
	this.UpdateNum(that[0])
}
