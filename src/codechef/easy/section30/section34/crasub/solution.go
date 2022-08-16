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
		s := readString(reader)
		res := solve(s)
		buf.WriteString(fmt.Sprintf("%s\n", res))
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

const INF = 1 << 30

func solve(s string) string {
	if count(s, '1') == 0 {
		return s
	}
	n := rindex(s, '1')

	suf := s[n:]
	s = s[:n]

	k := count(s, '0')

	if k&1 == 1 {
		if count(suf, '0') > 0 {
			return strings.Replace(s, "0", "", -1) + suf[:len(suf)-1]
		}
		if n > 0 && s[n-1] == '1' {
			// drop 1 from suf
			return strings.Replace(s, "0", "", -1)
		}
		if n > 1 && s[n-2] == '0' {
			// 00
			return strings.Replace(s, "0", "", -1) + "01"
		}

		x := index(s, '0')
		y := x
		for y < len(s) && s[y] == '0' {
			y++
		}
		if y-x != k-1 {
			return strings.Replace(s, "0", "", -1) + "01"
		}
		return strings.Replace(s, "0", "", -1)
	}
	if rindex(s, 0)-index(s, 0)+1 != k {
		return strings.Replace(s, "0", "", -1) + suf
	}

	if len(suf) >= 3 {
		return strings.Replace(s, "0", "", -1) + suf[:len(suf)-2]
	}
	if len(suf) == 2 {
		if n > 0 && s[n-1] == '1' {
			return strings.Replace(s, "0", "", -1)
		}
		return strings.Replace(s, "0", "", -1) + "01"
	}
	if strings.HasSuffix(s, "11") {
		s = strings.Replace(s, "0", "", -1)
		return s[:len(s)-1]
	}
	if n > 0 && s[n-1] == '1' {
		s = strings.Replace(s, "0", "", -1)
		return s[:len(s)-1] + "01"
	}
	return strings.Replace(s, "0", "", -1) + "001"
}

func index(s string, x byte) int {
	for i := 0; i < len(s); i++ {
		if s[i] == x {
			return i
		}
	}
	return len(s)
}

func rindex(s string, x byte) int {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == x {
			return i
		}
	}
	return -1
}

func count(s string, x byte) int {
	var res int
	for i := 0; i < len(s); i++ {
		if s[i] == x {
			res++
		}
	}
	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
