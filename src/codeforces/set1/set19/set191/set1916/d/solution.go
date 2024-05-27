package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(reader)
		res := solve(n)
		for _, x := range res {
			buf.WriteString(x)
			buf.WriteByte('\n')
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

var pre []string

func init() {
	pre = bruteForce(11)
}

func solve(n int) []string {
	if n == 1 {
		return []string{"1"}
	}
	if n == 3 {
		return []string{"169", "196", "961"}
	}

	if n < 11 {
		return bruteForce(n)[:n]
	}
	// n > 11
	var res []string
	for i := 0; i < n; i++ {
		res = append(res, appendZeros(pre[i], n-11))
	}
	return res
}

func bruteForce(n int) []string {
	// n <= 11
	// 找到最多的有相同digits的number
	occ := make(map[string][]int)
	// n is odd
	h := (n + 1) / 2
	// h = 1, num = 1
	// h = 2, num = 10
	// h = 3, num = 100
	num := 1
	for i := 0; i < h-1; i++ {
		num *= 10
	}

	lmt := num * 10
	num++
	for num < lmt {
		sq := num * num
		rs := rep(sq)
		if len(rs) == n {
			occ[rs] = append(occ[rs], sq)
		}
		num++
	}

	var ans []int

	for _, v := range occ {
		if len(v) > len(ans) {
			ans = v
		}
	}

	res := make([]string, len(ans))
	for i, num := range ans {
		res[i] = fmt.Sprintf("%d", num)
	}

	return res
}

func rep(num int) string {
	s := fmt.Sprintf("%d", num)
	buf := []byte(s)
	sort.Slice(buf, func(i, j int) bool {
		return buf[i] < buf[j]
	})
	return string(buf)
}

func appendZeros(s string, cnt int) string {
	buf := make([]byte, len(s)+cnt)
	copy(buf, s)
	for i := len(s); i < len(buf); i++ {
		buf[i] = '0'
	}
	return string(buf)
}
