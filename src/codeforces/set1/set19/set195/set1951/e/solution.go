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
		res := solve(s)

		if len(res) == 0 {
			buf.WriteString("NO\n")
		} else {
			buf.WriteString(fmt.Sprintf("YES\n%d\n", len(res)))
			for _, x := range res {
				buf.WriteString(x)
				buf.WriteByte(' ')
			}
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

func solve(s string) []string {
	p := manacher(s)

	n := len(s)

	if p[n-1] < n {
		return []string{s}
	}
	// 要么就是两端的情况
	// 要么就是三段的情况？
	// 假设存在4段的情况, a, b, c, d它们都不是回文
	// a和b如果组成了回文，那么a,b,c肯定不是回文
	// 如果 a + b 是偶数， 那么 a + b + c 也是回文 =》 c也是回文
	// 那么如何处理第三种情况呢？
	// 只能迭代 r, 然后检查 p[l+r] < r - l + 1
	// 且p[0+l-1] < l
	// 假设第二步不存在， 那么必然存在一个位置i, [0...i-1] [i...n-1]分别是回文
	// 且整个字符串是回文，这里不妨设前半部分更短
	// 是不是中间的肯定是 2/3个长的？
	// 假设不存在长度为2/3的非回文串，但是存在长度更大的回文串
	// 这是不可能的，因为不存在2的 =》 大家都一样，不存在3的，那么必然是010101
	for i := 0; i < n; i++ {
		if p[0+i] < i+1 && p[i+1+n-1] < n-1-i {
			return []string{s[:i+1], s[i+1:]}
		}
	}

	for i := 0; i < n; i++ {
		if p[0+i] < i+1 {
			// 前缀不是回文
			for l := 2; i+l+1 < n && l <= 4; l++ {
				if p[i+l+1+n-1] < n-1-i-l {
					// 后缀不是回文
					if p[i+1+i+l] < l {
						// 中间也不是回文
						return []string{s[:i+1], s[i+1 : i+l+1], s[i+l+1:]}
					}
				}
			}
		}
	}

	return nil
}

func manacher(s string) []int {
	buf := make([]byte, 2*len(s)+1)
	for i := 0; i < len(s); i++ {
		buf[i*2] = '#'
		buf[i*2+1] = s[i]
	}
	buf[2*len(s)] = '#'

	v := manacher_odd(string(buf))

	return v[1 : len(v)-1]
}

func manacher_odd(s string) []int {
	s = "&" + s + "^"
	n := len(s)
	p := make([]int, n)

	l, r := 1, 1

	for i := 1; i < n-1; i++ {
		p[i] = max(0, min(r-i, p[l+(r-i)]))
		for s[i-p[i]] == s[i+p[i]] {
			p[i]++
		}
		if i+p[i] > r {
			l = i - p[i]
			r = i + p[i]
		}
	}

	return p[1 : n-1]
}
