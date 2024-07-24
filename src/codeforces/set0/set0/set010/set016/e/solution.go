package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	a := make([][]float64, n)
	for i := 0; i < n; i++ {
		a[i] = readNFloats(reader, n)
	}
	res := solve(a)
	var buf bytes.Buffer
	for i := 0; i < n; i++ {
		buf.WriteString(fmt.Sprintf("%.7f ", res[i]))
	}
	buf.WriteByte('\n')
	fmt.Print(buf.String())
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
	if n == 0 {
		return res
	}
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

func readFloat64(bytes []byte, from int, val *float64) int {
	i := from
	var sign float64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var real int64

	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		real = real*10 + int64(bytes[i]-'0')
		i++
	}

	if i == len(bytes) || bytes[i] != '.' {
		*val = float64(real)
		return i
	}

	// bytes[i] == '.'
	i++

	var fraq float64
	var base float64 = 0.1
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		fraq += base * float64(bytes[i]-'0')
		base /= 10
		i++
	}

	*val = (float64(real) + fraq) * sign

	return i
}

func readNFloats(reader *bufio.Reader, n int) []float64 {
	s, _ := reader.ReadBytes('\n')
	res := make([]float64, n)
	var pos int
	for i := 0; i < n; i++ {
		pos = readFloat64(s, pos, &res[i]) + 1
	}
	return res
}

func solve(a [][]float64) []float64 {
	n := len(a)
	N := 1 << n

	f := make([]float64, N)
	f[N-1] = 1

	for s := N - 1; s > 0; s-- {
		k := bits.OnesCount(uint(s))
		if k == 1 {
			continue
		}

		c := 1 / float64(k*(k-1)/2)
		for x := s; x > 0; x &= x - 1 {
			i := bits.TrailingZeros(uint(x))
			t := s ^ 1<<i
			p := 0.
			for y := t; y > 0; y &= y - 1 {
				j := bits.TrailingZeros(uint(y))
				p += a[j][i]
			}
			f[t] += f[s] * p * c
		}
	}

	res := make([]float64, n)
	for i := 0; i < n; i++ {
		res[i] = f[1<<i]
	}

	return res
}
