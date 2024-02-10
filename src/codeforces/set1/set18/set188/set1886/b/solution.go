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
		P := readNNums(reader, 2)
		A := readNNums(reader, 2)
		B := readNNums(reader, 2)
		res := solve(P, A, B)
		buf.WriteString(fmt.Sprintf("%.8f\n", res))
	}

	buf.WriteByte('\n')
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
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func solve(P []int, A []int, B []int) float64 {
	O := []int{0, 0}
	check := func(w float64) bool {
		// 要么P & O in the same cycle
		// else, A & B have touched
		if inSameCycle(A, P, O, w) || inSameCycle(B, P, O, w) {
			return true
		}
		dab := distance(A, B)
		if inCycle(A, P, w) && inCycle(B, O, w) && lessOrEqual(dab, 2*w) {
			return true
		}

		if inCycle(B, P, w) && inCycle(A, O, w) && lessOrEqual(dab, 2*w) {
			return true
		}
		return false
	}

	var l, r float64 = 0, 1 << 30

	for i := 0; i < 60; i++ {
		mid := (l + r) / 2
		if check(mid) {
			r = mid
		} else {
			l = mid
		}
	}

	return (l + r) / 2
}

func inSameCycle(o, a, b []int, r float64) bool {
	return inCycle(o, a, r) && inCycle(o, b, r)
}

func lessOrEqual(a, b float64) bool {
	if a < b {
		return true
	}
	return math.Abs(b-a) <= 1e-6
}

func inCycle(o []int, p []int, r float64) bool {
	d := distance(o, p)
	return lessOrEqual(d, r)
}

func distance(a, b []int) float64 {
	dx := float64(b[0] - a[0])
	dy := float64(b[1] - a[1])
	s := dx*dx + dy*dy
	return math.Sqrt(s)
}
