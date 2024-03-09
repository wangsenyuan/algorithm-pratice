package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(reader, &n)

	a := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	res := solve(a)

	fmt.Println(res)
}

const mod = 998244353

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(a, b int) int {
	return a * b % mod
}

func pow(a, b int) int {
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

func sub(a, b int) int {
	return add(a, mod-b)
}

func inverse(a int) int {
	return pow(a, mod-2)
}

var i2 int

func init() {
	i2 = inverse(2)
}

func solve(a []int) int {
	// 假设这样一个序列b, 满足条件
	// 如果b的长度正好 = 3， 那么其中必须有奇数个偶数
	// 如果长度 = 4, 1, 0，1, 1 (或者全部是偶数）
	// 如果长度 > 4, 0, 1, 1, 0, 1, 1, 0 （或者全部是偶数)
	//              1, 0, 1, 1, 0, 1, 1, 0
	// 很明显就是存在一个pattern
	// dp[x][j] 表示第x%3个位置为0/1时的总计数
	// 如果a[i]是偶数，要怎么添加进去呢？
	// 如果最后两个数的sum是偶数，那么a[i]就可以加进去
	// dp[3][0] = dp[2][0]
	// res += dp[3][0]
	// 增加一个新的偶数进来时，前面两个数必须是偶数/奇数，
	// 增加一个新的奇数进来时，前面必须一个奇数，一个偶数
	// 00, 11, 01, 10
	// 0表示没有数字的状态
	// 1表示有一个0的状态
	// 2表示有一个1的状态
	// 3表示00
	// 4表示01
	// 5表示10
	// 6表示11
	// dp[x] = 前面两个数是这些状态时的奇数
	dp := make([]int, 7)
	dp[0] = 1
	var cnt0 int
	var res int
	for _, num := range a {
		x := num & 1
		if x == 1 {
			// 011/101 都是有效状态
			res = add(res, dp[4])
			res = add(res, dp[5])
			dp[6] = add(dp[6], dp[2])
			// 状态4 =》 6， 且是一个有效状态
			dp[6] = add(dp[6], dp[4])
			// 状态5 =》 4
			dp[4] = add(dp[4], dp[5])
			dp[4] = add(dp[4], dp[1])
			dp[2] = add(dp[2], dp[0])
		} else {
			cnt0++
			//res = add(res, dp[3])
			res = add(res, dp[6])
			// 状态3 => 新的一个状态3
			//dp[3] = add(dp[3], 1)
			// 状态1 => 状态3
			dp[3] = add(dp[3], dp[1])
			// 状态0 => 状态1
			dp[1] = add(dp[1], dp[0])
			// 状态6 => 状态5
			dp[5] = add(dp[5], dp[6])
			// 状态2 => 状态5, 状态4不能变成状态5，是因为010是一个无效的组合
			dp[5] = add(dp[5], dp[2])
		}
	}
	if cnt0 >= 3 {
		tmp := pow(2, cnt0)
		// 选一个
		tmp = sub(tmp, cnt0)
		// 选两个
		tmp = sub(tmp, mul(mul(cnt0, cnt0-1), i2))
		tmp = sub(tmp, 1)

		res = add(res, tmp)
	}
	return res
}
