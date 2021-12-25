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
	first := readNNums(scanner, 3)
	n, m, q := first[0], first[1], first[2]
	rect := make([][]int, n)
	for i := 0; i < n; i++ {
		rect[i] = readNNums(scanner, m)
	}
	cols := make([][]int, n)
	for i := 0; i < n; i++ {
		cols[i] = make([]int, m)
	}

	for i := 0; i < q; i++ {
		k, l := readTwoNums(scanner)
		fmt.Println(solve(n, m, rect, cols, k, l))
	}
}

func solve(n, m int, rect, cols [][]int, k, l int) int {
	need := k*l/2 + 1
	can := func(median int) bool {
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if i > 0 {
					cols[i][j] = cols[i-1][j]
				} else {
					cols[i][j] = 0
				}
				if rect[i][j] >= median {
					cols[i][j]++
				}
			}
		}

		for i := 0; i+k <= n; i++ {
			sum := 0
			for j := 0; j < l; j++ {
				sum += cols[i+k-1][j]
				if i > 0 {
					sum -= cols[i-1][j]
				}
			}
			if sum >= need {
				return true
			}
			for j := l; j < m; j++ {
				sum += cols[i+k-1][j]
				sum -= cols[i+k-1][j-l]
				if i > 0 {
					sum -= cols[i-1][j]
					sum += cols[i-1][j-l]
				}
				if sum >= need {
					return true
				}
			}
		}

		return false
	}

	left, right := 0, 10000001

	for left+1 < right {
		mid := left + (right-left)/2
		if can(mid) {
			left = mid
		} else {
			right = mid
		}
	}
	return left
}
