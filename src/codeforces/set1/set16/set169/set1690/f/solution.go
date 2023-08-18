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
		p := readNNums(reader, n)
		res := solve(s, p)
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

func solve(s string, p []int) int {
	// 单个cycle，需要知道，经过多少次能变成初始值
	// 如果都不相同，那么就是n(cycle的长度）次
	// 如果都相同, 就是1次
	// aba => baa => aab => aba
	// abab => baba => abab
	// 在cycle内，找出循环节
	// 200 ~ 20 * 20 / 2
	// 最多是 2 * 3 * ... * 20
	n := len(s)
	for i := 0; i < n; i++ {
		p[i]--
	}
	pos := make([]int, n)
	for i := 0; i < n; i++ {
		pos[p[i]] = i
	}

	vis := make([]bool, n)

	var cycles []string

	buf := make([]byte, n)

	for i := 0; i < n; i++ {
		if vis[i] {
			continue
		}
		var it int
		j := i
		for !vis[j] {
			buf[it] = s[j]
			it++
			vis[j] = true
			j = pos[j]
		}
		cycles = append(cycles, string(buf[:it]))
	}

	ans := 1

	for _, cur := range cycles {
		tmp := getMinCycleLength(cur)
		ans = lcm(ans, tmp)
	}

	return ans
}

func lcm(a, b int) int {
	g := gcd(a, b)
	return a * b / g
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func getMinCycleLength(s string) int {
	n := len(s)

	for i := 1; i < n; i++ {
		if s == s[i:]+s[:i] {
			return i
		}
	}

	return n
}
