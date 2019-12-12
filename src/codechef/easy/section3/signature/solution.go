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
		n, m := readTwoNums(scanner)

		A := make([]string, n)
		for i := 0; i < n; i++ {
			scanner.Scan()
			A[i] = scanner.Text()
		}
		B := make([]string, n)
		for i := 0; i < n; i++ {
			scanner.Scan()
			B[i] = scanner.Text()
		}
		res := solve(n, m, A, B)
		fmt.Println(res)
	}
}

const INF = 1 << 20

func solve(n, m int, A []string, B []string) int {

	b := func(i, j int) byte {
		if i < 0 || i >= n || j < 0 || j >= m {
			return '0'
		}
		return B[i][j]
	}

	check := func(x, y int) int {
		// A[i][j] = B[i+x][j+y]
		var res int

		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if A[i][j] != b(i+x, j+y) {
					res++
				}
			}
		}

		return res
	}

	best := INF

	for x := -n; x <= n; x++ {
		for y := -m; y <= m; y++ {
			best = min(best, check(x, y))
		}
	}

	return best
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
