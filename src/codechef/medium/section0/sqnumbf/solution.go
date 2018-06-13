package main

import (
	"fmt"
	"math"
)

func main() {
	var t int

	fmt.Scanf("%d", &t)

	for t > 0 {
		var n int

		fmt.Scanf("%d", &n)
		A := make([]int64, n)

		for i := 0; i < n; i++ {
			fmt.Scanf("%d", &A[i])
		}

		fmt.Println(solve(n, A))

		t--
	}
}

func solve(n int, A []int64) int64 {
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			g := gcd(A[i], A[j])
			if g > 1 {
				// g divides A[i] & A[j]
				return g
			}
		}
	}

	for i := 0; i < n; i++ {
		x := A[i]

		// x = p * p * q

		for p := int64(1); p*p*p <= x; p++ {
			if p > 1 && x%(p*p) == 0 {
				return p
			}

			if x%p == 0 {
				q := float64(x / p)
				sq := math.Sqrt(q)

				if sq > 1 && math.Abs(sq*sq-q) < 1e-6 {
					return int64(sq)
				}
			}
		}
	}

	return -1
}

func gcd(a, b int64) int64 {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
