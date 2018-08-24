package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
		line := readNNums(scanner, 5)
		res := solve(line[0], line[1], line[2], line[3], line[4])
		if len(res) == 0 {
			fmt.Println(-1)
		} else {
			for i := 0; i < len(res); i++ {
				fmt.Printf("%d", res[i])
				if i < len(res)-1 {
					fmt.Print(" ")
				} else {
					fmt.Println()
				}
			}
		}
	}
}

const MAX_N = 1e5

var A []int

func init() {
	A = make([]int, MAX_N+1)
}

func solve(N, S, K, m, M int) []int {
	if m == M {
		return nil
	}

	k := N % K

	check := func(stop int) []int {
		rem := S
		for i := 0; i < N; i++ {
			rem -= A[i]
		}
		if rem < 0 {
			return nil
		}

		goal := M
		for i := N - 1; i >= 0; i-- {
			if i == stop {
				goal--
			}
			theta := min(goal-A[i], rem)
			A[i] += theta
			rem -= theta
		}

		if rem > 0 {
			return nil
		}

		var med1 float64

		if N%2 == 1 {
			med1 = float64(A[N/2])
		} else {
			med1 = (float64(A[N/2]) + float64(A[N/2-1])) / 2
		}
		B := make([]int, N)
		n := N - 2

		copy(B, A[:k])
		copy(B[k:], A[k+1:K])
		copy(B[K-1:n], A[K+1:N])

		med2 := float64(B[n/2])
		if med1 == med2 {
			return nil
		}
		copy(B, A)
		return B
	}

	res := make([][]int, 0, 5)
	half := (N + 1) / 2
	for f := max(0, half-3); f <= min(N, half+3); f++ {
		for i := 0; i < f; i++ {
			A[i] = m
		}
		for i := f; i < N; i++ {
			A[i] = m + 1
		}
		B := check(f - 1)
		if len(B) > 0 {
			res = append(res, B)
		}
	}
	if len(res) == 0 {
		return nil
	}
	sort.Sort(IntsSlice(res))

	return res[0]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type IntsSlice [][]int

func (slice IntsSlice) Len() int {
	return len(slice)
}

func (slice IntsSlice) Less(i, j int) bool {
	n := len(slice[i])
	for k := 0; k < n; k++ {
		if slice[i][k] < slice[j][k] {
			return true
		} else if slice[i][k] > slice[j][k] {
			return false
		}
	}
	return false
}

func (slice IntsSlice) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
