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

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		N, M := readTwoNums(scanner)
		G := make([][]int, N)
		for i := 0; i < N; i++ {
			G[i] = readNNums(scanner, M)
		}
		res := solve(N, M, G)
		if res {
			fmt.Println("FIRST")
		} else {
			fmt.Println("SECOND")
		}
		if tc > 0 {
			scanner.Scan()
		}
	}
}

func solve(N, M int, G [][]int) bool {
	var ans int

	for i := 0; i < N; i++ {
		var prev int
		for j := M - 1; j >= 0; j-- {
			if G[i][j] > 0 {
				if G[i][j] <= prev {
					prev = G[i][j] - 1
				} else {
					prev = G[i][j]
				}
			}
		}
		ans ^= prev
	}

	return ans != 0
}
