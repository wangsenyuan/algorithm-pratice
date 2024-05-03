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

	n, r := readTwoNums(reader)
	x := readNNums(reader, n)

	res := solve(r, x)

	var buf bytes.Buffer

	for i := 0; i < n; i++ {
		buf.WriteString(fmt.Sprintf("%.8f ", res[i]))
	}
	fmt.Println(buf.String())
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadBytes('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return string(s[:i])
		}
	}
	return string(s)
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

func solve(r int, x []int) []float64 {
	R := float64(r)
	n := len(x)
	Y := make([]float64, n)

	V := 4 * R * R

	for i := 0; i < n; i++ {
		// 默认在R处
		Y[i] = R
		for j := i - 1; j >= 0; j-- {
			if abs(x[i]-x[j]) <= 2*r {
				dx := float64(x[i] - x[j])
				U := V - dx*dx
				dy := math.Sqrt(U)
				Y[i] = math.Max(Y[i], dy+Y[j])
			}
		}
	}

	return Y
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
