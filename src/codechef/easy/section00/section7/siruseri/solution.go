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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	tc := readNum(scanner)

	for tc > 0 {
		tc--
		A, B := readTwoNums(scanner)
		n := A * B
		C := readNNums(scanner, n)
		D := readNNums(scanner, n)
		fmt.Println(solve(A, B, C, D))
	}
}

func solve(A, B int, C []int, D []int) int {
	n := A * B

	win := make([]int, 0, n)
	lose := make([]int, 0, n)
	for i := 0; i < n; i++ {
		if C[i] > D[i] {
			win = append(win, C[i])
		} else {
			lose = append(lose, D[i])
		}
	}

	sort.Ints(win)
	sort.Ints(lose)

	var res int
	rem := A
	a, b := len(win)-1, len(lose)-1
	for a >= 0 && b >= 0 {
		if win[a] > lose[b] {
			res++
			a--
		} else {
			b--
		}
		b -= (B - 1)
		rem--
	}
	if b < 0 {
		res += rem
	}
	return res
}
