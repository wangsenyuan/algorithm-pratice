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

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		K := readNFloats(reader, n)
		C := readNFloats(reader, n)
		ok, f, X := solve(n, K, C)

		if !ok {
			buf.WriteString("-1\n")
		} else {
			buf.WriteString(fmt.Sprintf("%.3f", f))
			for i := 0; i < n; i++ {
				buf.WriteString(fmt.Sprintf(" %.3f", X[i]))
			}
			buf.WriteByte('\n')
		}
	}

	fmt.Print(buf.String())
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
		if s[i] == '\n' {
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

func solve(n int, K []float64, C []float64) (bool, float64, []float64) {
	// ai =sqrt(Ki * (Xi + Ci))
	// bi = 1/ sqrt(Ki)
	var sum1 float64
	var sum2 float64
	for i := 0; i < n; i++ {
		sum1 += K[i] * C[i]
		sum2 += 1 / K[i]
	}

	var prod = sum1 * sum2
	if prod < 0 {
		return false, 0, nil
	}
	F := math.Sqrt(prod)

	X := make([]float64, n)

	for i := 0; i < n; i++ {
		X[i] = sum1/(K[i]*K[i]*sum2) - C[i]
	}

	return true, F, X
}
