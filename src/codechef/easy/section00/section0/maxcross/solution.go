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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var n int
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &n)
	A := make([][]byte, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		A[i] = scanner.Bytes()
	}

	B := solve(n, A)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%d", B[i][j])
			if j < n-1 {
				fmt.Printf(" ")
			} else {
				fmt.Println()
			}
		}
	}
}

func solve(n int, A [][]byte) [][]int {

	B := make([][]int, n)
	X := make([][][]int, n)
	Y := make([][][]int, n)
	for i := 0; i < n; i++ {
		B[i] = make([]int, n)
		X[i] = make([][]int, n)
		Y[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			X[i][j] = make([]int, 4)
			Y[i][j] = make([]int, 4)
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if A[i][j] == 'X' {
				X[i][j][0] = 1
				if i > 0 {
					// from top
					X[i][j][0] += X[i-1][j][0]
				}
				X[i][j][1] = 1

				if j > 0 {
					// from left
					X[i][j][1] += X[i][j-1][1]
				}
				X[i][j][2] = 1

				if i > 0 && j > 0 {
					//from top-left dial
					X[i][j][2] += X[i-1][j-1][2]
				}
				X[i][j][3] = 1

				if i > 0 && j < n-1 {
					// from top-right dial
					X[i][j][3] += X[i-1][j+1][3]
				}
			}
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if A[i][j] == 'X' {
				Y[i][j][0] = 1
				if i < n-1 {
					Y[i][j][0] += Y[i+1][j][0]
				}
				Y[i][j][1] = 1
				if j < n-1 {
					Y[i][j][1] += Y[i][j+1][1]
				}
				Y[i][j][2] = 1
				if i < n-1 && j < n-1 {
					Y[i][j][2] += Y[i+1][j+1][2]
				}
				Y[i][j][3] = 1
				if i < n-1 && j > 0 {
					Y[i][j][3] += Y[i+1][j-1][3]
				}
			}
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if A[i][j] == '.' {
				continue
			}

			tmp := X[i][j][0] + Y[i][j][0] - 1
			for k := 1; k < 4; k++ {
				xy := X[i][j][k] + Y[i][j][k] - 1
				if xy > tmp {
					tmp = xy
				}
			}

			B[i][j] = tmp
		}
	}

	return B
}
