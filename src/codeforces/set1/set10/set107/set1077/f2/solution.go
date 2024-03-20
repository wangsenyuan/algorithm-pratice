package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, k, x := readThreeNums(reader)

	a := readNNums(reader, n)

	res := solve(a, k, x)

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

const oo = 1 << 62

func solve(a []int, k int, x int) int {
	n := len(a)
	// dp[i]表示在times之前，能够在i处获得的最大值
	dp := make([]int, n+1)

	for i := 0; i <= n; i++ {
		dp[i] = -oo
	}

	que := make([]int, n+1)

	for times := 0; times < x; times++ {
		var front, back int
		for l, r := n, n-1; r >= times; r-- {
			for r-l+1 <= k && l > 0 {
				l--
				// 要把l放到队列里面去
				for back > front && dp[l] > dp[que[back-1]] {
					back--
				}
				// dp[l] < dp[que[back-1]]
				que[back] = l
				back++
			}

			if que[front] == r {
				front++
			}

			if front < back {
				// 队列头部是最大的值
				dp[r] = max(dp[r], dp[que[front]]+a[r])
			}
			// 足够从0开始
			if r < k {
				dp[r] = max(dp[r], a[r])
			}
		}
	}

	if dp[n-k] < 0 {
		return -1
	}

	res := dp[n-k]
	for i := n - k; i < n; i++ {
		res = max(res, dp[i])
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
