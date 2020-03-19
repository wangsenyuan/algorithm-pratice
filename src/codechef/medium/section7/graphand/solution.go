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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	n, m := readTwoNums(scanner)
	E := make([][]int, m)
	for i := 0; i < m; i++ {
		E[i] = readNNums(scanner, 4)
	}
	line := readNNums(scanner, 3)
	fmt.Println(solve(n, m, E, line[0], line[1], line[2]))
}

const INF = int64(1) << 60

func solve(n int, m int, E [][]int, s, t, k int) int64 {
	s--
	t--
	edges := make([][]Edge, n)
	for i := 0; i < n; i++ {
		edges[i] = make([]Edge, 0, 3)
	}

	for _, e := range E {
		x, y, c, v := e[0]-1, e[1]-1, int64(e[2]), uint64(e[3])
		edges[x] = append(edges[x], Edge{y, v, c})
		edges[y] = append(edges[y], Edge{x, v, c})
	}

	dist := make([]int64, n)
	que := make([]int, n)
	bfs := func(mask uint64) int64 {
		for i := 0; i < n; i++ {
			dist[i] = INF
		}
		var front, end int
		que[end] = s
		end++
		dist[s] = 0

		for front < end {
			x := que[front]
			front++
			if x == t {
				return dist[x]
			}

			for _, e := range edges[x] {
				y, v, c := e.y, e.v, e.c

				if v&mask == mask {
					if dist[x]+c <= dist[y] {
						dist[y] = dist[x] + c
						que[end] = y
						end++
					}
				}

			}
		}

		return INF
	}
	var mask uint64 = 1
	K := uint64(k)
	ans := INF
	for K <= 1000000000 {
		ans = min(ans, bfs(K))
		if K&mask == mask {
			K += mask
		}
		mask <<= 1
	}

	if ans == INF {
		return -1
	}
	return ans
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}

type Edge struct {
	y int
	v uint64
	c int64
}
