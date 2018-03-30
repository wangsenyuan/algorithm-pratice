package main

import (
	"fmt"
	"math"
)

func main() {
	var A, B int

	fmt.Scanf("%d %d", &A, &B)
	fmt.Println(solve(A, B))
}

func solve(A, B int) int {
	divs := make([][]int, B+1)
	sb := int(math.Sqrt(float64(B)))

	for i := 1; i <= B; i++ {
		divs[i] = make([]int, 0, sb)
	}

	for i := 1; i <= sb; i++ {
		for j := i; j <= B; j += i {
			divs[j] = append(divs[j], i)
			if j/i > sb {
				divs[j] = append(divs[j], j/i)
			}
		}
	}
	var ans int
	for b := 1; b <= B; b++ {
		for _, x := range divs[b] {
			y := b / x
			if x >= y {
				if (x+y)%2 == 0 && (x-y)%2 == 0 {
					a := (x - y) / 2
					if a >= 1 && a <= A {
						ans++
					}
				}
			}
		}
	}
	return ans
}
