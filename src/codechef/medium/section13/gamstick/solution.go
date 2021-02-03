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
		s, _ := reader.ReadBytes('\n')
		var pos int
		var n, x1, y1, x2, y2 uint64
		pos = readUint64(s, pos, &n) + 1
		pos = readUint64(s, pos, &x1) + 1
		pos = readUint64(s, pos, &y1) + 1
		pos = readUint64(s, pos, &x2) + 1
		pos = readUint64(s, pos, &y2) + 1
		res := solve(n, x1, y1, x2, y2)
		buf.WriteString(res)
		buf.WriteByte('\n')
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

const Miron = "Miron"
const Slava = "Slava"
const Draw = "Draw"

func solve(n uint64, x1, y1, x2, y2 uint64) string {
	if y1 == y2 {
		return Draw
	}

	if y1 > y2 {
		y1 = n + 1 - y1
		y2 = n + 1 - y2
	}

	if x1 == x2 {
		pos := (y1 + y2) / 2
		if pos > (n - pos) {
			return Miron
		}
		if pos == (n - pos) {
			return Draw
		}
		return Slava
	}

	if y1+1 == y2 {
		if y1 > n-y1 {
			return Miron
		}
		return Draw
	}

	gap := y2 - y1 - 2

	y1 += (gap + 1) / 2
	y2 -= gap / 2
	if y1 > n-y1 {
		return Miron
	}

	if y1+1 < n-y1-1 {
		return Slava
	}

	return Draw
}
