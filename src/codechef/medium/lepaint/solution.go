package main

import (
	"bufio"
	"os"
	"fmt"
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

	t := readNum(scanner)

	for t > 0 {
		first := readNNums(scanner, 3)
		N, C, K := first[0], first[1], first[2]
		L := make([]int, N)
		R := make([]int, N)
		for i := 0; i < N; i++ {
			L[i], R[i] = readTwoNums(scanner)
		}

		res := solve(N, C, K, L, R)
		fmt.Printf("%.9f\n", res)

		t--
	}
}

func solve(N int, C int, K int, L []int, R []int) float64 {
	f := make([][]float64, K+1)
	for k := 0; k <= K; k++ {
		f[k] = make([]float64, C)
	}
	var ans float64

	for i := 0; i < N; i++ {
		for k := 0; k <= K; k++ {
			for c := 0; c < C; c++ {
				f[k][c] = 0
			}
		}
		f[0][1] = 1.0
		for k := 0; k < K; k++ {
			left, right := L[k]-1, R[k]-1
			for c := 0; c < C; c++ {
				if f[k][c] > 1e-10 {
					for x := 0; x < C; x++ {
						y := (x * c) % C
						if i >= left && i <= right {
							p := 1.0 / float64(C) / 2.0
							f[k+1][c] += f[k][c] * p
							f[k+1][y] += f[k][c] * p
						} else {
							p := 1.0 / float64(C)
							f[k+1][c] += f[k][c] * p
						}
					}
				}
			}
		}

		for c := 0; c < C; c++ {
			ans += float64(c) * f[K][c]
		}
	}

	return ans
}
