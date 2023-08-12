package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		h, m := readTwoNums(reader)
		s := readString(reader)
		res := solve(h, m, s)
		buf.WriteString(res)
		buf.WriteByte('\n')
	}

	fmt.Print(buf.String())
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

func solve(h int, m int, s string) string {
	hour, min := parse(s)

	checkMirror := func(a int, b int) bool {
		ok, c := mirror(a, 2)
		if !ok || c >= m {
			return false
		}
		ok, d := mirror(b, 2)
		if !ok || d >= h {
			return false
		}
		return true
	}

	format := func(a int, b int) string {
		x := fmt.Sprintf("%02d:", a)
		y := fmt.Sprintf("%02d", b)
		return x + y
	}

	for i := 0; i < h; i++ {
		a := (hour + i) % h
		for j := 0; j < m; j++ {
			b := j
			if i == 0 {
				b = (min + j) % m
			}
			if checkMirror(a, b) {
				return format(a, b)
			}
			if i == 0 && b == m-1 {
				break
			}
		}
	}
	return "00:00"
}

func mirror(t int, l int) (bool, int) {
	var res int
	for i := 0; i < l; i++ {
		r := t % 10
		if r == 0 || r == 1 || r == 8 {
			res = res*10 + r
		} else if r == 2 {
			res = res*10 + 5
		} else if r == 5 {
			res = res*10 + 2
		} else {
			return false, -1
		}
		t /= 10
	}
	return true, res
}

func parse(s string) (hour int, min int) {
	for i := 0; i < len(s); i++ {
		if s[i] == ':' {
			hour, _ = strconv.Atoi(s[:i])
			min, _ = strconv.Atoi(s[i+1:])
			return
		}
	}
	return
}
