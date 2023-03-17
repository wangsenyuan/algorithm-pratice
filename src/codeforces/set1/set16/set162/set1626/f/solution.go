package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	nums := readNNums(reader, 6)

	res := solve(nums[0], nums[1], nums[2], nums[3], nums[5], nums[4])

	fmt.Println(res)
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

const MOD = 998244353
const L = 720720

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

func pow(a int, b int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r = mul(r, a)
		}
		a = mul(a, a)
		b >>= 1
	}
	return r
}

func inverse(a int) int {
	return pow(a, MOD-2)
}

func div(a int, b int) int {
	return mul(a, inverse(b))
}

func solve(n int, a0 int, x int, y int, m int, k int) int {
	// a[i]的贡献
	// 假设 dp[i][j] = 前i个数选择了j次后的期望值
	// dp[i+1][j] = dp[i][j] 第i个数不选择
	//           + dp[i][j-1] + a[i] 第i个数选一次
	//           + dp[i][j-2] + a[i] + (a[i] - a[i] % i)
	arr := make([]int, n)
	arr[0] = a0
	for i := 1; i < n; i++ {
		arr[i] = int((int64(arr[i-1])*int64(x) + int64(y)) % int64(m))
	}
	total := pow(n, k)

	cur_coeff := div(total, n)

	coeff := mul(k, cur_coeff)

	dp := make([][]int, k)
	for i := 0; i < k; i++ {
		dp[i] = make([]int, L)
	}

	var ans int
	for i := 0; i < n; i++ {
		p, q := arr[i]/L, arr[i]%L
		dp[0][q]++
		ans = add(ans, mul(mul(p, L), coeff))
	}

	for i := 1; i <= k; i++ {
		for j := 0; j < L; j++ {
			cur := dp[i-1][j]
			if i < k {
				dp[i][j] = add(dp[i][j], mul(n-1, cur))
			}
			ans = add(ans, mul(j, mul(cur, cur_coeff)))
			if i < k {
				dp[i][j-(j%i)] = add(dp[i][j-(j%i)], cur)
			}
		}
		cur_coeff = div(cur_coeff, n)
	}

	return ans
}
