package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInt(bs []byte, from int, val *int) int {
	tmp := 0
	i := from
	for i < len(bs) && bs[i] != ' ' {
		tmp = tmp*10 + int(bs[i]-'0')
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

func readNNums(scanner *bufio.Scanner, n int) []int {
	scanner.Scan()
	bs := scanner.Bytes()
	res := make([]int, n)
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
		firstLine := readNNums(scanner, 4)
		N, M, X, K := firstLine[0], firstLine[1], firstLine[2], firstLine[3]
		scanner.Scan()
		S := scanner.Text()
		if solve(N, M, X, K, S) {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}
}

func solve(N, M, X, K int, S string) bool {
	var odd int
	for i := 0; i < K; i++ {
		if S[i] == 'O' {
			odd++
		}
	}

	even := K - odd

	finish := min(odd, X*(M+1)/2) + min(even, X*M/2)

	return finish >= N
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
