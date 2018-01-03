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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var t int
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &t)

	for i := 0; i < t; i++ {
		scanner.Scan()
		var n int
		scanner.Scan()
		readInt(scanner.Bytes(), 0, &n)
		gates := make([][]int, n)
		for j := 0; j < n; j++ {
			scanner.Scan()
			bs := scanner.Bytes()
			var tp int
			x := readInt(bs, 0, &tp)
			if tp == 0 {
				gates[j] = []int{0}
			} else {
				var a, b int
				x = readInt(bs, x+1, &a)
				readInt(bs, x+1, &b)
				gates[j] = []int{tp, a, b}
			}
		}
		fmt.Printf("%.5f\n", solve(n, gates))
	}
}

var EPS = 0.00000001

func solve(n int, gates [][]int) float64 {

	var g func(v int, p float64) float64

	g = func(v int, p float64) float64 {
		gate := gates[v]
		if gate[0] == 0 {
			//input
			return p
		}
		a, b := gate[1], gate[2]
		pa := g(a-1, p)
		pb := g(b-1, p)
		if gate[0] == 1 {
			return 1.0 - (1-pa)*(1-pb)
		}

		return pa * pb
	}

	check := func(p float64) bool {
		return g(n-1, p) >= 0.5-EPS
	}

	left, right := 0.0, 1.0
	for left+EPS < right {
		mid := (left + right) / 2.0
		if check(mid) {
			right = mid
		} else {
			left = mid
		}
	}
	return left
}
