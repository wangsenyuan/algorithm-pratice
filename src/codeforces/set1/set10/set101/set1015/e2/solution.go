package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	ok, res, _ := process(reader)

	if !ok {
		fmt.Println(-1)
		return
	}
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	for _, cur := range res {
		buf.WriteString(fmt.Sprintf("%d %d %d\n", cur[0], cur[1], cur[2]))
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

func process(reader *bufio.Reader) (ok bool, res [][]int, a []string) {
	n, m := readTwoNums(reader)
	a = make([]string, n)
	for i := range n {
		a[i] = readString(reader)[:m]
	}
	ok, res = solve(a)
	return
}

func solve(a []string) (bool, [][]int) {
	n := len(a)
	m := len(a[0])
	cnt := make([][]int, n)
	for i := 0; i < n; i++ {
		cnt[i] = make([]int, m)
	}

	dp := make([][][4]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][4]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if a[i][j] == '.' {
				continue
			}
			dp[i][j][0] = 1
			if i > 0 {
				dp[i][j][0] += dp[i-1][j][0]
			}
			dp[i][j][1] = 1
			if j > 0 {
				dp[i][j][1] += dp[i][j-1][1]
			}
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if a[i][j] == '.' {
				continue
			}
			dp[i][j][2] = 1
			if i+1 < n {
				dp[i][j][2] += dp[i+1][j][2]
			}
			dp[i][j][3] = 1
			if j+1 < m {
				dp[i][j][3] += dp[i][j+1][3]
			}
		}
	}

	add := func(r1 int, c1 int, r2 int, c2 int) {
		cnt[r1][c1]++
		if r2+1 < n {
			cnt[r2+1][c1]--
		}
		if c2+1 < m {
			cnt[r1][c2+1]--
		}
		if r2+1 < n && c2+1 < m {
			cnt[r2+1][c2+1]++
		}
	}

	var res [][]int

	process := func(i int, j int) {
		r := min(dp[i][j][0], dp[i][j][1], dp[i][j][2], dp[i][j][3])
		if r == 1 {
			return
		}
		res = append(res, []int{i + 1, j + 1, r - 1})
		add(i-r+1, j, i+r-1, j)
		add(i, j-r+1, i, j+r-1)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if a[i][j] == '*' {
				process(i, j)
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i > 0 {
				cnt[i][j] += cnt[i-1][j]
			}
			if j > 0 {
				cnt[i][j] += cnt[i][j-1]
			}
			if i > 0 && j > 0 {
				cnt[i][j] -= cnt[i-1][j-1]
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if a[i][j] == '*' && cnt[i][j] == 0 {
				return false, nil
			}
			if a[i][j] == '.' && cnt[i][j] > 0 {
				return false, nil
			}
		}
	}

	return true, res
}
