package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	s := readString(reader)
	n := readNInt64s(reader, 1)[0]
	res := solve(s, n)
	fmt.Printf("%d\n", res)
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

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func solve(s string, n int64) int64 {

	tmp := n

	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			tmp -= 1 << (len(s) - 1 - i)
		}
	}
	if tmp < 0 {
		return -1
	}
	// tmp no change
	var res int64
	for i := 0; i < len(s); i++ {
		x := s[i]
		if x == '?' {
			if tmp >= 1<<(len(s)-1-i) {
				x = '1'
				tmp -= 1 << (len(s) - 1 - i)
			} else {
				x = '0'
			}
		}
		res *= 2
		if x == '1' {
			res++
		}
	}

	return res
}

func solve1(s string, n int64) int64 {
	// change s[i] = ? to 0 or 1
	var digits []int
	for n > 0 {
		digits = append(digits, int(n&1))
		n >>= 1
	}

	for len(digits) < len(s) {
		digits = append(digits, 0)
	}
	reverse(digits)

	if len(digits) > len(s) {
		// change all ? to 1
		var res int64
		for i := 0; i < len(s); i++ {
			res *= 2
			if s[i] == '0' {
				continue
			}
			res++
		}
		return res
	}
	// len(digits) == len(s)

	dp := make([]int64, 2)
	// dp[0] 表示已经比之前的高位小的情况下
	// dp[1] 表示到目前为止还是相等的情况下
	dp[0] = -1
	dp[1] = 0

	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			// no change
			if digits[i] == 0 {
				dp[1] = 2 * dp[1]
				dp[0] = 2 * dp[0]
			} else {
				// digits[i] = 1
				// 当前数字无法改变，只能是0，如果已经比当前数字小了，*2,
				// 或者之前是相等的，但是当前 0 < 1
				dp[0] = max(2*dp[0], 2*dp[1])
				// not valid anymore
				dp[1] = -1
			}
		} else if s[i] == '1' {
			if digits[i] == 1 {
				dp[0] = 2*dp[0] + 1
				dp[1] = 2*dp[1] + 1
			} else {
				dp[0] = dp[0]*2 + 1
				// dp[1] is invalid
				dp[1] = -1
			}
		} else {
			// s[i] = ?
			if digits[i] == 1 {
				// 如果已经比数字小了，当前数字可以变成1，或者原来相等，当前数字变为0
				dp[0] = max(2*dp[0]+1, 2*dp[1])
				// 如果要保持相等，当前数字只能1
				dp[1] = 2*dp[1] + 1
			} else {
				//  如果要保持相等，当前数字只能0
				dp[1] = 2 * dp[1]
				//  如果已经比数字小了，当前数字可以变成1
				dp[0] = 2*dp[0] + 1
			}
		}
	}
	res := max(dp[0], dp[1])
	return max(res, -1)
}

func reverse(num []int) {
	for i, j := 0, len(num)-1; i < j; i, j = i+1, j-1 {
		num[i], num[j] = num[j], num[i]
	}
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}
