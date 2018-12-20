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
	var best int
	var pivot int
	pos := -1
	for i := 0; i < n-1; i++ {
		if A[i] <= A[i+1] {
			pivot = A[i]
			continue
		}
		prev := pivot
		var tmp int
		j := i + 1
		for j < n && A[j] < A[i] {
			if A[j] >= prev {
				tmp++
				prev = A[j]
			}
			j++
		}
		pivot = A[i]
		if tmp > best {
			best = tmp
			pos = i
		}
		i = j - 1
	}

	var ans int

	pivot = 0
	for i := 0; i < n; i++ {
		if i == pos {
			continue
		}
		if A[i] >= pivot {
			ans++
			pivot = A[i]
		}
	}
	var ans2 int
	pivot = 0
	for i := 1; i < n; i++ {
		if A[i] >= pivot {
			ans2++
			pivot = A[i]
		}
	}
	return max(ans, ans2)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
