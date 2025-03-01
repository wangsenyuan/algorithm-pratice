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
	for range tc {
		s := readString(reader)
		res := solve(s)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	buf.WriteTo(os.Stdout)
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

func solve(s string) int {
	s = trimPalindrome(s)
	if len(s) == 0 {
		return 0
	}

	n := len(s)
	// 要么重排前缀，要么重排后缀
	pref := make([][]int, n+1)
	for i := range n + 1 {
		pref[i] = make([]int, 26)
	}
	for i := 0; i < n; i++ {
		copy(pref[i+1], pref[i])
		x := int(s[i] - 'a')
		pref[i+1][x]++
	}
	// 始终可以？
	l := (n - 1) / 2
	r := l
	if n%2 == 0 {
		r++
	}
	l++
	r--
	for l > 0 && s[l-1] == s[r+1] {
		l--
		r++
	}

	get := func(l int, r int, cnt []int) {
		clear(cnt)
		if l <= r {
			for i := range 26 {
				cnt[i] = pref[r+1][i] - pref[l][i]
			}
		}
	}

	best := n
	cnt1 := make([]int, 26)
	cnt2 := make([]int, 26)
	for i := range n {
		// 如果只重拍0...i的话，是否得到回文？
		get(0, i, cnt1)
		if (i+1)*2 <= n {
			// 在前半段
			get(n-1-i, n-1, cnt2)
			// 如果两半部分相同
			if checkEquals(cnt1, cnt2) && i >= l-1 {
				best = i + 1
				break
			}
		} else {
			// 超过了一半
			get(i+1, n-1, cnt2)
			ok := true
			for j := range 26 {
				cnt1[j] -= cnt2[j]
				if cnt1[j] < 0 {
					ok = false
					break
				}
			}
			if ok && checkPalindrome(cnt1) == n&1 {
				best = i + 1
				break
			}
		}
	}

	for i := n - 1; i > 0; i-- {
		get(i, n-1, cnt1)
		d := n - i
		if d*2 <= n {
			// 在后半段
			get(0, d-1, cnt2)
			if checkEquals(cnt1, cnt2) && i <= r+1 {
				best = min(best, d)
				break
			}
		} else {
			get(0, i-1, cnt2)
			ok := true
			for j := range 26 {
				cnt1[j] -= cnt2[j]
				if cnt1[j] < 0 {
					ok = false
					break
				}
			}
			if ok && checkPalindrome(cnt1) == n&1 {
				best = min(best, d)
				break
			}
		}
	}

	return best
}

func checkEquals(a, b []int) bool {
	for i, x := range a {
		if b[i] != x {
			return false
		}
	}
	return true
}

func checkPalindrome(cnt []int) int {
	var odd int
	for i := range 26 {
		if cnt[i]%2 == 1 {
			odd++
		}
		if odd > 1 {
			return -1
		}
	}
	return odd
}

func trimPalindrome(s string) string {

	for i, j := 0, len(s)-1; i <= j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return s[i : j+1]
		}
	}
	return ""
}
