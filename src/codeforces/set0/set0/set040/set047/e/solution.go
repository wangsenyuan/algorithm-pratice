package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	var buf bytes.Buffer
	for _, cur := range res {
		buf.WriteString(fmt.Sprintf("%.6f %.6f\n", cur[0], cur[1]))
	}

	buf.WriteTo(os.Stdout)
}

func readFloat64(bs []byte, pos int, r *float64) int {
	var first float64
	var second float64
	exp := 1.0
	var dot bool
	for pos < len(bs) {
		if bs[pos] == '\n' || bs[pos] == '\r' || bs[pos] == ' ' {
			break
		}
		if bs[pos] == '.' {
			dot = true
		} else if !dot {
			first = first*10 + float64(bs[pos]-'0')
		} else {
			second = second*10 + float64(bs[pos]-'0')
			exp *= 10
		}
		pos++
	}
	*r = first + second/exp
	//fmt.Fprintf(os.Stderr, "%.6f ", *r)
	return pos
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
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func process(reader *bufio.Reader) [][]float64 {
	n, V := readTwoNums(reader)
	shots := make([]float64, n)
	for i := range n {
		s, _ := reader.ReadBytes('\n')
		readFloat64(s, 0, &shots[i])
	}
	m := readNum(reader)
	walls := make([][]float64, m)
	for i := range m {
		walls[i] = make([]float64, 2)
		s, _ := reader.ReadBytes('\n')
		pos := readFloat64(s, 0, &walls[i][0]) + 1
		readFloat64(s, pos, &walls[i][1])
	}
	return solve(V, shots, walls)
}

const g = 9.8

var pi_4 float64 = math.Pi / 4

const iter = 20

func solve(v int, shots []float64, walls [][]float64) [][]float64 {
	V := float64(v)

	search := func(x float64, y float64, l float64, r float64) float64 {
		for range iter {
			a := (l + r) / 2
			t := x / (V * math.Cos(a))
			h := V*math.Sin(a)*t - g*t*t/2
			if h > y {
				r = a
			} else {
				l = a
			}
		}
		return (l + r) / 2
	}

	calc := func(x, y float64) []float64 {
		// 对于墙 x, y, 找到角度范围，可以落在这堵墙上
		// 刚好落在它前面时
		// x = V * cos(a) * t
		// 0 = V * sin(a) * t - g * t * t / 2
		// V * sin(a) = g * t / 2
		// t = V * math.Sin(a) * 2 / g
		// x = V * cos(a) * V * sin(a) * 2 / g
		// x = V * V* sin(2*a) / g
		tmp := x * g / (V * V)
		if tmp > 1 {
			// can't reach it
			return nil
		}
		a := math.Asin(tmp) / 2
		// x = V * cos(a) * t
		// y = V * sin(a) * t - g * t * t / 2
		// t = x / (V * cos(a))
		b := search(x, y, a, pi_4)

		return []float64{a, b}
	}

	n := len(shots)

	type ball struct {
		angel float64
		id    int
	}

	balls := make([]ball, n)
	for i, x := range shots {
		balls[i] = ball{x, i}
	}

	slices.SortFunc(balls, func(a, b ball) int {
		if a.angel < b.angel {
			return -1
		}
		if a.angel > b.angel {
			return 1
		}
		return 0
	})

	open := make([][]int, n)
	end := make([][]int, n)

	slices.SortFunc(walls, func(a, b []float64) int {
		if a[0] < b[0] || math.Abs(a[0]-b[0]) < 1e-6 && a[0] > b[0] {
			// 只保留高的墙
			return -1
		}
		return 1
	})

	for i, wall := range walls {
		tmp := calc(wall[0], wall[1])
		if len(tmp) == 0 {
			continue
		}
		j := binarySearch(n, func(j int) bool {
			return balls[j].angel >= tmp[0]
		})

		k := binarySearch(n, func(j int) bool {
			return balls[j].angel > tmp[1]
		})

		if j == k {
			continue
		}
		open[j] = append(open[j], i)
		end[k-1] = append(end[k-1], i)
	}
	m := len(walls)

	ans := make([][]float64, n)

	tr := NewSegTree(m)

	far := func(a float64) []float64 {
		// x = V * cos(a) * t
		// 0 = V * sin(a) * t - g * t * t / 2
		// V * sin(a) = g * t / 2
		t := V * math.Sin(a) * 2 / g
		return []float64{V * math.Cos(a) * t, 0}
	}

	find := func(a float64) []float64 {
		i := tr.Get(0, m)
		if i == inf {
			// 没有墙能组织它
			return far(a)
		}
		// 被墙i阻止
		x := walls[i][0]
		t := x / (V * math.Cos(a))
		y := V*math.Sin(a)*t - g*t*t/2
		return []float64{x, y}
	}

	for i, cur := range balls {
		for _, j := range open[i] {
			tr.Update(j, j)
		}

		ans[cur.id] = find(cur.angel)

		for _, j := range end[i] {
			tr.Update(j, inf)
		}
	}

	return ans
}

func binarySearch(n int, f func(int) bool) int {
	l, r := 0, n
	for l < r {
		mid := (l + r) / 2
		if f(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}

type SegTree struct {
	size int
	arr  []int
}

const inf = 1 << 60

func NewSegTree(n int) *SegTree {
	arr := make([]int, 2*n)
	for i := 0; i < len(arr); i++ {
		arr[i] = inf
	}
	return &SegTree{n, arr}
}

func (seg *SegTree) Update(p int, v int) {
	p += seg.size
	seg.arr[p] = v
	for p > 1 {
		seg.arr[p>>1] = min(seg.arr[p], seg.arr[p^1])
		p >>= 1
	}
}

func (seg *SegTree) Get(l, r int) int {
	res := inf
	l += seg.size
	r += seg.size
	for l < r {
		if l&1 == 1 {
			res = min(res, seg.arr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = min(res, seg.arr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
