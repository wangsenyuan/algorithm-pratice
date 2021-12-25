package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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
		line := readNNums(scanner, 4)
		n, m, w, h := line[0], line[1], line[2], line[3]
		scanner.Scan()
		S := scanner.Text()
		points := make([][]int, m)
		for i := 0; i < m; i++ {
			points[i] = readNNums(scanner, 2)
		}
		res := solve(n, m, w, h, S, points)
		fmt.Printf("%.7f\n", res)
	}
}

func solve(n int, m int, w int, h int, S string, points [][]int) float64 {
	res := closePairDist(m, points)

	di := make([]float64, 4)

	for i := 0; i < 4; i++ {
		di[i] = math.MaxFloat64
	}

	// res := math.MaxFloat64

	for i := 0; i < m; i++ {
		p := points[i]
		di[0] = math.Min(di[0], float64(h-p[1]))
		di[1] = math.Min(di[1], float64(p[1]))
		di[2] = math.Min(di[2], float64(p[0]))
		di[3] = math.Min(di[3], float64(w-p[0]))

		// for j := i + 1; j < m; j++ {
		// 	tmp := calcDistance(p, points[j])

		// 	res = math.Min(res, tmp)
		// }
	}

	for i := 0; i < n; i++ {
		var j int
		if S[i] == 'U' {
			j = 0
		} else if S[i] == 'D' {
			j = 1
		} else if S[i] == 'R' {
			j = 2
		} else {
			j = 3
		}

		res = math.Min(res, 2*di[j])
		di[j] = di[j^1]
	}

	return res
}

func calcDistance(a, b []int) float64 {
	dx := float64(a[0]) - float64(b[0])
	dy := float64(a[1]) - float64(b[1])
	return math.Sqrt(dx*dx + dy*dy)
}

func closePairDist(m int, points [][]int) float64 {
	X := make([]Pair, m)
	Y := make([]Pair, m)

	for i := 0; i < m; i++ {
		x, y := points[i][0], points[i][1]
		X[i] = Pair{x, i}
		Y[i] = Pair{y, i}
	}

	sort.Sort(Pairs(X))
	sort.Sort(Pairs(Y))

	distance := func(i, j int) float64 {
		a := points[i]
		b := points[j]
		return calcDistance(a, b)
	}

	bruceForce := func(ps []Pair) float64 {
		res := math.MaxFloat64
		n := len(ps)
		for i := 0; i < n; i++ {
			a := ps[i].second
			for j := i + 1; j < n; j++ {
				b := ps[j].second
				res = math.Min(res, distance(a, b))
			}
		}

		return res
	}

	stripCloset := func(ps []Pair, res float64) float64 {
		for i := 0; i < len(ps); i++ {
			for j := i + 1; j < len(ps); j++ {
				dy := float64(points[ps[i].second][1]) - float64(points[ps[j].second][1])
				if dy > res {
					break
				}
				res = math.Min(distance(ps[i].second, ps[j].second), res)
			}
		}

		return res
	}

	var rec func(PX []Pair, PY []Pair) float64

	rec = func(PX []Pair, PY []Pair) float64 {
		n := len(PX)
		if n < 4 {
			return bruceForce(PX)
		}
		mid := n / 2
		LX := PX[:mid]
		RX := PX[mid:]

		// LX, LY := make([]Pair, mid), make([]Pair, mid)
		LY, RY := make([]Pair, mid), make([]Pair, n-mid)

		x0 := PX[mid].first
		var j, k int
		for i := 0; i < n; i++ {
			p := PY[i]
			if points[p.second][0] < x0 {
				LY[j] = p
				j++
			} else {
				RY[k] = p
				k++
			}
		}

		dl := rec(LX, LY)
		dy := rec(RX, RY)
		dd := math.Min(dl, dy)

		strip := make([]Pair, 0, n)

		for i := 0; i < n; i++ {
			p := PY[i]
			d := math.Abs(float64(x0) - float64(points[p.second][0]))
			if d < dd {
				strip = append(strip, p)
			}
		}

		ds := stripCloset(strip, dd)

		return math.Min(ds, dd)
	}

	return rec(X, Y)
}

type Pair struct {
	first, second int
}

type Pairs []Pair

func (ps Pairs) Len() int {
	return len(ps)
}

func (ps Pairs) Less(i, j int) bool {
	return ps[i].first < ps[j].first
}

func (ps Pairs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
