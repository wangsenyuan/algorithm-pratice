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
	scanner.Scan()
	x := readInt(scanner.Bytes(), 0, &a)
	readInt(scanner.Bytes(), x+1, &b)
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	n, m := readTwoNums(scanner)
	grid := make([][]int, n)
	for i := 0; i < n; i++ {
		grid[i] = readNNums(scanner, m)
	}
	fmt.Println(solve(n, m, grid))
}

var dd = []int{-1, 0, 1, 0, -1}

func solve(n int, m int, grid [][]int) int {

	var dfs func(x int, y int, col0 int, col1 int, label int) int

	used := make([]bool, n*m)
	visited := make([]int, n*m)

	dfs = func(x int, y int, col0 int, col1 int, label int) int {
		visited[x*m+y] = label
		used[x*m+y] = true

		var cnt int
		for i := 0; i < 4; i++ {
			u, v := x+dd[i], y+dd[i+1]
			if u >= 0 && u < n && v >= 0 && v < m && visited[u*m+v] != label && (grid[u][v] == col0 || grid[u][v] == col1) {
				cnt += dfs(u, v, col0, col1, label)
			}
		}
		return cnt + 1
	}

	var best int
	label := 1
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			used[i*m+j] = true
			for k := 0; k < 4; k++ {
				x, y := i+dd[k], j+dd[k+1]
				if x >= 0 && x < n && y >= 0 && y < m && visited[x*m+y] != label && !used[x*m+y] {
					best = max(best, dfs(i, j, grid[i][j], grid[x][y], label))
				}
			}
			label++
		}
	}

	return best
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
