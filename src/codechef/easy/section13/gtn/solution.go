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

	query := func(i int) int {
		fmt.Printf("%d %d\n", 1, i)
		return readNum(scanner)
	}

	tc := readNum(scanner)

	for tc > 0 {
		N, X := readTwoNums(scanner)
		i := solve(N, X, query)
		fmt.Printf("2 %d\n", i)
		scanner.Scan()
	}
}

const MAX_N = 1000002

func solve(N, X int, query func(int) int) int {
	if N == -1 {
		N = search(MAX_N, func(i int) bool {
			return query(i) == 0
		})

		N--
	}
	n := N
	// search(odd) >= X
	if N%2 == 0 {
		n--
	}
	// if N is 4, n is 3, and h is 1
	// if N is 5, n is 5, and h is 2
	// make n odd
	h := n / 2

	var last int

	i := search(h+1, func(i int) bool {
		j := 2*i + 1
		x := query(j)
		if x >= X {
			last = x
		}
		return x >= X
	})

	i = 2*i + 1
	if last == X {
		return i
	}

	// then last > X
	// and query(i - 2) < X
	// need to check i - 1
	i--

	ii := i & (-i)

	S := query(i)

	for ii > 1 {
		ii >>= 1
		s := query(i + ii)
		S -= s
	}
	if S == X {
		return i
	}
	return -1
}

func search(n int, fn func(int) bool) int {
	var left int
	right := n

	for left < right {
		mid := (left + right) / 2
		if fn(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return right
}
