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

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, k := readTwoNums(reader)
		x := readNNums(reader, n)
		readNNums(reader, n)
		res := solve(x, k)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func solve(x []int, k int) int {
	// 一个观察是，平台的左端点可以某个point的x坐标重合
	// 如果 max(x) - min(x) > 2 * k, 平台最好是没有重合的
	//  max(x) - min(x) <= 2 *  k, 则全部的可以被save
	// 现在只考虑没有重合的情况，如果以i的左端点为起点，是可以计算出有多少落在它里面的
	// dp[i] 表示这个值， dp[i] = max(cur, dp[i+1])
	sort.Ints(x)
	n := len(x)
	if x[n-1]-x[0] <= 2*k {
		return n
	}

	dp := make([]int, n)
	for i := 0; i < n; i++ {
		j := sort.Search(i+1, func(j int) bool {
			return x[i]-k <= x[j]
		})
		dp[i] = i - j + 1
		if i > 0 {
			dp[i] = max(dp[i], dp[i-1])
		}
	}
	var cur int
	var ans int
	for i := n - 1; i > 0; i-- {
		j := sort.Search(n, func(j int) bool {
			return x[j]-k > x[i]
		})
		j--
		cur = max(cur, j-i+1)
		ans = max(ans, cur+dp[i-1])
	}
	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
