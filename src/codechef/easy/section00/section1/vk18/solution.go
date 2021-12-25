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
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

var f []int
var g []int

func init() {
	N := 1000001
	f = make([]int, 2*N+1)
	for i := 0; i < len(f); i++ {
		var x, y int
		j := i
		for j > 0 {
			d := j % 10

			if d%2 == 0 {
				x += d
			} else {
				y += d
			}

			j /= 10
		}
		f[i] = abs(x - y)
	}
	g = make([]int, N)
	g[1] = f[2]
	S := g[1]
	for i := 2; i < N; i++ {
		S += f[2*i-1] + f[2*i] - f[i]
		g[i] = g[i-1] + 2*S - f[2*i]
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var t int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &t)

	for i := 0; i < t; i++ {
		scanner.Scan()
		var n int
		readInt(scanner.Bytes(), 0, &n)
		fmt.Println(solve(n))
	}
}

func solve(n int) int {
	return g[n]
}
