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
		line := readNNums(reader, 6)
		res := solve(line[0], line[1], line[2], line[3], line[4], line[5])
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

func solve(N int, M int, X1, Y1 int, X2, Y2 int) int64 {
	//n,  x is for column
	// m, y is for row
	var res int64

	n, m := int64(M), int64(N)
	x2 := int64(X2)
	x1 := int64(X1)
	y2 := int64(Y2)
	y1 := int64(Y1)

	if abs(x2-x1)*2 > m {
		// abs(x2 - x1) - (m - abs(x2 - x1))
		res += int64(max(0, n-2*abs(y2-y1))) * int64(2*abs(x2-x1)-m)
	}

	if abs(y2-y1)*2 > n {
		res += int64(max(0, m-2*abs(x2-x1))) * int64(2*abs(y2-y1)-n)
	}

	if abs(x2-x1)*2 > m && abs(y2-y1)*2 > n {
		res -= int64(2*abs(y2-y1)-n) * int64(2*abs(x2-x1)-m)
	}
	// point one is the left one
	res += 2 * int64(abs(y2-y1)) * int64(abs(x2-x1))

	return res
}

func abs(num int64) int64 {
	if num < 0 {
		return -num
	}
	return num
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}
