package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	var f, s = 1, 0
	for n > 0 {
		n--
		var x, y uint64
		bb, _ := reader.ReadBytes('\n')
		pos := readUint64(bb, 0, &x)
		readUint64(bb, pos+1, &y)

		if f == s {
			continue
		}

		a := chk(x, y)
		b := lck(x, y)

		if s == 1 {
			a ^= 1
			b ^= 1
		}
		s, f = a, b
	}

	fmt.Printf("%d %d\n", s, f)
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

func chk(s, e uint64) int {
	if s == e {
		return 0
	}
	if s+1 == e {
		return 1
	}
	if e&1 == 1 {
		return 1 - int(s&1)
	}
	if s <= e/4 {
		return chk(s, e/4)
	}
	if s > (e/4)*2 {
		return int((e - s) & 1)
	}
	return 1
}

func lck(s, e uint64) int {
	if e < 2*s {
		return 1
	}
	return chk(s, e/2)
}
