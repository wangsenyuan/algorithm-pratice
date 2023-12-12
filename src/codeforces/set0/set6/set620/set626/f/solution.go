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

	tc := 1

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, k := readTwoNums(reader)
		A := readNNums(reader, n)
		res := solve(A, k)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
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

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

const MOD = 1e9 + 7

func add(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func mul(a, b int) int {
	return int(int64(a) * int64(b) % MOD)
}

func solve(A []int, k int) int {
	// dp[i][k] 前i个元素，且不平衡度切好是k时的计数
	n := len(A)
	sort.Ints(A)
	// 一个组不需要时连续的sub array
	// 如果当前A[i]放入前面的group
	// 如果它做为最大值时, 如果前面有某个（未关闭的集合）最小值时a， 那么这时它的贡献就是A[i] - a
	// 但是中间的这些非边界的元素怎么分配？
	// 如果某个数做为左边界元素，那么它的共享时 -A[i],如果是右边界元素，它的贡献是A[i]
	// 不考虑其他的元素， 假设选中了m个元素，x1,x2....xm
	// -A[x1] + A[x2] - A[x3] + ... + A[xm]
	// dp[k][i] 表示当前和为k，且有i个未关闭group时的计数
	// 如果当前做为最大值 dp[k+A[j]][i-1]
	// 如果当前最为最小值 dp[k-A[j]][i+1]
	// 或者当前做为集合中的某个非端点 dp[k][i] * i
	dp := make([][][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([][]int, n+1)
		for j := 0; j <= n; j++ {
			dp[i][j] = make([]int, k+1)
		}
	}

	dp[0][1][0] = 1
	dp[0][0][0] = 1
	for i := 1; i < n; i++ {
		for j := 0; j <= n; j++ {
			for x := 0; x <= k; x++ {
				dp[i&1][j][x] = 0
			}
		}
		for j := 0; j <= i; j++ {
			t := j * (A[i] - A[i-1])
			for x := t; x <= k; x++ {
				// 不参与最大值、最小值(中间值),所以x和j都不变，有j个选择（j个还未关闭的group）+1，它单独做为一个group
				dp[i&1][j][x] = add(dp[i&1][j][x], mul(dp[(i-1)&1][j][x-t], j+1))
				if j < n {
					// A[i]做为一个新的最小值
					dp[i&1][j+1][x] = add(dp[i&1][j+1][x], dp[(i-1)&1][j][x-t])
				}
				if j > 0 {
					// A[i]做为最大值，有j个选择去匹配，并未关闭的group -1
					dp[i&1][j-1][x] = add(dp[i&1][j-1][x], mul(dp[(i-1)&1][j][x-t], j))
				}
			}
		}
	}

	var res int

	for x := 0; x <= k; x++ {
		res = add(res, dp[(n-1)&1][0][x])
	}

	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	return a + b - min(a, b)
}
