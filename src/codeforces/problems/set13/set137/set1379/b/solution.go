package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		s, _ := reader.ReadBytes('\n')
		var l, r int
		var m uint64
		pos := readInt(s, 0, &l)
		pos = readInt(s, pos+1, &r)
		readUint64(s, pos+1, &m)
		a, b, c := solve(l, r, int64(m))
		fmt.Printf("%d %d %d\n", a, b, c)
	}
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

func solve(l, r int, m int64) (a int, b int, c int) {
	for a = l; a <= r; a++ {
		N := (m + int64(r-l)) / int64(a)
		if N > 0 {
			//find a n
			x := m - N*int64(a)
			// x = b - c
			if x == 0 {
				b = l
				c = l
				return
			}
			if x > 0 && x <= int64(r-l) {
				b = int(x) + l
				c = l
				return
			}

			if x < 0 && x >= int64(l-r) {
				b = int(x) + r
				c = r
				return
			}
		}
	}

	return
}
