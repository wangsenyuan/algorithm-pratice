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

func readUint64Num(scanner *bufio.Scanner) (a uint64) {
	scanner.Scan()
	readUInt64(scanner.Bytes(), 0, &a)
	return
}

func readTwoUint64Nums(scanner *bufio.Scanner) (a uint64, b uint64) {
	scanner.Scan()
	x := readUInt64(scanner.Bytes(), 0, &a)
	readUInt64(scanner.Bytes(), x+1, &b)
	return
}

func readNUint64Nums(scanner *bufio.Scanner, n int) []uint64 {
	res := make([]uint64, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readUInt64(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	t := readNum(scanner)

	for t > 0 {
		n := readNum(scanner)
		A := readNUint64Nums(scanner, n)
		res := solve(n, A)
		if res {
			fmt.Println("Henry")
		} else {
			fmt.Println("Derek")
		}
		t--
	}
}

func solve(n int, A []uint64) bool {
	x := g(A[0])
	for i := 1; i < n; i++ {
		x = x ^ g(A[i])
	}
	return x > 0
}

var gs = []uint64{0, 1, 2, 2, 3, 3, 0, 0, 0, 0, 0, 0}

func g(x uint64) uint64 {
	if x < 12 {
		return gs[x]
	}
	return g(x / 12)
}
