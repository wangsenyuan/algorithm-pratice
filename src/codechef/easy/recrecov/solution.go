package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	for i < len(bytes) && bytes[i] == ' ' {
		i++
	}
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readOneNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	scanner.Scan()
	x := readInt(scanner.Bytes(), 0, &a)
	readInt(scanner.Bytes(), x+1, &b)
	return
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	t := readOneNum(scanner)

	for i := 0; i < t; i++ {
		n, m := readTwoNums(scanner)
		pairs := make([][]int, m)
		for j := 0; j < m; j++ {
			a, b := readTwoNums(scanner)
			pairs[j] = []int{a, b}
		}
		fmt.Println(solve(n, m, pairs))
	}
}

func solve(n int, m int, pairs [][]int) int {
	out := make([][]int, n)
	for i := 0; i < m; i++ {
		a, b := pairs[i][0]-1, pairs[i][1]-1
		if out[a] == nil {
			out[a] = make([]int, 0, 10)
		}

		out[a] = append(out[a], b)
	}

	var fn func(v int) bool
	visited := make([]bool, n)
	p := make([]int, n)

	for i := 0; i < n; i++ {
		p[i] = -1
	}

	fn = func(v int) bool {
		if out[v] == nil {
			return false
		}
		for _, w := range out[v] {
			if visited[w] {
				continue
			}
			visited[w] = true
			if p[w] == -1 || fn(p[w]) {
				p[w] = v
				return true
			}
		}
		return false
	}
	var res int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			visited[j] = false
		}
		if fn(i) {
			res++
		}
	}
	return n - res
}

func solve1(n int, m int, pairs [][]int) int {
	s := 2 * n
	t := s + 1

	graph := make([][]int, t+1)
	for i := 0; i <= t; i++ {
		graph[i] = make([]int, t+1)
	}

	for i := 0; i < n; i++ {
		// 2 * a is the out, 2 * a + 1 is the in
		// from s to a-out
		graph[s][2*i] = 1
		// from a-in to t
		graph[2*i+1][t] = 1
	}

	for i := 0; i < m; i++ {
		a, b := pairs[i][0]-1, pairs[i][1]-1
		// from a-out to b-in
		graph[2*a][2*b+1] = 1
	}

	que := make([]int, t+1)
	parent := make([]int, t+1)
	dist := make([]int, t+1)
	bfs := func(source int, sink int) bool {
		for i := 0; i <= t; i++ {
			dist[i] = -1
		}
		head, tail := 0, 0
		que[tail] = source
		tail++
		parent[source] = -1
		dist[source] = 0
		for head < tail {
			tmp := tail
			for head < tmp {
				v := que[head]
				head++
				if v == sink {
					return true
				}

				for i := 0; i <= t; i++ {
					if graph[v][i] > 0 && dist[i] == -1 {
						dist[i] = dist[v] + 1
						parent[i] = v
						que[tail] = i
						tail++
					}
				}
			}
		}

		return false
	}
	var res int
	for bfs(s, t) {
		minPath := math.MaxInt32

		for x := t; x != s; x = parent[x] {
			tmp := graph[parent[x]][x]
			if tmp < minPath {
				minPath = tmp
			}
		}
		res += minPath
		for x := t; x != s; x = parent[x] {
			px := parent[x]
			graph[px][x] -= minPath
			graph[x][px] += minPath
		}
	}

	return n - res
}
