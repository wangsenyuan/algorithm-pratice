package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/bits"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	a := readNNums(reader, n)
	res := solve(a)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d ", x))
	}

	buf.WriteByte('\n')

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

const inf = 1 << 60

func solve(a []int) []int {
	n := len(a)
	if n == 1 {
		return []int{0}
	}

	h := (n + 1) / 2
	dp := make([][]int, h+1)
	fp := make([][]int, h+1)
	for i := range h + 1 {
		dp[i] = make([]int, 3)
		fp[i] = make([]int, 3)
		for j := range 3 {
			dp[i][j] = inf
			fp[i][j] = inf
		}
	}
	dp[0][0] = 0
	dp[1][2] = max(0, a[1]-(a[0]-1))
	dp[1][1] = max(0, a[0]-(a[1]-1))
	if n > 2 {
		dp[1][1] += max(0, a[2]-(a[1]-1))
	}

	for i := 2; i < n; i++ {
		for j := 0; j <= h; j++ {
			for state := 0; state < 3; state++ {
				// 不放置房子
				fp[j][(state<<1)&3] = min(fp[j][(state<<1)&3], dp[j][state])
				if state == 1 || j+1 > h || dp[j][state] == inf {
					// 上一个hill有房子时，不能在当前位置放置房子
					continue
				}
				cur := dp[j][state]
				if state == 2 {
					cur -= max(0, a[i-1]-(a[i-2]-1))
					cur += max(0, a[i-1]-min(a[i-2], a[i])+1)
				} else {
					// state == 0
					cur += max(0, a[i-1]-(a[i]-1))
				}
				if i+1 < n {
					cur += max(0, a[i+1]-(a[i]-1))
				}
				next := (state<<1)&3 | 1
				fp[j+1][next] = min(fp[j+1][next], cur)
			}
		}
		for j := 0; j <= h; j++ {
			for state := 0; state < 3; state++ {
				dp[j][state] = fp[j][state]
				fp[j][state] = inf
			}
		}
	}
	ans := make([]int, h)
	for i := 0; i < h; i++ {
		ans[i] = slices.Min(dp[i+1])
	}
	return ans
}

func solve1(a []int) []int {
	n := len(a)

	if n == 1 {
		return []int{0}
	}
	if n == 2 {
		if a[0] != a[1] {
			return []int{0}
		}
		// 如果要造一个房子，至少要弄低其中的一个
		return []int{1}
	}

	// n >= 3

	h := (n + 1) / 2
	dp := make([][]int, h+1)
	fp := make([][]int, h+1)
	for i := range h + 1 {
		dp[i] = make([]int, 8)
		fp[i] = make([]int, 8)
		for j := range 8 {
			dp[i][j] = inf
			fp[i][j] = inf
		}
	}

	// 先把前3个位置计算出来
	for state := 0; state < 8; state++ {
		cnt := bits.OnesCount(uint(state))
		if state&2 == 2 && state&5 > 0 {
			// 这个状态是无效的
			if cnt <= h {
				dp[cnt][state] = inf
			}
		} else {
			if state&2 == 2 {
				// 中间有房子，它需要在后面被处理
				dp[cnt][state] = max(0, a[0]-(a[1]-1)) + max(0, a[2]-(a[1]-1))
			} else if state == 5 {
				// 两头有房子
				dp[cnt][state] = max(0, a[1]-min(a[0], a[2])+1)
			} else if state == 4 {
				// 只需要中间的房子比第一个低
				dp[cnt][state] = max(0, a[1]-(a[0]-1))
			} else if state == 1 {
				// 只需要中间的房子比第三个低
				dp[cnt][state] = max(0, a[1]-(a[2]-1))
			} else {
				dp[cnt][state] = 0
			}
			if state&1 == 1 && n > 3 {
				dp[cnt][state] += max(0, a[3]-(a[2]-1))
			}
		}
	}

	for i := 3; i < n; i++ {
		for j := 0; j <= h; j++ {
			for state := 0; state < 8; state++ {
				// 往下一个状态迁移
				// 如果在当前位置不放置房子
				fp[j][(state<<1)&7] = min(fp[j][(state<<1)&7], dp[j][state])
				if state&1 == 1 || j+1 > h {
					// 上一个hill有房子时，不能在当前位置放置房子
					continue
				}
				// 如果在当前位置放置房子
				cur := dp[j][state]

				if state&2 == 2 {
					// 要把中间的贡献去掉
					cur -= max(0, a[i-1]-(a[i-2]-1))
				}
				next := (state<<1)&7 | 1
				if state&2 == 2 {
					// 前面的中间位置有房子, 所以要考虑a[i-2]位置处
					// 这个需要取消掉，否则就重复计算了
					cur += max(0, a[i-1]-min(a[i-2], a[i])+1)
				} else {
					// 只需要考虑a[i-1]位置处
					cur += max(0, a[i-1]-(a[i]-1))
				}
				if i+1 < n {
					cur += max(0, a[i+1]-(a[i]-1))
				}
				fp[j+1][next] = min(fp[j+1][next], cur)
			}
		}
		for j := 0; j <= h; j++ {
			for state := 0; state < 8; state++ {
				dp[j][state] = fp[j][state]
				fp[j][state] = inf
			}
		}
	}

	ans := make([]int, h)

	for i := 1; i <= h; i++ {
		ans[i-1] = inf
		for state := 0; state < 8; state++ {
			ans[i-1] = min(ans[i-1], dp[i][state])
		}
	}

	return ans
}
