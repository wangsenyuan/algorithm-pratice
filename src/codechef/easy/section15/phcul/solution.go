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

func fillNNums(scanner *bufio.Scanner, n int, res []int) {
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var tmp int64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readNInt64Nums(scanner *bufio.Scanner, n int) []int64 {
	res := make([]int64, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt64(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		x, y := readTwoNums(scanner)
		line0 := readNNums(scanner, 3)
		n, m, k := line0[0], line0[1], line0[2]
		A := readNNums(scanner, 2*n)
		B := readNNums(scanner, 2*m)
		C := readNNums(scanner, 2*k)
		res := solve(n, A, m, B, k, C, x, y)
		fmt.Printf("%.7f\n", res)
	}
}

func solve(n int, A []int, m int, B []int, k int, C []int, x int, y int) float64 {
	// for each point in A, find the min distance to C
	cps := make([]*Point, k)

	for i := 0; i < k; i++ {
		e, f := C[2*i], C[2*i+1]
		p := new(Point)
		p.x = e
		p.y = f
		cps[i] = p
	}

	// sort.Sort(Points(cps))

	aa := make([]float64, n)

	for i := 0; i < n; i++ {
		aa[i] = findClosestDistance(A[2*i], A[2*i+1], cps)
	}

	bb := make([]float64, m)

	for i := 0; i < m; i++ {
		bb[i] = findClosestDistance(B[2*i], B[2*i+1], cps)
	}

	best := math.MaxFloat64

	for i := 0; i < n; i++ {
		d1 := distance(x, y, A[2*i], A[2*i+1])
		for j := 0; j < m; j++ {
			d2 := distance(A[2*i], A[2*i+1], B[2*j], B[2*j+1])
			d3 := bb[j]

			d := d1 + d2 + d3

			if d < best {
				best = d
			}
		}
	}

	for j := 0; j < m; j++ {
		d1 := distance(x, y, B[2*j], B[2*j+1])
		for i := 0; i < n; i++ {
			d2 := distance(B[2*j], B[2*j+1], A[2*i], A[2*i+1])
			d3 := aa[i]
			d := d1 + d2 + d3
			if d < best {
				best = d
			}
		}
	}

	return best
}

func distance(a, b, c, d int) float64 {
	dx := float64(c) - float64(a)
	dy := float64(d) - float64(b)

	dd := dx*dx + dy*dy

	return math.Sqrt(dd)
}

func findClosestDistance(x, y int, ps []*Point) float64 {
	dist := int64(math.MaxInt64)

	for i := 0; i < len(ps); i++ {
		dx := int64(ps[i].x) - int64(x)
		dy := int64(ps[i].y) - int64(y)
		d := dx*dx + dy*dy

		if d < dist {
			dist = d
		}
	}

	return math.Sqrt(float64(dist))
}

type Point struct {
	x int
	y int
}

// type Points []*Point

// func (ps Points) Len() int {
// 	return len(ps)
// }

// func (ps Points) Less(i, j int) bool {
// 	return ps[i].x < ps[j].x || (ps[i].x == ps[j].x && ps[i].y < ps[j].y)
// }

// func (ps Points) Swap(i, j int) {
// 	ps[i], ps[j] = ps[j], ps[i]
// }
