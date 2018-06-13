package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
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

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwo(scanner *bufio.Scanner) (a int, b int) {
	scanner.Scan()
	x := readInt(scanner.Bytes(), 0, &a)
	readInt(scanner.Bytes(), x+1, &b)
	return
}

func readThree(scanner *bufio.Scanner) (a int, b int, c int) {
	scanner.Scan()
	x := readInt(scanner.Bytes(), 0, &a)
	x = readInt(scanner.Bytes(), x+1, &b)
	readInt(scanner.Bytes(), x+1, &c)
	return
}

func readFour(scanner *bufio.Scanner) (a int, b int, c int, d int) {
	scanner.Scan()
	x := readInt(scanner.Bytes(), 0, &a)
	x = readInt(scanner.Bytes(), x+1, &b)
	x = readInt(scanner.Bytes(), x+1, &c)
	x = readInt(scanner.Bytes(), x+1, &d)
	return
}

func readNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := -1
	scanner.Scan()
	bs := scanner.Bytes()
	for i := 0; i < n; i++ {
		x = readInt(bs, x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	t := readNum(scanner)

	for i := 0; i < t; i++ {
		n, m := readTwo(scanner)
		A := make([][]int, n)
		for j := 0; j < n; j++ {
			A[j] = readNums(scanner, m)
		}
		fmt.Println(solve(n, m, A))
	}
}

func solve(n, m int, A [][]int) int {
	var max int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if max < A[i][j] {
				max = A[i][j]
			}
		}
	}
	rows := make([][]int, max+1)
	cols := make([][]int, max+1)

	for i := 0; i <= max; i++ {
		rows[i] = make([]int, n)
		cols[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			a := A[i][j]
			rows[a][i]++
			cols[a][j]++
		}
	}
	maxRow := make([]map[int]bool, max+1)
	maxCol := make([]map[int]bool, max+1)

	for i := 1; i <= max; i++ {
		maxRow[i] = make(map[int]bool)
		maxCol[i] = make(map[int]bool)
	}

	bestRow := make([]int, max+1)
	bestCol := make([]int, max+1)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			a := A[i][j]
			if rows[a][i] > rows[a][bestRow[a]] {
				maxRow[a] = make(map[int]bool)
				maxRow[a][i] = true
				bestRow[a] = i
			} else if rows[a][i] == rows[a][bestRow[a]] {
				maxRow[a][i] = true
			}

			if cols[a][j] > cols[a][bestCol[a]] {
				maxCol[a] = make(map[int]bool)
				maxCol[a][j] = true
				bestCol[a] = j
			} else if cols[a][j] == cols[a][bestCol[a]] {
				maxCol[a][j] = true
			}
		}
	}

	var ans int
	for x := 1; x <= max; x++ {
		for i := range maxRow[x] {
			a := rows[x][i]
			for j := range maxCol[x] {
				b := cols[x][j]
				c := a + b
				if A[i][j] == x {
					c--
				}
				if c > ans {
					ans = c
				}
			}
		}
	}

	return ans
}
