package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, m := readTwoNums(reader)
	s := readString(reader)
	queries := make([][]int, m)
	for i := 0; i < m; i++ {
		queries[i] = readNNums(reader, 2)
	}
	res := solve(s, queries)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}
	fmt.Print(buf.String())
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

func solve(s string, queries [][]int) []int {
	// abcabc
	n := len(s)
	// dp[i][0] = 在第i位开始, [i, i + 3, i + 6, ...] 为一组，这组里面为a的计数
	dp := make([][]int, n+5)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 3)
	}

	for i := n - 1; i >= 0; i-- {
		copy(dp[i], dp[i+3])
		dp[i][int(s[i]-'a')]++
	}

	check := func(l int, r int, a, b, c int) int {
		ln := r - l + 1
		cnt := []int{ln / 3, ln / 3, ln / 3}
		if ln%3 > 0 {
			cnt[0]++
		}
		if ln%3 > 1 {
			cnt[1]++
		}
		res := cnt[0] - (dp[l][a] - dp[l+cnt[0]*3][a])
		res += cnt[1] - (dp[l+1][b] - dp[l+1+cnt[1]*3][b])
		res += cnt[2] - (dp[l+2][c] - dp[l+2+cnt[2]*3][c])

		return res
	}

	ans := make([]int, len(queries))

	for i, cur := range queries {
		l, r := cur[0]-1, cur[1]-1
		if l == r {
			continue
		}
		if r-l+1 <= 2 {
			if s[l] == s[r] {
				ans[i] = 1
			}
			continue
		}
		ans[i] = r - l + 1
		for a := 0; a < 3; a++ {
			for b := 0; b < 3; b++ {
				if a == b {
					continue
				}
				c := 3 - a - b
				ans[i] = min(ans[i], check(l, r, a, b, c))
			}
		}
	}
	return ans
}
