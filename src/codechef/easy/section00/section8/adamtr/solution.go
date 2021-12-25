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
		tc--
		n := readNum(scanner)
		A := make([][]int, n)
		for i := 0; i < n; i++ {
			A[i] = readNNums(scanner, n)
		}
		B := make([][]int, n)
		for i := 0; i < n; i++ {
			B[i] = readNNums(scanner, n)
		}
		res := solve(n, A, B)
		if res {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}

func solve(n int, A [][]int, B [][]int) bool {
	grid := make([][]int, n)

	for i := 0; i < n; i++ {
		grid[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if A[i][j] == B[i][j] && A[j][i] == B[j][i] && A[i][j] == A[j][i] {
				//no edge between i, j
				continue
			}

			if A[i][j] == B[i][j] && A[j][i] == B[j][i] {
				//same color
				grid[i][j] = 1
				grid[j][i] = 1
				continue
			}

			if A[i][j] == B[j][i] && A[j][i] == B[i][j] {
				//different color
				grid[i][j] = -1
				grid[j][i] = -1
				continue
			}
			//not able to get B
			return false
		}
	}

	vis := make([]bool, n)
	colors := make([]int, n)
	var dfs func(u int, color int) bool

	dfs = func(u int, color int) bool {
		if vis[u] {
			return colors[u] == color
		}
		colors[u] = color
		vis[u] = true

		for v := 0; v < n; v++ {
			if grid[u][v] != 0 {
				res := dfs(v, grid[u][v]*color)
				if !res {
					return false
				}
			}
		}
		return true
	}

	for i := 0; i < n; i++ {
		if !vis[i] {
			if !dfs(i, 1) {
				return false
			}
		}
	}
	return true
}
