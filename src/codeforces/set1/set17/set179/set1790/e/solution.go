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
		res := solve(n)
		if len(res) == 0 {
			buf.WriteString("-1\n")
		} else {
			buf.WriteString(fmt.Sprintf("%d %d\n", res[0], res[1]))
		}
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

const H = 31

func solve(x int) []int {
	// a + b = a | b + a & b = a ^ b + a & b + a & b
	// a & b = (a ^ b) / 2

	y := x >> 1

	if x&1 == 1 || x&y > 0 {
		return nil
	}

	return []int{x | y, y}
}

func solve1(x int) []int {
	// a ^ b = (a + b) / 2
	// (a + b) = 2 * (a ^ b) = 2 * x
	// 如果x的最高位 = 1 =》 a + b的最高位也为1
	// 假设 a >= b
	//   两种情况 a[d] = 1, b[d] = 0, 或者 a[d-1] = 1, b[d-1] = 1
	// 如果从低位开始呢？
	// 如果 x[0] = 1, a[0], b[0] = 0, 0 是有可能的
	//               a[0], b[0] 中有一个1是不行的
	//               a[0], b[0] = 1， 1 貌似是可行的,
	// 如果 x[0] = 0 呢？a[0], b[0] 至少有一个1
	// 这里必须要有进位才行的
	// 一个1，一个进位
	// 还是要从高往低位
	var a, b int
	for d := H - 1; d >= 0; d-- {
		u := (x >> d) & 1
		if u == 1 {
			// a[d] = 1, and a[d-1] = 1, b[d-1] = 1
			if d == 0 || (x>>(d-1))&1 == 1 {
				return nil
			}
			a |= 1 << d
			b |= 1 << (d - 1)
			a |= 1 << (d - 1)
			d--
		}
	}
	if (a+b)&1 == 1 {
		return nil
	}
	if (a+b)>>1 != x {
		return nil
	}
	return []int{a, b}
}
