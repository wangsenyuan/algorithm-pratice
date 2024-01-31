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
		s := readString(reader)[:n]
		res := solve(s, k)
		s = fmt.Sprintf("%v", res)
		buf.WriteString(s[1 : len(s)-1])
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

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
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

func solve(s string, k int) []int {
	n := len(s)
	pref := make([][]int, n+1)
	suf := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		pref[i] = make([]int, n+1)
		suf[i] = make([]int, n+1)
	}

	for l := 0; l < n; l++ {
		var cnt1 int
		for r := l + 1; r <= n; r++ {
			cnt1 += int(s[r-1] - '0')
			pref[r][cnt1] = max(pref[r][cnt1], r-l)
			suf[l][cnt1] = max(suf[l][cnt1], r-l)
		}
	}

	for r := 0; r <= n; r++ {
		for cnt := 0; cnt <= n; cnt++ {
			if r > 0 {
				pref[r][cnt] = max(pref[r][cnt], pref[r-1][cnt])
			}
			if cnt > 0 {
				pref[r][cnt] = max(pref[r][cnt], pref[r][cnt-1])
			}
		}
	}

	for l := n; l >= 0; l-- {
		for cnt := 0; cnt <= n; cnt++ {
			if l+1 <= n {
				suf[l][cnt] = max(suf[l][cnt], suf[l+1][cnt])
			}
			if cnt > 0 {
				suf[l][cnt] = max(suf[l][cnt], suf[l][cnt-1])
			}
		}
	}
	// dp[i] 表示长度为i的l1，能够得到的最大的l0

	dp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = -(1 << 30)
	}

	for l := 0; l < n; l++ {
		var cnt0 int
		for r := l; r <= n; r++ {
			if r > l && s[r-1] == '0' {
				cnt0++
			}
			if cnt0 > k {
				break
			}
			dp[r-l] = max(dp[r-l], max(pref[l][k-cnt0], suf[r][k-cnt0]))
		}
	}

	ans := make([]int, n)
	for i := 0; i <= n; i++ {
		for a := 1; a <= n; a++ {
			ans[a-1] = max(ans[a-1], i+a*dp[i])
		}
	}

	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
