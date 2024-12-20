package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

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

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer

	tc := readNum(reader)
	for tc > 0 {
		tc--
		readString(reader)
		s := readString(reader)
		t := readString(reader)
		cnt, res := solve(s, t)
		buf.WriteString(fmt.Sprintf("%d\n", cnt))

		for i := 0; i < len(res); i++ {
			if i+1 < len(res) {
				buf.WriteString(fmt.Sprintf("%d ", res[i]))
			} else {
				buf.WriteString(fmt.Sprintf("%d\n", res[i]))
			}
		}
	}
	fmt.Print(buf.String())
}

func solve(s string, t string) (int, []int) {
	lsh := func(s string, k int) string { return s[k:] + strings.Repeat("0", k) }
	rsh := func(s string, k int) string { return strings.Repeat("0", k) + s[:len(s)-k] }
	xor := func(s, t string) string {
		x := []byte(s)
		for i := range x {
			x[i] ^= t[i] ^ '0'
		}
		return string(x)
	}

	i := strings.IndexByte(s, '1')
	j := strings.IndexByte(t, '1')
	if i < 0 && j < 0 {
		return 0, nil
	}
	if i < 0 || j < 0 {
		return -1, nil
	}

	var ans []int

	if i > j {
		ans = append(ans, i-j)
		s = xor(s, lsh(s, i-j))
		i = j
	}

	k := i + 1
	for s[i+1:] != t[i+1:] {
		for s[k] == t[k] {
			k++
		}
		ans = append(ans, i-k)
		s = xor(s, rsh(s, k-i))
	}

	r := strings.LastIndexByte(s, '1')
	k = i
	for s != t {
		for s[k] == t[k] {
			k--
		}
		ans = append(ans, r-k)
		s = xor(s, lsh(s, r-k))
	}

	return len(ans), ans
}
