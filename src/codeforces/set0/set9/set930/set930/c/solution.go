package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	segs := make([][]int, n)
	for i := 0; i < n; i++ {
		segs[i] = readNNums(reader, 2)
	}
	res := solve(n, m, segs)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\r' || s[i] == '\n' {
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

func solve(n int, m int, segs [][]int) int {
	cnt := make([]int, m+2)
	for _, seg := range segs {
		l, r := seg[0], seg[1]
		cnt[l]++
		cnt[r+1]--
	}

	prev := make(fenwick, n+10)
	// 前面覆盖最多的点
	dp := make([]int, m+1)
	for i := 1; i <= m; i++ {
		cnt[i] += cnt[i-1]
		tmp := prev.pre(cnt[i]) + 1
		prev.update(cnt[i], tmp)
		dp[i] = tmp
	}
	suf := make(fenwick, n+10)
	var res int
	for i := m; i > 0; i-- {
		tmp := suf.pre(cnt[i])
		res = max(res, tmp+dp[i])
		suf.update(cnt[i], tmp+1)
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type fenwick []int

func (f fenwick) update(i int, val int) {
	for i++; i < len(f); i += i & -i {
		f[i] = max(f[i], val)
	}
}
func (f fenwick) pre(i int) (res int) {
	for i++; i > 0; i &= i - 1 {
		res = max(res, f[i])
	}
	return
}
