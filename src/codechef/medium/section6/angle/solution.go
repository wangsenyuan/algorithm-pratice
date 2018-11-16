package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	tc := readNum(scanner)
	for tc > 0 {
		tc--
		line := readNNums(scanner, 3)
		n, p, q := line[0], line[1], line[2]
		A := readNNums(scanner, n)
		x, y, z := solve(n, p, q, A)
		if x < 0 {
			fmt.Println(-1)
		} else {
			fmt.Printf("%d %d %d\n", x, y, z)
		}
	}
}

func solve(n int, p, q int, A []int) (int, int, int) {
	sort.Ints(A)

	P := int64(p)
	Q := int64(q)

	var u, v, w int
	var best float64 = 1
	var found bool
	for i := 0; i < n; i++ {
		x := int64(A[i])
		for j := i + 1; j < n; j++ {
			y := int64(A[j])

			C := (x*x+y*y)*Q - 2*x*y*P

			k := sort.Search(n, func(k int) bool {
				z := int64(A[k])
				return z*z*Q > C
			})
			k--
			for k > 0 && (k == j || k == i) {
				k--
			}
			if k < 0 {
				//not found
				continue
			}
			//A[k] * A[k] * Q <= C
			if A[k] <= abs(A[i]-A[j]) {
				continue
			}
			tmp := calculate(A[i], A[j], A[k])
			if tmp < best {
				best = tmp
				found = true
				u, v, w = k, i, j
			}
		}
	}

	if !found {
		return -1, -1, -1
	}

	return u + 1, v + 1, w + 1
}

func calculate(a, b, c int) float64 {
	A := float64(a)
	B := float64(b)
	C := float64(c)
	return (A*A + B*B - C*C) / (2 * A * B)
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
