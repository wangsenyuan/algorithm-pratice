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

	readNum(reader)
	S := readString(reader)
	C := readNNums(reader, 26)
	n := readNum(reader)
	P := make([]string, n)
	for i := 0; i < n; i++ {
		P[i] = readString(reader)
	}

	res := solve(P, S, C)

	var buf bytes.Buffer

	for _, cur := range res {
		if cur {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
		}
	}

	fmt.Print(buf.String())
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func solve(P []string, S string, C []int) []bool {
	sort.Strings(P)

	n := len(P)

	cnt := make([]int, 26)

	check_last := func(p string, s string) bool {
		m := len(s)
		// s is a prefix of p
		for i := m; i < len(p); i++ {
			x := int(p[i] - 'a')
			for j := 25; j > x; j-- {
				if cnt[j] > 0 {
					return true
				}
			}
			if cnt[x] == 0 {
				return false
			}
			cnt[x]--
		}
		// p == s
		for j := 0; j < 25; j++ {
			if cnt[j] > 0 {
				return true
			}
		}
		return false
	}

	check := func(i int, s string) bool {
		// now append
		if i+1 == n {
			// p is the last one, only check whether can make s to be larger than p
			return check_last(P[i], s)
		}
		m := len(s)
		// s is a prefix of P[i], s < P[i+1],
		// 要么 some s[j] < P[i+1][j], or len(s) < len(P[i+1]) and s is prefix of P[i+1]

		for j := 0; j < m; j++ {
			if s[j] < P[i+1][j] {
				return check_last(P[i], s)
			}
		}
		var buf bytes.Buffer
		buf.WriteString(s)
		// s is both a prefix of P[i], P[i+1]
		for j := m; j < len(P[i]) && j < len(P[i+1]); j++ {
			// s is prefix of both P[i] & P[i+1], now we meet a P[i][j] < P[i][j+1]
			if P[i][j] < P[i+1][j] {
				for x := int(P[i][j]-'a') + 1; x <= int(P[i+1][j]-'a'); x++ {
					if cnt[x] > 0 {
						return true
					}
				}
				if cnt[int(P[i][j]-'a')] == 0 {
					return false
				}
				cnt[int(P[i][j]-'a')]--
				return check_last(P[i], buf.String())
			}
			// P[i][j] == P[i+1][j]
			if cnt[int(P[i][j]-'a')] == 0 {
				return false
			}
			cnt[int(P[i][j]-'a')]--
			buf.WriteByte(P[i][j])
		}
		// buf == P[i] && P[i] is a prefix of P[i+1]
		hi := int(P[i+1][buf.Len()] - 'a')

		for j := 0; j <= hi; j++ {
			if cnt[j] > 0 {
				return true
			}
		}

		return false
	}
	ans := make([]bool, n+1)

	ans[0] = S <= P[0]

	for i := 0; i < n; i++ {
		if i+1 < n && (S > P[i+1] || P[i] == P[i+1]) {
			continue
		}
		// S <= P[i+1] && P[i] < P[i+1]
		if S > P[i] {
			ans[i+1] = true
			continue
		}
		// S <= P[i] < P[i+1]
		if len(S) > len(P[i]) || S != P[i][:len(S)] {
			ans[i+1] = false
			continue
		}
		// len(S) <= len(P) and S is prefix of P[i]
		copy(cnt, C)
		ans[i+1] = check(i, S)
	}

	return ans
}
