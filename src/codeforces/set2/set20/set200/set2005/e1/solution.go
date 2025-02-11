package main

import (
	"bufio"
	"bytes"
	"os"
)

func main() {
	var buf bytes.Buffer
	reader := bufio.NewReader(os.Stdin)
	tc := readNum(reader)
	for range tc {
		res := process(reader)
		buf.WriteString(res)
		buf.WriteByte('\n')
	}
	buf.WriteTo(os.Stdout)
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

func process(reader *bufio.Reader) string {
	l, n, m := readThreeNums(reader)
	a := readNNums(reader, l)
	b := make([][]int, n)
	for i := range n {
		b[i] = readNNums(reader, m)
	}
	return solve(a, b)
}

type pair struct {
	first  int
	second int
}

func solve(a []int, b [][]int) string {
	pos := make([][]pair, 8)
	n := len(b)
	m := len(b[0])
	for i, row := range b {
		for j, x := range row {
			pos[x] = append(pos[x], pair{i, j})
		}
	}
	sum := make([][]int, n+1)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
		sum[i] = make([]int, m+1)
	}

	k := len(a)

	for u := k - 1; u >= 0; u-- {
		x := a[u]
		// y := a[u+1]
		for _, cur := range pos[x] {
			r, c := cur.first, cur.second
			if sum[r+1][c+1] == 0 {
				// 这个位置是一个获胜位置, 因为在这个子区域，全部是失败的位置
				dp[r][c] = 1
			} else {
				dp[r][c] = 0
			}
		}
		for i := n - 1; i >= 0; i-- {
			for j := m - 1; j >= 0; j-- {
				sum[i][j] = dp[i][j]
				sum[i][j] += sum[i+1][j]
				sum[i][j] += sum[i][j+1]
				sum[i][j] -= sum[i+1][j+1]
				dp[i][j] = 0
			}
		}
	}
	if sum[0][0] > 0 {
		return "T"
	}
	return "N"
}
