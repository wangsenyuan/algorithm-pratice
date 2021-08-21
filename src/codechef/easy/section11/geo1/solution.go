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
		n, q := readTwoNums(reader)
		P := make([][]int, n)
		for i := 0; i < n; i++ {
			P[i] = readNNums(reader, 2)
		}
		Q := make([][]int, q)
		for i := 0; i < q; i++ {
			Q[i] = readNNums(reader, 2)
		}
		res := solve(P, Q)
		for _, cur := range res {
			buf.WriteString(fmt.Sprintf("%.3f\n", cur))
		}
	}
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

func solve(P [][]int, Q [][]int) []float64 {
	var cotTheta, side, init float64
	n := len(P)
	for i := 0; i < len(P); i++ {
		i1 := (i + n - 1) % n
		i2 := (i + 1) % n
		a := dist(toFloatCoord(P[i]), toFloatCoord(P[i1]))
		b := dist(toFloatCoord(P[i]), toFloatCoord(P[i2]))
		c := dist(toFloatCoord(P[i1]), toFloatCoord(P[i2]))
		ang := math.Acos((a*a+b*b-c*c)/(2*a*b)) / 2
		cotTheta += 1 / math.Tan(ang)
		side += a
		init += area(P[i], P[i2])
	}

	init /= 2

	res := make([]float64, len(Q))

	for i, cur := range Q {
		ti := float64(cur[0]) * float64(cur[1])
		res[i] = init + side*ti + cotTheta*ti*ti
	}

	return res
}

func area(a, b []int) float64 {
	return float64(a[0])*float64(b[1]) - float64(b[0])*float64(a[1])
}

func toFloatCoord(a []int) []float64 {
	return []float64{float64(a[0]), float64(a[1])}
}

func dist(a, b []float64) float64 {
	dx := (b[0] - a[0]) * (b[0] - a[0])
	dy := (b[1] - a[1]) * (b[1] - a[1])

	return math.Sqrt(dx + dy)
}
