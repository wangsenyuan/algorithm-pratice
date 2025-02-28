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
		res := process(reader)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func process(reader *bufio.Reader) int {
	n, m := readTwoNums(reader)
	a := readNNums(reader, n)
	b := readNNums(reader, m)
	return solve(a, b)
}

const inf = 1 << 60

func solve(a []int, b []int) int {
	n := len(a)
	m := len(b)
	sum := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		sum[i] = sum[i+1] + a[i]
	}

	dp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = inf
	}
	que := make([]int, n+1)
	for k := m; k > 0; k-- {
		var head, tail int
		for i := n - 1; i >= 0; i-- {
			for tail < head && sum[i]-sum[que[tail]] > b[k-1] {
				tail++
			}
			if tail < head {
				dp[i] = min(dp[i], dp[que[tail]]+m-k)
			}
			if sum[i] <= b[k-1] {
				// 可以直接到结尾
				dp[i] = min(dp[i], m-k)
			}
			for head > tail && dp[i] <= dp[que[head-1]] {
				// i离前面的数更近，所以sum也越小
				head--
			}
			que[head] = i
			head++
		}
	}
	if dp[0] == inf {
		return -1
	}
	return dp[0]
}
