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
		n, k := readTwoNums(reader)
		a := readString(reader)[:n]
		b := readString(reader)[:n]
		res := solve(a, b, k)
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

func solve(a string, b string, k int) int {
	// [l...r] 越长越好
	// [l...r] 长度为x  共x * (1 + x) / 2
	// 这个是个dp问题呐，不能贪心
	n := len(a)

	// 这个是个mask问题
	// 先找出所有a中的不同的字符
	// 假设有x个，然后是处理恰好有k个字符的情况
	// 然后暴力计算，把这k个字符改成和b中的一致，
	var buf []byte
	cnt := make([]int, 26)
	for i := 0; i < n; i++ {
		x := int(a[i] - 'a')
		cnt[x]++
		if cnt[x] == 1 {
			buf = append(buf, a[i])
		}
	}

	if len(buf) <= k {
		// all change
		return n * (n + 1) / 2
	}
	// len(buf) > k
	m := len(buf)
	pos := make([]int, 26)
	for i := 0; i < m; i++ {
		pos[int(buf[i]-'a')] = i
	}

	inSet := func(c byte, state int) bool {
		x := pos[int(c-'a')]
		return (state>>x)&1 == 1
	}

	check := func(state int) int {
		// 直到遇到一个不一样的字符
		var cur int
		for i := 0; i < n; {
			j := i
			for i < n && (a[i] == b[i] || inSet(a[i], state)) {
				i++
			}
			if i == j {
				i++
				continue
			}
			ln := i - j
			cur += ln * (ln + 1) / 2
		}
		return cur
	}

	var ans int
	for state := 0; state < (1 << m); state++ {
		dc := digitCount(state)
		if dc == k {
			ans = max(ans, check(state))
		}
	}

	return ans
}

func digitCount(num int) int {
	var res int
	for num > 0 {
		res++
		num -= num & -num
	}
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func search(n int, f func(i int) bool) int {
	l, r := 0, n
	for l < r {
		mid := (l + r) / 2
		if f(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}
