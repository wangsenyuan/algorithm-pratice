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
		nums := readNNums(reader, 4)
		res := solve(nums[0], nums[1], nums[2], nums[3])

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

func solve(n int, a int, b int, k int) int {
	// cur -> next if abs(cur - next) < abs(cur - b)
	// dp[x][y] = 进过y步后到达x层时的计数
	// dp[x1][y+1] += dp[x][y] if abs(x - x1) < abs(x - b) 这个可以用差分进行处理
	dp := make([]int, n+2)
	dp[a] = 1
	diff := make([]int, n+2)

	for i := 0; i < k; i++ {
		for j := 1; j <= n; j++ {
			dis := abs(b-j) - 1
			x := max(1, j-dis)
			y := min(n, j+dis)
			if x < j {
				diff[x] = modAdd(diff[x], dp[j])
				diff[j] = modSub(diff[j], dp[j])
			}
			if j <= y {
				diff[j+1] = modAdd(diff[j+1], dp[j])
				diff[y+1] = modSub(diff[y+1], dp[j])
			}
		}

		for j := 1; j <= n; j++ {
			diff[j] = modAdd(diff[j], diff[j-1])
		}
		copy(dp, diff)
		for j := 1; j <= n; j++ {
			diff[j] = 0
		}
	}

	var res int

	for j := 1; j <= n; j++ {
		res = modAdd(res, dp[j])
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

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

const MOD = 1000000007

func modMul(a, b int) int {
	return int(int64(a) * int64(b) % MOD)
}

func modAdd(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func modSub(a, b int) int {
	return modAdd(a, MOD-b)
}

func pow(a, b int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r = modMul(r, a)
		}
		a = modMul(a, a)
		b >>= 1
	}
	return r
}
