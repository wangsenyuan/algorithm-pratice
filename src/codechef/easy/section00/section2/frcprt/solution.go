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
		tc--
		n, m := readTwoNums(scanner)
		A := make([][]byte, n)
		for i := 0; i < n; i++ {
			scanner.Scan()
			A[i] = scanner.Bytes()
		}
		scanner.Scan()
		S := scanner.Bytes()
		res := solve(n, m, A, S)
		for i := 0; i < n; i++ {
			fmt.Println(string(res[i][:m]))
		}
	}
}

const MAX_N = 100

var X [][]byte
var POS []int

func init() {
	X = make([][]byte, MAX_N)
	for i := 0; i < MAX_N; i++ {
		X[i] = make([]byte, MAX_N)
	}
	POS = make([]int, MAX_N)
}

func solve(N, M int, G [][]byte, S []byte) [][]byte {
	for i := 0; i < N; i++ {
		copy(X[i], G[i])
	}

	fns := []func([][]byte, int, int){up, right, down, left}

	index := []int{'U': 0, 'R': 1, 'D': 2, 'L': 3}

	var v, h bool
	T := make([]byte, 5)
	var idx int
	for i := len(S) - 1; i > 0; i-- {
		if !v && (S[i] == 'R' || S[i] == 'L') {
			T[idx] = S[i]
			idx++
			v = true
		}
		if !h && (S[i] == 'U' || S[i] == 'D') {
			T[idx] = S[i]
			idx++
			h = true
		}
	}
	T[idx] = S[0]
	idx++

	for i := idx - 1; i >= 0; i-- {
		j := index[T[i]]
		fns[j](X, N, M)
	}

	return X
}

func up(X [][]byte, N, M int) {

	for j := 0; j < M; j++ {
		POS[j] = 0
		for i := 0; i < N; i++ {
			if X[i][j] == '1' {
				X[i][j] = '0'
				X[POS[j]][j] = '1'
				POS[j]++
			}
		}
	}
}

func right(X [][]byte, N, M int) {
	for i := 0; i < N; i++ {
		POS[i] = M - 1
		for j := M - 1; j >= 0; j-- {
			if X[i][j] == '1' {
				X[i][j] = '0'
				X[i][POS[i]] = '1'
				POS[i]--
			}
		}
	}
}

func down(X [][]byte, N, M int) {
	for j := 0; j < M; j++ {
		POS[j] = N - 1
		for i := N - 1; i >= 0; i-- {
			if X[i][j] == '1' {
				X[i][j] = '0'
				X[POS[j]][j] = '1'
				POS[j]--
			}
		}
	}
}

func left(X [][]byte, N, M int) {
	for i := 0; i < N; i++ {
		POS[i] = 0
		for j := 0; j < M; j++ {
			if X[i][j] == '1' {
				X[i][j] = '0'
				X[i][POS[i]] = '1'
				POS[i]++
			}
		}
	}
}
