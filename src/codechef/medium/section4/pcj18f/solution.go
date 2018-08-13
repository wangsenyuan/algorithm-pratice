package main

import (
	"bufio"
	"fmt"
	"math"
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

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	sign := int64(1)
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := int64(0)
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	n, m := readTwoNums(scanner)

	points := make([][]int, n)
	for i := 0; i < n; i++ {
		points[i] = readNNums(scanner, 2)
	}

	res := solve(n, m, points)

	fmt.Printf("%.10f", res)
}

func solve(n int, m int, points [][]int) float64 {
	M := int64(m)
	side := M*M + 1
	ds := math.Sqrt(float64(side))
	INF := float64(math.MaxFloat32)
	a, b := INF, -INF
	c, d := INF, -INF

	f := float64(m)
	for i := 0; i < n; i++ {
		x, y := float64(points[i][0]), float64(points[i][1])
		u := (y - f*x) / ds
		v := (x + f*y) / ds
		a = min(a, u)
		b = max(b, u)

		c = min(c, v)
		d = max(d, v)
	}

	return ((b - a) + (d - c)) * 2
}

func min(a, b float64) float64 {
	if a <= b {
		return a
	}
	return b
}

func max(a, b float64) float64 {
	if a >= b {
		return a
	}
	return b
}
