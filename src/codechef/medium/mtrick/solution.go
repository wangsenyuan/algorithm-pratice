package main

import (
	"bufio"
	"os"
	"fmt"
)

func readInt64(bytes []byte, from int, val *uint64) int {
	i := from
	var tmp uint64 = 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readNum(scanner *bufio.Scanner) (a uint64) {
	scanner.Scan()
	readInt64(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a uint64, b uint64) {
	scanner.Scan()
	x := readInt64(scanner.Bytes(), 0, &a)
	readInt64(scanner.Bytes(), x+1, &b)
	return
}

func readNNums(scanner *bufio.Scanner, n int) []uint64 {
	res := make([]uint64, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt64(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	t := readNum(scanner)

	for t > 0 {
		n := int(readNum(scanner))
		L := readNNums(scanner, n)
		X := readNNums(scanner, 3)
		A, B, C := X[0], X[1], X[2]
		scanner.Scan()

		res := solve(n, L, A, B, C, scanner.Bytes())
		fmt.Printf("%d", res[0])

		for i := 1; i < n; i++ {
			fmt.Printf(" %d", res[i])
		}
		fmt.Println()
		t--
	}
}

func solve(n int, L []uint64, A, B, C uint64, S []byte) []uint64 {
	var addition uint64 = 0
	var multiply uint64 = 1
	dir := 1
	left, right := 0, n-1
	res := make([]uint64, n)
	for j := 0; j < n; j++ {
		if S[j] == 'R' {
			left, right = right, left
			dir = -dir
		} else if S[j] == 'A' {
			addition = (addition + A) % C
		} else {
			addition = mul(addition, B, C)
			multiply = mul(multiply, B, C)
		}
		res[j] = (mul(L[left], multiply, C) + addition) % C
		left += dir
	}

	return res
}

func mul(a, b, c uint64) uint64 {
	if b == 0 {
		return 0
	}

	ret := mul(a, b>>1, c)
	ret = (ret + ret) % c
	if b&1 == 1 {
		return (ret + a) % c
	}
	return ret
}
