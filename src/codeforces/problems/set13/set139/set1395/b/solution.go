package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	line := readNNums(reader, 4)
	n, m, x, y := line[0], line[1], line[2], line[3]
	res := solve(n, m, x, y)
	fmt.Print(res)
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

func solve(n, m int, x, y int) string {
	var buf bytes.Buffer
	u, v := x, y
	for v <= m {
		buf.WriteString(fmt.Sprintf("%d %d\n", u, v))
		v++
	}
	u, v = x, y-1
	for v > 0 {
		buf.WriteString(fmt.Sprintf("%d %d\n", u, v))
		v--
	}

	u, v = x-1, 1
	for u > 0 {
		// left most
		v = move(v, m, &buf, u)
		u--
	}

	u = x + 1

	for u <= n {
		// left most
		v = move(v, m, &buf, u)
		u++
	}
	return buf.String()
}

func move(v int, m int, buf *bytes.Buffer, u int) int {
	if v == 1 {
		for v <= m {
			buf.WriteString(fmt.Sprintf("%d %d\n", u, v))
			v++
		}
		v--
	} else {
		for v > 0 {
			buf.WriteString(fmt.Sprintf("%d %d\n", u, v))
			v--
		}
		v++
	}
	return v
}
