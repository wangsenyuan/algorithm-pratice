package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	nums := readNInt64s(reader, 2)
	res := solve(nums[0], nums[1])

	fmt.Println(res)
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
	if len(s) == 0 || len(s) == 1 && s[0] == '\n' {
		return readNInt64s(reader, n)
	}
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

func solve(n int64, k int64) int64 {
	digits := getDigits(n)

	m := len(digits)
	dp := make([]map[int64]int64, m)
	for i := 0; i < m; i++ {
		dp[i] = make(map[int64]int64)
	}

	var dfs func(i int, prod int64, lmt bool, num bool) int64

	dfs = func(i int, prod int64, lmt bool, num bool) int64 {
		if i == m {
			if num && prod <= k {
				return 1
			}
			return 0
		}
		var res int64
		if !lmt && num {
			if v, ok := dp[i][prod]; ok {
				return v
			}
			// 只有在有效的情况下，才缓存结果
			defer func() {
				dp[i][prod] = res
			}()
		}
		if !num {
			res += dfs(i+1, 1, false, false)
		}
		end := 9
		if lmt {
			end = digits[i]
		}
		var start int
		if !num {
			start = 1
		}
		for x := start; x <= end; x++ {
			res += dfs(i+1, prod*int64(x), lmt && x == end, true)
		}
		return res
	}

	return dfs(0, 1, true, false)
}

func solve1(n int64, k int64) int64 {
	// 如果数字中有0, 那么就必然是满足条件的，因为乘积为0
	// 0必须特殊处理。
	// 反过来算，哪些是不符合条件的
	digits := getDigits(n)

	m := len(digits)

	p9 := make([]int64, m+1)
	p9[0] = 1
	for i := 1; i <= m; i++ {
		p9[i] = 9 * p9[i-1]
	}

	dp := make([][]map[int64]int64, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]map[int64]int64, 2)
		for j := 0; j < 2; j++ {
			dp[i][j] = make(map[int64]int64)
		}
	}

	var dfs func(i int, now int64, eq int) int64

	dfs = func(i int, now int64, eq int) int64 {
		if i == m {
			if now > k {
				return 1
			}
			return 0
		}

		if now > k && eq == 0 {
			// 只有在已经小于n的情况下，才能随意取值，但是不能出现0
			return p9[m-i]
		}

		if v, ok := dp[i][eq][now]; ok {
			return v
		}
		end := digits[i]

		var res int64

		if eq == 1 {
			if digits[i] > 0 {
				res += dfs(i+1, now*int64(digits[i]), 1)
			}
		} else {
			end = 10
		}

		for x := 1; x < end; x++ {
			res += dfs(i+1, now*int64(x), 0)
		}

		dp[i][eq][now] = res

		return res
	}
	res := n - dfs(0, 1, 1)

	for i := 1; i < m; i++ {
		res -= dfs(i, 1, 0)
	}

	return res
}

func getDigits(num int64) []int {
	var arr []int
	for num > 0 {
		arr = append(arr, int(num%10))
		num /= 10
	}
	reverse(arr)
	return arr
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func bruteForce(n int64, k int64) int64 {
	var res int64
	for i := 1; i <= int(n); i++ {
		ds := getDigits(int64(i))
		var prod int64 = 1
		for _, x := range ds {
			prod *= int64(x)
		}
		if prod <= k {
			res++
		}
	}
	return res
}
