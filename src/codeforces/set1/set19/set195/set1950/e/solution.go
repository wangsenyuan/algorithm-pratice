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
		readNum(reader)
		s := readString(reader)
		res := solve(s)
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

func solve(s string) int {
	n := len(s)
	flag := make([]bool, n)
	check := func(k int) bool {
		for i := 0; i < k; i++ {
			flag[i] = false
		}
		for i := 0; i < n; i++ {
			if s[i] != s[i%k] {
				flag[i%k] = true
			}
		}
		diff := -1
		for i := 0; i < k; i++ {
			if flag[i] {
				if diff >= 0 {
					return false
				}
				diff = i
			}
		}
		if diff < 0 {
			return true
		}
		if n/k == 2 {
			return true
		}

		var expect byte

		if s[diff] == s[diff+k] || s[diff] == s[diff+2*k] {
			// s[diff]是多的那方
			expect = s[diff]
		} else {
			expect = s[diff+k]
		}
		var cnt int
		for i := diff; i < n; i += k {
			if s[i] != expect {
				cnt++
			}
		}
		return cnt == 1
	}
	res := n
	for k := 1; k <= n/k; k++ {
		if n%k == 0 {
			if check(k) {
				res = min(res, k)
			}
			if k*k != n && check(n/k) {
				res = min(res, n/k)
			}
		}
	}
	return res
}
