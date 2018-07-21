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

	N, M := readTwoNums(scanner)
	edges := make([][]int, M)
	for i := 0; i < M; i++ {
		edges[i] = readNNums(scanner, 2)
	}
	res := solve(N, M, edges)

	if res {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func solve(N, M int, edges [][]int) bool {
	conn := make([][]int, N)
	in := make([][]int, N)
	for i := 0; i < N; i++ {
		conn[i] = make([]int, 0, 3)
		in[i] = make([]int, 0, 3)
	}

	for i, edge := range edges {
		u, v := edge[0]-1, edge[1]-1
		conn[u] = append(conn[u], v)
		conn[v] = append(conn[v], u)
		in[u] = append(in[u], i)
		in[v] = append(in[v], i)
	}

	disc := make([]int, N)
	low := make([]int, N)
	visited := make([]bool, N)
	bridge := make([]bool, M)

	var dfs1 func(u int, p int, time *int)

	dfs1 = func(u int, p int, time *int) {
		*time++
		disc[u] = *time
		low[u] = *time
		var cnt int
		for i := 0; i < len(conn[u]); i++ {
			v := conn[u][i]
			ii := in[u][i]
			if v == p {
				if cnt == 0 {
					cnt++
					continue
				}
			}
			if disc[v] > 0 {
				low[u] = min(low[u], disc[v])
			} else {
				dfs1(v, u, time)
				if low[v] > disc[u] {
					bridge[ii] = true
				}
				low[u] = min(low[u], low[v])
			}
		}
	}
	dfs1(0, 0, new(int))

	var cnt int
	comp := make([]int, N)
	var dfs2 func(u int)

	dfs2 = func(u int) {
		comp[u] = cnt
		visited[u] = true

		for i := 0; i < len(conn[u]); i++ {
			v := conn[u][i]
			ii := in[u][i]
			if visited[v] {
				continue
			}
			if bridge[ii] {
				continue
			}

			dfs2(v)
		}
	}

	for i := 0; i < N; i++ {
		if !visited[i] {
			dfs2(i)
			cnt++
		}
	}

	edgeMap := make([]map[int]int, cnt)
	for i := 0; i < cnt; i++ {
		edgeMap[i] = make(map[int]int)
	}

	deg := make([]int, cnt)

	for u := 0; u < N; u++ {
		for i := 0; i < len(conn[u]); i++ {
			v := conn[u][i]
			ii := in[u][i]
			if bridge[ii] {
				if edgeMap[comp[u]][comp[v]] == 0 {
					deg[comp[u]]++
					deg[comp[v]]++
					edgeMap[comp[u]][comp[v]]++
					edgeMap[comp[v]][comp[u]]++
				}
			}
		}
	}

	for i := 0; i < cnt; i++ {
		if deg[i] > 2 {
			return false
		}
	}

	return true
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
