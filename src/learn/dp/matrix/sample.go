package main

import (
	"fmt"
	"math"
)

func main() {
	d := []int{5, 2, 3, 4, 6, 7, 8}
	l, p := multiple(d)
	fmt.Println(l)
	output(p, 1, len(d)-1)
}

func output(p [][]int, i int, j int) {
	if i == j {
		fmt.Print("A")
		return
	}

	k := p[i][j]
	fmt.Print("(")
	output(p, i, k)
	output(p, k+1, j)
	fmt.Print(")")
}

func multiple(d []int) (int, [][]int) {
	n := len(d) - 1

	m := make([][]int, n+1)
	p := make([][]int, n+1)
	for i := range m {
		m[i] = make([]int, n+1)
		p[i] = make([]int, n+1)
	}

	for diagonal := 1; diagonal <= n-1; diagonal++ {
		for i := 1; i <= n-diagonal; i++ {
			j := i + diagonal
			m[i][j] = math.MaxInt32
			x := i
			for k := i; k < j; k++ {
				if m[i][k]+m[k+1][j]+d[i-1]*d[k]*d[j] < m[i][j] {
					x = k
					m[i][j] = m[i][k] + m[k+1][j] + d[i-1]*d[k]*d[j]
				}
			}
			p[i][j] = x
		}
	}

	return m[1][n], p
}
