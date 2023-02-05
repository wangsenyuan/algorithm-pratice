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
		blocks := readNNums(reader, n)
		res := solve(blocks)

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

func solve(blocks []int) []int {
	n := len(blocks)
	// 对于有同一color的两个block, 能够叠在同一个x的必要条件是
	// (i - j) % 2 = 1
	// 这个是不是充分条件呢？
	// yes, 如果它们距离 % 2 = 1, 他们就可以叠在一起
	dp := make([]int, n)
	pos := make([][]int, n)
	for i := 0; i < n; i++ {
		pos[i] = make([]int, 2)
		for j := 0; j < 2; j++ {
			pos[i][j] = -1
		}
	}

	for i := 0; i < n; i++ {
		dp[i] = 1
		color := blocks[i] - 1
		if pos[color][1-(i%2)] >= 0 {
			dp[i] = dp[pos[color][1-(i%2)]] + 1
		}
		pos[color][i%2] = i
	}

	res := make([]int, n)

	for i := 0; i < n; i++ {
		color := blocks[i] - 1
		res[color] = max(res[color], dp[i])
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
