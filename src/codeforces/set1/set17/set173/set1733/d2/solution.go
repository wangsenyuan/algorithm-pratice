package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, x, y := readThreeNums(reader)
		a := readString(reader)[:n]
		b := readString(reader)[:n]
		res := solve(x, y, a, b)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

const INF = 1 << 60

func solve(x int, y int, a string, b string) int64 {
	if xor(a) != xor(b) {
		return -1
	}
	var diff []int
	n := len(a)

	for i := 1; i <= n; i++ {
		if a[i-1] != b[i-1] {
			diff = append(diff, i)
		}
	}

	if len(diff) == 0 {
		return 0
	}

	if len(diff) == 2 && diff[0]+1 == diff[1] {
		return int64(min(y*2, x))
	}
	// dp[i] = min(dp[i-2] + x * (diff[i] - diff[i-1]), dp[i-1] + y) for even i
	// dp[i] = min(dp[i-2] + x * (diff[i] - diff[i-1]), dp[i-1]) for odd i
	var u int64
	var v int64
	for i := 1; i < len(diff); i++ {
		w := u + int64(x)*int64(diff[i]-diff[i-1])
		if i&1 == 1 {
			w = min(w, v+int64(y))
		} else {
			w = min(w, v)
		}
		u, v = v, w
	}
	return v
}

func solve2(x int, y int, a string, b string) int64 {
	if x > y {
		return solve1(x, y, a, b)
	}
	n := len(a)

	dp := make([][][]int64, 2)
	for d := 0; d < 2; d++ {
		dp[d] = make([][]int64, n)
		for i := 0; i < n; i++ {
			dp[d][i] = make([]int64, n+1)
			for j := 0; j <= n; j++ {
				dp[d][i][j] = INF
			}
		}
	}

	var res int

	if a[0] != b[0] {
		res = 1
		dp[1][0][1] = 0
	} else {
		dp[0][0][0] = 0
	}

	for i := 1; i < n; i++ {
		var c int
		if a[i] != b[i] {
			c = 1
		}
		res ^= c
		if c == 0 {
			for j := i + 1; j >= 0; j-- {
				dp[0][i][j] = min(dp[0][i-1][j], dp[1][i-1][j])
				// 使用前面的某个1，和当前的0，进行flip
				dp[1][i][j] = min(dp[0][i-1][j]+int64(y), dp[1][i-1][j]+int64(x))
				if j > 1 {
					// 建i-1 (0) 和 当前0进行flip, 增加2个1
					// 建 < i - 1的某个位置的0进行flip, 增加2个1
					dp[1][i][j] = min(dp[1][i][j], min(dp[0][i-1][j-2]+int64(x), dp[1][i-1][j-2]+int64(y)))
				}
			}
		} else {
			//dp[0][i][0] = min(dp[0][i-1][1]+int64(y), dp[1][i-1][1]+int64(x))
			for j := i + 1; j >= 0; j-- {
				// 消除掉前面的某个（紧邻的，或者远离的）1
				if j <= i {
					dp[0][i][j] = min(dp[0][i][j], min(dp[0][i-1][j+1]+int64(y), dp[1][i-1][j+1]+int64(x)))
				}
				// 翻转前面某个（紧邻、远离）的0
				if j > 0 {
					dp[0][i][j] = min(dp[0][i][j], min(dp[0][i-1][j-1]+int64(x), dp[1][i-1][j-1]+int64(y)))
					dp[1][i][j] = min(dp[0][i-1][j-1], dp[1][i-1][j-1])
				}
			}
		}
	}

	if res != 0 {
		return -1
	}

	return dp[0][n-1][0]
}
func solve1(x int, y int, a string, b string) int64 {
	if a == b {
		return 0
	}

	n := len(a)
	// n <= 3000
	// x >= y
	// 一次操作 a[i] = 1 - a[i], a[j] = 1 - a[j] 同时改变两个，不变的是？ xor值是不变的
	// 00 => 11, 01 => 10
	ax := xor(a)
	bx := xor(b)

	if ax != bx {
		return -1
	}

	if n == 2 {
		return int64(x)
	}

	if n == 3 {
		// 001 vs 100
		if a[0] != b[0] {
			if a[2] != b[2] {
				return int64(y)
			}
		}
		// 010  001
		return int64(x)
	}

	var diff []int

	for i := 0; i < n; i++ {
		if a[i] != b[i] {
			diff = append(diff, i)
		}
	}

	// 然后我们来操作使得a = b,
	// count of diff % 2 == 0
	// 1100 vs 0000 => min(x, 2 * y)
	var res int64

	if x >= y {

		if len(diff) > 2 || diff[1] > diff[0]+1 {
			res = int64(len(diff)/2) * int64(y)
			// 0 pair half
			// 1 pair half + 1
			// ...
		} else {
			res = int64(min(2*y, x))
		}
	} else {
		// x < y 有点麻烦，就是相邻的不一致进行匹配是更优的方案
		// 假设 u = diff[i], v = diff[i+1]
		// 如果 u + 1 < v , 这时的cost一种选择是y,
		// 还有 连续 flip(u...v) = (v - u) * x
		// 需要dp
		m := len(diff)
		dp := make([][]int64, m)

		cost := func(u int, v int) int64 {
			return min(int64(y), int64(v-u)*int64(x))
		}

		for i := 0; i < m; i++ {
			dp[i] = make([]int64, m)
			for j := 0; j < m; j++ {
				dp[i][j] = INF
			}
			dp[i][i] = INF
			if i+1 < m {
				dp[i][i+1] = cost(diff[i], diff[i+1])
			}
		}
		for j := 1; j < m; j++ {
			for i := j - 1; i >= 0; i -= 2 {
				// 必须是两个进行处理
				if i+1 < j-1 {
					dp[i][j] = min(dp[i][j], int64(y)+dp[i+1][j-1])
				}
				if i+2 < j {
					dp[i][j] = min(dp[i][j], cost(diff[i], diff[i+1])+dp[i+2][j])
				}
				if i < j-2 {
					dp[i][j] = min(dp[i][j], dp[i][j-2]+cost(diff[j-1], diff[j]))
				}
			}
		}
		res = dp[0][m-1]
	}
	return res
}

func xor(a string) int {
	var res int
	for i := 0; i < len(a); i++ {
		res ^= int(a[i] - '0')
	}
	return res
}

type Num interface {
	int | int64
}

func min[T Num](a, b T) T {
	if a <= b {
		return a
	}
	return b
}
