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
		var A, B uint64
		s, _ := reader.ReadBytes('\n')
		pos := readUint64(s, 0, &A)
		readUint64(s, pos+1, &B)
		res := solve(int64(A), int64(B))
		if len(res) == 0 {
			buf.WriteString("-1\n")
		} else {
			for _, cur := range res {
				buf.WriteString(fmt.Sprintf("%d %d\n", cur[0], cur[1]))
			}
		}
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
		if s[i] == '\n' || s[i] == '\r' {
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

func solve(X, Y int64) [][]int64 {
	// [a, b], [c, d]
	// a + b = X and c + d = X, won't work
	// same for a * b = Y, c * d = Y
	// so the only case is a + b = X, c * d = Y
	// case one b < c
	b := (X + 1) / 2
	var d int64
	c := b + 1
	for c <= Y/c {
		if Y%c == 0 {
			d = Y / c
			if d >= c {
				break
			}
		}
		c++
	}
	if c*d == Y && d > 0 {
		return [][]int64{{X - b, b}, {c, d}}
	}
	a := X / 2
	c = 1
	for c <= Y/c {
		if Y%c == 0 {
			d = Y / c
			if d < c {
				break
			}
			if d < a {
				break
			}
		}
		c++
	}
	if c*d != Y || d < c || d > a {
		return nil
	}

	return [][]int64{{c, d}, {a, X - a}}
}
