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
	for range tc {
		n, k := readTwoNums(reader)
		buf.WriteString(fmt.Sprintf("%d\n", solve(n, k)))
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

func solve(n int, k int) int {

	var dfs func(l int, r int) int

	var res int

	dfs = func(l int, r int) int {
		if r-l+1 < k {
			return 0
		}
		m := (l + r) / 2

		if (r-l+1)%2 == 0 {
			cnt := dfs(l, m)
			// (m1 + m2 + 1) / 2 = m
			res = (2*m + 1) * cnt
			return cnt * 2
		}
		cnt := dfs(l, m-1)
		res = m + 2*m*cnt
		return 2*cnt + 1
	}

	dfs(1, n)

	return res
}

func solve1(n int, k int) int {
	// 这里n的大小是1e9,
	// 所以不能真的全部都访问一遍
	// 找到第一个位置（这个比较容易）
	// 两个区间始终是相同长度的，且就是独立的
	// 所以可以直接 * 2
	cur := 1
	mul := n + 1
	var sum int
	for n >= k {
		if n&1 == 1 {
			sum += cur
		}
		n >>= 1
		cur <<= 1
	}

	return mul * sum / 2
}
