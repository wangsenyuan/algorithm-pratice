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
		n := readNum(scanner)
		A := make([]uint64, n)
		scanner.Scan()
		var pos int
		for i := 0; i < n; i++ {
			pos = readUint64(scanner.Bytes(), pos, &A[i]) + 1
		}
		var lower, upper uint64
		scanner.Scan()
		pos = readUint64(scanner.Bytes(), 0, &lower)
		readUint64(scanner.Bytes(), pos+1, &upper)
		res := solve(n, A, lower, upper)
		fmt.Println(res)
	}
}

func solve(n int, A []uint64, lower uint64, upper uint64) uint64 {
	x := uint64(1)

	for i := 0; i < n; i++ {
		x = lcm(x, A[i])
	}

	if x > upper {
		return upper - lower + 1
	}

	cnt1 := upper / x
	cnt2 := (lower - 1) / x
	return upper - lower + 1 - (cnt1 - cnt2)
}

func gcd(a, b uint64) uint64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b uint64) uint64 {
	g := gcd(a, b)
	return (a * b) / g
}
