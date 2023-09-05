package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	res := solve(n)

	fmt.Println(res)
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

const inf = 1 << 30

func solve(n int) int {
	// x = b * b - a * a
	// 随着b增加，x也会增加
	// 给定b的时候，随着a增加 x会减小
	// 对于给定的b，一共有 b - 1个数, 1, 2, ... b - 1 个 a
	// x = b * b - a * a = (b - a) * (b + a)
	// 如果 f1 = b-a 和 f2 = b + a	是x的两个因子
	// 那么 b = (f1 + f2) / 2
	//     a = (f1 - f2) / 2
	// 显然f1和f2需要有相同的parity
	// 3, 5, 7, 8, 9, 11
	if n == 1 {
		return 3
	}
	if n == 2 {
		return 5
	}
	if n == 3 {
		return 7
	}
	r := n % 3
	n -= 3
	m := ((n+2)/3)*4 + 8
	if r == 0 {
		return m - 1
	}
	if r == 2 {
		return m - 3
	}
	return m - 4
}
