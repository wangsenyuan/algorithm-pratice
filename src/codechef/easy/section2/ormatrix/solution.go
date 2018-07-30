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

	tc := readNum(scanner)

	for tc > 0 {
		n, m := readTwoNums(scanner)
		A := make([][]byte, n)
		for i := 0; i < n; i++ {
			scanner.Scan()
			A[i] = scanner.Bytes()
		}
		res := solve(n, m, A)
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				fmt.Printf("%d ", res[i][j])
			}
			fmt.Println()
		}
		tc--
	}
}

const MAX_N = 1000

var B [][]int

func init() {
	B = make([][]int, MAX_N)
	for i := 0; i < MAX_N; i++ {
		B[i] = make([]int, MAX_N)
	}
}

func solve(N, M int, A [][]byte) [][]int {
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			B[i][j] = -1
		}
	}

	var ones int
	row := make([]int, N)
	col := make([]int, M)
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if A[i][j] == '1' {
				B[i][j] = 0
				ones++
				row[i]++
				col[j]++
			}
		}
	}

	if ones == 0 {
		// all -1
		return B
	}

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if row[i] > 0 || col[j] > 0 {
				if A[i][j] == '1' {
					B[i][j] = 0
				} else {
					B[i][j] = 1
				}
			} else {
				B[i][j] = 2
			}
		}
	}

	return B
}
