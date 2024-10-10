package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		x := readString(reader)
		xx := strings.Split(x, " ")
		s := xx[0]
		c := xx[1]
		res := solve(s, c)
		buf.WriteString(res)
		buf.WriteByte('\n')
	}

	fmt.Print(buf.String())
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadBytes('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return string(s[:i])
		}
	}
	return string(s)
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

const NA = "---"

func solve(s string, c string) string {
	// 如果能在s[0] < c[0], find swap it here

	pos := make([][]int, 26)

	for i, x := range []byte(s) {
		j := int(x - 'A')
		pos[j] = append(pos[j], i)
	}

	id := make([]int, 26)

	buf := []byte(s)

	for i := 0; i < len(s) && i < len(c); i++ {
		x := int(s[i] - 'A')
		y := int(c[i] - 'A')
		if x < y {
			// 不需要swap
			return s
		}
		// x >= y
		// 看是否能找到一个更小的
		for z := 0; z < y; z++ {
			if id[z] < len(pos[z]) {
				// found one
				buf[i], buf[pos[z][id[z]]] = buf[pos[z][id[z]]], buf[i]
				return string(buf)
			}
		}
		// 没有找到
		if x > y {
			// 完蛋了，必须交换一个y到这里
			for id[y] < len(pos[y]) {
				buf[i], buf[pos[y][id[y]]] = buf[pos[y][id[y]]], buf[i]
				if string(buf) < c {
					return string(buf)
				}
				buf[i], buf[pos[y][id[y]]] = buf[pos[y][id[y]]], buf[i]
				id[y]++
			}
			return NA
		}
		// x == y
		id[x]++
	}

	if len(s) < len(c) {
		return s
	}

	return NA
}
