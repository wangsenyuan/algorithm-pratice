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
		var x, y, n uint64
		pos = readUint64(s, pos, &x) + 1
		pos = readUint64(s, pos, &y) + 1
		readUint64(s, pos, &n)
		res := solve(x, y, n)
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

func solve(X, Y, N uint64) uint64 {
	if X == Y {
		return 0
	}

	calc := func(prev uint64, l uint64) uint64 {
		if (X>>l<<l)^prev < (Y>>l<<l)^prev {
			return 1 << l
		}
		if (X>>l<<l)^prev > (Y>>l<<l)^prev {
			return 0
		}
		return 1 << (l - 1)
	}
	var prev uint64
	var res uint64
	for i := 29; i >= 0; i-- {
		if N&(1<<uint(i)) > 0 {
			res += calc(prev, uint64(i))
			prev |= 1 << uint64(i)
		}
	}

	res += calc(prev, 0)

	return res
}

func min(a, b uint64) uint64 {
	if a <= b {
		return a
	}
	return b
}
