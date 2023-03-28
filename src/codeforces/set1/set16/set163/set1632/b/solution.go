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
		res := solve(n)
		for i := 0; i < n; i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
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

func solve(n int) []int {
	// pi ^ pi+1 最小
	// 从高位开始，应该将相同高位的值要放在一起
	// 但是必然存在某个高位i不匹配任何值
	// 那么这时候应该使的它与自己所有的低位进行匹配，
	// 剩下的可以随意匹配
	var h int
	for (1 << (h + 1)) < n {
		h++
	}
	// max (pi ^ pi+1) = 1 << h
	// h就是匹配的值
	res := make([]int, n)

	if (1<<h)+1 <= n-1 {
		// (1 << h) + 1 放在中间
		// 然后1
		// 左边全部是 x >= 1 << h的数
		x := n - 1
		for x >= (1 << h) {
			res[x] = x
			x--
		}
		x++
		res[x], res[x+1] = res[x+1], res[x]
		i := x - 1
		res[i] = 1
		i--
		if i >= 0 {
			res[i] = 0
			i--
		}
		x = 2
		for i >= 0 {
			res[i] = x
			i--
			x++
		}
	} else {
		res[0] = n - 1
		res[1] = 0
		for i, x := 2, 1; i < n; i++ {
			res[i] = x
			x++
		}
	}
	return res
}
