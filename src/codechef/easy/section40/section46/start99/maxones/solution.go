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
		n := readNum(reader)
		s := readString(reader)[:n]
		res := solve(s)
		buf.WriteString(res)
		buf.WriteByte('\n')
	}

	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
}

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')
	if len(s) == 0 || len(s) == 1 && s[0] == '\n' {
		return readNInt64s(reader, n)
	}
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
	if len(bs) == 0 || bs[0] == '\n' {
		return readNum(reader)
	}
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

func solve(s string) string {
	// s[i] = s[i-1] ^ s[i+1]
	// 感觉这个貌似很简单呐
	// 如果是00001 =》 0111111
	n := len(s)
	var sum int
	for i := 0; i < n; i++ {
		sum += int(s[i] - '0')
	}
	if sum == 0 || sum == n {
		return s
	}
	m := (n + 6) / 3 * 3
	bad := make([]byte, m)
	for i := 0; i < m; i += 3 {
		bad[i] = '0'
		bad[i+1] = '1'
		bad[i+2] = '1'
	}
	bs := string(bad)
	if s == bs[:n] || s == bs[1:n+1] || s == bs[2:n+2] {
		return s
	}

	ans := make([]byte, n)
	for i := 0; i < n; i++ {
		ans[i] = '1'
	}
	if s[0] == '0' && s[n-1] == '0' {
		ans[0] = '0'
		ans[n-1] = '0'
		return string(ans)
	}
	if s[0] == '0' {
		ans[0] = '0'
		return string(ans)
	}
	if s[n-1] == '0' {
		ans[n-1] = '0'
		return string(ans)
	}
	ans[1] = '0'
	return string(ans)
}
