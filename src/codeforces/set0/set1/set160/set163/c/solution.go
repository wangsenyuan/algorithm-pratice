package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	line := readNNums(reader, 4)
	n := line[0]
	l := line[1]
	v1 := line[2]
	v2 := line[3]
	A := readNNums(reader, n)
	res := solve(l, v1, v2, A)
	var buf bytes.Buffer

	for i := 0; i <= n; i++ {
		buf.WriteString(fmt.Sprintf("%.10f\n", res[i]))
	}
	fmt.Print(buf.String())
}

func readFloat64(bs []byte, pos int, r *float64) int {
	var first float64
	var second float64
	exp := 1.0
	var dot bool
	for pos < len(bs) && (bs[pos] == '.' || isDigit(bs[pos])) {
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

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func solve(L int, v1 int, v2 int, A []int) []float64 {
	sort.Ints(A)
	n := len(A)
	a := make([]float64, n*2)
	//copy(a, A)
	for i := 0; i < n; i++ {
		a[i] = float64(A[i])
		a[i+n] = a[i] + 2*float64(L)
	}
	dist := float64(v2) * float64(L) / float64(v1+v2)

	//var m int
	d := make([]float64, 0, 2*n+2)
	d = append(d, 0)
	d = append(d, 2*float64(L))
	for i := 0; i < n; i++ {
		d = append(d, a[i])
		x := a[i] - dist
		if x < 0 {
			x += 2 * float64(L)
		}
		d = append(d, x)
	}
	sort.Float64s(d)
	m := len(d)

	ans := make([]float64, n+1)

	k := func(x float64) int {
		return sort.SearchFloat64s(a, x+dist) - sort.SearchFloat64s(a, x)
	}
	for i := 0; i < m-1; i++ {
		j := k((d[i] + d[i+1]) / 2)
		if j <= n {
			ans[j] += (d[i+1] - d[i]) / (2 * float64(L))
		}
	}

	return ans
}
