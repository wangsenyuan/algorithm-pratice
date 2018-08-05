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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for x := 1; x <= tc; x++ {
		n := readNum(scanner)
		flies := make([][]int, n)
		for i := 0; i < n; i++ {
			flies[i] = readNNums(scanner, 6)
		}
		d, t := solve(n, flies)
		fmt.Printf("Case #%d: %.8f %.8f\n", x, d, t)
	}
}

func solve(n int, flies [][]int) (float64, float64) {
	var X, Y, Z int
	var VX, VY, VZ int

	for i := 0; i < n; i++ {
		fly := flies[i]
		X += fly[0]
		Y += fly[1]
		Z += fly[2]
		VX += fly[3]
		VY += fly[4]
		VZ += fly[5]
	}
	x, y, z := float64(X)/float64(n), float64(Y)/float64(n), float64(Z)/float64(n)

	vx := float64(VX) / float64(n)
	vy := float64(VY) / float64(n)
	vz := float64(VZ) / float64(n)
	d := dist(vx, vy, vz)

	if math.Abs(d) < 1e-10 {
		return dist(x, y, z), 0.0
	}

	// position after one second
	t := -dot([]float64{x, y, z}, []float64{vx, vy, vz}) / (d * d)
	if t < 0.0 || math.Abs(t) < 1e-6 {
		return dist(x, y, z), 0.0
	}
	XT, YT, ZT := x+vx*t, y+vy*t, z+vz*t
	return math.Sqrt(XT*XT + YT*YT + ZT*ZT), t
}

func dist(x, y, z float64) float64 {
	X := float64(x)
	Y := float64(y)
	Z := float64(z)
	return math.Sqrt(X*X + Y*Y + Z*Z)
}

func dot(a, b []float64) float64 {
	A := float64(a[0]) * float64(b[0])
	B := float64(a[1]) * float64(b[1])
	C := float64(a[2]) * float64(b[2])
	return A + B + C
}
