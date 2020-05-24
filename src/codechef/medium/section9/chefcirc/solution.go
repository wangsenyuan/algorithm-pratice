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
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func fillNNums(bs []byte, n int, res []int) {
	x := 0
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)

	points := make([][]float64, n)

	for i := 0; i < n; i++ {
		points[i] = make([]float64, 2)
		fmt.Scanf("%f %f", &points[i][0], &points[i][1])
	}

	res := solve(n, m, points)

	fmt.Printf("%.6f\n", res)
}

func solve(n int, m int, points [][]float64) float64 {

	var left float64
	var right float64

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			d := distance(points[i], points[j])
			// d /= 2
			if d > right {
				right = d
			}
		}
	}

	for math.Abs(right-left) > EPS {
		mid := (left + right) / 2

		cnt := count(mid, n, points)

		if cnt >= m {
			right = mid
		} else {
			left = mid
		}
	}

	return right

}

const EPS = 1e-8

func count(R float64, n int, points [][]float64) int {
	if R <= 0 {
		return 0
	}
	L := 2 * R

	countInPoints := func(center []float64) int {
		var cnt int
		for i := 0; i < n; i++ {
			d := distance(points[i], center)
			if d < R || math.Abs(d-R) < EPS {
				cnt++
			}
		}
		return cnt
	}

	var best int = 1

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			d := distance(points[i], points[j])
			if d > L || d < EPS {
				continue
			}

			center := getCenter(points[i], points[j], R)

			cnt := countInPoints(center)
			if cnt > best {
				best = cnt
			}
		}
	}

	return best
}

func distance(a, b []float64) float64 {
	dx := float64(b[0] - a[0])
	dy := float64(b[1] - a[1])

	d := dx*dx + dy*dy

	return math.Sqrt(d)
}

func getCenter(a, b []float64, r float64) []float64 {
	//
	x1, y1 := float64(a[0]), float64(a[1])
	x2, y2 := float64(b[0]), float64(b[1])
	rr := float64(r) * float64(r)

	q := math.Sqrt(((x2 - x1) * (x2 - x1)) + ((y2 - y1) * (y2 - y1)))

	x3 := (x1 + x2) / 2
	y3 := (y1 + y2) / 2
	basex := math.Sqrt(rr-((q/2)*(q/2))) * ((y1 - y2) / q)
	basey := math.Sqrt(rr-((q/2)*(q/2))) * ((x2 - x1) / q)
	return []float64{x3 + basex, y3 + basey}
}
