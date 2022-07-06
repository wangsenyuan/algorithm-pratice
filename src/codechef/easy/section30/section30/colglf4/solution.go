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
		line := readNNums(reader, 6)
		res := solve(line[0], line[1], line[2], line[3], line[4], line[5])
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func solve(N int, E int, H int, A, B, C int) int64 {
	// x + y + z = N
	// 2 * x + z <= E
	// 3 * y + z <= H
	// x * A + y * B + z * C

	var best int64 = math.MaxInt64

	for z := 0; z <= N && z <= E && z <= H; z++ {
		x := (E - z) / 2
		y := (H - z) / 3

		if A <= B {
			// more x more better
			x = min(N-z, x)
			y = min(y, max(0, N-z-x))
			if x+y+z == N {
				best = min2(best, int64(x)*int64(A)+int64(y)*int64(B)+int64(z)*int64(C))
			}
		} else {
			y = min(N-z, y)
			x = min(x, max(0, N-z-y))
			if x+y+z == N {
				best = min2(best, int64(x)*int64(A)+int64(y)*int64(B)+int64(z)*int64(C))
			}
		}
	}

	if best == math.MaxInt64 {
		return -1
	}

	return best
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min2(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}
