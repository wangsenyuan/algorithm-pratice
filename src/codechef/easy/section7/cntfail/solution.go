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
		A := readNNums(scanner, n)
		fmt.Println(solve(n, A))
	}
}

func solve(n int, A []int) int {
	if n == 0 {
		return 0
	}
	var first = A[0]
	var second = -1
	for i := 1; i < n; i++ {
		if A[i] != first {
			second = A[i]
			break
		}
	}

	if second == -1 {
		//all same
		if first == 0 {
			//all fail
			return n
		}
		if first == n-1 {
			//all pass
			return 0
		}
		// invalid
		return -1
	}

	if first < second {
		first, second = second, first
	}

	if first != second+1 {
		return -1
	}

	cnt1, cnt2 := 0, 0
	for i := 0; i < n; i++ {
		if A[i] == first {
			cnt1++
		} else if A[i] == second {
			cnt2++
		}
	}

	if cnt1+cnt2 != n {
		return -1
	}

	if cnt2 != first {
		// the failed should count all passed
		return -1
	}
	return cnt1
}
