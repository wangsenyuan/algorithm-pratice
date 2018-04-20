package main

import (
	"math"
	"bufio"
	"os"
	"fmt"
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

	t := readOneNum(scanner)

	for t > 0 {
		first := readNNums(scanner, 3)
		n, s, y := first[0], first[1], first[2]

		V := readNNums(scanner, n)
		D := readNNums(scanner, n)
		P := readNNums(scanner, n)
		C := readNNums(scanner, n)
		fmt.Printf("%.6f\n", solve(n, s, y, V, D, P, C))
		t--
	}
}

const INF = 1 << 30
const delta = 1e-6

func solve(N int, S int, Y int, V []int, D []int, P []int, C []int) float64 {
	tt := float64(Y) / float64(S)

	check := func(T float64) bool {
		var curr float64

		for i := 0; i < N; i++ {
			v, d, p, c := V[i], D[i], P[i], C[i]
			front := drive(v, d, p, curr)
			if d > 0 {
				// forward
				tail := front - float64(c)
				if tail <= 0 {
					if front < 0 {
						dt := distance(0, front) / float64(v)
						if dt < tt || math.Abs(dt-tt) < delta {
							curr += distance(0, tail) / float64(v)
						}
					} else {
						curr += distance(0, tail) / float64(v)
					}
				}
			} else {
				// backward
				tail := front + float64(c)
				if tail >= 0 {
					if front > 0 {
						dt := distance(0, front) / float64(v)
						if dt < tt || math.Abs(dt-tt) < delta {
							curr += distance(0, tail) / float64(v)
						}
					} else {
						curr += distance(0, tail) / float64(v)
					}
				}
			}
			curr += tt
			if curr > T {
				return false
			}
		}
		return true
	}

	var left float64
	var right float64 = INF

	for math.Abs(right-left) > delta {
		mid := (left + right) / 2
		if check(mid) {
			right = mid
		} else {
			left = mid
		}
	}
	return left
}

func drive(v, d, p int, t float64) float64 {
	if d == 0 {
		d = -1
	}
	return float64(p) + float64(v*d)*t
	// if d > 0 {
	// 	return front, front - float64(c)
	// }
	// return front, front + float64(c)
}

func distance(a, b float64) float64 {
	return math.Abs(a - b)
}
