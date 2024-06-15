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
		n, l, r := readThreeNums(reader)
		s := readString(reader)[:n]
		res := solve(s, l, r)
		for i := 0; i < len(res); i++ {
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

func solve(s string, l int, r int) []int {
	// l == r
	n := len(s)

	z := zFunc(s)
	// k <= sqrt(n)时，以k为主，查f(k)
	// 当k >= sqrt(n)时，以f(k)为主，查k

	check := func(expect int) int {
		var cnt int = 1
		for i := expect; i < n; i++ {
			if z[i] >= expect {
				cnt++
				i += expect - 1
			}
		}
		return cnt
	}

	find := func(k int) int {
		l, r := 0, n
		for l < r {
			mid := (l + r + 1) / 2
			if check(mid) >= k {
				l = mid
			} else {
				r = mid - 1
			}
		}
		return r
	}

	ans := make([]int, n+1)

	for k := 1; k <= n/k; k++ {
		ans[k] = find(k)
	}

	for i := 1; i <= n/i; i++ {
		k := 1
		for j := i; j < n; j++ {
			if z[j] >= i {
				k++
				j += i - 1
			}
		}
		ans[k] = max(ans[k], i)
	}

	for i := n - 1; i >= 1; i-- {
		ans[i] = max(ans[i], ans[i+1])
	}

	return ans[l : r+1]
}

func zFunc(s string) []int {
	n := len(s)
	l, r := 0, 0
	z := make([]int, n)
	for i := 1; i < n; i++ {
		if i <= r {
			z[i] = min(r-i+1, z[i-l])
		}
		for i+z[i] < n && s[i+z[i]] == s[z[i]] {
			z[i]++
		}
		if i+z[i]-1 > r {
			r = i + z[i] - 1
			l = i
		}
	}
	return z
}
