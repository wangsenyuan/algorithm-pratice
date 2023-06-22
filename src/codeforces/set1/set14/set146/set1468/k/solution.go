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
		s := readString(reader)
		res := solve(s)
		buf.WriteString(fmt.Sprintf("%d %d\n", res[0], res[1]))
	}
	fmt.Print(buf.String())
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

const INF = 1_000_000_000

func solve(cmd string) []int {
	n := len(cmd)

	check := func(x0, y0 int) bool {
		if x0 == 0 && y0 == 0 {
			return false
		}
		var x, y int
		for i := 0; i < n; i++ {
			if cmd[i] == 'L' && (y != y0 || (x-1 > x0 || x < x0)) {
				x--
			}
			if cmd[i] == 'R' && (y != y0 || (x+1 < x0 || x0 < x)) {
				x++
			}
			if cmd[i] == 'D' && (x != x0 || (y-1 > y0 || y < y0)) {
				y--
			}
			if cmd[i] == 'U' && (x != x0 || (y+1 < y0 || y0 < y)) {
				y++
			}
		}

		return x == 0 && y == 0
	}
	var x, y int
	for i := 0; i < n; i++ {
		x0, y0 := x, y
		if cmd[i] == 'L' {
			x0--
		} else if cmd[i] == 'R' {
			x0++
		} else if cmd[i] == 'D' {
			y0--
		} else {
			y0++
		}

		if check(x0, y0) {
			return []int{x0, y0}
		}

		x, y = x0, y0
	}

	return []int{0, 0}
}
