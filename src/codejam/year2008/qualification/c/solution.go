package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp
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

func readFloat64(bs []byte, from int, val *float64) int {
	i := from
	sign := 1
	if bs[i] == '-' {
		sign = -1
		i++
	}

	var dec float64
	for i < len(bs) && bs[i] != '.' && bs[i] != ' ' {
		dec = dec*10 + float64(bs[i]-'0')
		i++
	}
	*val = dec

	if i == len(bs) || bs[i] == ' ' {
		//no fraction
		return i
	}
	i++
	var frac float64
	base := 1.0
	for i < len(bs) && bs[i] != ' ' {
		frac = frac*10 + float64(bs[i]-'0')
		base *= 10
		i++
	}
	*val += frac / base
	return i * sign
}

func readNFloat64s(scanner *bufio.Scanner, n int) []float64 {
	res := make([]float64, n)

	pos := 0
	scanner.Scan()
	bs := scanner.Bytes()
	//fmt.Printf("[debug] %s\n", string(bs))
	for i := 0; i < n; i++ {
		for pos < len(bs) && bs[pos] == ' ' {
			pos++
		}

		pos = readFloat64(bs, pos, &res[i])
	}
	return res
}

func main() {
	f, err := os.Open("./C-large-practice.in")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	tc := readNum(scanner)

	for i := 1; i <= tc; i++ {
		line := readNFloat64s(scanner, 5)
		f, R, t, r, g := line[0], line[1], line[2], line[3], line[4]
		res := solve(f, R, t, r, g)
		fmt.Printf("Case #%d: %.6f\n", i, res)
	}
}

func solve(f, R, t, r, g float64) float64 {
	if g <= 2*f {
		// no hole at all
		return 1.0
	}

	EPS := 1e-6 * 1e-6

	var intersection func(x1, y1, x2, y2 float64) float64

	intersection = func(x1, y1, x2, y2 float64) float64 {
		if x1*x1+y1*y1 > 1 {
			return 0
		}
		if x2*x2+y2*y2 < 1 {
			return (x2 - x1) * (y2 - y1)
		}

		if (x2-x1)*(y2-y1) < EPS {
			return (x2 - x1) * (y2 - y1) / 2
		}

		mx := (x2 + x1) / 2
		my := (y2 + y1) / 2

		return intersection(x1, y1, mx, my) +
			intersection(mx, y1, x2, my) +
			intersection(mx, my, x2, y2) +
			intersection(x1, my, mx, y2)
	}
	Rx := R - t - f
	var area float64
	for x0 := 0.0; x0 <= Rx; x0 += g + 2*r {
		for y0 := 0.0; y0 <= Rx; y0 += g + 2*r {
			x1 := (x0 + r + f) / Rx
			y1 := (y0 + r + f) / Rx
			x2 := (x0 + r + g - f) / Rx
			y2 := (y0 + r + g - f) / Rx
			area += intersection(x1, y1, x2, y2)
		}
	}

	area *= 4.0
	rr := R / Rx
	cycle := math.Pi * rr * rr

	return 1.0 - area/cycle
}
