package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	pos := readInt(scanner.Bytes(), 0, &a)
	readInt(scanner.Bytes(), pos+1, &b)
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
	return res[:n]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var N, K, Q int
	scanner.Scan()
	pos := readInt(scanner.Bytes(), 0, &N)
	pos = readInt(scanner.Bytes(), pos+1, &K)
	readInt(scanner.Bytes(), pos+1, &Q)
	F := readNNums(scanner, K)
	X := make([]int, Q)
	for i := 0; i < Q; i++ {
		X[i] = readNum(scanner)
	}

	res := solve(N, K, Q, F, X)

	for _, ans := range res {
		if ans {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}

const MAX = 2e6 + 5

func solve(N, K, Q int, F []int, X []int) []bool {
	can := make([]bool, MAX)
	can[0] = true

	sort.Ints(F)
	res := make([]bool, Q)
	if F[0] == 1 {
		for i := 0; i < Q; i++ {
			res[i] = true
		}
		return res
	}

	for i := 0; i < K; i++ {
		val := F[i]
		if can[val] {
			continue
		}
		for j := 0; j+val < MAX; j++ {
			if can[j] {
				can[j+val] = true
			}
		}
	}

	for i := 0; i < Q; i++ {
		res[i] = can[X[i]]
	}
	return res
}
