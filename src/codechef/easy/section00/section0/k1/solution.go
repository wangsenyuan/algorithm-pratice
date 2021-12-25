package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	t := toInt(scanner.Bytes())

	for i := 0; i < t; i++ {
		scanner.Scan()
		n := toInt(scanner.Bytes())
		scanner.Scan()
		a, b, c := toIntTrip(scanner.Bytes())
		points := make([][]int, n)
		for j := 0; j < n; j++ {
			scanner.Scan()
			x, y := toIntPair(scanner.Bytes())
			points[j] = []int{x, y}
		}

		res := solve(a, b, c, points, n)
		fmt.Printf("%.6f\n", res)
	}
}

func toIntTrip(buf []byte) (int, int, int) {
	i, a, b, c := 0, 0, 0, 0
	sign := 1
	for buf[i] != ' ' {
		if buf[i] == '-' {
			sign = -1
			i++
			continue
		}
		a = a*10 + int(buf[i]-'0')
		i++
	}
	a *= sign

	i++
	sign = 1
	for buf[i] != ' ' {
		if buf[i] == '-' {
			sign = -1
			i++
			continue
		}
		b = b*10 + int(buf[i]-'0')
		i++
	}
	b *= sign
	i++
	sign = 1
	for i < len(buf) {
		if buf[i] == '-' {
			sign = -1
			i++
			continue
		}
		c = c*10 + int(buf[i]-'0')
		i++
	}
	c *= sign
	return a, b, c
}

func toIntPair(buf []byte) (int, int) {
	i, a, b := 0, 0, 0
	sign := 1
	for buf[i] != ' ' {
		if buf[i] == '-' {
			sign = -1
			i++
			continue
		}
		a = a*10 + int(buf[i]-'0')
		i++
	}
	a *= sign
	i++

	sign = 1
	for i < len(buf) {
		if buf[i] == '-' {
			sign = -1
			i++
			continue
		}
		b = b*10 + int(buf[i]-'0')
		i++
	}
	b *= sign
	return a, b
}

func toInt(buf []byte) (n int) {
	for _, v := range buf {
		n = n*10 + int(v-'0')
	}
	return
}

func dist(a, b, c, d float64) float64 {
	return math.Sqrt((c-a)*(c-a) + (d-b)*(d-b))
}

func solve(A int, B int, C int, points [][]int, n int) float64 {
	if B == 0 {
		return verticalLine(A, C, points, n)
	}

	a, b := -float64(A)/float64(B), -float64(C)/float64(B)

	var sx, sy float64
	for i := 0; i < n; i++ {
		sx += float64(points[i][0])
		sy += float64(points[i][1])
	}
	x := (sx - float64(n)*a*b + a*sy) / ((a*a + 1.0) * float64(n))
	y := a*x + b

	var res float64
	for i := 0; i < n; i++ {
		res += dist(x, y, float64(points[i][0]), float64(points[i][1]))
	}
	return res
}

func verticalLine(A int, C int, points [][]int, n int) float64 {
	x := -1.0 * float64(C) / float64(A)
	var s float64
	for i := 0; i < n; i++ {
		s += float64(points[i][1])
	}
	y := s / float64(n)
	var res float64
	for i := 0; i < n; i++ {
		res += dist(x, y, float64(points[i][0]), float64(points[i][1]))
	}
	return res
}
