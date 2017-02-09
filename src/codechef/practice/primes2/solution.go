package main

import (
	"fmt"
	"sort"
)

var MAX_N = 1000000

func main() {

	p1 := sieve(MAX_N)
	p2, p3 := squareAndCube(p1)

	var n int
	for {
		fmt.Scanf("%d", &n)
		if n <= 0 {
			break
		}
		a, b, c := findAnswer(n, p1, p2, p3)
		fmt.Printf("%d %d %d\n", a, b, c)
	}
}

func findAnswer(n int, p1, p2, p3 []int) (int, int, int) {
	for i := 0; i < len(p3); i++ {
		a := p3[i]
		a3 := a * a * a
		if a3 >= n {
			break
		}
		for j := 0; j < len(p2); j++ {
			b := p2[j]
			b2 := b * b
			if a3+b2 >= n {
				break
			}
			c := n - a3 - b2
			k := sort.Search(len(p1), func(k int) bool {
				return p1[k] >= c
			})

			if k < len(p1) && p1[k] == c {
				return c, b, a
			}
		}
	}

	return 0, 0, 0
}

func squareAndCube(p1 []int) ([]int, []int) {
	p2, p3 := make([]int, 0, 1000), make([]int, 0, 100)

	for i := 0; i < len(p1); i++ {
		num := p1[i]
		if num*num >= MAX_N {
			break
		}
		p2 = append(p2, num)
		if num*num*num < MAX_N {
			p3 = append(p3, num)
		}
	}

	return p2, p3
}
func sieve(n int) []int {
	flags := make([]bool, n+1)
	ps := make([]int, 0, n)

	flags[2] = false

	for i := 2; i <= n; i++ {
		if flags[i] {
			continue
		}
		ps = append(ps, i)
		for j := 2 * i; j <= n; j += i {
			flags[j] = true
		}
	}

	return ps
}
