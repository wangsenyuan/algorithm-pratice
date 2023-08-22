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
		s := make([]string, 2*n+1)
		for i := 0; i < len(s); i++ {
			s[i] = readString(reader)
		}
		res := solve(s)
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

func solve(x []string) string {
	// 如何从单个字母开始测试？
	// 完全没有头绪。。。
	// 唯一能确定的，就是初始字母，肯定是x中长度为1的字符
	// 如果只有一个，答案是显然的
	// 但是如果是多个，考虑2个的情况，如何确定呢？
	// 如果一个数，它是被替换进来，那么它的奇偶性发生了变化，被替换出去，奇偶性也发生了变化
	// 所以不用管，它是替换还是被替换，只关心奇偶性即可
	var flag int
	for _, cur := range x {
		for i := 0; i < len(cur); i++ {
			a := int(cur[i] - 'a')
			flag ^= 1 << a
		}
	}
	// 只有一位为1
	for i := 0; i < 26; i++ {
		if (flag>>i)&1 == 1 {
			c := byte('a' + i)
			return fmt.Sprintf("%c", c)
		}
	}
	return ""
}
