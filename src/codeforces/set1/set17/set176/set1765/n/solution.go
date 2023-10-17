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
		s := readString(reader)
		k := readNum(reader)
		res := solve(s, k)
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
func solve(s string, k int) string {
	// 如果从i开始 s[i]是第一个数
	// s[i] != 0
	// i - 1 <= k
	// n - i >= k - (i - 1)
	// n - i + i - 1 >= k
	// n - 1 >= k  当 i > 0 时
	// 只要保证后面的数字能够被删完，就可以
	pos := make([][]int, 10)
	for i := 0; i < 10; i++ {
		pos[i] = make([]int, 0, 1)
	}
	n := len(s)
	for i := 1; i <= n; i++ {
		x := int(s[i-1] - '0')
		pos[x] = append(pos[x], i)
	}

	cur := make([]int, 10)

	buf := make([]byte, n-k)
	var prev int
	for it := 0; it < len(buf); it++ {
		if k == 0 {
			buf[it] = s[prev]
			prev++
			continue
		}
		var start int
		if it == 0 {
			start = 1
		}
		pick := -1
		for x := start; x <= 9; x++ {
			if cur[x] < len(pos[x]) {
				i := pos[x][cur[x]]
				// 对于位置i，首先后面要有足够的东西来删
				// n - i >= k
				// 但是它前面的部分肯定都被删除了
				if i-prev-1 <= k && n-i >= k-(i-prev-1) {
					pick = i
					break
				}
			}
		}
		buf[it] = s[pick-1]

		for x := 0; x <= 9; x++ {
			for cur[x] < len(pos[x]) && pos[x][cur[x]] <= pick {
				cur[x]++
			}
		}
		k -= (pick - prev - 1)
		prev = pick
	}

	return string(buf)
}
