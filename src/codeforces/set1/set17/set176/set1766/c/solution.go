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
		wall := make([]string, 2)
		wall[0] = readString(reader)[:n]
		wall[1] = readString(reader)[:n]
		res := solve(wall)

		if res {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
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

func solve(walls []string) bool {
	m := len(walls[0])
	// 如果有断开的，答案为false
	// 假设都是连在一起的
	// 两个单个的b中间的个数
	// 如果是偶数，这两个B要在同一行，如果是奇数，要在不同行
	// 前面刷子在哪一行

	findBlackCell := func(i int) int {
		if walls[0][i] == 'B' {
			return 0
		}
		return 1
	}

	check := func(j int, i int) bool {
		cnt := i - j - 1
		x := findBlackCell(j)
		y := findBlackCell(i)
		if cnt&1 == 0 {
			return x == y
		}
		return x != y
	}

	var i int
	for i < m && walls[0][i] == walls[1][i] {
		i++
	}

	for i < m {
		// j is single
		j := i
		i++
		for i < m && walls[0][i] == walls[1][i] {
			// 只能都是black
			i++
		}
		if i < m && !check(j, i) {
			return false
		}
	}
	return true
}
