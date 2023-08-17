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
		n, k := readTwoNums(reader)
		s := readString(reader)[:n]
		res := solve(s, k)
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

func solve(s string, k int) string {
	//在能将s[0]变成a的前提下，先将尽可能靠近它的数，变成s[0]
	n := len(s)

	var it int
	for it < n && int(s[it]-'a') <= k {
		it++
	}
	// it == n or s[it] - 'a' > k
	if it == n {
		return all_a(n)
	}
	buf := []byte(s)
	if it == 0 {
		// s[0] 也不能到a
		buf[0] = changeByte(buf[0], k)
		for i := 1; i < n; i++ {
			if buf[i] <= s[0] && buf[i] > buf[0] {
				buf[i] = buf[0]
			}
		}
		return string(buf)
	}
	x := 0
	for i := 0; i < it; i++ {
		if s[i] > s[x] {
			x = i
		}
		buf[i] = 'a'
	}
	for i := it + 1; i < n; i++ {
		if buf[i] <= s[x] {
			buf[i] = 'a'
		}
	}

	k -= int(s[x] - 'a')
	buf[it] = changeByte(buf[it], k)
	for i := it + 1; i < n; i++ {
		if buf[i] <= s[it] && buf[i] > buf[it] {
			buf[i] = buf[it]
		}
	}
	return string(buf)
}

func changeByte(x byte, d int) byte {
	return byte(int(x-'a') - d + 'a')
}

func all_a(n int) string {
	res := make([]byte, n)
	for i := 0; i < n; i++ {
		res[i] = 'a'
	}
	return string(res)
}
