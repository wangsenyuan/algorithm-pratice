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
	n, k := readTwoNums(scanner)
	A := readNNums(scanner, n)
	fmt.Println(solve(n, A, k))
}

const MAX_N = 100010

var left [MAX_N]int
var right [MAX_N]int

func solve(n int, A []int, k int) int {

	check := func(mid int) bool {
		for i := 0; i < MAX_N; i++ {
			left[i] = 0
			right[i] = 0
		}

		for i := 0; i <= mid; i++ {
			left[A[i]]++
		}
		for i := mid; i < n; i++ {
			right[A[i]]++
		}

		for i := 0; i < MAX_N-1; i++ {
			left[i+1] += left[i] / k
			left[i] %= k
			right[i+1] += right[i] / k
			right[i] %= k
		}

		for i := MAX_N - 1; i >= 0; i-- {
			if left[i] < right[i] {
				return true
			} else if left[i] > right[i] {
				return false
			}
		}
		return false
	}

	i, j := 0, n-1
	var ans int
	for i <= j {
		mid := (i + j) / 2
		if check(mid) {
			ans = mid
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	return ans + 1
}
