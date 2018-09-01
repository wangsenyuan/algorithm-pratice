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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--

		var N, K int
		scanner.Scan()
		pos := readInt(scanner.Bytes(), 0, &N)
		readInt(scanner.Bytes(), pos+1, &K)
		for i := 0; i < K; i++ {
			scanner.Scan()
			pos = readInt(scanner.Bytes(), 0, &dumps[2*i])
			readInt(scanner.Bytes(), pos+1, &dumps[2*i+1])
		}
		fmt.Println(solve(N, K, dumps))
	}
}

const MAX_N = 1e6 + 10

var A []bool
var B []bool
var SO []int64
var SE []int64
var dumps []int

func init() {
	A = make([]bool, 2*MAX_N)
	B = make([]bool, 2*MAX_N)
	SO = make([]int64, 2*MAX_N)
	SE = make([]int64, 2*MAX_N)
	dumps = make([]int, 2*MAX_N)
}

func solve(N, K int, dumps []int) int64 {
	for i := 0; i < 2*MAX_N; i++ {
		A[i] = false
		B[i] = false
	}

	for i := 0; i < K; i++ {
		x, y := dumps[2*i], dumps[2*i+1]
		A[MAX_N+x-y] = true
		B[x+y] = true
	}

	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	counta := func(diff int) int64 {
		return int64(N - abs(diff))
	}

	countb := func(sum int) int64 {
		return int64(N - abs(N+1-sum))
	}

	var sum int64

	for i := -N + 1; i < N; i++ {
		if A[MAX_N+i] {
			sum += counta(i)
		}
	}

	for i := 2; i <= 2*N; i++ {
		if B[i] {
			sum += countb(i)
		}
	}
	for i := 2; i <= 2*N; i++ {
		SO[i] = 0
		SE[i] = 0
	}
	for i := 2; i <= 2*N; i++ {
		if i%2 == 1 {
			SO[i] = SO[i-2] + boolToInt(B[i])
		} else {
			SE[i] = SE[i-2] + boolToInt(B[i])
		}
	}

	for i := -N + 1; i < N; i++ {
		if A[i+MAX_N] {
			minpl := 2 + abs(i)
			maxpl := 2*N - abs(i)

			if minpl%2 == 0 {
				sum -= SE[maxpl] - SE[minpl-2]
			} else {
				sum -= SO[maxpl] - SO[minpl-2]
			}
		}
	}

	return int64(N)*int64(N) - sum
}

func boolToInt(v bool) int64 {
	if v {
		return 1
	}
	return 0
}
