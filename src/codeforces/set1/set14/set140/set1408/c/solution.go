package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, l := readTwoNums(reader)
		nums := readNNums(reader, n)
		res := solve(l, nums)
		buf.WriteString(fmt.Sprintf("%.8f\n", res))
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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func solve(L int, A []int) float64 {
	n := len(A)
	// A is increasing
	var p1, p2 float64 = 0, float64(L)
	s1, s2 := 1.0, 1.0
	l, r := 0, n-1
	var took float64

	for p1 < p2 {
		if l > r {
			took += (p2 - p1) / (s1 + s2)
			break
		}
		a := (float64(A[l]) - p1) / s1
		b := (p2 - float64(A[r])) / s2
		if a <= b {
			// a <= b
			// first car arrive earlier at l flag
			took += a
			p1 = float64(A[l])
			l++
			p2 -= s2 * a
			s1 += 1.0
		} else {
			took += b
			p2 = float64(A[r])
			r--
			p1 += s1 * b
			s2 += 1.0
		}
	}

	return took
}
