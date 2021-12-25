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

func readUInt64(bytes []byte, from int, val *uint64) int {
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
	t := readNum(scanner)

	for t > 0 {
		scanner.Scan()
		var A, B uint64
		pos := readUInt64(scanner.Bytes(), 0, &A)
		readUInt64(scanner.Bytes(), pos+1, &B)
		fmt.Println(solve(A, B))
		t--
	}
}

func solve(A, B uint64) int {
	// assume B > 0; there is no answer when B == 0
	if A == B {
		return 0
	}

	if A > 1 && B == 1 {
		return -1
	}
	if B == 0 {
		return -1
	}

	a := bitsCount(A)
	b := bitsCount(B - 1)
	if a > b {
		return 2
	}
	return b - a + 1
}

func bitsCount(x uint64) int {
	var res int
	for x > 0 {
		res += int(x & 1)
		x >>= 1
	}

	return res
}
