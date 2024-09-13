package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	a := readNNums(reader, n)

	res := solve(a)

	fmt.Println(res)
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadBytes('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return string(s[:i])
		}
	}
	return string(s)
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

func solve(a []int) int {
	n := len(a)
	vis := make([]bool, n+1)
	for _, i := range a {
		if i > 0 {
			vis[i] = true
		}
	}

	cnt := []int{0, 0}

	for i := 1; i <= n; i++ {
		if !vis[i] {
			cnt[i&1]++
		}
	}

	if cnt[0] == 0 {
		// 没有空位
		return calc(a)
	}

	// dp[i][0/1]表示上一个parity是0/1时，且使用了i个0时的最小值
	dp := make([][]int, cnt[0]+2)
	fp := make([][]int, cnt[0]+2)
	for i := 0; i <= cnt[0]+1; i++ {
		dp[i] = make([]int, 2)
		fp[i] = make([]int, 2)
		for j := 0; j < 2; j++ {
			dp[i][j] = n
			fp[i][j] = n
		}
	}
	if a[0] == 0 {
		// 第一位放置一个1时，没有消耗偶数
		dp[0][1] = 0
		// 第一位放置0时，消耗了一个偶数
		dp[1][0] = 0
	} else {
		// 必须使用a[0]
		dp[0][a[0]&1] = 0
	}

	for i := 1; i < n; i++ {
		for j := 0; j <= cnt[0]; j++ {
			fp[j][0] = n
			fp[j][1] = n
		}
		if a[i] == 0 {
			// 在这里必须决定使用偶数，还是奇数
			for j := 0; j <= cnt[0]; j++ {
				for y := 0; y <= 1; y++ {
					fp[j+check(y == 0)][y] = min(fp[j+check(y == 0)][y], dp[j][y])
					fp[j+check(y == 1)][y^1] = min(fp[j+check(y == 1)][y^1], dp[j][y]+1)
				}
			}
		} else {
			x := a[i] & 1
			for j := 0; j <= cnt[0]; j++ {
				for y := 0; y <= 1; y++ {
					fp[j][x] = min(fp[j][x], dp[j][y]+(x^y))
				}
			}
		}

		dp, fp = fp, dp
	}

	return min(dp[cnt[0]][0], dp[cnt[0]][1])
}

func check(b bool) int {
	if b {
		return 1
	}
	return 0
}

func calc(a []int) int {
	var res int
	for i := 1; i < len(a); i++ {
		res += (a[i-1] & 1) ^ (a[i] & 1)
	}
	return res
}
