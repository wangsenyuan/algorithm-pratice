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
		s := readString(reader)[:n]
		res := solve(s)
		s = fmt.Sprintf("%v", res)
		buf.WriteString(s[1 : len(s)-1])
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

func solve(s string) []int64 {
	n := len(s)
	// 如果只能变更一个人的direction, 假设变更了i的方向
	// 对于其他人，没有影响
	// 对i来说, 如果他原来L, 那么变更R后， 变化是 n - 1 - i - i = n - 1 - 2 * i
	// 要让这个值更大，所以基本是把左边的改成向右，把右边的更改成向左
	// 越靠左的要越要修改
	var cur int64
	for i := 0; i < n; i++ {
		if s[i] == 'L' {
			cur += int64(i)
		} else {
			cur += int64(n - 1 - i)
		}
	}
	ans := make([]int64, n)
	h := n / 2
	l, r := 0, n-1
	for k := 0; k < n; k++ {
		ans[k] = cur
		for l < h && s[l] == 'R' {
			l++
		}
		for r >= h && s[r] == 'L' {
			r--
		}
		var a, b int64
		if l < h {
			// a from change s[l]
			a = int64(n - 1 - 2*l)
		}
		if r >= h {
			b = int64(2*r - (n - 1))
		}
		if a >= b && a > 0 {
			// better to change s[l]
			cur += a
			l++
		} else if b > a && b > 0 {
			cur += b
			r--
		}
		ans[k] = cur
	}

	return ans
}
