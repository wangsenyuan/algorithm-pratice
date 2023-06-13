package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer
	for tc > 0 {
		tc--
		x := readNum(reader)
		tv := make([]int, 4)
		vv := make([]int, 4)
		for i := 0; i < 4; i++ {
			tv[i], vv[i] = readTwoNums(reader)
		}
		res := solve(x, tv, vv)
		buf.WriteString(fmt.Sprintf("%.12f\n", res))
	}

	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

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

	for len(bs) == 0 || bs[0] == '\n' || bs[0] == '\r' {
		bs, _ = reader.ReadBytes('\n')
	}

	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func solve(t int, tv []int, vv []int) float64 {
	poly := NewPolynomial(tv, vv)
	roots := poly.FindRoots(float64(t))
	return integrate(poly.coeff, roots)
}

type Polynomial struct {
	mat   [][]float64
	coeff []float64
	sz    int
}

func NewPolynomial(tv []int, vv []int) Polynomial {
	sz := len(tv)
	mat := make([][]float64, sz)
	for i := 0; i < sz; i++ {
		mat[i] = make([]float64, sz+1)
	}

	for i := 0; i < sz; i++ {
		var p int64 = 1
		for j := sz - 1; j >= 0; j-- {
			mat[i][j] = float64(p)
			p *= int64(tv[i])
		}
		mat[i][sz] = float64(vv[i])
	}

	coeff := gauss_jordan(mat)

	return Polynomial{mat, coeff, sz}
}

func gauss_jordan(mat [][]float64) []float64 {
	n := len(mat)
	coeff := make([]float64, n)

	for i := 0; i < n; i++ {
		d := mat[i][i]
		for j := i; j <= n; j++ {
			mat[i][j] /= d
		}
		for r := 0; r < n; r++ {
			if r == i {
				continue
			}
			d = mat[r][i]
			for j := i; j <= n; j++ {
				mat[r][j] -= mat[i][j] * d
			}
		}
	}

	for i := 0; i < n; i++ {
		if !floatEq(mat[i][n], 0) {
			coeff[i] = mat[i][n]
		}
	}

	return coeff
}

func (poly Polynomial) Search(l float64, r float64) float64 {
	dl := fx(poly.coeff, l)
	dr := fx(poly.coeff, r)
	if dl*dr >= 0 {
		return -1
	}
	mid := (l + r) / 2
	dm := fx(poly.coeff, mid)
	if dl < 0 {
		for !floatEq(dm, 0) {
			if dm < 0 {
				l = mid
			} else {
				r = mid
			}
			mid = (l + r) / 2
			dm = fx(poly.coeff, mid)
		}
	} else {
		for !floatEq(dm, 0) {
			if dm < 0 {
				r = mid
			} else {
				l = mid
			}
			mid = (l + r) / 2
			dm = fx(poly.coeff, mid)
		}
	}
	return mid
}

func (poly Polynomial) FindRoots(t float64) []float64 {
	a := 3 * poly.coeff[0]
	b := 2 * poly.coeff[1]
	c := 1 * poly.coeff[2]
	d := b*b - 4*a*c
	var roots []float64
	if d < 0 || floatEq(d, 0) {
		r := poly.Search(0, t)
		if r > 0 {
			roots = append(roots, r)
		}
		roots = append(roots, t)
		return roots
	}
	r1 := (-b - math.Sqrt(d)) / (2 * a)
	r2 := (-b + math.Sqrt(d)) / (2 * a)
	if r1 > r2 {
		r1, r2 = r2, r1
	}
	// r1 <= r2
	if r1 > 0 {
		r := poly.Search(0, math.Min(r1, t))
		if r > 0 {
			roots = append(roots, r)
		}
	}
	if r1 < t && r2 > 0 {
		r := poly.Search(math.Max(r1, 0), math.Min(r2, t))
		if r > 0 {
			roots = append(roots, r)
		}
	}
	if r2 < t {
		r := poly.Search(math.Max(r2, 0), t)
		if r > 0 {
			roots = append(roots, r)
		}
	}
	roots = append(roots, t)
	return roots
}

func integrate(coeff []float64, roots []float64) float64 {
	n := len(coeff)
	intCoeff := make([]float64, n+1)
	for i := 0; i < n; i++ {
		intCoeff[i] = coeff[i] / float64(n-i)
	}
	var t1 float64
	var dist float64

	for _, it := range roots {
		dist += math.Abs(fx(intCoeff, it) - fx(intCoeff, t1))
		t1 = it
	}

	return dist
}

func floatEq(a, b float64) bool {
	return math.Abs(a-b) < 1e-9
}

func fx(c []float64, x float64) float64 {
	n := len(c)
	r := c[n-1]
	for i, j := n-2, x; i >= 0; i, j = i-1, j*x {
		r += c[i] * j
	}

	return r
}
