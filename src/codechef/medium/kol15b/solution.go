package main

import (
	"math"
	"bufio"
	"os"
	"fmt"
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
	t := readNum(scanner)

	for t > 0 {
		n, m := readTwoNums(scanner)
		grid := make([][]int, n)
		for i := 0; i < n; i++ {
			grid[i] = readNNums(scanner, m)
		}
		fmt.Println(solve(n, m, grid))
		t--
	}
}

func solve(n int, m int, grid [][]int) int {
	a, b, c, d := make([][]int, n), make([][]int, n), make([][]int, n), make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, m)
		b[i] = make([]int, m)
		c[i] = make([]int, m)
		d[i] = make([]int, m)
	}

	for j := 0; j < m; j++ {
		for i := 0; i < n; i++ {
			a[i][j] = grid[i][j]
			if j > 0 && a[i][j-1] < 0 {
				a[i][j] += a[i][j-1]
			}
		}
	}

	for j := m - 1; j >= 0; j-- {
		for i := 0; i < n; i++ {
			b[i][j] = grid[i][j]
			if j < m-1 && b[i][j+1] < 0 {
				b[i][j] += b[i][j+1]
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			c[i][j] = grid[i][j]
			if i > 0 && c[i-1][j] < 0 {
				c[i][j] += c[i-1][j]
			}
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := 0; j < m; j++ {
			d[i][j] = grid[i][j]
			if i < n-1 && d[i+1][j] < 0 {
				d[i][j] += d[i+1][j]
			}
		}
	}

	ans := math.MaxInt32

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			tmp := a[i][j] + b[i][j] + c[i][j] + d[i][j] - 3*grid[i][j]
			if tmp < ans {
				ans = tmp
			}
		}
	}

	return ans
}
