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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		scanner.Scan()
		S := scanner.Text()
		fmt.Println(solve(S))
		tc--
	}
}

func solve(S string) int64 {
	n := len(S)

	A := make([][]int64, n)
	C := make([][]int64, n)
	for i := 0; i < n; i++ {
		A[i] = make([]int64, 26)
		if i > 0 {
			copy(A[i], A[i-1])
		}
		C[i] = make([]int64, 26)
		copy(C[i], A[i])
		for j := 1; j < 26; j++ {
			C[i][j] += C[i][j-1]
		}
		A[i][int(S[i]-'a')]++
	}

	B := make([][]int64, n)
	D := make([][]int64, n)
	for i := n - 1; i >= 0; i-- {
		B[i] = make([]int64, 26)
		if i < n-1 {
			copy(B[i], B[i+1])
		}
		D[i] = make([]int64, 26)
		copy(D[i], B[i])
		for j := 24; j >= 0; j-- {
			D[i][j] += D[i][j+1]
		}
		B[i][int(S[i]-'a')]++
	}

	count := func(i, j int) int64 {
		var less int64
		if j > 0 {
			less = C[i][j-1]
		}
		var more int64
		if j < 25 {
			more = D[i][j+1]
		}
		return less + more
	}

	var total int64

	for i := 0; i < n; i++ {
		j := int(S[i] - 'a')
		if j > 0 {
			total += C[i][j-1]
		}
	}

	best := total
	for i := 0; i < n; i++ {
		k := int(S[i] - 'a')
		old := count(i, k)
		for j := 0; j < 26; j++ {
			x := int64(abs(j - k))
			y := total + count(i, j) - old
			if x+y < best {
				best = x + y
			}
		}
	}
	return best
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
