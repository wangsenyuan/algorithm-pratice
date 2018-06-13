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

func readNum(scanner *bufio.Scanner) int {
	var x int
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &x)
	return x
}

func readTripNums(scanner *bufio.Scanner) (a int, b int, c int) {
	scanner.Scan()
	bs := scanner.Bytes()
	pos := readInt(bs, 0, &a)
	pos = readInt(bs, pos+1, &b)
	readInt(bs, pos+1, &c)
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	scanner.Scan()
	bs := scanner.Bytes()
	x := -1
	for i := 0; i < n; i++ {
		x = readInt(bs, x+1, &res[i])
	}

	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	t := readNum(scanner)

	for i := 0; i < t; i++ {
		m, n, k := readTripNums(scanner)
		grid := make([][]int, m)
		for j := 0; j < m; j++ {
			grid[j] = readNNums(scanner, n)
		}
		cnts := solve(m, n, k, grid)
		for a := 0; a < m-k+1; a++ {
			for b := 0; b < n-k+1; b++ {
				fmt.Printf("%d", grid[a][b])
				if cnts[a][b] > 1 {
					fmt.Printf("(%d)", cnts[a][b])
				}
				if b < n-k {
					fmt.Print(" ")
				} else {
					fmt.Println()
				}
			}
		}
		if i < t-1 {
			scanner.Scan()
		}
	}
}

func solve(m, n, k int, grid [][]int) [][]int {
	rows := make([][][]int, m)
	for i := 0; i < m; i++ {
		rows[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			rows[i][j] = make([]int, 2)
		}
	}
	// rows[i][j][0] = max element from (i, j) to (i, j + k - 1)
	// rows[i][j][0] = count of the max elements
	stack := make([][]int, n)
	for i := 0; i < n; i++ {
		stack[i] = make([]int, 2)
	}
	for i := 0; i < m; i++ {
		head, tail := 0, 0
		for j := 0; j < n; j++ {
			for tail > 0 && stack[tail-1][0] < grid[i][j] {
				tail--
			}
			if tail > 0 && stack[tail-1][0] == grid[i][j] {
				stack[tail-1][1]++
			} else {
				stack[tail][0] = grid[i][j]
				stack[tail][1] = 1
				tail++
			}
			if j+1 >= k {
				rows[i][j+1-k][0] = stack[head][0]
				rows[i][j+1-k][1] = stack[head][1]
				if grid[i][j+1-k] == stack[head][0] {
					stack[head][1]--
					if stack[head][1] == 0 {
						head++
					}
				}
			}
		}
	}
	stack = make([][]int, m)
	for i := 0; i < m; i++ {
		stack[i] = make([]int, 2)
	}

	cnts := make([][]int, m)
	for i := 0; i < m; i++ {
		cnts[i] = make([]int, n)
	}

	for j := 0; j < n-k+1; j++ {
		head, tail := 0, 0
		for i := 0; i < m; i++ {
			for tail > 0 && stack[tail-1][0] < rows[i][j][0] {
				tail--
			}
			if tail > 0 && stack[tail-1][0] == rows[i][j][0] {
				stack[tail-1][1] += rows[i][j][1]
			} else {
				stack[tail][0] = rows[i][j][0]
				stack[tail][1] = rows[i][j][1]
				tail++
			}
			if i+1 >= k {
				grid[i+1-k][j] = stack[head][0]
				cnts[i+1-k][j] = stack[head][1]
				if rows[i+1-k][j][0] == stack[head][0] {
					stack[head][1] -= rows[i+1-k][j][1]
					if stack[head][1] == 0 {
						head++
					}
				}
			}
		}
	}
	return cnts
}

func solve1(m, n, k int, grid [][]int) [][]int {
	cnts := make([][]int, m)
	for i := 0; i < m; i++ {
		cnts[i] = make([]int, n)
	}
	for j := 0; j < n; j++ {
		for i := 0; i < m-k+1; i++ {
			max, cnt := 0, 0
			for a := i; a < i+k; a++ {
				if grid[a][j] > max {
					max, cnt = grid[a][j], 1
				} else if grid[a][j] == max {
					cnt++
				}
			}
			grid[i][j] = max
			cnts[i][j] = cnt
		}
	}

	for i := 0; i < m-k+1; i++ {
		for j := 0; j < n-k+1; j++ {
			max, cnt := grid[i][j], cnts[i][j]
			for a := j + 1; a < j+k; a++ {
				if grid[i][a] > max {
					max, cnt = grid[i][a], cnts[i][a]
				} else if grid[i][a] == max {
					cnt += cnts[i][a]
				}
			}
			grid[i][j] = max
			cnts[i][j] = cnt
		}
	}

	return cnts
}
