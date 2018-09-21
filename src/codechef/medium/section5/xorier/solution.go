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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		n := readNum(scanner)
		fillNNums(scanner, n, A)
		fmt.Println(solve(n, A))
	}
}

const MAX_NUM = 1000002
const MAX_N = 1e5 + 10

var freq []int
var A []int

func init() {
	freq = make([]int, MAX_NUM)
	A = make([]int, MAX_N)
}

func solve(n int, A []int) int {
	for i := 0; i < MAX_NUM; i++ {
		freq[i] = 0
	}

	var odd, even int

	for i := 0; i < n; i++ {
		if A[i]&1 == 1 {
			odd++
		} else {
			even++
		}
		freq[A[i]]++
	}

	var ans int
	for i := 0; i < n; i++ {
		if A[i]&1 == 1 {
			ans += odd
		} else {
			ans += even
		}
		ans -= freq[A[i]^2]
		ans -= freq[A[i]]
	}
	return ans >> 1
}
