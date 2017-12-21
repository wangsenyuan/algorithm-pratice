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
		scanner.Scan()
	}
}

func solve(m, n, k int, grid [][]int) [][]int {
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
