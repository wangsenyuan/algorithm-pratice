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

	A := readNNums(scanner, n)

	queries := make([][]int, m)

	for i := 0; i < m; i++ {
		queries[i] = readNNums(scanner, 2)
	}

	res := solve(n, A, queries)

	for i := 0; i < m; i++ {
		fmt.Println(res[i])
	}
}

func solve(n int, A []int, queries [][]int) []int {
	// f[x][i][0] = bit column x, before i (including) has how many 0 digits
	// f[x][i][1] = bit column x, before i (including) has how many 1 digits
	f := make([][][]int, 31)
	for x := 0; x < 31; x++ {
		f[x] = make([][]int, n)
		for i := 0; i < n; i++ {
			f[x][i] = make([]int, 2)
		}
	}

	for i := 0; i < n; i++ {
		a := A[i]
		for x := 0; x < 31; x++ {
			if i > 0 {
				f[x][i][0] = f[x][i-1][0]
				f[x][i][1] = f[x][i-1][1]
			}
			d := (a >> uint(x)) & 1
			f[x][i][d]++
		}
	}

	m := len(queries)
	res := make([]int, m)

	for j := 0; j < m; j++ {
		left, right := queries[j][0]-1, queries[j][1]-1
		var tmp int
		for x := 0; x < 31; x++ {
			ones := f[x][right][1]
			zeros := f[x][right][0]
			if left > 0 {
				ones -= f[x][left-1][1]
				zeros -= f[x][left-1][0]
			}
			if ones < zeros {
				tmp |= 1 << uint(x)
			}
		}
		res[j] = tmp
	}
	return res
}
