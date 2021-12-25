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

const MAX_N = 1000010

var A [MAX_N]int

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		N, K := readTwoNums(scanner)
		fillNNums(scanner, N, A[:N])
		fmt.Println(solve(N, K, A[:N]))
	}
}

func solve(n int, k int, A []int) int {
	var prev int
	var j int
	var best int
	var dist int
	var bl bool
	for i := 0; i < n; i++ {
		if A[i] > k {
			if !bl {
				bl = true
				prev = A[i]
				j = i
				dist++
				best = max(best, dist)
			} else if prev == A[i] {
				dist++
				best = max(best, dist)
				j = i
			} else {
				dist = i - j
				best = max(best, dist)
				prev = A[i]
				j = i
			}
		} else {
			dist++
			if bl {
				best = max(best, dist)
			}
		}
	}
	return best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func solve1(n int, k int, A []int) int {
	var j int
	gt := make(map[int]int)
	var best int
	for i := 0; i < n; i++ {
		if A[i] > k {
			gt[A[i]]++
		}
		for j < i && len(gt) > 1 {
			if A[j] > k {
				gt[A[j]]--
				if gt[A[j]] == 0 {
					delete(gt, A[j])
				}
			}
			j++
		}
		if len(gt) == 1 && i-j+1 > best {
			best = i - j + 1
		}
	}
	return best
}
