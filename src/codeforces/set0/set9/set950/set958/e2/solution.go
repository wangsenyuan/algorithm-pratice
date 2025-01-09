package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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
	k, n := readTwoNums(reader)
	a := readNNums(reader, n)
	return solve(a, k)
}

const inf = 1 << 60

func solve(a []int, k int) int {
	n := len(a)
	type pair struct {
		first  int
		second int
	}
	sort.Ints(a)
	arr := make([]pair, n-1)
	for i := range n - 1 {
		arr[i] = pair{a[i+1] - a[i], a[i]}
	}
	if len(arr) > 3*k {
		slices.SortFunc(arr, func(x, y pair) int {
			return x.first - y.first
		})
		arr = arr[:3*k]
	}
	slices.SortFunc(arr, func(x, y pair) int {
		return x.second - y.second
	})

	var b []int
	for i := 0; i < len(arr); i++ {
		if len(b) == 0 || b[len(b)-1] < arr[i].second {
			b = append(b, arr[i].second)
		}
		b = append(b, arr[i].second+arr[i].first)
	}

	n = len(b)

	dp := make([][]int, k+1)
	fp := make([][]int, k+1)
	for i := 0; i <= k; i++ {
		dp[i] = make([]int, 2)
		fp[i] = make([]int, 2)
		for j := 0; j < 2; j++ {
			dp[i][j] = inf
			fp[i][j] = inf
		}
	}
	// dp[j][0] 表示前j个，且第i个为被使用的情况
	// dp[j][1] 表示前j个，且第i个被使用的情况
	dp[0][0] = 0

	for i := 1; i < n; i++ {
		for j := 0; j <= k; j++ {
			fp[j][0] = min(dp[j][0], dp[j][1])
			if j > 0 {
				fp[j][1] = dp[j-1][0] + b[i] - b[i-1]
			}
		}
		for j := 0; j <= k; j++ {
			copy(dp[j], fp[j])
			fp[j][0] = inf
			fp[j][1] = inf
		}
	}

	return min(dp[k][0], dp[k][1])
}
