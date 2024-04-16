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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
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

const X = 200

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

func solve(a []int) int {
	dp := make([][][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([][]int, 2)
		for j := 0; j < 2; j++ {
			dp[i][j] = make([]int, X+1)
		}
	}

	if a[0] == -1 {
		for x := 1; x <= X; x++ {
			dp[0][0][x] = 1
		}
	} else {
		dp[0][0][a[0]] = 1
	}

	n := len(a)
	for i := 1; i < n; i++ {
		// reset
		for j := 0; j < 2; j++ {
			for x := 0; x <= X; x++ {
				dp[i&1][j][x] = 0
			}
		}
		if a[i] != -1 {
			var tmp int
			for x := 1; x < a[i]; x++ {
				tmp = add(tmp, dp[(i-1)&1][0][x])
				// 如果a[i-1] = x, 且它满足 a[i-1] <= a[i-2]
				tmp = add(tmp, dp[(i-1)&1][1][x])
			}
			// 如果 j = 0, 那么当前的y要比前面的x更大
			// 0 => 要求 a[i] > x
			// 如果a[i-1] = a[i] 也可以满足前面的条件
			dp[i&1][0][a[i]] = tmp
			tmp = 0
			for x := a[i] + 1; x <= X; x++ {
				// a[i-1] >= a[i]时，前面的必须满足 a[i-1] <= a[i-2]
				tmp = add(tmp, dp[(i-1)&1][1][x])
			}
			// 如果 a[i-1] = a[i], 也满足 a[i] <= a[i-1]的条件
			tmp = add(tmp, dp[(i-1)&1][0][a[i]])
			tmp = add(tmp, dp[(i-1)&1][1][a[i]])

			dp[i&1][1][a[i]] = tmp
		} else {
			// 当前可以选择所有的数
			var sum int
			// 如果选择比前面的数大
			for x := 1; x <= X; x++ {
				dp[i&1][0][x] = sum
				// 前面是x，且还不满足 a[i-1] <= a[i-2]的条件
				sum = add(sum, dp[(i-1)&1][0][x])
				// 前面是x，但已经满足 a[i-1] <= a[i-2]的条件
				sum = add(sum, dp[(i-1)&1][1][x])
			}
			sum = 0
			for x := X; x > 0; x-- {
				dp[i&1][1][x] = add(sum, add(dp[(i-1)&1][0][x], dp[(i-1)&1][1][x]))
				// 因为 a[i] < a[i-1], 所以 a[i-1] <= a[i-2]
				sum = add(sum, dp[(i-1)&1][1][x])
			}
		}
	}
	var res int

	for x := 1; x <= X; x++ {
		res = add(res, dp[(n-1)&1][1][x])
	}

	return res
}
