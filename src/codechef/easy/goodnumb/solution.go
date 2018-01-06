package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	for i < len(bytes) && bytes[i] == ' ' {
		i++
	}
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readOneNum(scanner *bufio.Scanner) (a int) {
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

// var factors [][]int
var F []int

const N = 100000
const M = 304128

func init() {
	set := make([]bool, N+1)
	squareFree := make([]bool, N+1)
	// primes := make([]int, N)
	sum := make([]int, N+1)
	// sz := 0
	for i := 2; i <= N; i++ {
		set[i] = true
		squareFree[i] = true
		sum[i] = 1
	}
	squareFree[1] = true

	primeFactorsCnt := make([]int, M+1)

	seen := make([]bool, N+1)
	for i := 2; i <= N; i++ {
		if set[i] {
			// primes[sz] = i
			// sz++
			for j := 1; j*i <= N; j++ {
				if j > 1 {
					set[j*i] = false
				}
				if j%i == 0 {
					squareFree[j*i] = false
				}
				x := i * j
				if !seen[x] {
					seen[x] = true
					for k := 1; x*k <= N; k++ {
						sum[x*k] += x
					}
				}
			}
			for j := 2; j*i <= M; j++ {
				primeFactorsCnt[j*i]++
			}
		}
	}

	F = make([]int, N+1)
	for i := 1; i <= N; i++ {
		F[i] = F[i-1]
		if squareFree[i] {
			if set[primeFactorsCnt[sum[i]]] {
				F[i] += sum[i]
			}
		}
	}
	// fmt.Println(maxSum)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	t := readOneNum(scanner)

	for i := 0; i < t; i++ {
		L, R := readTwoNums(scanner)
		fmt.Println(solve(L, R))
	}
}

func solve(L, R int) int {
	return F[R] - F[L-1]
}
