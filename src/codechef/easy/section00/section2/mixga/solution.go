package main

import (
	"bufio"
	"fmt"
	"os"
)

const MAX_N = 30010

var input []int

func init() {
	input = make([]int, MAX_N)
}

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
	// res := make([]int, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt(scanner.Bytes(), x+1, &input[i])
	}
	return input[:n]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		N, K := readTwoNums(scanner)
		A := readNNums(scanner, N)
		fmt.Println(solve(N, A, K))
	}
}

func solve(n int, A []int, K int) int {
	var X int

	for i := 0; i < n; i++ {
		if A[i] == 0 {
			continue
		}
		if i&1 == 0 {
			X++
		} else {
			if X == 0 {
				X++
			} else {
				X--
			}
		}
	}
	if X >= K {
		return 1
	}
	return 2
}
