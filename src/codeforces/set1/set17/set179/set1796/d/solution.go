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

	for tc > 0 {
		tc--
		n, k, x := readThreeNums(reader)
		a := readNNums(reader, n)
		res := solve(a, k, x)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

const inf = 1 << 60

func solve(a []int, k int, x int) int {
	// dp[i][j] 表示到i为止，且修改次数位j次时的最大值？
	// 似乎有点撞墙
	// 首先，观察一下
	// 1， 修改k个元素，是不是连续的？
	//    假设不是连续的,存在某个i，它的两边都有修改，它没有修改 （ k > 1)的时候
	//   如果最终的结果里面时包括这个i所在的一段的,
	//    并假设x > 0, 因为x > 0 , 所以，它两边修改的位置，肯定也时包含
	//    如果不包含，通过调整，让i增加x，会得到更优的结果
	//    假设 x < 0, 那么至少这段是不能包括修改的部分的，并且+x的部分，最好集中在两头
	// 分两种情况，一种是 x >= 0, 另外一种 x < 0
	// x < 0 的，修改两头的值(0....k) (k....0)后判断最大值即可
	// x >= 0, 就是修改连续的一段，然后取最大值
	// 两种状态都要处理
	var ans = solve1(a, k, x)
	return max(ans, solve2(a, k, x))
}

func solve2(a []int, k int, x int) int {
	//x  < 0
	n := len(a)
	var best int = -inf
	for l := 0; l <= k; l++ {
		r := k - l
		for i := 0; i < n; i++ {
			if i < l || i >= n-r {
				a[i] += x
			} else {
				a[i] -= x
			}
		}
		best = max(best, getBestSum(a))

		for i := 0; i < n; i++ {
			if i < l || i >= n-r {
				a[i] -= x
			} else {
				a[i] += x
			}
		}
	}

	return best
}

func getBestSum(a []int) int {
	var res int
	var sum int
	for _, num := range a {
		sum += num
		res = max(res, sum)
		if sum < 0 {
			sum = 0
		}
	}
	return res
}

func solve1(a []int, k int, x int) int {
	// x >= 0
	// dp[i] = 不修改时，以i结尾且连续的最大值
	// 这个有问题呐， 修改了以i结尾的k个连续的值，其他的部分要变成-x
	// 还需要一个fp
	n := len(a)
	// k <= n
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			dp[i] = a[i] - x
		} else {
			dp[i] = max(dp[i-1]+a[i]-x, a[i]-x)
		}
	}

	fp := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		if i == n-1 {
			fp[i] = a[i] - x
		} else {
			fp[i] = max(fp[i+1]+a[i]-x, a[i]-x)
		}
	}

	prev := make([]int, k+1)
	var best int
	for i := k - 1; i < n; i++ {
		// 修改 i-k...i
		if i-k < 0 {
			prev[0] = 0
		} else {
			prev[0] = max(dp[i-k], 0)
		}
		for j := 0; j < k; j++ {
			// dp[i-k] 可能 < 0
			prev[j+1] = max(prev[j], 0) + a[i-k+1+j] + x
		}
		suf := fp[i+1]
		for j := k - 1; j >= 0; j-- {
			suf = max(suf, 0) + a[i-k+1+j] + x
			best = max(best, prev[j]+suf)
		}
	}

	return best
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
