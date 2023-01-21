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
		dest := readNNums(reader, 2)
		C := readNNums(reader, 6)
		res := solve(dest, C)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

var dx = []int64{1, 1, 0, -1, -1, 0}
var dy = []int64{1, 0, -1, -1, 0, 1}

func solve(dest []int, C []int) int64 {
	x, y := int64(dest[1]), int64(dest[0])

	var calc func(i, j int) int64
	calc = func(i, j int) int64 {
		if i == j {
			if dx[i] != 0 {
				a := x / dx[i]
				if a > 0 && a*dx[i] == x && a*dy[i] == y {
					return a * int64(C[i])
				}
				return math.MaxInt64
			}
			// dy[i] != 0
			b := y / dy[i]
			if b > 0 && b*dx[i] == x && b*dy[i] == y {
				return b * int64(C[i])
			}
			return math.MaxInt64
		}
		if i > j {
			if j+3 == i {
				return math.MaxInt64
			}
		} else {
			if i+3 == j {
				return math.MaxInt64
			}
		}
		// a, b => a * dx[i], a * dy[i] => x1, y1
		// x1 + b * dx[j], y1+ b * dy[j] => x, y
		// a * dx[i] + b * dx[j] = x
		// a * dy[i] + b * dy[j] = y
		var a, b int64
		if dx[i] == 0 {
			b = x / dx[j]
			a = (y - b*dy[j]) / dy[i]

		}
		if dx[j] == 0 {
			a = x / dx[i]
			b = (y - a*dy[i]) / dy[j]

		}
		if dy[i] == 0 {
			b = y / dy[j]
			a = (x - b*dx[j]) / dx[i]

		}
		// dy[j] == 0
		if dy[j] == 0 {
			a = y / dy[i]
			b = (x - a*dx[i]) / dx[j]

		}
		if a < 0 || b < 0 {
			return math.MaxInt64
		}
		return a*int64(C[i]) + b*int64(C[j])
	}

	var best int64 = math.MaxInt64

	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			tmp := calc(i, j)
			if tmp < best {
				best = tmp
			}
		}
	}
	return best
}
