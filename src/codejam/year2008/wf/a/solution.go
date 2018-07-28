package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for i := 1; i <= tc; i++ {
		n := readNum(scanner)
		A := make([]int, n)
		B := make([]int, n)
		C := make([]int, n)

		for j := 0; j < n; j++ {
			line := readNNums(scanner, 3)
			A[j] = line[0]
			B[j] = line[1]
			C[j] = line[2]
		}
		res := solve(n, A, B, C)
		fmt.Printf("Case #%d: %d\n", i, res)
	}

}

func solve(n int, A []int, B []int, C []int) int {
	v := make([]int, 10001)
	h := make([]int, 10001)

	var best int
	for c := 0; c <= 10000; c++ {
		for i := 0; i < 10001; i++ {
			v[i] = 0
			h[i] = 0
		}

		for i := 0; i < n; i++ {
			if C[i] <= c && A[i]+B[i]+c <= 10000 {
				v[A[i]]++
				h[B[i]]++
			}
		}

		var cnt int

		for a := -1; a < 10000-c; a++ {
			cnt += v[a+1]
			if 10000-c-a < 10001 {
				cnt -= h[10000-c-a]
			}
			if cnt > best {
				best = cnt
			}
		}
	}
	return best
}

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
