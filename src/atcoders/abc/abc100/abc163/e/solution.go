package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
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
	n := readNum(reader)
	a := readNNums(reader, n)
	return solve(a)
}

const inf = 1 << 60

func solve(a []int) int {
	// n <= 2000
	n := len(a)

	type pair struct {
		first  int
		second int
	}

	arr := make([]pair, n)
	for i := 0; i < n; i++ {
		arr[i] = pair{a[i], i}
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i].first > arr[j].first || arr[i].first == arr[j].first && arr[i].second < arr[j].second
	})

	// 本质上，是把大的数，从两头到中间放置
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			dp[i][j] = -inf
		}
	}

	dp[0][0] = 0

	for l := 0; l < n; l++ {
		for i := l; i < n; i++ {
			r := i - l
			cur := arr[i]
			p := cur.second
			v := cur.first
			// 把第i个数放在位置l+1, 或者r-1处
			dp[l+1][r] = max(dp[l+1][r], dp[l][r]+v*(p-l))
			dp[l][r+1] = max(dp[l][r+1], dp[l][r]+v*((n-(r+1))-p))
		}

	}

	var res int
	for i := 0; i <= n; i++ {
		res = max(res, dp[i][n-i])
	}
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
