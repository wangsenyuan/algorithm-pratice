package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	words := make([]string, n)
	for i := 0; i < n; i++ {
		words[i] = readString(reader)
	}

	res := solve(words)

	fmt.Println(res)
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
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

const oo = 1 << 60

func solve(words []string) int {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) > len(words[j])
	})
	n := len(words)
	next := make([][]int, n)
	for i := 0; i < n; i++ {
		next[i] = kmp(words[i])
	}
	var rem int
	var m int
	for i := 0; i < n; i++ {
		if (rem>>i)&1 == 1 {
			continue
		}
		words[m] = words[i]
		next[m] = next[i]
		m++
		for j := i + 1; j < n; j++ {
			ok, _ := check(words[i], words[j], next[j])
			if ok {
				rem |= 1 << j
			}
		}
	}

	n = m
	fp := make([][]int, n)
	for i := 0; i < n; i++ {
		fp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			_, w := check(words[i], words[j], next[j])
			fp[i][j] = len(words[j]) - w[len(words[i])-1]
		}
	}
	T := 1 << n
	dp := make([][]int, T)

	for i := 0; i < T; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = oo
		}
	}

	for i := 0; i < n; i++ {
		dp[1<<i][i] = len(words[i])
	}

	for state := 1; state < T; state++ {
		for i := 0; i < n; i++ {
			if (state>>i)&1 == 1 {
				for j := 0; j < n; j++ {
					if (state>>j)&1 == 0 {
						tmp := state | (1 << j)
						dp[tmp][j] = min(dp[tmp][j], dp[state][i]+fp[i][j])
					}
				}
			}
		}
	}

	res := oo
	for i := 0; i < n; i++ {
		res = min(res, dp[T-1][i])
	}
	return res
}

func check(s string, p string, next []int) (ok bool, w []int) {
	w = make([]int, len(s))
	for i, j := 0, 0; i < len(s); i++ {
		for j > 0 && s[i] != p[j] {
			j = next[j-1]
		}
		if s[i] == p[j] {
			j++
		}
		w[i] = j
		if j == len(p) {
			ok = true
			j = next[j-1]
		}
	}
	return
}

func kmp(pattern string) []int {
	n := len(pattern)

	lps := make([]int, n)

	for i := 1; i < n; i++ {
		j := lps[i-1]
		for j > 0 && pattern[i] != pattern[j] {
			j = lps[j-1]
		}
		lps[i] = j
		if pattern[i] == pattern[j] {
			lps[i]++
		}
	}
	return lps
}

func solve1(words []string) int {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) > len(words[j])
	})
	n := len(words)

	rem := make([]bool, n)
	var m int
	var ln int
	for i := 0; i < n; i++ {
		if rem[i] {
			continue
		}
		words[m] = words[i]
		m++
		ln += len(words[i])
		for j := i + 1; j < n; j++ {
			if contains(words[i], words[j]) {
				rem[j] = true
			}
		}
	}
	n = m
	d := make([][]int, n)

	for i := 0; i < n; i++ {
		d[i] = make([]int, n)
	}

	buf := make([]byte, ln)
	for i := 0; i < n; i++ {
		copy(buf, words[i])
		p := len(words[i])
		for j := 0; j < n; j++ {
			if j != i {
				copy(buf[p:], words[j])
				p += len(words[j])
			}
		}
		cur := len(words[i])
		za := zFunction(string(buf[:p]))
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			d[j][i] = len(words[i])
			for k := 0; k < len(words[j]); k++ {
				if za[cur+k] >= len(words[j])-k {
					d[j][i] = k + len(words[i]) - len(words[j])
					break
				}
			}
			cur += len(words[j])
		}
	}
	T := 1 << n
	dp := make([][]int, T)
	for i := 0; i < T; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = oo
		}
	}

	for i := 0; i < n; i++ {
		dp[1<<i][i] = len(words[i])
	}
	for state := 1; state < T-1; state++ {
		for i := 0; i < n; i++ {
			if (state>>i)&1 == 0 {
				continue
			}
			for j := 0; j < n; j++ {
				if (state>>j)&1 == 1 {
					continue
				}
				dp[state|(1<<j)][j] = min(dp[state|(1<<j)][j], dp[state][i]+d[i][j])
			}
		}
	}
	res := oo
	for i := 0; i < n; i++ {
		res = min(res, dp[T-1][i])
	}
	return res
}

func contains(s, x string) bool {
	z := zFunction(x + s)
	for i := len(x); i < len(x)+len(s); i++ {
		if z[i] >= len(x) {
			return true
		}
	}
	return false
}

func zFunction(s string) []int {
	n := len(s)
	z := make([]int, n)
	for i, l, r := 1, 0, 0; i < n; i++ {
		if i <= r {
			z[i] = min(r-i+1, z[i-l])
		}
		for i+z[i] < n && s[z[i]] == s[i+z[i]] {
			z[i]++
		}
		if i+z[i]-1 > r {
			l = i
			r = i + z[i] - 1
		}
	}
	return z
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
