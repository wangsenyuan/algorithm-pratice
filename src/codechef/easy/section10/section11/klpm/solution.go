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

	scanner.Scan()
	S := scanner.Text()
	fmt.Println(solve(S))
}

func solve(S string) int64 {
	n := len(S)

	P := make([][]int, n)

	for i := 0; i < n; i++ {
		P[i] = make([]int, n)
	}

	for l := 1; l <= n; l++ {
		for i := 0; i+l <= n; i++ {
			j := i + l - 1
			if S[i] == S[j] {
				if i == j || i+1 == j {
					P[i][j] = 1
				} else {
					P[i][j] = P[i+1][j-1]
				}
			}
		}
	}

	LR := make([][]int, n)
	for i := 0; i < n; i++ {
		LR[i] = make([]int, n)
		for j := i; j < n; j++ {
			if j == 0 {
				LR[i][j] = P[i][j]
			} else {
				LR[i][j] = P[i][j] + LR[i][j-1]
			}
		}
	}

	RL := make([][]int, n)
	for i := 0; i < n; i++ {
		RL[i] = make([]int, n)
	}

	for j := 0; j < n; j++ {
		for i := j; i >= 0; i-- {
			if i == n-1 {
				RL[i][j] = P[i][j]
			} else {
				RL[i][j] = P[i][j] + RL[i+1][j]
			}
		}
	}

	var res int64

	D := make([][]int64, n)

	for i := 0; i < n; i++ {
		D[i] = make([]int64, n)
	}

	for j := 0; j < n; j++ {
		for i := j - 1; i >= 0; i-- {
			if S[i] == S[j] {
				D[i][j] = 1 + D[i+1][j-1] + int64(LR[i+1][j-1]) + int64(RL[i+1][j-1])
			}
			res += D[i][j]
		}
	}

	return res
}
