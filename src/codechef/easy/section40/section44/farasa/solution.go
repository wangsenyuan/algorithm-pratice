package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	A := readNNums(reader, n)
	res := solve(A)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
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
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

const PI = math.Pi
const TWO_PI = 2 * PI

// data index from 1
func fft(arr []float64, _n int, sign int) []float64 {
	n := _n * 2
	data := make([]float64, n+1)
	copy(data, arr)
	j := 1
	for i := 1; i < n; i += 2 {
		if j > i {
			data[j], data[i] = data[i], data[j]
			data[j+1], data[i+1] = data[i+1], data[j+1]
		}
		m := n >> 1
		for m >= 2 && j > m {
			j -= m
			m >>= 1
		}
		j += m
	}

	mm := 2

	for n > mm {
		step := 2 * mm
		theta := TWO_PI / float64(sign*mm)
		temp := math.Sin(0.5 * theta)
		pr := -2.0 * temp * temp
		pi := math.Sin(theta)
		wr := 1.0
		wi := 0.0

		for m := 1; m < mm; m += 2 {
			for i := m; i <= n; i += step {
				j = i + mm
				x := wr*data[j] - wi*data[j+1]
				y := wr*data[j+1] + wi*data[j]
				data[j] = data[i] - x
				data[j+1] = data[i+1] - y
				data[i] += x
				data[i+1] += y
			}
			temp = wr
			wr = temp*pr - wi*pi + wr
			wi = wi*pr + temp*pi + wi
		}
		mm = step
	}
	return data
}

const N3 = 1 << 23

func method3(A []int) int {
	var sum int
	for _, num := range A {
		sum += num
	}

	X := make([]float64, N3)
	Y := make([]float64, N3)

	X[1] = 1
	Y[2*sum+1] = 1

	var t int
	for i := 1; i <= len(A); i++ {
		t += A[i-1]
		X[2*t+1] = 1
		Y[2*(sum-t)+1] = 1
	}

	lim := sum
	for sum&(sum-1) != 0 {
		sum &= sum - 1
	}

	sum <<= 1

	x := fft(X, sum*2, 1)
	y := fft(Y, sum*2, 1)

	for i := 0; i < sum*2; i++ {
		p := 2*i + 1
		q := p + 1
		a := x[p]*y[p] - x[q]*y[q]
		b := x[p]*y[q] + x[q]*y[p]
		x[p] = a
		x[q] = b
	}

	x = fft(x, sum*2, -1)

	var res int
	for i := lim + 1; i <= lim*2; i++ {
		if x[2*i+1] > 1e-6 {
			res++
		}
	}
	return res - 1
}

func method2(A []int) int {
	s := make([]int, 20010)
	t := make([]bool, 20000010)
	s[1] = A[0]
	n := len(A)
	for i := 2; i <= n; i++ {
		s[i] = s[i-1] + A[i-1]
	}

	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			t[s[i]-s[j]] = true
		}
	}

	var res int
	for i := 0; i < len(t); i++ {
		if t[i] {
			res++
		}
	}
	return res - 1
}

func method1(A []int) int {
	s := make([]int64, 2010)
	t := make([]int64, 2001000)
	s[0] = 0
	s[1] = int64(A[0])
	n := len(A)
	for i := 2; i <= n; i++ {
		s[i] = s[i-1] + int64(A[i-1])
	}

	var k int
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			t[k] = s[i] - s[j]
			k++
		}
	}
	t = t[:k]
	sort.Slice(t, func(i, j int) bool {
		return t[i] < t[j]
	})

	var res int
	for i := 1; i <= len(t); i++ {
		if i == len(t) || t[i] > t[i-1] {
			res++
		}
	}
	return res - 1
}

func solve(A []int) int {
	n := len(A)
	if n <= 2000 {
		return method1(A)
	}
	if n <= 20000 {
		return method2(A)
	}
	return method3(A)
}
