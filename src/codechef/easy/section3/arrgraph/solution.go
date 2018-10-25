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
		fillNNums(scanner, n, ARR[:n])
		res := solve(n, ARR[:n])
		fmt.Println(res)
		for i := 0; i < n; i++ {
			if i < n-1 {
				fmt.Printf("%d ", ARR[i])
			} else {
				fmt.Printf("%d\n", ARR[i])
			}
		}
	}
}

const MAX_N = 50

var ARR [MAX_N]int
var SET [MAX_N]int
var CNT [MAX_N]int

func solve(n int, A []int) int {
	size := n
	for i := 0; i < n; i++ {
		SET[i] = i
		CNT[i] = 1
	}

	var find func(v int) int
	find = func(v int) int {
		if v != SET[v] {
			SET[v] = find(SET[v])
		}
		return SET[v]
	}

	union := func(i, j int) {
		pi := find(i)
		pj := find(j)
		if pi == pj {
			return
		}
		if CNT[pi] >= CNT[pj] {
			CNT[pi] += CNT[pj]
			SET[pj] = pi
		} else {
			CNT[pj] += CNT[pi]
			SET[pi] = pj
		}
		size--
	}

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if gcd(A[i], A[j]) == 1 {
				union(i, j)
			}
		}
	}

	if size == 1 {
		return 0
	}

	flag := true

	for i := 0; i < n && flag; i++ {
		flag = A[i] == 47
	}

	if flag {
		A[0] = 41
	} else {
		A[0] = 47
	}

	return 1
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
