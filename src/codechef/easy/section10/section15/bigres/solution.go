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

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var tmp int64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readNInt64Nums(scanner *bufio.Scanner, n int) []int64 {
	res := make([]int64, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt64(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		n := readNum(scanner)
		X := make([]int, n)
		P := make([]int, n)

		for i := 0; i < n; i++ {
			X[i], P[i] = readTwoNums(scanner)
		}
		fmt.Println(solve(n, X, P))
	}
}

func solve(n int, X []int, P []int) int {
	sort.Ints(X)
	sort.Ints(P)

	if n == 2 {
		//any way
		return (X[1] - X[0]) * (P[0] + P[1])
	}

	// two rules
	// 1. put highest columns at two ends of the farthest points
	// 2. don't put highest point at ends (X[0] or X[n-1])
	D := make([]int, n-1)

	for i := 0; i < n-1; i++ {
		D[i] = X[i+1] - X[i]
	}

	DD := make([]int, n)
	DD[0] = D[0]
	for i := 1; i < n-1; i++ {
		DD[i] = D[i] + D[i-1]
	}

	DD[n-1] = D[n-2]

	sort.Ints(DD)

	var res int

	for i := 0; i < n; i++ {
		res += DD[i] * P[i]
	}

	return res
}

type Seg struct {
	dist  int
	index int
}

type Segs []Seg

func (ss Segs) Len() int {
	return len(ss)
}

func (ss Segs) Less(i, j int) bool {
	return ss[i].dist > ss[j].dist
}

func (ss Segs) Swap(i, j int) {
	ss[i], ss[j] = ss[j], ss[i]
}
