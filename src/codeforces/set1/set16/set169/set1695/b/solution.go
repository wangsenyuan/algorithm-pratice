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
		a := readNNums(reader, n)
		res := solve(a)
		buf.WriteString(fmt.Sprintf("%s\n", res))
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

const inf = 1 << 30

func solve(a []int) string {
	n := len(a)
	// 如果 n 是偶数，似乎mike始终能获胜?
	// 那么mike 和 bob操作的pile是互不干扰的
	//此时双方最优的策略就是1个1个的取，直到先遇到取完的那方输
	// 就是在a[i] a[i+1] 判断大小，
	if n&1 == 0 {
		x, y := inf, inf
		for i := 0; i < n; i += 2 {
			x = min(x, a[i])
			y = min(y, a[i+1])
		}
		if x > y {
			return "Mike"
		}
		if x < y {
			return "Joe"
		}

		for i := 0; i < n; i += 2 {
			tmp := a[i] - x
			if tmp == 0 {
				return "Joe"
			}
			tmp = a[i+1] - x
			if tmp == 0 {
				return "Mike"
			}
		}

		return "Joe"
	}
	// 第一圈mike从一开始，第二圈joe开始
	// mike 只要把第一个取完即可，joe就输了
	return "Mike"
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
