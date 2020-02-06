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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		n, m := readTwoNums(scanner)
		grid := make([][]int, n)
		for i := 0; i < n; i++ {
			grid[i] = readNNums(scanner, m)
		}
		fmt.Println(solve(n, m, grid))
	}
}

var dd = []int{-1, 0, 1, 0, -1}

func solve(n, m int, grid [][]int) int {

	L := n * m

	que := make([]int, 2*L)
	dist := make([]int, L)

	bfs := func(x, y int) int {
		for i := 0; i < L; i++ {
			dist[i] = -1
		}
		dist[x*m+y] = 0
		front, end := L, L
		que[end] = x*m + y
		end++

		var res int

		for front < end {
			cur := que[front]
			res = max(res, dist[cur])
			front++
			u, v := cur/m, cur%m

			for i := 0; i < 4; i++ {
				a, b := u+dd[i], v+dd[i+1]
				if a >= 0 && a < n && b >= 0 && b < m && dist[a*m+b] < 0 {
					if grid[a][b] == grid[u][v] {
						// same
						dist[a*m+b] = dist[cur]
						front--
						que[front] = a*m + b
					} else {
						dist[a*m+b] = dist[cur] + 1
						que[end] = a*m + b
						end++
					}
				}
			}
		}
		return res
	}

	ans := L

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			ans = min(ans, bfs(i, j))
		}
	}

	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
