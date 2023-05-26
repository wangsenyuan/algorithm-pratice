package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := 1
	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, k := readTwoNums(reader)
		V := readNNums(reader, n)
		res := solve(V, k)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}
func solve(V []int, K int) int64 {
	n := len(V)
	// V <= 1200, K <= 1200
	// dp[i][j][k] 表示到i为止时，如果在i处购买了j个商品，且总共购买了k件商品时的最大价值
	// dp[i+1][x][y] = max(dp[i][j][k] + x * v[i], and x <= j + 1, k + x = y
	//  这样子时 N * K * K的复杂性
	//  这里另外一个关键的信息时 x <= j + 1
	//  dp[i][x] 表示从i开始，有x次购买机会，且在i处只能购买1个时的最大值
	// 还是不行，如果在i处购买1，那么在i+1处，就可以购买2，一定是涨上去更优吗？
	// 假设某个j， V[j]非常大，我们希望能买到最多的它
	// 那么为了买到x个, V[j]， 必须花费 (1 + x) * x / 2 个机会
	// dp[i][x] 表示在i处剩余x次机会时的最大总价值
	//  在 i处购买j个产品时, (1 + j) * j / 2 <= x
	//   剩余的就是 dp[i-1][x - s]
	sq := make([]int, K+1)

	for i := 1; (i+1)*i/2 <= K; i++ {
		j := (i + 1) * i / 2
		sq[j] = i
	}

	for i := 1; i <= K; i++ {
		if sq[i] == 0 {
			sq[i] = sq[i-1]
		}
	}

	dp := make([]int64, K+1)
	fp := make([]int64, K+1)

	kk := sq[K] + 1

	for i := n - 1; i >= 0; i-- {
		for j := 0; j <= K; j++ {
			var sum int64
			for k := i; k < min(n, i+kk); k++ {
				sum += int64(V[k])
				l := k - i + 1
				if l > j {
					break
				}
				fp[j] = max(fp[j], dp[j-l]+sum)
			}
			if j > 0 {
				fp[j] = max(fp[j], fp[j-1])
			}
		}
		copy(dp, fp)
	}

	return dp[K]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}
