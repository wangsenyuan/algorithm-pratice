package main

import (
	"bufio"
	"bytes"
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
		mat := make([][]int, n)

		for i := 0; i < n; i++ {
			mat[i] = readNNums(scanner, m)
		}
		can := solve(n, m, mat)
		if !can {
			fmt.Println(-1)
		} else {
			output(n, m, mat)
		}
		t--
	}
}

func output(n, m int, mat [][]int) {
	var buf bytes.Buffer

	for i := 0; i < n; i++ {
		buf.WriteByte(byte(mat[i][0] + '0'))
		for j := 1; j < m; j++ {
			buf.WriteByte(' ')
			buf.WriteByte(byte(mat[i][j] + '0'))
		}
		buf.WriteByte('\n')
	}

	fmt.Print(buf.String())
}

func solve(n, m int, mat [][]int) bool {
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if mat[i][j] < 0 {
				up, left := 1, 1
				if i > 0 {
					up = mat[i-1][j]
				}
				if j > 0 {
					left = mat[i][j-1]
				}
				cand := max(up, left)
				mat[i][j] = cand
			}

			if i > 0 && mat[i][j] < mat[i-1][j] {
				return false
			}

			if j > 0 && mat[i][j] < mat[i][j-1] {
				return false
			}
		}
	}

	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
