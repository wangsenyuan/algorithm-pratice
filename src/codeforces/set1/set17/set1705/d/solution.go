package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		s := readString(reader)
		x := readString(reader)
		res := solve(s[:n], x[:n])
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func solve(s string, t string) int {
	// s[i-1] != s[i+1], then can toggle it s[i]
	// no way to change s[0] & s[n-1], otherwise seems always has a solution?
	n := len(s)

	if s == t {
		return 0
	}

	if s[0] != t[0] || s[n-1] != t[n-1] {
		return -1
	}

	if !can_change(s) {
		return -1
	}

	// 每次变化影响的是什么呢？
	// 在变化前 s[i] = s[i-1], 变化后肯定 s[i] = s[i+1]， 且 s[i] != s[i-1], 反之亦然
	// 也就是说无论怎么变化 s[i] != s[i-1] 的个数不会减少，
	// 似乎就是在移动 s[i] != s[i-1]的关系到正确的位置
	// 构造两个新的字符串 a[i] = 1, when s[i+1] != s[i]
	//             和  b[i] = 1, when t[i+1] != t[i]

	var a []int

	for i := 1; i < n; i++ {
		if s[i] != s[i-1] {
			a = append(a, i)
		}
	}

	var b []int

	for i := 1; i < n; i++ {
		if t[i] != t[i-1] {
			b = append(b, i)
		}
	}

	if len(a) != len(b) {
		return -1
	}

	var res int

	for i := 0; i < len(a); i++ {
		res += abs(a[i] - b[i])
	}
	return res
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func can_change(s string) bool {
	for i := 1; i+1 < len(s); i++ {
		if s[i-1] != s[i+1] {
			return true
		}
	}
	return false
}
