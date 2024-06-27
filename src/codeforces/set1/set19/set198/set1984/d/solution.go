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

type Pair struct {
	first  int
	second int
}

func solve(s string) int {
	n := len(s)
	var other []Pair
	for i := 0; i < n; i++ {
		if s[i] != 'a' {
			other = append(other, Pair{int(s[i]), i})
		}
	}

	// all 'a'
	if len(other) == 0 {
		return n - 1
	}

	m := len(other)

	t := make([]byte, m)

	for i := 0; i < m; i++ {
		t[i] = byte(other[i].first)
	}

	check := func(d int) bool {
		for i := 0; i < m; i++ {
			if t[i] != t[i%d] {
				return false
			}
		}
		return true
	}

	sum := make([]int, n+1)
	for i := 0; i < n; i++ {
		sum[i+1] = sum[i]
		if s[i] != 'a' {
			sum[i+1]++
		}
	}

	process := func(d int) int {
		x := s[other[0].second : other[d-1].second+1]
		p := kmp(x)
		var cnt []int
		prev := 0
		for i, j := 0, 0; i < n; i++ {
			for j > 0 && s[i] != x[j] {
				j = p[j-1]
			}
			if s[i] == x[j] {
				j++
			}
			if j == len(x) {
				// found one
				if sum[i+1-len(x)] != sum[prev] {
					return 0
				}

				cnt = append(cnt, i+1-len(x)-prev)
				prev = i + 1
				j = 0
			}
		}

		if sum[n] != sum[prev] {
			return 0
		}
		cnt = append(cnt, n-prev)
		lc := len(cnt)
		first := cnt[0]
		last := cnt[lc-1]
		if lc == 2 {
			// 至少有两个
			return (first + 1) * (last + 1)
		}

		mid := cnt[1]
		for i := 1; i+1 < lc; i++ {
			mid = min(mid, cnt[i])
		}
		// u + v <= mid
		// u <= first
		// v <= last
		// 这个计数是多少呢？
		var res int
		for u := 0; u <= first && u <= mid; u++ {
			v := min(last, mid-u)
			res += (v + 1)
		}
		return res
	}

	var ans int

	for d := 1; d <= m/d; d++ {
		if m%d == 0 {
			if check(d) {
				// other[:d] is a proper partition of other
				ans += process(d)
			}
			if m != d*d && check(m/d) {
				ans += process(m / d)
			}
		}
	}

	return ans
}

func kmp(s string) []int {
	n := len(s)
	p := make([]int, n)
	for i := 1; i < n; i++ {
		j := p[i-1]
		for j > 0 && s[i] != s[j] {
			j = p[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		p[i] = j
	}
	return p
}
