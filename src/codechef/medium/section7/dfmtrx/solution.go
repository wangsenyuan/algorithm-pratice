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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// writer := bufio.NewWriter(os.Stdout)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		n := readNum(scanner)

		res := solve(n)

		if !res {
			fmt.Println("Boo")
		} else {
			fmt.Println("Hooray")
			for i := 0; i < n; i++ {
				for j := 0; j < n; j++ {
					if j < n-1 {
						fmt.Printf("%d ", matrix[i][j])
					} else {
						fmt.Printf("%d\n", matrix[i][j])
					}
				}
			}
		}
		// writer.Flush()
	}
}

const MAX_N = 2001

var matrix [][]int

func init() {
	matrix = make([][]int, MAX_N)
	for i := 0; i < MAX_N; i++ {
		matrix[i] = make([]int, MAX_N)
	}
}

func solve(n int) bool {
	if n == 1 {
		matrix[0][0] = 1
		return true
	}

	if n&1 == 1 {
		return false
	}

	for i := 0; i < n; i++ {
		matrix[i][i] = 1
	}

	r := roundRobin(n)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n/2; j++ {
			matrix[r[i][j][0]][r[i][j][1]] = i + 2
			matrix[r[i][j][1]][r[i][j][0]] = i + n + 1
		}
	}

	return true
}

func roundRobin(n int) [][][]int {
	r := make([][][]int, n-1)
	for i := 0; i < n-1; i++ {
		r[i] = make([][]int, 0, n/2)
	}

	for i := 0; i < n-1; i++ {
		for j := i; j < n-1; j++ {
			k := (i + j) % (n - 1)
			var a int
			if j == i {
				a = n - 1
			} else {
				a = j
			}
			b := i
			r[k] = append(r[k], []int{a, b})
		}
	}
	return r
}
